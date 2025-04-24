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
	DeleteFile = "delete_file"
)

var DeleteFileOptions = []mcp.ToolOption{
	mcp.WithDescription("删除文件"),
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
		"commitMessage", mcp.Description("提交说明"),
		mcp.Required()),
	mcp.WithString(
		"branch", mcp.Description("分支名称"),
		mcp.Required()),
}

var DeleteFileTool = func() mcp.Tool {
	return mcp.NewTool(DeleteFile, DeleteFileOptions...)
}()

func DeleteFileFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	filePath := request.Params.Arguments["filePath"].(string)
	commitMessage := request.Params.Arguments["commitMessage"].(string)
	branch := request.Params.Arguments["branch"].(string)

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
		"branch":        branch,
		"commitMessage": commitMessage,
	}

	yunxiaoClient := utils.NewYunxiaoClient("DELETE", apiUrl, utils.WithQueries(queryParams))

	response := &types.DeleteFileResponse{}

	return yunxiaoClient.HandleMCPResult(response)
}
