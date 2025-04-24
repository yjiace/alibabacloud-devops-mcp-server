package flow

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-devops-mcp-server/types"
	"github.com/aliyun/alibaba-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListPipelines = "list_pipelines"
)

// ListPipelinesResponse 定义流水线列表的返回结构
type ListPipelinesResponse []types.PipelineBaseInfo

var ListPipelinesOptions = []mcp.ToolOption{
	mcp.WithDescription("获取流水线列表"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithNumber(
		"createStartTime", mcp.Description("创建开始时间，时间戳（毫秒）")),
	mcp.WithNumber(
		"createEndTime", mcp.Description("创建结束时间，时间戳（毫秒）")),
	mcp.WithNumber(
		"executeStartTime", mcp.Description("执行开始时间，时间戳（毫秒）")),
	mcp.WithNumber(
		"executeEndTime", mcp.Description("执行结束时间，时间戳（毫秒）")),
	mcp.WithString(
		"pipelineName", mcp.Description("流水线名称")),
	mcp.WithString(
		"statusList", mcp.Description("流水线运行状态，多个逗号分割，SUCCESS,RUNNING,FAIL,CANCELED,WAITING")),
	mcp.WithNumber(
		"perPage", mcp.Description("每页数据条数，默认10，最大支持30")),
	mcp.WithNumber(
		"page", mcp.Description("当前页，默认1")),
}

var ListPipelinesTool = func() mcp.Tool {
	return mcp.NewTool(ListPipelines, ListPipelinesOptions...)
}()

func ListPipelinesFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/flow/organizations/%s/pipelines", organizationId)

	queryParams := make(map[string]string)

	if createStartTime, ok := request.Params.Arguments["createStartTime"]; ok && createStartTime != nil {
		queryParams["createStartTime"] = fmt.Sprintf("%v", createStartTime)
	}

	if createEndTime, ok := request.Params.Arguments["createEndTime"]; ok && createEndTime != nil {
		queryParams["createEndTime"] = fmt.Sprintf("%v", createEndTime)
	}

	if executeStartTime, ok := request.Params.Arguments["executeStartTime"]; ok && executeStartTime != nil {
		queryParams["executeStartTime"] = fmt.Sprintf("%v", executeStartTime)
	}

	if executeEndTime, ok := request.Params.Arguments["executeEndTime"]; ok && executeEndTime != nil {
		queryParams["executeEndTime"] = fmt.Sprintf("%v", executeEndTime)
	}

	if pipelineName, ok := request.Params.Arguments["pipelineName"].(string); ok && pipelineName != "" {
		queryParams["pipelineName"] = pipelineName
	}

	if statusList, ok := request.Params.Arguments["statusList"].(string); ok && statusList != "" {
		queryParams["statusList"] = statusList
	}

	if perPage, ok := request.Params.Arguments["perPage"]; ok && perPage != nil {
		queryParams["perPage"] = fmt.Sprintf("%v", perPage)
	}

	if page, ok := request.Params.Arguments["page"]; ok && page != nil {
		queryParams["page"] = fmt.Sprintf("%v", page)
	}

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl, utils.WithQueries(queryParams))

	response := ListPipelinesResponse{}

	return yunxiaoClient.HandleMCPResult(&response)
}
