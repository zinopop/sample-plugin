module gomod.alauda.cn/katanomi-plugin-sample

go 1.16

require (
	github.com/cloudevents/sdk-go/v2 v2.5.0
	github.com/emicklei/go-restful/v3 v3.8.0
	github.com/katanomi/core v0.2.1-0.20220322074511-159920df7cb0
	github.com/katanomi/pkg v0.3.1-0.20220620154556-338f8e9d3857
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.19.0
	go.uber.org/zap v1.18.1
	k8s.io/api v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v0.20.7
	knative.dev/eventing v0.24.0
	knative.dev/pkg v0.0.0-20210730172132-bb4aaf09c430
	sigs.k8s.io/controller-runtime v0.8.3
)

replace (
	github.com/opencontainers/image-spec => github.com/opencontainers/image-spec v1.0.2-0.20210819154149-5ad6f50d6283
	go.uber.org/zap => github.com/katanomi/zap v1.18.2 // indirect
	sigs.k8s.io/controller-runtime v0.8.3 => github.com/katanomi/controller-runtime v0.8.3-1
)
