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
	GetFileBlobs = "get_file_blobs"
)

var GetFileBlobsOptions = []mcp.ToolOption{
	mcp.WithDescription("查询文件内容"),
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
		"ref", mcp.Description("指定引用名，一般为分支名，可为分支名、标签名和CommitSHA"),
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
