package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListRepositories = "list repositories"
)

var ListRepositoriesOptions = []mcp.ToolOption{
	mcp.WithDescription("list repositories"),
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
}

var ListRepositoriesTool = func() mcp.Tool {
	return mcp.NewTool(ListRepositories, ListRepositoriesOptions...)
}()

func ListRepositoriesFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories", organizationId)

	giteeClient := utils.NewYunxiaoClient("GET", apiUrl)
	repositories := make([]types.Repository, 0)
	return giteeClient.HandleMCPResult(&repositories)
}
