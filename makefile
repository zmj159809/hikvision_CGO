# 定义变量
BINARY_NAME = testfile  # 生成的二进制文件名
GO_CMD = go           # Go 命令
GO_BUILD = $(GO_CMD) build  # 构建命令
SRC = example/main.go  # 源文件路径

# 默认目标（直接运行 `make` 时执行）
.DEFAULT_GOAL := build

# 构建目标：生成二进制文件
build:
        $(GO_BUILD) -o $(BINARY_NAME) $(SRC)

# 运行目标：构建并执行二进制文件
run: build
        ./$(BINARY_NAME)

# 清理目标：删除生成的二进制文件
clean:
        rm -f $(BINARY_NAME)

# 声明伪目标（避免与同名文件冲突）
.PHONY: build run clean
~