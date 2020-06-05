# Executor

##  概述

一个将Docker和Libvirt进行封装的工具，可以使用同样的方法部署容器和虚拟机。

目前完成了机器、卷、网络、镜像的封装

## 安装

[安装Docker](https://docs.docker.com/engine/install/)

[安装Libvirt](https://help.ubuntu.com/community/KVM/Installation)

如果需要编译proto文件生成服务端和客户端代码，则还需要安装以下依赖

- protoc
- protoc-gen-go

使用makefile可以编译整个项目

- make 编译，运行
- make cmd 编译命令行工具
- make build 编译整个项目
- make proto 根据protobuf文件生成go代码
- make clean 清理编译生成的可执行文件和protobuf生成代码
- make help  显示帮助信息
