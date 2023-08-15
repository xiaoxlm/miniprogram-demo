FROM dockerproxy.com/library/golang:1.19-buster AS build-env

FROM build-env AS builder

ARG GOPROXY=https://goproxy.cn,direct
WORKDIR /go/src
COPY ./ ./

# build
RUN make build WORKSPACE=miniprogram

# runtime
FROM alpine
COPY --from=builder /go/src/cmd/miniprogram/miniprogram /go/bin/app1

EXPOSE 80
COPY --from=builder /go/src/cmd/miniprogram/openapi.json /go/bin/openapi.json

ARG PROJECT_NAME
ARG PROJECT_VERSION
ENV GOENV=DEV PROJECT_NAME=${PROJECT_NAME} PROJECT_VERSION=${PROJECT_VERSION}

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/app1"]
