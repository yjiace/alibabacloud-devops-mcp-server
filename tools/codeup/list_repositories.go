package codeup

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-devops-mcp-server/types"
	"github.com/aliyun/alibaba-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListRepositories = "list_repositories"
)

var ListRepositoriesOptions = []mcp.ToolOption{
	mcp.WithDescription("查询代码库列表"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithNumber(
		"page", mcp.Description("页码，默认从1开始，一般不要超过150页")),
	mcp.WithNumber(
		"perPage", mcp.Description("每页大小，默认20，取值范围【1，100】")),
	mcp.WithString(
		"orderBy", mcp.Description("排序字段，可选值包括 {created_at, name, path, last_activity_at}，默认值为 created_at"),
		mcp.Enum("created_at", "name", "path", "last_activity_at")),
	mcp.WithString(
		"sort", mcp.Description("排序方式，可选值包括{asc, desc}，默认值为 desc"),
		mcp.Enum("asc", "desc")),
	mcp.WithString(
		"search", mcp.Description("搜索关键字，用于模糊匹配代码库路径")),
	mcp.WithBoolean(
		"archived", mcp.Description("是否归档")),
}

var ListRepositoriesTool = func() mcp.Tool {
	return mcp.NewTool(ListRepositories, ListRepositoriesOptions...)
}()

func ListRepositoriesFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories", organizationId)

	// 添加查询参数
	queries := make(map[string]string)

	// 添加可选参数
	if page, ok := request.Params.Arguments["page"]; ok && page != nil {
		queries["page"] = fmt.Sprintf("%v", page)
	}

	if perPage, ok := request.Params.Arguments["perPage"]; ok && perPage != nil {
		queries["perPage"] = fmt.Sprintf("%v", perPage)
	}

	if orderBy, ok := request.Params.Arguments["orderBy"]; ok && orderBy != nil {
		queries["orderBy"] = orderBy.(string)
	}

	if sort, ok := request.Params.Arguments["sort"]; ok && sort != nil {
		queries["sort"] = sort.(string)
	}

	if search, ok := request.Params.Arguments["search"]; ok && search != nil {
		queries["search"] = search.(string)
	}

	if archived, ok := request.Params.Arguments["archived"]; ok && archived != nil {
		queries["archived"] = fmt.Sprintf("%v", archived)
	}

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl, utils.WithQueries(queries))
	repositories := make([]types.Repository, 0)
	return yunxiaoClient.HandleMCPResult(&repositories)
}
