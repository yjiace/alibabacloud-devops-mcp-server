package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateChangeRequest = "create_change_request"
)

var CreateChangeRequestOptions = []mcp.ToolOption{
	mcp.WithDescription("创建合并请求"),
	mcp.WithString("organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"), mcp.Required()),
	mcp.WithString("repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径，例如: 2835387 或 codeup-org-id%2Fcodeup-demo"), mcp.Required()),
	mcp.WithString("title", mcp.Description("标题，不超过256个字符"), mcp.Required()),
	mcp.WithString("description", mcp.Description("描述，不超过10000个字符")),
	mcp.WithString("sourceBranch", mcp.Description("源分支名称"), mcp.Required()),
	mcp.WithNumber("sourceProjectId", mcp.Description("源库ID"), mcp.Required()),
	mcp.WithString("targetBranch", mcp.Description("目标分支名称"), mcp.Required()),
	mcp.WithNumber("targetProjectId", mcp.Description("目标库ID"), mcp.Required()),
	mcp.WithArray("reviewerUserIds", mcp.Description("评审人用户ID列表")),
	mcp.WithArray("workItemIds", mcp.Description("关联工作项ID列表")),
	mcp.WithString("createFrom", mcp.Description("创建来源：WEB - 页面创建；COMMAND_LINE - 命令行创建；默认为WEB"), mcp.Enum("WEB", "COMMAND_LINE")),
}

var CreateChangeRequestTool = func() mcp.Tool {
	return mcp.NewTool(CreateChangeRequest, CreateChangeRequestOptions...)
}()

func CreateChangeRequestFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/changeRequests", organizationId, repositoryId)

	// 准备payload，确保类型正确
	payload := map[string]interface{}{
		"title":           request.Params.Arguments["title"],
		"sourceBranch":    request.Params.Arguments["sourceBranch"],
		"targetBranch":    request.Params.Arguments["targetBranch"],
		"sourceProjectId": request.Params.Arguments["sourceProjectId"],
		"targetProjectId": request.Params.Arguments["targetProjectId"],
	}

	// 添加可选参数
	if val, ok := request.Params.Arguments["description"]; ok && val != nil {
		payload["description"] = val
	}

	if val, ok := request.Params.Arguments["reviewerUserIds"]; ok && val != nil {
		payload["reviewerUserIds"] = val
	}

	if val, ok := request.Params.Arguments["workItemIds"]; ok && val != nil {
		payload["workItemIds"] = val
	}

	if val, ok := request.Params.Arguments["createFrom"]; ok && val != nil {
		payload["createFrom"] = val
	} else {
		payload["createFrom"] = "WEB" // 默认为WEB
	}

	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(payload))
	changeRequest := &types.ChangeRequest{}
	return yunxiaoClient.HandleMCPResult(changeRequest)
}
