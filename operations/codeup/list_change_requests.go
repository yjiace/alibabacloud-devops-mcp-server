package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListChangeRequests = "list_change_requests"
)

var ListChangeRequestsOptions = []mcp.ToolOption{
	mcp.WithDescription("list change requests"),
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
}

var ListChangeRequestsTool = func() mcp.Tool {
	return mcp.NewTool(ListChangeRequests, ListChangeRequestsOptions...)
}()

func ListChangeRequestsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/changeRequests", organizationId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)
	changeRequests := make([]types.ChangeRequest, 0)
	return yunxiaoClient.HandleMCPResult(&changeRequests)
}
