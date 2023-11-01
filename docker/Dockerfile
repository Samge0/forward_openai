# 使用 golang 官方镜像提供 Go 运行环境，并且命名为 buidler 以便后续引用
FROM samge/forward_openai:base

# 设置工作目录
WORKDIR /app

# 将代码复制到容器中的工作目录
COPY . .
RUN cp config.dev.json config.json

# 下载依赖包
RUN go mod download

# 构建应用程序
RUN go build -o main .

# 暴露应用程序的端口
EXPOSE 8080

# 运行应用程序
CMD ["./main"]