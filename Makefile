.PHONY: build proto clean help cmd statbuild

Executor_TARGET = executor-server

BUILD_TIME = $(shell date -R)
GIT_COMMIT = $(shell git rev-parse --short HEAD)
GO_VERSION = $(word 3, $(shell go version))
OS_ARCH    = $(word 4, $(shell go version))

GROUP = executor
GROUP_EXIST = $(findstring $(GROUP):x,$(shell cat /etc/group))

all: build run

build: build-executor stat

build-executor:
	go build -v -ldflags \
	  "-X 'github.com/yanlingqiankun/executor/daemon.GitCommit=${GIT_COMMIT}'\
	  -X 'github.com/yanlingqiankun/executor/daemon.GoVersion=${GO_VERSION}' \
	  -X 'github.com/yanlingqiankun/executor/daemon.OSArch=${OS_ARCH}' \
	  -X 'github.com/yanlingqiankun/executor/daemon.BuildTime=${BUILD_TIME}'" -o ${Executor_TARGET}

cmd:
	cd cmd && go build -v -o ../executor

proto:
	protoc --go_out=plugins=grpc,paths=source_relative:. pb/*.proto

clean:
	go clean --cache -i .
	cd cmd && go clean -i .

cleanall:
	go clean --cache -r -i .
	cd cmd && go clean -r -i .
	rm -f pb/*.pb.go

run:
	sudo -E ./executor-server

stat:
	@echo "生成文件大小: " $(shell echo "scale=3; $(shell stat -c %s executor-server)/1024/1024" |bc) "MB"


install:
ifeq ($(GROUP_EXIST),)
	sudo groupadd -f $(GROUP)
	sudo usermod -a -G $(GROUP) $(shell whoami)
	newgrp $(GROUP)
endif

help:
	@echo "=======使用说明======="
	@echo "make: 编译、运行executor-server"
	@echo "make build: 编译executor-server"
	@echo "make proto: 编译protobuf生成go文件"
	@echo "make install: 设置用户组"
	@echo "make clean: 清理二进制文件和编译缓存"
	@echo "make cleanall: 清理编译缓存，import库，proto文件"
	@echo "make run：运行编译生成的executor-server"
