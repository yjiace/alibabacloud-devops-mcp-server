package codeup

import (
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-devops-mcp-server/types"
	"github.com/aliyun/alibabacloud-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListChangeRequests = "list_change_requests"
)

var ListChangeRequestsOptions = []mcp.ToolOption{
	mcp.WithDescription("查询合并请求列表"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithNumber(
		"page", mcp.Description("页码")),
	mcp.WithNumber(
		"perPage", mcp.Description("每页大小")),
	mcp.WithString(
		"projectIds", mcp.Description("代码库ID或者组织ID与仓库名称的组合列表，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F），多个以逗号分隔")),
	mcp.WithString(
		"authorIds", mcp.Description("创建者用户ID列表，多个以逗号分隔")),
	mcp.WithString(
		"reviewerIds", mcp.Description("评审人用户ID列表，多个以逗号分隔")),
	mcp.WithString(
		"state", mcp.Description("合并请求筛选状态：opened，merged，closed，默认为null，即查询全部状态"),
		mcp.Enum("opened", "merged", "closed")),
	mcp.WithString(
		"search", mcp.Description("标题关键字搜索")),
	mcp.WithString(
		"orderBy", mcp.Description("排序字段，仅支持：created_at - 创建时间；updated_at - 更新时间，默认排序字段"),
		mcp.Enum("created_at", "updated_at")),
	mcp.WithString(
		"sort", mcp.Description("排序方式：asc - 升序；desc - 降序，默认排序方式"),
		mcp.Enum("asc", "desc")),
	mcp.WithString(
		"createdBefore", mcp.Description("起始创建时间，时间格式为ISO 8601")),
	mcp.WithString(
		"createdAfter", mcp.Description("截止创建时间，时间格式为ISO 8601")),
}

var ListChangeRequestsTool = func() mcp.Tool {
	return mcp.NewTool(ListChangeRequests, ListChangeRequestsOptions...)
}()

func ListChangeRequestsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/changeRequests", organizationId)

	// 添加查询参数
	queries := make(map[string]string)

	// 添加可选参数
	if page, ok := request.Params.Arguments["page"]; ok && page != nil {
		queries["page"] = fmt.Sprintf("%v", page)
	}

	if perPage, ok := request.Params.Arguments["perPage"]; ok && perPage != nil {
		queries["perPage"] = fmt.Sprintf("%v", perPage)
	}

	if projectIds, ok := request.Params.Arguments["projectIds"]; ok && projectIds != nil {
		queries["projectIds"] = projectIds.(string)
	}

	if authorIds, ok := request.Params.Arguments["authorIds"]; ok && authorIds != nil {
		queries["authorIds"] = authorIds.(string)
	}

	if reviewerIds, ok := request.Params.Arguments["reviewerIds"]; ok && reviewerIds != nil {
		queries["reviewerIds"] = reviewerIds.(string)
	}

	if state, ok := request.Params.Arguments["state"]; ok && state != nil {
		queries["state"] = state.(string)
	}

	if search, ok := request.Params.Arguments["search"]; ok && search != nil {
		queries["search"] = search.(string)
	}

	if orderBy, ok := request.Params.Arguments["orderBy"]; ok && orderBy != nil {
		queries["orderBy"] = orderBy.(string)
	}

	if sort, ok := request.Params.Arguments["sort"]; ok && sort != nil {
		queries["sort"] = sort.(string)
	}

	if createdBefore, ok := request.Params.Arguments["createdBefore"]; ok && createdBefore != nil {
		queries["createdBefore"] = createdBefore.(string)
	}

	if createdAfter, ok := request.Params.Arguments["createdAfter"]; ok && createdAfter != nil {
		queries["createdAfter"] = createdAfter.(string)
	}

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl, utils.WithQueries(queries))
	changeRequests := make([]types.ChangeRequest, 0)
	return yunxiaoClient.HandleMCPResult(&changeRequests)
}
