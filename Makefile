# Component name inside cmd/
C ?= sample-plugin
# Full image address for docker build
IMG ?= registry.cn-beijing.aliyuncs.com/alauda-yuzheng/katanomi-plugin-$(C)
# Tag to be used in the docker image
TAG ?= dev-$(shell git config --get user.email | sed -e "s/@/-/")-$(shell git rev-parse --short HEAD)-$(shell date +%Y-%m-%d-%H-%I-%s)
# Architecture to be used during binary building
ARCH ?= amd64
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true,preserveUnknownFields=false"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

# need to download gitversion in advance
GITNAME = $(shell git config --get user.name | sed 's/ //g')
BUNDLE_VERSION ?= $(shell gitversion /output json /showvariable MajorMinorPatch)-$(shell gitversion /output json /showvariable ShortSha)-$(GITNAME)
CONTROLLER_VERSION ?= $(BUNDLE_VERSION)


all: lint test

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

manifests: generate controller-gen base-rbac ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.

base-rbac: ## Generate api rbac objects listed inside cmd/plugin
	$(CONTROLLER_GEN) rbac:roleName='plugin-role' paths="./cmd/base" output:rbac:artifacts:config=charts/base/rbac
	find ./charts/base/rbac -name 'role.yaml' -exec sed -i '' 's/name: plugin-role/name: {{ .Values.plugin }}-role/' "{}" \;

generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

lint: golangcilint ## Run golangci-lint against code.
	$(GOLANGCILINT) run

ENVTEST_ASSETS_DIR=$(shell pwd)/testbin
test: fmt vet goimports ## Run tests.  manifests generate
	mkdir -p ${ENVTEST_ASSETS_DIR}
	test -f ${ENVTEST_ASSETS_DIR}/setup-envtest.sh || curl -sSLo ${ENVTEST_ASSETS_DIR}/setup-envtest.sh https://raw.githubusercontent.com/kubernetes-sigs/controller-runtime/v0.8.3/hack/setup-envtest.sh
	source ${ENVTEST_ASSETS_DIR}/setup-envtest.sh; fetch_envtest_tools $(ENVTEST_ASSETS_DIR); setup_envtest_env $(ENVTEST_ASSETS_DIR); go test ./... -coverprofile cover.out

##@ Deployment

build:  ## Builds a binary using C=<name> i.e  make build C=controller
	GOOS=linux GOARCH=$(ARCH) go build -o bin/$(ARCH)/$(C) cmd/$(C)/main.go

docker: ## Builds a docker image using C=<name> i.e  make build C=controller
	docker build -t $(IMG):$(TAG) -f ./images/Dockerfile --build-arg component=$(C) .
	docker push $(IMG):$(TAG)

deploy: ko manifests certmanager ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	helm template ./charts/$(C) --namespace katanomi-system \
		--set replicas=1 \
		--set namespace=katanomi-system \
		--set ko.enable=true | $(KO) apply -P ${LOCAL} -f -

undeploy: ko ## Undeploy controller from the K8s cluster specified in ~/.kube/config.
	helm template ./charts/${CHART_VERSION} --namespace katanomi-system --set global.namespace=katanomi-system  --set ko=true | $(KO) delete -f -

apply: kustomize ko certmanager ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/default | $(KO) apply -P ${LOCAL} -f -

unapply: ## Undeploy controller from the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/default | $(KO) delete -f -

certmanager: ## Install certmanager v1.4.0 from github manifest to the K8s cluster specified in ~/.kube/config.
	$(call installyaml,cert-manager,https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml,cert-manager)

##@ Bundle
.PHONY: list-latest-charts-values
list-latest-charts-values: ## 列出所有工具最新版本chart的values文件路径
	@cat charts/charts.latest.values.list | awk '{ printf("./charts/%s", $$0) }'| xargs echo

##@ Automated Testing

e2e: ginkgo ## Executes e2e tests inside test/e2e folder
	$(GINKGO) -progress -v -tags e2e ./test/e2e

##@ Setup

CONTROLLER_GEN = $(shell pwd)/bin/controller-gen
controller-gen: ## Download controller-gen locally if necessary.
	## this is a necessary evil already reported by knative community https://github.com/kubernetes-sigs/controller-tools/ issue 560
	## once the issue is fixed we can move to use the original package. the original line uses go-get-tools with sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1
	$(call go-get-fork,$(CONTROLLER_GEN),https://github.com/katanomi/controller-tools,cmd/controller-gen,controller-gen,master)

KUSTOMIZE = $(shell pwd)/bin/kustomize
kustomize: ## Download kustomize locally if necessary.
	$(call go-get-tool,$(KUSTOMIZE),sigs.k8s.io/kustomize/kustomize/v3@v3.8.7)

.PHONY: opm
# find or download opm
# download opm if necessary
opm: ## 安装 opm
ifeq (, $(shell which opm))
	@{ \
	set -e ;\
	OPM_TMP_DIR=$$(mktemp -d) ;\
	cd $$OPM_TMP_DIR ;\
	git clone --branch v1.15.2 --depth 1 https://github.com/operator-framework/operator-registry.git && cd operator-registry ;\
	export CGO_ENABLED=1 ;\
	make build ;\
	cp ./bin/* $(GOBIN)/ ;\
	rm -rf $$OPM_TMP_DIR ;\
	}
OPM=$(GOBIN)/opm
else
OPM=$(shell which opm)
endif

KO = $(shell pwd)/bin/ko
ko: ## Download ko locally if necessary.
	$(call go-get-tool,$(KO),github.com/google/ko@v0.8.3)

GOIMPORTS = $(shell pwd)/bin/goimports
goimports: ## Download goimports locally if necessary.
	$(call go-get-tool,$(GOIMPORTS),golang.org/x/tools/cmd/goimports)
	$(GOIMPORTS) -w -l $(shell find . -path '.git' -prune -path './vendor' -prune -o -path './examples' -prune -o -name '*.pb.go' -prune -o -type f -name '*.go' -print)

GINKGO = $(shell pwd)/bin/ginkgo
ginkgo: ## Download ginkgo locally if necessary
	$(call go-get-tool,$(GINKGO),github.com/onsi/ginkgo/ginkgo@v1.16.4)

GOLANGCILINT = $(shell pwd)/bin/golangci-lint
golangcilint: ## Download golangci-lint locally if necessary
	$(call go-get-tool,$(GOLANGCILINT),github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1)

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef

# go-get-fork is a "go-get-tool" like command to get temporary module forks.
define go-get-fork
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
echo "Cloning $(2)" ;\
git clone $(2) $(4) ;\
cd $(4) ;\
git checkout $(5) ;\
GOBIN=$(PROJECT_DIR)/bin go install ./$(3);\
rm -rf $$TMP_DIR ;\
}
endef

# installyaml will check if a given namespace is present, if not will apply a yaml file and wait for a deployment to rollout
define installyaml
kubectl get ns $(1) > /dev/null || { \
set -ex ;\
kubectl apply -f $(2) ;\
kubectl -n $(1) rollout status deploy/$(3) --timeout=10m ;\
}
endef
