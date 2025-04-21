package projex

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	SearchWorkitems = "search_workitems"
)

// 定义工作项列表的返回结构
type SearchWorkitemsResponse []types.WorkItem

var SearchWorkitemsOptions = []mcp.ToolOption{
	mcp.WithDescription("搜索工作项"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"category", mcp.Description("搜索的工作项类型，例如 Req，多值用逗号隔开"),
		mcp.Required()),
	mcp.WithString(
		"conditions", mcp.Description("过滤条件，JSON格式，例如：{\"conditionGroups\":[[{\"fieldIdentifier\":\"status\",\"operator\":\"CONTAINS\",\"value\":[\"100005\",\"100010\",\"154395\"],\"toValue\":null,\"className\":\"status\",\"format\":\"list\"}]]}")),
	mcp.WithString(
		"orderBy", mcp.Description("排序字段，默认为gmtCreate")),
	mcp.WithNumber(
		"page", mcp.Description("分页参数，第几页")),
	mcp.WithNumber(
		"perPage", mcp.Description("分页参数，每页大小")),
	mcp.WithString(
		"sort", mcp.Description("排序方式，desc或asc")),
	mcp.WithString(
		"spaceId", mcp.Description("空间ID")),
}

var SearchWorkitemsTool = func() mcp.Tool {
	return mcp.NewTool(SearchWorkitems, SearchWorkitemsOptions...)
}()

func SearchWorkitemsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/projex/organizations/%s/workitems:search", organizationId)

	payload := make(map[string]interface{})

	payload["category"] = request.Params.Arguments["category"]

	if conditions, ok := request.Params.Arguments["conditions"]; ok {
		payload["conditions"] = conditions
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

	if spaceId, ok := request.Params.Arguments["spaceId"]; ok {
		payload["spaceId"] = spaceId
	}

	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(payload))

	response := SearchWorkitemsResponse{}

	return yunxiaoClient.HandleMCPResult(&response)
}
