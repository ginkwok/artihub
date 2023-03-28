# 基础镜像
FROM golang:1.20

# 设置工作目录
WORKDIR /app

# 设置 Go 环境
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct

# 将源代码复制到容器中
COPY . .

# 设置环境变量
ENV APP_ENV production

# 构建应用
RUN go build -o app .

# 暴露端口
EXPOSE 8080

# 启动应用
CMD ["./app"]
