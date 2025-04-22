package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateFile = "create_file"
)

var CreateFileOptions = []mcp.ToolOption{
	mcp.WithDescription("创建代码库文件"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"repositoryId", mcp.Description("代码库ID或者URL-Encoder编码的全路径，例如: 2835387 或 codeup-org-id%2Fcodeup-demo"),
		mcp.Required()),
	mcp.WithString(
		"branch", mcp.Description("分支名称"),
		mcp.Required()),
	mcp.WithString(
		"filePath", mcp.Description("文件路径，非空，一般不要使用包含特殊符号的文件路径"),
		mcp.Required()),
	mcp.WithString(
		"content", mcp.Description("文件内容"),
		mcp.Required()),
	mcp.WithString(
		"commitMessage", mcp.Description("提交信息，非空，不超过102400个字符"),
		mcp.Required()),
	mcp.WithString(
		"encoding", mcp.Description("编码规则，可选值{text, base64}，默认为text")),
}

var CreateFileTool = func() mcp.Tool {
	return mcp.NewTool(CreateFile, CreateFileOptions...)
}()

func CreateFileFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	branch := request.Params.Arguments["branch"].(string)
	filePath := request.Params.Arguments["filePath"].(string)
	content := request.Params.Arguments["content"].(string)
	commitMessage := request.Params.Arguments["commitMessage"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/files", organizationId, repositoryId)

	requestBody := types.CreateFileRequest{
		Branch:        branch,
		FilePath:      filePath,
		Content:       content,
		CommitMessage: commitMessage,
		Encoding:      "text", // 默认使用text编码
	}

	if encoding, ok := request.Params.Arguments["encoding"].(string); ok && encoding != "" {
		requestBody.Encoding = encoding
	}

	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(requestBody))

	response := types.CreateFileResponse{}

	return yunxiaoClient.HandleMCPResult(&response)
}
