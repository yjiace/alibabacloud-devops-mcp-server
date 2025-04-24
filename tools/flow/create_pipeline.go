package flow

import (
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreatePipeline = "create_pipeline"
)

// PipelineCreateRequest 创建流水线的请求
type PipelineCreateRequest struct {
	Content string `json:"content"`
	Name    string `json:"name"`
}

var CreatePipelineOptions = []mcp.ToolOption{
	mcp.WithDescription("创建流水线"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"name", mcp.Description("流水线名称，最大支持60个字符"),
		mcp.Required()),
	mcp.WithString(
		"content", mcp.Description("流水线YAML描述，可参考YAML流水线的帮助文档编写"),
		mcp.Required()),
}

var CreatePipelineTool = func() mcp.Tool {
	return mcp.NewTool(CreatePipeline, CreatePipelineOptions...)
}()

func CreatePipelineFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	name := request.Params.Arguments["name"].(string)
	content := request.Params.Arguments["content"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/flow/organizations/%s/pipelines", organizationId)

	payload := PipelineCreateRequest{
		Content: content,
		Name:    name,
	}

	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(payload))

	var response int64

	return yunxiaoClient.HandleMCPResult(&response)
}
