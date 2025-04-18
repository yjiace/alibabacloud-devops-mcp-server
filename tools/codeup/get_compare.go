package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetCompare = "get_compare"
)

var GetCompareOptions = []mcp.ToolOption{
	mcp.WithDescription("get repository"),
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("repository id"),
		mcp.Required()),
	mcp.WithString(
		"from", mcp.Description("from"),
		mcp.Required()),
	mcp.WithString(
		"to", mcp.Description("to"),
		mcp.Required()),
	mcp.WithString(
		"sourceType", mcp.Description("source type")),
	mcp.WithString(
		"targetType", mcp.Description("target type")),
	mcp.WithString(
		"straight", mcp.Description("straight")),
}

var GetCompareTool = func() mcp.Tool {
	return mcp.NewTool(GetCompare, GetCompareOptions...)
}()

func GetCompareFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	from := request.Params.Arguments["from"].(string)
	to := request.Params.Arguments["to"].(string)
	
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/compares", organizationId, repositoryId)
	
	// 准备查询参数
	queries := map[string]string{
		"from": from,
		"to":   to,
	}
	
	// 添加可选参数
	if val, ok := request.Params.Arguments["sourceType"]; ok && val != nil {
		queries["sourceType"] = val.(string)
	}
	
	if val, ok := request.Params.Arguments["targetType"]; ok && val != nil {
		queries["targetType"] = val.(string)
	}
	
	if val, ok := request.Params.Arguments["straight"]; ok && val != nil {
		queries["straight"] = val.(string)
	}
	
	// 创建客户端并传入查询参数
	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl, utils.WithQueries(queries))
	
	compare := &types.Compare{}
	return yunxiaoClient.HandleMCPResult(compare)
}
