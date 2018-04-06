FROM alpine:3.7

RUN apk install -U openssl curl && \
  curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get > get_helm.sh && \
  sh get_helm.sh && \
  rm -rf get_helm.sh
