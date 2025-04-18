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
	mcp.WithDescription("create change request"),
	mcp.WithString("organizationId", mcp.Description("organization id"), mcp.Required()),
	mcp.WithString("repositoryId", mcp.Description("repository id"), mcp.Required()),
	mcp.WithString("title", mcp.Description("title, max 256 chars"), mcp.Required()),
	mcp.WithString("description", mcp.Description("description, max 10000 chars")),
	mcp.WithString("sourceBranch", mcp.Description("source branch"), mcp.Required()),
	mcp.WithNumber("sourceProjectId", mcp.Description("source project id"), mcp.Required()),
	mcp.WithString("targetBranch", mcp.Description("target branch"), mcp.Required()),
	mcp.WithNumber("targetProjectId", mcp.Description("target project id"), mcp.Required()),
	mcp.WithArray("reviewerUserIds", mcp.Description("reviewer user id list")),
	mcp.WithArray("workItemIds", mcp.Description("related work item id list")),
	mcp.WithString("createFrom", mcp.Description("creation source: WEB or COMMAND_LINE"), mcp.Enum("WEB", "COMMAND_LINE")),
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
