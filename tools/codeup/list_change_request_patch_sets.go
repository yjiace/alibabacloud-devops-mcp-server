package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListChangeRequestPatchSets = "list_change_requests_patch_sets"
)

var ListChangeRequestPatchSetsOptions = []mcp.ToolOption{
	mcp.WithDescription("查询合并请求版本列表"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径，例如: 2835387 或 codeup-org-id%2Fcodeup-demo"),
		mcp.Required()),
	mcp.WithString(
		"localId", mcp.Description("局部ID，表示代码库中第几个合并请求"),
		mcp.Required()),
}

var ListChangeRequestPatchSetsTool = func() mcp.Tool {
	return mcp.NewTool(ListChangeRequestPatchSets, ListChangeRequestPatchSetsOptions...)
}()

func ListChangeRequestPatchSetsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	localId := request.Params.Arguments["localId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/changeRequests/%s/diffs/patches", organizationId, repositoryId, localId)

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)
	patchSets := make([]types.PatchSet, 0)
	return yunxiaoClient.HandleMCPResult(&patchSets)
}
