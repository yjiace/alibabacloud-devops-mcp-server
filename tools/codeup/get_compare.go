package codeup

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-devops-mcp-server/types"
	"github.com/aliyun/alibaba-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
	"net/url"
	"strings"
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
		"repositoryId", mcp.Description("代码库ID或者组织ID与仓库名称的组合，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F）"),
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

	// 自动处理repositoryId中未编码的斜杠
	if strings.Contains(repositoryId, "/") {
		// 发现未编码的斜杠，自动进行URL编码
		parts := strings.SplitN(repositoryId, "/", 2)
		if len(parts) == 2 {
			encodedRepoName := url.QueryEscape(parts[1])
			// 移除编码中的+号（空格被编码为+，但我们需要%20）
			encodedRepoName = strings.ReplaceAll(encodedRepoName, "+", "%20")
			repositoryId = parts[0] + "%2F" + encodedRepoName
		}
	}

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
