package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"net/url"
	"strings"
)

const (
	ListFiles = "list_files"
)

var ListFilesOptions = []mcp.ToolOption{
	mcp.WithDescription("查询文件树"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者组织ID与仓库名称的组合，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F）"),
		mcp.Required()),
	mcp.WithString(
		"path", mcp.Description("指定查询的路径，例如需要查询 src/main 目录下的文件")),
	mcp.WithString(
		"ref", mcp.Description("指定引用名，一般为分支名，可为分支名、标签名和CommitSHA，若不传值，则为当前代码库的默认分支，如 master")),
	mcp.WithString(
		"type", mcp.Description("文件树获取方式：DIRECT - 仅获取当前目录，默认方式；RECURSIVE - 递归查找当前路径下的所有文件；FLATTEN - 扁平化展示（如果是目录，递归查找，直到子目录包含文件或多个目录为止）"),
		mcp.Enum("DIRECT", "RECURSIVE", "FLATTEN")),
}

var ListFilesTool = func() mcp.Tool {
	return mcp.NewTool(ListFiles, ListFilesOptions...)
}()

func ListFilesFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/files/tree",
		organizationId, repositoryId)

	queryParams := make(map[string]string)

	if path, ok := request.Params.Arguments["path"].(string); ok && path != "" {
		// 确保path已被URL编码
		if strings.Contains(path, "/") {
			path = url.PathEscape(path)
		}
		queryParams["path"] = path
	}

	if ref, ok := request.Params.Arguments["ref"].(string); ok && ref != "" {
		queryParams["ref"] = ref
	}

	if fileType, ok := request.Params.Arguments["type"].(string); ok && fileType != "" {
		queryParams["type"] = fileType
	}

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl, utils.WithQueries(queryParams))

	response := []types.FileInfo{}

	return yunxiaoClient.HandleMCPResult(&response)
}
