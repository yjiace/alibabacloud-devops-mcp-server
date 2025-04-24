package projex

import (
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-devops-mcp-server/types"
	"github.com/aliyun/alibabacloud-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetWorkItem = "get_work_item"
)

var GetWorkItemOptions = []mcp.ToolOption{
	mcp.WithDescription("获取工作项详情"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"workItemId", mcp.Description("工作项唯一标识，必填参数"),
		mcp.Required()),
}

var GetWorkItemTool = func() mcp.Tool {
	return mcp.NewTool(GetWorkItem, GetWorkItemOptions...)
}()

func GetWorkItemFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	workItemId := request.Params.Arguments["workItemId"].(string)

	//// 临时开启调试模式
	//oldDebug := utils.Debug
	//utils.Debug = true
	//defer func() {
	//	utils.Debug = oldDebug
	//}()

	apiUrl := fmt.Sprintf("/oapi/v1/projex/organizations/%s/workitems/%s", organizationId, workItemId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)
	workItem := &types.WorkItem{}
	return yunxiaoClient.HandleMCPResult(workItem)
}
