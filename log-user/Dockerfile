FROM golang:1.19-alpine
#为镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

#工作目录
WORKDIR /project/go-docker/mylog

#下载依赖
ADD log_user /log_user
RUN go mod download
RUN go build -o /project/go-docker/mylog/build/log_user .


ENTRYPOINT ["/project/go-docker/mylog/build/log_user"]