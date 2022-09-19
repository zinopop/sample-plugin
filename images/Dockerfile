FROM build-harbor.alauda.cn/ops/alpine:3.14.0 as builder

ARG component=controller
COPY ./bin /workspace/bin
COPY ./charts /workspace/charts
COPY .git /workspace/.git

RUN apk update && apk add git
RUN cd /workspace && git rev-parse HEAD > /bin/HEAD

RUN ARCH= &&dpkgArch="$(arch)" \
  && case "${dpkgArch}" in \
  x86_64) ARCH='amd64';; \
  aarch64) ARCH='arm64';; \
  *) echo "unsupported architecture"; exit 1 ;; \
  esac \
  && cp /workspace/bin/${ARCH}/${component} /bin/${component}

RUN set -eux ; \
  if [[ "${component}" = "operator" ]] ; then \
  cp -rf /workspace/charts /bin/charts ; \
  else \
  mkdir -p /bin/charts/ && touch /bin/charts/.keep ; \
  fi

FROM build-harbor.alauda.cn/ops/alpine:3.14.0

RUN apk add openssl
ARG commit_id=dev
ARG app_version=dev
ENV COMMIT_ID=${commit_id}
ENV APP_VERSION=${app_version}

ARG component=controller
ENV KO_DATA_PATH=/kodata

COPY attributes/${component}.yaml /kodata/attributes/${component}.yaml
COPY integrationclasses/${component}.yaml /kodata/integrationclasses/${component}.yaml
COPY --from=builder /bin/${component} /${component}
COPY --from=builder /bin/charts /kodata/charts
COPY --from=builder /bin/HEAD /kodata/HEAD

# variable substitution is not possible using ENTRYPOINT unless it is used in shell form
# https://stackoverflow.com/questions/37904682/how-do-i-use-docker-environment-variable-in-entrypoint-array#37904830
RUN echo "#!/bin/sh " > /entrypoint.sh && \
  echo "/${component} \$@" >> /entrypoint.sh && \
  chmod +x /entrypoint.sh && cat /entrypoint.sh && \
  # ls -la /kodata/tasks && ls -la /kodata/stages &&  \
  ls -la /kodata/charts

USER 1000:1000

ENTRYPOINT [ "/entrypoint.sh" ]
