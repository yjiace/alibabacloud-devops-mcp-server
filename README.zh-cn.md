<p align="center"><a href="README.md">English</a> | 中文<br></p>

# alibabacloud-devops-mcp-server
[云效](https://www.aliyun.com/product/yunxiao)mcp-server工具为 AI 助手提供了与云效平台交互的能力，能够让 AI 助手可以读取项目中工作项的内容，在理解需求后自动编写代码，并提交代码合并请求。企业研发团队可以使用它协助代码审查、优化任务管理、完成构建、部署等任务，从而专注于更重要的创新和产品交付。

## 功能特性

alibabacloud-devops-mcp-server提供了以下功能，让AI助手能够：

* **代码仓库管理**：查询代码仓库及其分支、创建分支
* **文件操作**：创建、更新、删除和获取代码文件内容
* **代码评审**：创建和管理代码评审流程
* **项目管理**：搜索项目、获取项目详情
* **流水线管理**：获取流水线详情、获取流水线列表、运行流水线、获取最近一次流水线运行信息、获取流水线运行实例、获取流水线运行实例列表、查询/运行流水线部署任务
* **制品仓库管理**：查看制品仓库信息、查询制品信息、查看单个制品信息
* **应用交付**：创建和管理部署单、应用、编排、变量组、模板、标签、全局变量和部署资源

## 工具列表

alibabacloud-devops-mcp-server集成了多种工具，包括：

### 组织管理
- `get_current_organization_Info`: 获取当前用户所在组织信息
- `get_user_organizations`: 获取当前用户加入的组织列表
- `get_organization_role`: 获取组织角色信息
- `get_organization_departments`: 获取组织中的部门列表
- `get_organization_department_info`: 获取组织中某个部门的信息
- `get_organization_department_ancestors`: 获取组织中部门的上级部门
- `get_organization_members`: 获取组织成员列表
- `get_organization_member_info`: 获取组织成员信息
- `get_organization_member_info_by_user_id`: 通过用户ID获取组织成员信息
- `search_organization_members`: 搜索组织成员
- `list_organization_roles`: 列出组织角色
- `get_organization_role`: 获取组织角色信息

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
- `list_change_request_comments`: 查询合并请求评论列表
- `get_compare`: 代码比较
- `get_repository`: 获取仓库详情
- `list_repositories`: 获取仓库列表

### 项目管理工具

- `get_project`: 获取项目详情
- `search_projects`: 搜索项目
- `get_sprint`: 获取迭代详情
- `list_sprints`: 获取项目中的迭代列表
- `get_work_item`: 获取工作项详情
- `search_workitems`: 搜索工作项
- `get_work_item_types`: 获取工作项类型
- `create_work_item`: 创建工作项
- `list_all_work_item_types`: 列出组织中所有工作项类型
- `list_work_item_types`: 列出项目空间中工作项类型
- `get_work_item_type`: 获取特定工作项类型的详细信息
- `list_work_item_relation_work_item_types`: 列出可关联到特定工作项的工作项类型
- `get_work_item_type_field_config`: 获取工作项类型的字段配置
- `get_work_item_workflow`: 获取工作项类型的工作流信息
- `list_work_item_comments`: 列出特定工作项的评论
- `create_work_item_comment`: 为特定工作项创建评论
- `list_current_user_effort_records`: [项目管理] 获取用户的实际工时明细，结束时间和开始时间的间隔不能大于6个月
- `list_effort_records`: [项目管理] 获取实际工时明细
- `create_effort_record`: [项目管理] 登记实际工时
- `list_estimated_efforts`: [项目管理] 获取预计工时明细
- `create_estimated_effort`: [项目管理] 登记预计工时
- `update_effort_record`: [项目管理] 更新登记实际工时
- `update_estimated_effort`: [项目管理] 更新登记预计工时

###  流水线工具
- `get_pipeline` - 获取流水线详情
- `list_pipelines` - 获取流水线列表
- `smart_list_pipelines` - 智能查询流水线（支持自然语言时间）
- `create_pipeline_run` - 运行流水线
- `get_latest_pipeline_run` - 获取最新运行信息
- `get_pipeline_run` - 获取运行详情
- `list_pipeline_runs` - 获取运行历史
- `list_pipeline_jobs_by_category` - 获取流水线任务
- `list_pipeline_job_historys` - 获取任务历史
- `execute_pipeline_job_run` - 手动运行任务
- `get_pipeline_job_run_log` - 获取任务日志
- `list_service_connections` - 获取服务连接列表
- `create_pipeline_from_description`: 根据自然语言描述生成流水线 YAML 并创建流水线
- `update_pipeline`: 更新流水线YAML内容

### 制品仓库工具

- `list_package_repositories`: 查看制品仓库信息
- `list_artifacts`: 查询制品信息
- `get_artifact`: 查看单个制品信息

### 应用交付工具

- `create_change_order`: [应用交付] 创建部署单
- `list_change_order_versions`: [应用交付] 查看部署单版本列表
- `get_change_order`: [应用交付] 读取部署单使用的物料和工单状态
- `list_change_order_job_logs`: [应用交付] 查询环境部署单日志
- `find_task_operation_log`: [应用交付] 查询部署任务执行日志，其中通常包含下游部署引擎的调度细节信息
- `execute_job_action`: [应用交付] 操作环境部署单
- `list_change_orders_by_origin`: [应用交付] 根据创建来源查询部署单
- `list_applications`: [应用交付] 分页获取组织中的应用列表
- `get_application`: [应用交付] 根据应用名获取应用详情
- `create_application`: [应用交付] 创建应用
- `update_application`: [应用交付] 更新应用
- `get_latest_orchestration`: [应用交付] 获取环境的最新编排
- `list_app_orchestration`: [应用交付] 列出应用编排
- `create_app_orchestration`: [应用交付] 创建应用编排
- `delete_app_orchestration`: [应用交付] 删除应用编排
- `get_app_orchestration`: [应用交付] 获取应用编排
- `update_app_orchestration`: [应用交付] 更新应用编排
- `get_env_variable_groups`: [应用交付] 获取环境的变量组
- `create_variable_group`: [应用交付] 创建变量组
- `delete_variable_group`: [应用交付] 删除变量组
- `get_variable_group`: [应用交付] 获取变量组
- `update_variable_group`: [应用交付] 更新变量组
- `get_app_variable_groups`: [应用交付] 获取应用的变量组
- `get_app_variable_groups_revision`: [应用交付] 获取应用变量组的版本
- `search_app_templates`: [应用交付] 搜索应用模板
- `create_app_tag`: [应用交付] 创建应用标签
- `update_app_tag`: [应用交付] 更新应用标签
- `delete_app_tag`: [应用交付] 删除应用标签
- `search_app_tags`: [应用交付] 搜索应用标签
- `update_app_tag_bind`: [应用交付] 更新应用标签绑定
- `create_global_var`: [应用交付] 创建全局变量组
- `delete_global_var`: [应用交付] 删除全局变量组
- `get_global_var`: [应用交付] 获取全局变量组
- `update_global_var`: [应用交付] 更新全局变量组
- `list_global_vars`: [应用交付] 列出全局变量组
- `get_machine_deploy_log`: [应用交付] 获取机器部署日志
- `add_host_list_to_host_group`: [应用交付] 添加主机列表到主机组
- `add_host_list_to_deploy_group`: [应用交付] 添加主机列表到部署组
- `delete_host_list_from_deploy_group`: [应用交付] 从部署组删除主机列表
- `delete_host_list_from_host_group`: [应用交付] 从主机组删除主机列表

## 用法



### 先决条件
* node 版本  >= 18.0.0
* 阿里云[云效](https://www.aliyun.com/product/yunxiao)个人访问令牌，[点击前往](https://help.aliyun.com/zh/yunxiao/developer-reference/obtain-personal-access-token)，授予组织管理、项目协作、代码管理、流水线、制品仓库、应用交付、测试管理下所有api的读写权限。令牌的到期时间注意选择一个长期有效的时间。
  
  ![个人令牌授权页面](https://agent-install-beijing.oss-cn-beijing.aliyuncs.com/alibabacloud-devops-mcp-server/img_8.jpg)

### 在 Smithery.ai 中使用云效 MCP 服务

云效 MCP 服务已部署到 Smithery.ai 中，可以按照下列命令安装使用 [Smithery](https://smithery.ai/server/@aliyun/alibabacloud-devops-mcp-server):

```bash
npx -y @smithery/cli install @aliyun/alibabacloud-devops-mcp-server --client claude
```

### 通过 MCP市场 安装[云效](https://www.aliyun.com/product/yunxiao) MCP 服务
通义灵码内置的MCP市场中已经提供了云效的MCP服务，在通义灵码中进入MCP市场并且找到「云效DevOps」，直接安装即可。

![MCP市场安装云效MCP服务](https://agent-install-beijing.oss-cn-beijing.aliyuncs.com/alibabacloud-devops-mcp-server/img_9.png)

### 在Cursor、Claude code/iFlow等终端中配置mcp server。通过npx的方式来运行mcp server。
```json
{
  "mcpServers": {
    "yunxiao": {
      "command": "npx",
      "args": [
        "-y",
        "alibabacloud-devops-mcp-server"
      ],
      "env": {
        "YUNXIAO_ACCESS_TOKEN": "<YOUR_TOKEN>"
      }
    }
  }
}
```
### 通过 Docker 容器运行 MCP 服务器
1.Docker build
```shell
docker build -t alibabacloud/alibabacloud-devops-mcp-server .
```
2.配置 MCP 服务器
```json
{
  "mcpServers": {
    "yunxiao": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e",
        "YUNXIAO_ACCESS_TOKEN",
        "alibabacloud/alibabacloud-devops-mcp-server"
      ],
      "env": {
        "YUNXIAO_ACCESS_TOKEN": "<YOUR_TOKEN>"
      }
    }
  }
}
```
### 通过 docker compose 运行MCP 服务器
1. 环境设置
```shell
cd alibabacloud-devops-mcp-server
cp .env.example
```

2. Docker Compose构建
```shell
docker compose up -d
```
3. 配置 MCP 服务器
```json
{
  "mcpServers": {
    "yunxiao": {
      "url":"http://localhost:3000/sse"
    }
  }
}
```

### SSE 模式下使用独立令牌
在 SSE 模式下运行时，每个用户都可以通过查询参数或请求头传递自己的令牌：

1. 通过查询参数：
```
http://localhost:3000/sse?yunxiao_access_token=USER_SPECIFIC_TOKEN
```

2. 通过请求头：
```
x-yunxiao-token: USER_SPECIFIC_TOKEN
```

这允许多个用户共享同一个 SSE 服务，同时使用各自独立的令牌进行身份验证。

## 联系我们
如有任何疑问或疑虑，请通过钉钉群联系我们：134400004101

![Alibaba Cloud Devops MCP Server Group](https://agent-install-beijing.oss-cn-beijing.aliyuncs.com/alibabacloud-devops-mcp-server/1750147152464.png)

## 相关链接
- [阿里云云效](https://www.aliyun.com/product/yunxiao)
- [MCP 市场](https://modelscope.cn/mcp/servers/@aliyun/alibabacloud-devops-mcp-server)
- [使用场景示例](https://mp.weixin.qq.com/s/KQsN6dQlnNeCNATC-QD7pg)
