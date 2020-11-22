#!/bin/bash
cd $WORKSPACE

export GOPROXY=https://goproxy.io

go mod tidy
# 打印依赖，部署成功后查看版本依赖是否如预期
cat ./go.mod

# linux环境编译
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
# 构建docker镜像，项目中需要在当前目录下有dockerfile，否则构建失败
docker build -t demo .
#rm -rf main
#docker rm http_test -f
docker run --name http_test -p 9999:9999 -d demo