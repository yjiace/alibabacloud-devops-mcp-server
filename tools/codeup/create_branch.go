package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateBranch = "create_branch"
)

var CreateBranchOptions = []mcp.ToolOption{
	mcp.WithDescription("create branch"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径，例如: 2835387 或 codeup-org-id%2Fcodeup-demo"),
		mcp.Required()),
	mcp.WithString(
		"branch", mcp.Description("创建的分支名称"),
		mcp.Required()),
	mcp.WithString(
		"ref", mcp.Description("来源分支名称"),
		mcp.Required()),
}

var CreateBranchTool = func() mcp.Tool {
	return mcp.NewTool(CreateBranch, CreateBranchOptions...)
}()

func CreateBranchFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	branch := request.Params.Arguments["branch"].(string)
	ref := request.Params.Arguments["ref"].(string)

	// 构建API URL
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/branches",
		organizationId, repositoryId)

	// 创建客户端
	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl)

	// 添加查询参数
	yunxiaoClient.Query = make(map[string]string)
	yunxiaoClient.Query["branch"] = branch
	yunxiaoClient.Query["ref"] = ref

	// 创建响应对象
	branchInfo := &types.Branch{}

	return yunxiaoClient.HandleMCPResult(branchInfo)
}
