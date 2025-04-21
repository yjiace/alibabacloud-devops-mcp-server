package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	DeleteBranch = "delete_branch"
)

var DeleteBranchOptions = []mcp.ToolOption{
	mcp.WithDescription("delete branch"),
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("repository id or URL-encoded full path"),
		mcp.Required()),
	mcp.WithString(
		"branchName", mcp.Description("branch name (use URL-encoding for special characters like 'feature/branch')"),
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
	
	// 构建API URL
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/branches/%s", 
		organizationId, repositoryId, branchName)
	
	// 创建客户端
	yunxiaoClient := utils.NewYunxiaoClient("DELETE", apiUrl)
	
	// 创建响应对象
	response := &DeleteBranchResponse{}
	
	return yunxiaoClient.HandleMCPResult(response)
} 