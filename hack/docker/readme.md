
# 开发环境
docker run -itd  --name nezha-mysql -e MYSQL_ROOT_PASSWORD="gin-vue-nezha@123" mariadb:10.3 


docker run -itd \
--name nezha-dev   \
--restart=always \
-p8003:8000 \
--link nezha-mysql:nezha-mysql \
-v ~/whj/dockerhub/whj/nezha:/app \
nezha-backend-builder 
#bash -c cd /app/cmd/ &&  go run main.go  -migrate=true

docker run -it \
--entrypoint=sh   \
--rm \
-v ~/whj/dockerhub/whj/nezha/vue-admin:/app  \
-p80:9528 \
nezha-frontend-builder \

# "ln -s /workspace/node_modules/   /app/node_modules && cd /app  && npm run dev"


# 运行环境
docker run -itd --name nezha-mysql -e MYSQL_ROOT_PASSWORD="gin-vue-nezha@123"   mariadb:10.3

docker run -it -p 8005:8000--name nezha-backend 

docker run -it -p 80:80 --name nezha-frontend 
