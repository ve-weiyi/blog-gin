.PHONY: help deps swag run build clean test lint

BIN_NAME := blog-gin

help:
	@echo "可用命令："
	@echo "  make deps      - 安装依赖"
	@echo "  make swag      - 生成 Swagger 文档"
	@echo "  make run       - 启动服务（开发模式）"
	@echo "  make build     - 编译二进制文件"
	@echo "  make clean     - 清理编译文件"
	@echo "  make test      - 运行测试"
	@echo "  make lint      - 运行代码检查"

deps:
	go mod tidy
	go mod download

swag:
	swag init -g main.go -d ./,api/ --parseDependency

run:
	go run main.go -f config.yaml

build:
	go build -o bin/$(BIN_NAME) main.go
	@echo "编译完成 -> bin/$(BIN_NAME)"

clean:
	rm -rf bin/
	@echo "清理完成！"

test:
	go test ./... -v

lint:
	go vet ./...
