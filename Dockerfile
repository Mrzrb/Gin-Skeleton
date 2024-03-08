FROM golang:alpine as builder
WORKDIR /app
COPY . .

ENV GOPROXY=https://goproxy.cn,direct
# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build  -o /app/main .

# 第二阶段：最小化的运行时环境
FROM alpine
# 从第一阶段的构建镜像中复制构建好的应用程序到最小化的运行时镜像中
RUN apk update && apk add tzdata
COPY --from=builder /app/main /main
COPY --from=builder /app/config /config
# 设置容器启动时运行的命令
CMD ["/main"]
