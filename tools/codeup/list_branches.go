package codeup

import (
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-devops-mcp-server/types"
	"github.com/aliyun/alibabacloud-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
	"net/url"
	"strings"
)

const (
	ListBranches = "list_branches"
)

var ListBranchesOptions = []mcp.ToolOption{
	mcp.WithDescription("查询分支列表"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者组织ID与仓库名称的组合，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F）"),
		mcp.Required()),
	mcp.WithNumber(
		"page", mcp.Description("页码")),
	mcp.WithNumber(
		"perPage", mcp.Description("每页大小")),
	mcp.WithString(
		"sort", mcp.Description("排序方式：name_asc - 名称升序，name_desc - 名称降序；updated_asc - 更新时间升序；updated_desc - 更新时间降序"),
		mcp.Enum("name_asc", "name_desc", "updated_asc", "updated_desc")),
	mcp.WithString(
		"search", mcp.Description("查询条件")),
}

var ListBranchesTool = func() mcp.Tool {
	return mcp.NewTool(ListBranches, ListBranchesOptions...)
}()

func ListBranchesFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)

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

	// 构建基本URL
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/branches",
		organizationId, repositoryId)

	// 创建客户端
	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)

	// 添加查询参数
	yunxiaoClient.Query = make(map[string]string)

	// 添加可选参数
	if page, ok := request.Params.Arguments["page"]; ok && page != nil {
		yunxiaoClient.Query["page"] = fmt.Sprintf("%v", page)
	}

	if perPage, ok := request.Params.Arguments["perPage"]; ok && perPage != nil {
		yunxiaoClient.Query["perPage"] = fmt.Sprintf("%v", perPage)
	}

	if sort, ok := request.Params.Arguments["sort"]; ok && sort != nil {
		yunxiaoClient.Query["sort"] = sort.(string)
	}

	if search, ok := request.Params.Arguments["search"]; ok && search != nil {
		yunxiaoClient.Query["search"] = search.(string)
	}

	// 设置响应数据类型为Branch数组
	branches := make([]types.Branch, 0)
	return yunxiaoClient.HandleMCPResult(&branches)
}
