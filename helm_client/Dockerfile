FROM alpine:3.7

RUN apk add -U openssl curl && \
  curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get > get_helm.sh && \
  sh get_helm.sh --version v2.8.0 && \
  rm -rf get_helm.sh
