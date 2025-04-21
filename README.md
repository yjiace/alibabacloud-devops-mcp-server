# mcp-yunxiao

mcp-yunxiao是一个为阿里云云效(Yunxiao)平台构建的API集成工具。该工具基于模型上下文协议(Model Context Protocol，MCP)协议]，为AI助手提供了与云效平台交互的能力。

## 功能特性

mcp-yunxiao提供了以下功能，让AI助手能够：

* **代码仓库管理**：创建、查询、管理代码仓库及其分支
* **文件操作**：创建、更新、删除和获取代码文件内容
* **代码评审**：创建和管理代码评审流程
* **项目管理**：搜索项目、获取项目详情
* **迭代管理**：查询迭代列表及详情
* **流水线管理**：创建、查询、运行流水线及获取运行记录

## 快速开始

### 先决条件

* Go 1.24.0或以上版本
* 阿里云云效个人访问令牌，[点击前往](https://help.aliyun.com/zh/yunxiao/developer-reference/obtain-personal-access-token?spm=a2c4g.11186623.help-menu-150040.d_5_0_1.5dc72af2GnT64i)

### 安装

1. 克隆仓库：

```bash
git clone git@gitlab.alibaba-inc.com:yunxiao-packages/mcp-yunxiao.git
cd mcp-yunxiao
```

2. 安装依赖：

```bash
make deps
```

3. 构建应用：

```bash
make build
```

### 使用方法

#### 直接运行

```bash
./bin/mcp-yunxiao
```

或者使用Docker：

```bash
make docker
make docker-run
```

## 跨平台构建

MCP-Yunxiao支持多种操作系统和架构：

* Windows (AMD64, ARM64)
* macOS (AMD64, ARM64)
* Linux (AMD64, ARM64)

构建所有平台的二进制文件：

```bash
make build-all
```

构建特定平台：

```bash
make build-windows
make build-darwin
make build-linux
```

## 工具列表

mcp-yunxiao集成了多种工具，包括：

### 代码管理工具

- `create_branch`: 创建分支
- `delete_branch`: 删除分支
- `get_branch`: 获取分支信息
- `list_branches`: 获取分支列表
- `create_file`: 创建文件
- `delete_file`: 删除文件
- `get_file_blobs`: 获取文件内容
- `list_files`: 查询文件树
- `update_file`: 更新文件内容
- `create_change_request`: 创建合并请求
- `create_change_request_comment`: 创建合并请求评论
- `get_change_request`: 查询合并请求
- `list_change_request_patch_sets`: 查询合并请求版本列表
- `list_change_request`: 查询合并请求列表
- `get_compare`: 代码比较
- `get_repository`: 获取仓库详情
- `list_repositories`: 获取仓库列表

### 项目管理工具

- `get_project`: 获取项目详情
- `search_projects`: 搜索项目
- `get_sprint`: 获取迭代详情
- `get_work_item`: 获取工作项详情

### 流水线工具

- `create_pipeline`: 创建流水线
- `get_pipeline`: 获取流水线详情
- `list_pipelines`: 获取流水线列表
- `get_pipeline_run`: 获取流水线运行实例
- `get_latest_pipeline_run`: 获取最近一次流水线运行信息
- `list_pipeline_runs`: 获取流水线运行实例列表

## 使用实例
1、获取工作项详情：
![img.png](img/img.png)

![img.png](img/img_1.png)

2、根据工作项描述实现编码 & 创建 Pull Request
![img.png](img/img_2.png)

![img.png](img/img_3.png)
