package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetBranch = "get_branch"
)

var GetBranchOptions = []mcp.ToolOption{
	mcp.WithDescription("查询分支信息"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径，例如: 2835387 或 codeup-org-id%2Fcodeup-demo"),
		mcp.Required()),
	mcp.WithString(
		"branchName", mcp.Description("分支名称（若有特殊符号，可使用URL-Encoder进行编码处理），例如: master 或 feature%2Fdev"),
		mcp.Required()),
}

var GetBranchTool = func() mcp.Tool {
	return mcp.NewTool(GetBranch, GetBranchOptions...)
}()

func GetBranchFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	branchName := request.Params.Arguments["branchName"].(string)
	
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/branches/%s", 
		organizationId, repositoryId, branchName)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)
	branch := &types.Branch{}
	return yunxiaoClient.HandleMCPResult(branch)
} 