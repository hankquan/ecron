BINARY_NAME=./output/ecron
VERSION := $(shell git describe --tags --always)
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S)

.PHONY: build

build:
	go build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o $(BINARY_NAME) main.go

# 修改版本
# release 打包发布到仓库