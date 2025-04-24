package flow

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-devops-mcp-server/types"
	"github.com/aliyun/alibaba-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListPipelineRuns = "list_pipeline_runs"
)

// ListPipelineRunsResponse 定义流水线运行实例列表的返回结构
type ListPipelineRunsResponse []types.PipelineRunSimple

var ListPipelineRunsOptions = []mcp.ToolOption{
	mcp.WithDescription("获取流水线运行实例列表"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"pipelineId", mcp.Description("流水线ID"),
		mcp.Required()),
	mcp.WithNumber(
		"perPage", mcp.Description("每页数据条数，默认10，最大支持30")),
	mcp.WithNumber(
		"page", mcp.Description("当前页，默认1")),
	mcp.WithNumber(
		"startTime", mcp.Description("执行开始时间，时间戳（毫秒）")),
	mcp.WithNumber(
		"endTime", mcp.Description("执行结束时间，时间戳（毫秒）")),
	mcp.WithString(
		"status", mcp.Description("运行状态，FAIL(运行失败)、SUCCESS(运行成功)、RUNNING(运行中)")),
	mcp.WithNumber(
		"triggerMode", mcp.Description("触发方式，1:人工触发 2:定时触发 3:代码提交触发 5:流水线触发 6:WEBHOOK触发")),
}

var ListPipelineRunsTool = func() mcp.Tool {
	return mcp.NewTool(ListPipelineRuns, ListPipelineRunsOptions...)
}()

func ListPipelineRunsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	pipelineId := request.Params.Arguments["pipelineId"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/flow/organizations/%s/pipelines/%s/runs", organizationId, pipelineId)

	queryParams := make(map[string]string)

	// 添加可选参数
	if perPage, ok := request.Params.Arguments["perPage"]; ok && perPage != nil {
		queryParams["perPage"] = fmt.Sprintf("%v", perPage)
	}

	if page, ok := request.Params.Arguments["page"]; ok && page != nil {
		queryParams["page"] = fmt.Sprintf("%v", page)
	}

	if startTime, ok := request.Params.Arguments["startTime"]; ok && startTime != nil {
		queryParams["startTime"] = fmt.Sprintf("%v", startTime)
	}

	if endTime, ok := request.Params.Arguments["endTime"]; ok && endTime != nil {
		queryParams["endTme"] = fmt.Sprintf("%v", endTime) // 注意：API文档中这个参数名是endTme而不是endTime
	}

	if status, ok := request.Params.Arguments["status"].(string); ok && status != "" {
		queryParams["status"] = status
	}

	if triggerMode, ok := request.Params.Arguments["triggerMode"]; ok && triggerMode != nil {
		queryParams["triggerMode"] = fmt.Sprintf("%v", triggerMode)
	}

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl, utils.WithQueries(queryParams))

	response := ListPipelineRunsResponse{}

	return yunxiaoClient.HandleMCPResult(&response)
}
