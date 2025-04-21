package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"net/url"
)

const (
	DeleteFile = "delete_file"
)

var DeleteFileOptions = []mcp.ToolOption{
	mcp.WithDescription("删除代码库文件"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径"),
		mcp.Required()),
	mcp.WithString(
		"filePath", mcp.Description("文件路径，需使用URL-Encoder编码进行处理"),
		mcp.Required()),
	mcp.WithString(
		"branch", mcp.Description("分支名称"),
		mcp.Required()),
	mcp.WithString(
		"commitMessage", mcp.Description("提交信息"),
		mcp.Required()),
}

var DeleteFileTool = func() mcp.Tool {
	return mcp.NewTool(DeleteFile, DeleteFileOptions...)
}()

func DeleteFileFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	filePath := request.Params.Arguments["filePath"].(string)
	branch := request.Params.Arguments["branch"].(string)
	commitMessage := request.Params.Arguments["commitMessage"].(string)

	// URL编码文件路径
	encodedFilePath := url.PathEscape(filePath)

	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/files/%s",
		organizationId, repositoryId, encodedFilePath)

	queryParams := map[string]string{
		"branch":        branch,
		"commitMessage": commitMessage,
	}

	yunxiaoClient := utils.NewYunxiaoClient("DELETE", apiUrl, utils.WithQueries(queryParams))

	response := types.DeleteFileResponse{}

	return yunxiaoClient.HandleMCPResult(&response)
}
