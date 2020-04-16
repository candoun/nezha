# docker build -t nezha-frontend-builder -f   hack/build/frontend/builder.Dockerfile . 
FROM node:10.18-alpine3.9
WORKDIR /workspace

COPY vue-admin/package.json  /workspace

# COPY vue-admin/package-lock.json  /workspace

RUN export http_proxy=http://10.17.171.129:8080 && \
    export https_proxy=http://10.17.171.129:8080 && \
    npm config set loglevel=http && \
    npm install --registry=https://registry.npm.taobao.org node-sass 

RUN npm install --registry=http://nexus.stock.paic.com.cn:30001/repository/npm-group/

