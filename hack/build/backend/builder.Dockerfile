# docker build -t nezha-backend-builder -f   hack/build/backend/builder.Dockerfile . 
FROM golang:1.13.6
WORKDIR /go/src/github.com/aguncn/nezha

COPY cmd/main.go /go/src/github.com/aguncn/nezha

RUN export GO111MODULE=on && \
    export http_proxy=http://proxy_ip:8080 && \
    export https_proxy=http://proxy_ip:8080 && \
    export GOPROXY=https://goproxy.cn && \
    go install /go/src/github.com/aguncn/nezha/main.go

