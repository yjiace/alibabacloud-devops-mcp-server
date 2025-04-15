package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetChangeRequest = "get_change_request"
)

var GetChangeRequestOptions = []mcp.ToolOption{
	mcp.WithDescription("get change request"),
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("repository id"),
		mcp.Required()),
	mcp.WithString(
		"localId", mcp.Description("change request local id"),
		mcp.Required()),
}

var GetChangeRequestTool = func() mcp.Tool {
	return mcp.NewTool(GetChangeRequest, GetChangeRequestOptions...)
}()

func GetChangeRequestFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	localId := request.Params.Arguments["localId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/changeRequests/%s", organizationId, repositoryId, localId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)
	changeRequest := &types.ChangeRequest{}
	return yunxiaoClient.HandleMCPResult(changeRequest)
}
