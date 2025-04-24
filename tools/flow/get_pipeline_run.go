package flow

import (
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-devops-mcp-server/types"
	"github.com/aliyun/alibabacloud-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetPipelineRun = "get_pipeline_run"
)

var GetPipelineRunOptions = []mcp.ToolOption{
	mcp.WithDescription("获取流水线运行实例"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"pipelineId", mcp.Description("流水线ID"),
		mcp.Required()),
	mcp.WithString(
		"pipelineRunId", mcp.Description("流水线运行ID"),
		mcp.Required()),
}

var GetPipelineRunTool = func() mcp.Tool {
	return mcp.NewTool(GetPipelineRun, GetPipelineRunOptions...)
}()

func GetPipelineRunFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	pipelineId := request.Params.Arguments["pipelineId"].(string)
	pipelineRunId := request.Params.Arguments["pipelineRunId"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/flow/organizations/%s/pipelines/%s/runs/%s", organizationId, pipelineId, pipelineRunId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)

	response := types.PipelineRunInfo{}

	return yunxiaoClient.HandleMCPResult(&response)
}
