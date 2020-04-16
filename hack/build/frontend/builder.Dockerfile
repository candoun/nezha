# docker build -t nezha-frontend-builder -f   hack/build/frontend/builder.Dockerfile . 
FROM node:10.18-alpine3.9
WORKDIR /workspace

COPY vue-admin/package.json  /workspace

# COPY vue-admin/package-lock.json  /workspace

RUN export http_proxy=http://proxy_ip:8080 && \
    export https_proxy=http://proxy_ip:8080 && \
    npm config set loglevel=http && \
    npm install --registry=https://registry.npm.taobao.org node-sass 

RUN npm install 

