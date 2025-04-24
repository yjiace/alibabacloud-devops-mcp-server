package codeup

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-devops-mcp-server/types"
	"github.com/aliyun/alibaba-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
	"net/url"
	"strings"
)

const (
	CreateBranch = "create_branch"
)

var CreateBranchOptions = []mcp.ToolOption{
	mcp.WithDescription("创建分支"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者组织ID与仓库名称的组合，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F）"),
		mcp.Required()),
	mcp.WithString(
		"branch", mcp.Description("要创建的分支名称"),
		mcp.Required()),
	mcp.WithString(
		"ref", mcp.Description("源分支名称，新分支基于哪个分支创建"),
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

	// 自动处理repositoryId中未编码的斜杠
	if strings.Contains(repositoryId, "/") {
		// 发现未编码的斜杠，自动进行URL编码
		parts := strings.SplitN(repositoryId, "/", 2)
		if len(parts) == 2 {
			encodedRepoName := url.QueryEscape(parts[1])
			// 移除编码中的+号（空格被编码为+，但我们需要%20）
			encodedRepoName = strings.ReplaceAll(encodedRepoName, "+", "%20")
			repositoryId = parts[0] + "%2F" + encodedRepoName
		}
	}

	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/branches", organizationId, repositoryId)

	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl)

	yunxiaoClient.Query = make(map[string]string)
	yunxiaoClient.Query["branch"] = branch
	yunxiaoClient.Query["ref"] = ref

	branchInfo := &types.Branch{}
	return yunxiaoClient.HandleMCPResult(branchInfo)
}
