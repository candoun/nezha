# docker build -t nezha-frontend -f   hack/build/frontend/Dockerfile . 

FROM nezha-frontend-builder as builder
WORKDIR /workspace

COPY vue-admin /workspace

RUN cd /workspace && \
    npm run build:prod


FROM openresty/openresty:1.15.8.1-1-centos

COPY --from=builder /workspace/dist/ /usr/local/openresty/nginx/html/

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    sed -i '/index  index.html index.htm;/a\        try_files $uri $uri/ /index.html;' /etc/nginx/conf.d/default.conf

CMD ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]

