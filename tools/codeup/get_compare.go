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
	mcp.WithDescription("查询代码比较内容"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径，例如: 2835387 或 codeup-org-id%2Fcodeup-demo"),
		mcp.Required()),
	mcp.WithString(
		"from", mcp.Description("可为CommitSHA、分支名或者标签名"),
		mcp.Required()),
	mcp.WithString(
		"to", mcp.Description("可为CommitSHA、分支名或者标签名"),
		mcp.Required()),
	mcp.WithString(
		"sourceType", mcp.Description("可选值：branch、tag；若是commit比较，可不传；若是分支比较，则需传入：branch，亦可不传，但需要确保不存在分支或tag重名的情况；若是tag比较，则需传入：tag；若是存在分支和标签同名的情况，则需要严格传入branch或者tag")),
	mcp.WithString(
		"targetType", mcp.Description("可选值：branch、tag；若是commit比较，可不传；若是分支比较，则需传入：branch，亦可不传，但需要确保不存在分支或Tag重名的情况；若是tag比较，则需传入：tag；若是存在分支和标签同名的情况，则需要严格传入branch或者tag")),
	mcp.WithString(
		"straight", mcp.Description("是否使用Merge-Base：straight=false，表示使用Merge-Base；straight=true，表示不使用Merge-Base；默认为false，即使用Merge-Base")),
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
