BINARY_NAME=./output/ecron

.PHONY: build

build:
	go build -o $(BINARY_NAME) main.go

# 修改版本
# release 打包发布到仓库