package flow

import (
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-devops-mcp-server/types"
	"github.com/aliyun/alibabacloud-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetPipeline = "get_pipeline"
)

var GetPipelineOptions = []mcp.ToolOption{
	mcp.WithDescription("获取流水线详情"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"pipelineId", mcp.Description("流水线ID"),
		mcp.Required()),
}

var GetPipelineTool = func() mcp.Tool {
	return mcp.NewTool(GetPipeline, GetPipelineOptions...)
}()

func GetPipelineFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	pipelineId := request.Params.Arguments["pipelineId"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/flow/organizations/%s/pipelines/%s", organizationId, pipelineId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)

	response := types.PipelineInfo{}

	return yunxiaoClient.HandleMCPResult(&response)
}
