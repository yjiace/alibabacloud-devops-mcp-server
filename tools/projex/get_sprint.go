package projex

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-devops-mcp-server/types"
	"github.com/aliyun/alibaba-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetSprint = "get_sprint"
)

var GetSprintOptions = []mcp.ToolOption{
	mcp.WithDescription("获取迭代详情"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"projectId", mcp.Description("项目唯一标识"),
		mcp.Required()),
	mcp.WithString(
		"id", mcp.Description("迭代唯一标识"),
		mcp.Required()),
}

var GetSprintTool = func() mcp.Tool {
	return mcp.NewTool(GetSprint, GetSprintOptions...)
}()

func GetSprintFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	projectId := request.Params.Arguments["projectId"].(string)
	sprintId := request.Params.Arguments["id"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/projex/organizations/%s/projects/%s/sprints/%s", organizationId, projectId, sprintId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)

	response := types.SprintInfo{}

	return yunxiaoClient.HandleMCPResult(&response)
}
