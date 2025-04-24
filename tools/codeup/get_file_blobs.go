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
	GetFileBlobs = "get_file_blobs"
)

var GetFileBlobsOptions = []mcp.ToolOption{
	mcp.WithDescription("查询文件内容"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者组织ID与仓库名称的组合，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F）"),
		mcp.Required()),
	mcp.WithString(
		"filePath", mcp.Description("文件路径，需使用URL-Encoder编码进行处理"),
		mcp.Required()),
	mcp.WithString(
		"ref", mcp.Description("指定引用名，一般为分支名，可为分支名、标签名和CommitSHA，若不传值，则为当前代码库的默认分支，如 master"),
		mcp.Required()),
}

var GetFileBlobsTool = func() mcp.Tool {
	return mcp.NewTool(GetFileBlobs, GetFileBlobsOptions...)
}()

func GetFileBlobsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	filePath := request.Params.Arguments["filePath"].(string)
	ref := request.Params.Arguments["ref"].(string)

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

	// 确保filePath已被URL编码
	if strings.Contains(filePath, "/") {
		filePath = url.PathEscape(filePath)
	}

	encodedFilePath := url.PathEscape(filePath)

	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/files/%s",
		organizationId, repositoryId, encodedFilePath)

	queryParams := map[string]string{
		"ref": ref,
	}

	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl, utils.WithQueries(queryParams))

	response := types.FileContent{}

	return yunxiaoClient.HandleMCPResult(&response)
}
