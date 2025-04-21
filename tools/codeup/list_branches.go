package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListBranches = "list_branches"
)

var ListBranchesOptions = []mcp.ToolOption{
	mcp.WithDescription("list repository branches"),
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("repository id or URL-encoded full path"),
		mcp.Required()),
	mcp.WithNumber(
		"page", mcp.Description("page number")),
	mcp.WithNumber(
		"perPage", mcp.Description("items per page")),
	mcp.WithString(
		"sort", mcp.Description("sort order: name_asc, name_desc, updated_asc, updated_desc"),
		mcp.Enum("name_asc", "name_desc", "updated_asc", "updated_desc")),
	mcp.WithString(
		"search", mcp.Description("search query")),
}

var ListBranchesTool = func() mcp.Tool {
	return mcp.NewTool(ListBranches, ListBranchesOptions...)
}()

func ListBranchesFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	
	// 构建基本URL
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/branches", 
		organizationId, repositoryId)

	// 创建客户端
	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)
	
	// 添加查询参数
	yunxiaoClient.Query = make(map[string]string)
	
	// 添加可选参数
	if page, ok := request.Params.Arguments["page"]; ok && page != nil {
		yunxiaoClient.Query["page"] = fmt.Sprintf("%v", page)
	}
	
	if perPage, ok := request.Params.Arguments["perPage"]; ok && perPage != nil {
		yunxiaoClient.Query["perPage"] = fmt.Sprintf("%v", perPage)
	}
	
	if sort, ok := request.Params.Arguments["sort"]; ok && sort != nil {
		yunxiaoClient.Query["sort"] = sort.(string)
	}
	
	if search, ok := request.Params.Arguments["search"]; ok && search != nil {
		yunxiaoClient.Query["search"] = search.(string)
	}
	
	// 设置响应数据类型为Branch数组
	branches := make([]types.Branch, 0)
	return yunxiaoClient.HandleMCPResult(&branches)
} 