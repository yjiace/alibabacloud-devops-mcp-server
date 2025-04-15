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
	mcp.WithDescription("list change request patch sets"),
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("repository id"),
		mcp.Required()),
	mcp.WithString(
		"localId", mcp.Description("local id"),
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
