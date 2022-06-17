FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /data/app/fginx/server

# 将代码复制到容器中
COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env
RUN go mod tidy
RUN go build -o server .

# 声明服务端口
EXPOSE 8888

# 启动容器时运行的命令
ENTRYPOINT ./server -c config.docker.yaml