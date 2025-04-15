package projex

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetWorkItem = "get_work_item"
)

var GetWorkItemOptions = []mcp.ToolOption{
	mcp.WithDescription("list repositories"),
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
	mcp.WithString(
		"workItemId", mcp.Description("work item id"),
		mcp.Required()),
}

var GetWorkItemTool = func() mcp.Tool {
	return mcp.NewTool(GetWorkItem, GetWorkItemOptions...)
}()

func GetWorkItemFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	workItemId := request.Params.Arguments["workItemId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/projex/organizations/%s/workitems/%s", organizationId, workItemId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)
	workItem := &types.WorkItem{}
	return yunxiaoClient.HandleMCPResult(workItem)
}
