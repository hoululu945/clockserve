# 使用一个基础镜像作为起点
FROM golang:1.20.5

# 设置工作目录
WORKDIR /app

# 复制项目文件到容器中
COPY . .

# 运行 go mod tidy 命令来整理依赖关系
RUN go mod tidy

# 构建 Go 项目
RUN go build -o main

# 设置容器启动命令
CMD ["./main"]