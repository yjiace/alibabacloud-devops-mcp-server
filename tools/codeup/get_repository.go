package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetRepository = "get_repository"
)

var GetRepositoryOptions = []mcp.ToolOption{
	mcp.WithDescription("get repository"),
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("repository id"),
		mcp.Required()),
}

var GetRepositoryTool = func() mcp.Tool {
	return mcp.NewTool(GetRepository, GetRepositoryOptions...)
}()

func GetRepositoryFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s", organizationId, repositoryId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)
	repository := &types.Repository{}
	return yunxiaoClient.HandleMCPResult(repository)
}
