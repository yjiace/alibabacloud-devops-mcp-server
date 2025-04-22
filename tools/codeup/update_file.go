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
	UpdateFile = "update_file"
)

var UpdateFileOptions = []mcp.ToolOption{
	mcp.WithDescription("更新文件内容"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径，例如: 2835387 或 codeup-org-id%2Fcodeup-demo"),
		mcp.Required()),
	mcp.WithString(
		"filePath", mcp.Description("文件路径，需使用URL-Encoder编码进行处理"),
		mcp.Required()),
	mcp.WithString(
		"branch", mcp.Description("分支名称"),
		mcp.Required()),
	mcp.WithString(
		"content", mcp.Description("文件内容"),
		mcp.Required()),
	mcp.WithString(
		"commitMessage", mcp.Description("提交信息，不能为空，不超过102400个字符"),
		mcp.Required()),
	mcp.WithString(
		"encoding", mcp.Description("编码规则，可选值{text, base64}，默认为text")),
}

var UpdateFileTool = func() mcp.Tool {
	return mcp.NewTool(UpdateFile, UpdateFileOptions...)
}()

func UpdateFileFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	filePath := request.Params.Arguments["filePath"].(string)
	branch := request.Params.Arguments["branch"].(string)
	content := request.Params.Arguments["content"].(string)
	commitMessage := request.Params.Arguments["commitMessage"].(string)

	encodedFilePath := url.QueryEscape(filePath)

	//// 临时开启调试模式
	//oldDebug := utils.Debug
	//utils.Debug = true
	//defer func() {
	//	utils.Debug = oldDebug
	//}()

	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/files/%s",
		organizationId, repositoryId, encodedFilePath)

	requestBody := types.UpdateFileRequest{
		Branch:        branch,
		CommitMessage: commitMessage,
		Content:       content,
		Encoding:      "text", // 默认使用text编码
	}

	if encoding, ok := request.Params.Arguments["encoding"].(string); ok && encoding != "" {
		requestBody.Encoding = encoding
	}

	yunxiaoClient := utils.NewYunxiaoClient("PUT", apiUrl, utils.WithPayload(requestBody))

	response := types.CreateFileResponse{} // 使用CreateFileResponse，因为返回结构相同

	return yunxiaoClient.HandleMCPResult(&response)
}
