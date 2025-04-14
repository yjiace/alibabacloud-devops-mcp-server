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
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
	mcp.WithString("repositoryId", mcp.Description("repository id"), mcp.Required()),
	mcp.WithString("title", mcp.Description("title"), mcp.Required()),
	mcp.WithString("description", mcp.Description("description")),
	mcp.WithString("sourceBranch", mcp.Description("source branch"), mcp.Required()),
	mcp.WithString("sourceProjectId", mcp.Description("source project id"), mcp.Required()),
	mcp.WithString("targetBranch", mcp.Description("target branch"), mcp.Required()),
	mcp.WithString("targetProjectId", mcp.Description("target project id"), mcp.Required()),
}

var CreateChangeRequestTool = func() mcp.Tool {
	return mcp.NewTool(CreateChangeRequest, CreateChangeRequestOptions...)
}()

func CreateChangeRequestFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/changeRequests", organizationId, repositoryId)

	giteeClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	changeRequest := &types.ChangeRequest{}
	return giteeClient.HandleMCPResult(changeRequest)
}
