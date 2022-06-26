# 命令字段
.PHONY: dev

# 加载配置文件，导入PROJECT_NAME参数
include ./docker/dev/.env
PWD := $(shell pwd)
DEV_NAME = $(PROJECT_NAME)-$(DEV_USER)

# 启动后端服务开发环境
dev:
	cd docker/dev && docker-compose -p "$(DEV_NAME)" down && docker-compose -p "$(DEV_NAME)" up --force-recreate