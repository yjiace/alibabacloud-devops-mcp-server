package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"net/url"
	"strings"
)

const (
	DeleteBranch = "delete_branch"
)

var DeleteBranchOptions = []mcp.ToolOption{
	mcp.WithDescription("删除分支"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者组织ID与仓库名称的组合，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F）"),
		mcp.Required()),
	mcp.WithString(
		"branchName", mcp.Description("分支名称(使用URL-Encoder进行处理，例如: feature%2Fdev)"),
		mcp.Required()),
}

var DeleteBranchTool = func() mcp.Tool {
	return mcp.NewTool(DeleteBranch, DeleteBranchOptions...)
}()

// DeleteBranchResponse 定义删除分支的返回结构
type DeleteBranchResponse struct {
	BranchName string `json:"branchName"`
}

func DeleteBranchFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

	// 构建API URL
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/branches/%s",
		organizationId, repositoryId, branchName)

	// 创建客户端
	yunxiaoClient := utils.NewYunxiaoClient("DELETE", apiUrl)

	// 创建响应对象
	response := &DeleteBranchResponse{}

	return yunxiaoClient.HandleMCPResult(response)
}
