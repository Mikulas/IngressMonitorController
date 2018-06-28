FROM golang:1.10-alpine3.7 AS build-env
RUN apk update && \
    apk add \
        curl \
        git \
    && \
    curl https://glide.sh/get | sh

ADD src /go/src/github.com/mangoweb/ingress-monitor
ADD vendor /go/src/github.com/mangoweb/ingress-monitor/vendor
RUN cd /go/src/github.com/mangoweb/ingress-monitor && \
    go build -o ingress-monitor


FROM stakater/base-alpine:3.7
LABEL author="stakater"

COPY --from=build-env /go/src/github.com/mangoweb/ingress-monitor/ingress-monitor /

ENTRYPOINT [ "/ingress-monitor" ]
