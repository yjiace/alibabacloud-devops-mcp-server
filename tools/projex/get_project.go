package projex

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetProject = "get_project"
)

var GetProjectOptions = []mcp.ToolOption{
	mcp.WithDescription("获取项目详情"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"id", mcp.Description("项目唯一标识"),
		mcp.Required()),
}

var GetProjectTool = func() mcp.Tool {
	return mcp.NewTool(GetProject, GetProjectOptions...)
}()

func GetProjectFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	projectId := request.Params.Arguments["id"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/projex/organizations/%s/projects/%s", organizationId, projectId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)

	projectInfo := &types.ProjectInfo{}

	return yunxiaoClient.HandleMCPResult(projectInfo)
}
