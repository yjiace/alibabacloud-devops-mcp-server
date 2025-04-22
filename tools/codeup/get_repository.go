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
	mcp.WithDescription("查询代码库"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径，例如: 2835387 或 codeup-org-id%2Fcodeup-demo"),
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
