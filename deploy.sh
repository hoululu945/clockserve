#!/bin/bash

# 设置错误处理
set -e

# 定义变量
BUILD_DIR="build"
DEPLOY_DIR="deploy"
EXECUTABLE="your_executable_name"  # 替换为编译后的可执行文件名称
SERVICE_NAME="$EXECUTABLE"           # 将服务名称设置为可执行文件名称

# 检查并创建构建目录
mkdir -p "$BUILD_DIR"

# 构建 Go 项目
echo "开始构建 Go 项目..."
if [ ! -f go.mod ]; then
    go mod init serve
fi
go mod tidy
go build -o "$BUILD_DIR/$EXECUTABLE" main.go

# 检查构建物是否存在
if [ ! -f "$BUILD_DIR/$EXECUTABLE" ]; then
    echo "构建失败，未找到可执行文件 $BUILD_DIR/$EXECUTABLE"
    exit 1
fi

# 复制构建物到部署目录
echo "开始复制构建物到部署目录..."
mkdir -p "$DEPLOY_DIR"

# 停止服务
echo "停止服务 $SERVICE_NAME..."
systemctl stop your_executable_name || echo "服务 your_executable_name 停止失败，可能未运行。"

# 强制停止进程（如果服务未能停止）
pkill -f "$EXECUTABLE" || echo "未找到正在运行的进程，继续复制..."

# 删除旧的可执行文件（可选）
if [ -f "$DEPLOY_DIR/$EXECUTABLE" ]; then
    echo "删除旧的可执行文件..."
    rm "$DEPLOY_DIR/$EXECUTABLE"
fi

# 复制新的可执行文件
cp "$BUILD_DIR/$EXECUTABLE" "$DEPLOY_DIR/"

# 启动服务
echo "启动服务 $SERVICE_NAME..."
nohup "$DEPLOY_DIR/$EXECUTABLE" > "$DEPLOY_DIR/logs.txt" 2>&1 &

echo "部署完成！"