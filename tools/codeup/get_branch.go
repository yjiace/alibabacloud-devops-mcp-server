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
	GetBranch = "get_branch"
)

var GetBranchOptions = []mcp.ToolOption{
	mcp.WithDescription("查询分支信息"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者组织ID与仓库名称的组合，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F）"),
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

	// 自动处理branchName中未编码的斜杠
	if strings.Contains(branchName, "/") {
		branchName = url.PathEscape(branchName)
	}

	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/branches/%s",
		organizationId, repositoryId, branchName)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)
	branch := &types.Branch{}
	return yunxiaoClient.HandleMCPResult(branch)
}
