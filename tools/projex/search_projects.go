package projex

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	SearchProjects = "search_projects"
)

var SearchProjectsOptions = []mcp.ToolOption{
	mcp.WithDescription("搜索项目"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"conditions", mcp.Description("过滤条件，JSON格式，例如：{\"conditionGroups\":[[{\"className\":\"string\",\"fieldIdentifier\":\"name\",\"format\":\"input\",\"operator\":\"BETWEEN\",\"toValue\":null,\"value\":[\"test\"]}]]}")),
	mcp.WithString(
		"extraConditions", mcp.Description("额外的过滤条件，例如我管理的、我参与的，我收藏的等")),
	mcp.WithString(
		"orderBy", mcp.Description("排序字段，默认为gmtCreate，支持：gmtCreate(创建时间)，name(名称)")),
	mcp.WithNumber(
		"page", mcp.Description("分页参数，第几页")),
	mcp.WithNumber(
		"perPage", mcp.Description("分页参数，每页大小，0-200，默认值20")),
	mcp.WithString(
		"sort", mcp.Description("排序方式，默认为desc，可选：desc(降序)，asc(升序)")),
}

var SearchProjectsTool = func() mcp.Tool {
	return mcp.NewTool(SearchProjects, SearchProjectsOptions...)
}()

func SearchProjectsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)

	// 构建API URL
	apiUrl := fmt.Sprintf("/oapi/v1/projex/organizations/%s/projects:search", organizationId)

	// 创建请求体
	payload := make(map[string]interface{})

	// 添加可选参数
	if conditions, ok := request.Params.Arguments["conditions"]; ok {
		payload["conditions"] = conditions
	}

	if extraConditions, ok := request.Params.Arguments["extraConditions"]; ok {
		payload["extraConditions"] = extraConditions
	}

	if orderBy, ok := request.Params.Arguments["orderBy"]; ok {
		payload["orderBy"] = orderBy
	}

	if page, ok := request.Params.Arguments["page"]; ok {
		payload["page"] = page
	}

	if perPage, ok := request.Params.Arguments["perPage"]; ok {
		payload["perPage"] = perPage
	}

	if sort, ok := request.Params.Arguments["sort"]; ok {
		payload["sort"] = sort
	}

	// 创建客户端并设置请求体
	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(payload))
	projectInfo := &types.ProjectInfo{}
	return yunxiaoClient.HandleMCPResult(projectInfo)
}
