package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListFiles = "list_files"
)

var ListFilesOptions = []mcp.ToolOption{
	mcp.WithDescription("查询文件树"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径"),
		mcp.Required()),
	mcp.WithString(
		"path", mcp.Description("指定查询的路径，例如需要查询src/main目录下的文件")),
	mcp.WithString(
		"ref", mcp.Description("指定引用名，一般为分支名，可为分支名、标签名和CommitSHA，若不传值，则为当前代码库的默认分支")),
	mcp.WithString(
		"type", mcp.Description("文件树获取方式：DIRECT-仅获取当前目录(默认)；RECURSIVE-递归查找当前路径下的所有文件；FLATTEN-扁平化展示")),
}

var ListFilesTool = func() mcp.Tool {
	return mcp.NewTool(ListFiles, ListFilesOptions...)
}()

func ListFilesFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/files/tree",
		organizationId, repositoryId)

	queryParams := make(map[string]string)

	if path, ok := request.Params.Arguments["path"].(string); ok && path != "" {
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
