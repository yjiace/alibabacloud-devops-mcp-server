package flow

import (
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-devops-mcp-server/types"
	"github.com/aliyun/alibabacloud-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetLatestPipelineRun = "get_latest_pipeline_run"
)

var GetLatestPipelineRunOptions = []mcp.ToolOption{
	mcp.WithDescription("获取最近一次流水线运行信息"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"pipelineId", mcp.Description("流水线ID"),
		mcp.Required()),
}

var GetLatestPipelineRunTool = func() mcp.Tool {
	return mcp.NewTool(GetLatestPipelineRun, GetLatestPipelineRunOptions...)
}()

func GetLatestPipelineRunFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	pipelineId := request.Params.Arguments["pipelineId"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/flow/organizations/%s/pipelines/%s/runs/latestPipelineRun", organizationId, pipelineId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)

	response := types.PipelineRunInfo{}

	return yunxiaoClient.HandleMCPResult(&response)
}
