package projex

import (
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-devops-mcp-server/types"
	"github.com/aliyun/alibabacloud-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListSprints = "list_sprints"
)

var ListSprintsOptions = []mcp.ToolOption{
	mcp.WithDescription("获取迭代列表"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"id", mcp.Description("项目唯一标识"),
		mcp.Required()),
	mcp.WithString(
		"status", mcp.Description("过滤的状态，TODO、DOING、ARCHIVED，分别对应未开始，进行中和已完成，多个状态使用逗号分隔")),
	mcp.WithNumber(
		"page", mcp.Description("分页参数，第几页")),
	mcp.WithNumber(
		"perPage", mcp.Description("分页参数，每页大小")),
}

var ListSprintsTool = func() mcp.Tool {
	return mcp.NewTool(ListSprints, ListSprintsOptions...)
}()

func ListSprintsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	projectId := request.Params.Arguments["id"].(string)

	// 构建API URL
	apiUrl := fmt.Sprintf("/oapi/v1/projex/organizations/%s/projects/%s/sprints", organizationId, projectId)

	// 创建查询参数
	queryParams := make(map[string]string)

	// 添加可选参数
	if status, ok := request.Params.Arguments["status"].(string); ok && status != "" {
		// 状态参数可能是逗号分隔的多个值
		queryParams["status"] = status
	}

	if page, ok := request.Params.Arguments["page"]; ok && page != nil {
		queryParams["page"] = fmt.Sprintf("%v", page)
	}

	if perPage, ok := request.Params.Arguments["perPage"]; ok && perPage != nil {
		queryParams["perPage"] = fmt.Sprintf("%v", perPage)
	}

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl, utils.WithQueries(queryParams))
	sprints := make([]types.SprintInfo, 0)
	return yunxiaoClient.HandleMCPResult(sprints)
}
