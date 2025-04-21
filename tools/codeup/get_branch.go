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
	mcp.WithDescription("get branch information"),
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("repository id or URL-encoded full path"),
		mcp.Required()),
	mcp.WithString(
		"branchName", mcp.Description("branch name (use URL-encoding for special characters)"),
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