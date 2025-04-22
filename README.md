# mcp-yunxiao

mcp-yunxiao is an API integration tool built for Alibaba Cloud Yunxiao platform. This tool is based on the Model Context Protocol (MCP), providing AI assistants with the ability to interact with the Yunxiao platform.

## Features

mcp-yunxiao provides the following functionalities, enabling AI assistants to:

* **Code Repository Management**: Create, query, and manage code repositories and their branches
* **File Operations**: Create, update, delete, and retrieve code file contents
* **Code Review**: Create and manage code review processes
* **Project Management**: Search projects, get project details
* **Sprint Management**: Query sprint lists and details
* **Pipeline Management**: Create, query, run pipelines, and get execution records

## Quick Start

### Prerequisites

* Go 1.24.0 or higher
* Alibaba Cloud Yunxiao personal access token, [Get it here](https://help.aliyun.com/zh/yunxiao/developer-reference/obtain-personal-access-token?spm=a2c4g.11186623.help-menu-150040.d_5_0_1.5dc72af2GnT64i)

### Installation

1. Clone the repository:

```bash
git clone git@gitlab.alibaba-inc.com:yunxiao-packages/mcp-yunxiao.git
cd mcp-yunxiao
```

2. Install dependencies:

```bash
make deps
```

3. Build the application:

```bash
make build
```

### Usage

#### Direct execution

```bash
./bin/mcp-yunxiao
```

Or using Docker:

```bash
make docker
make docker-run
```

## Cross-Platform Build

MCP-Yunxiao supports multiple operating systems and architectures:

* Windows (AMD64, ARM64)
* macOS (AMD64, ARM64)
* Linux (AMD64, ARM64)

Build binaries for all platforms:

```bash
make build-all
```

Build for specific platforms:

```bash
make build-windows
make build-darwin
make build-linux
```

## Tool List

mcp-yunxiao integrates various tools, including:

### Code Management Tools

- `create_branch`: Create a branch
- `delete_branch`: Delete a branch
- `get_branch`: Get branch information
- `list_branches`: Get branch list
- `create_file`: Create a file
- `delete_file`: Delete a file
- `get_file_blobs`: Get file content
- `list_files`: Query file tree
- `update_file`: Update file content
- `create_change_request`: Create a merge request
- `create_change_request_comment`: Create a merge request comment
- `get_change_request`: Query merge request
- `list_change_request_patch_sets`: Query merge request version list
- `list_change_request`: Query merge request list
- `get_compare`: Compare code
- `get_repository`: Get repository details
- `list_repositories`: Get repository list

### Project Management Tools

- `get_project`: Get project details
- `search_projects`: Search projects
- `get_sprint`: Get sprint details
- `get_work_item`: Get work item details

### Pipeline Tools

- `create_pipeline`: Create a pipeline
- `get_pipeline`: Get pipeline details
- `list_pipelines`: Get pipeline list
- `get_pipeline_run`: Get pipeline run instance
- `get_latest_pipeline_run`: Get the latest pipeline run information
- `list_pipeline_runs`: Get pipeline run instance list

## Usage Examples
1. Get work item details:
![img.png](img/img.png)

![img.png](img/img_1.png)

2. Implement code based on work item description & create Pull Request
![img.png](img/img_2.png)

![img.png](img/img_3.png) 