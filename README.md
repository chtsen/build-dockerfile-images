# Docker Image Build and Push Workflow

本仓库包含一个 GitHub Actions 工作流，用于构建并推送 Docker 镜像到私有 Docker 注册表。该工作流支持多平台构建，并可以在手动触发或者针对 `main` 分支的拉取请求时自动触发。

## 工作流详情

### Workflow Dispatch 输入

- **directory**: 构建 Docker 镜像的目录。如果未指定，则工作流将在整个仓库中搜索 `Dockerfile`。
- **platforms**: 要构建的平台（例如，`linux/amd64,linux/arm64`）。如果未指定，则默认平台为 `linux/amd64` 和 `linux/arm64`。

### 环境变量

- **DOCKER_REGISTRY**: 要推送镜像的 Docker 注册表（例如 `registry.cn-hangzhou.aliyuncs.com`）。
- **DOCKER_ORG**: 要使用的 Docker 组织（例如 `chtsen-sysnc`）。

### Secrets

- **DOCKER_ALIYUN_USER**: 私有 Docker 注册表的用户名。
- **DOCKER_ALIYUN_PW**: 私有 Docker 注册表的密码。

## 工作流如何运行

1. **检出仓库**: 工作流将仓库检出到运行器。
2. **设置 Docker Buildx**: 设置 Docker Buildx 以进行多平台构建。
3. **登录到私有 Docker 注册表**: 使用提供的 Secrets 登录到指定的 Docker 注册表。
4. **获取 Dockerfile 列表**: 工作流搜索指定目录（或整个仓库，如果未指定目录）中的 Dockerfile。
5. **构建并推送 Docker 镜像**: 对于找到的每个 Dockerfile，工作流：
   - 提取目录名称作为镜像名称。
   - 从 Dockerfile 的 `FROM` 行提取标签。
   - 使用提取的标签构建并推送 Docker 镜像到指定的 Docker 注册表。

## 使用示例

### 手动触发工作流

您可以在 GitHub Actions 页面上使用 "Run workflow" 按钮手动触发工作流。如果需要，可以指定 `directory` 和 `platforms` 输入。

### 拉取请求触发

工作流还将自动在针对 `main` 分支的拉取请求时触发。

## 示例 `workflow_dispatch` 输入

- **directory**: `./bitnami/redis`
- **platforms**: `linux/amd64,linux/arm64`

这些输入将导致工作流在 `./bitnami/redis` 目录中搜索 Dockerfile，并为 `linux/amd64` 和 `linux/arm64` 平台构建镜像。
