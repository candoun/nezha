# docker build -t nezha-backend -f   hack/build/backend/Dockerfile . 

FROM nezha-backend-builder as builder
WORKDIR /build

COPY ./  ./
RUN go build -x  -o /build/nezha ./cmd/main.go



FROM centos:7

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY config/config.yml /apps/nezha/config/config.yml
COPY --from=builder /build/nezha /apps/nezha/bin/nezha

WORKDIR /apps/nezha/bin/

CMD ["./nezha"]

