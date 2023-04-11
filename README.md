# Honeysheet-server

# 部署本地开发环境

1. 在 docker/dev 目录下创建 .env 文件。
2. 拷贝 docker/default.env 文件的内容到 .env 文件。
3. 执行`docker network create ${DEV_USER}`命令创建 docker 网络。

> 把 `${DEV_USER}` 替换为 .env 文件中 DEV_USER 的值。

4. 执行 `make dev` 启动项目