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
	mcp.WithDescription("删除分支"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径，例如: 2835387 或 codeup-org-id%2Fcodeup-demo"),
		mcp.Required()),
	mcp.WithString(
		"branchName", mcp.Description("分支名称（若有特殊符号，如feature/branch，可使用URL-Encoder进行编码处理）"),
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