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
	ListChangeRequestComments = "list_change_request_comments"
)

var ListChangeRequestCommentsOptions = []mcp.ToolOption{
	mcp.WithDescription("查询合并请求评论列表"),
	mcp.WithString("organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"), mcp.Required()),
	mcp.WithString("repositoryId", mcp.Description("代码库ID或者组织ID与仓库名称的组合，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F）"), mcp.Required()),
	mcp.WithString("localId", mcp.Description("局部ID，表示代码库中第几个合并请求"), mcp.Required()),
	mcp.WithArray("patchSetBizIds", mcp.Description("关联版本ID列表，每个评论会关联一个版本，表示该评论发表在第几个版本，其中对于全局评论，关联的是最新的合并源版本")),
	mcp.WithString("commentType", mcp.Description("评论类型"), mcp.DefaultString("GLOBAL_COMMENT"), mcp.Enum("GLOBAL_COMMENT", "INLINE_COMMENT")),
	mcp.WithString("state", mcp.Description("评论状态"), mcp.DefaultString("OPENED"), mcp.Enum("OPENED", "DRAFT")),
	mcp.WithBoolean("resolved", mcp.Description("是否标记已解决"), mcp.DefaultBool(false)),
	mcp.WithString("filePath", mcp.Description("文件名称，只有行内评论才有")),
}

var ListChangeRequestCommentsTool = func() mcp.Tool {
	return mcp.NewTool(ListChangeRequestComments, ListChangeRequestCommentsOptions...)
}()

func ListChangeRequestCommentsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	localId := request.Params.Arguments["localId"].(string)

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

	// 构建API URL
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/changeRequests/%s/comments/list",
		organizationId, repositoryId, localId)

	// 准备payload
	payload := map[string]interface{}{}

	// 处理patchSetBizIds参数，确保有默认空数组
	if val, exists := request.Params.Arguments["patchSetBizIds"]; exists && val != nil {
		payload["patchSetBizIds"] = val
	} else {
		payload["patchSetBizIds"] = []interface{}{}
	}

	// 添加其他可选参数
	if val, exists := request.Params.Arguments["commentType"]; exists && val != nil {
		payload["commentType"] = val
	} else {
		payload["commentType"] = "GLOBAL_COMMENT" // 默认值
	}

	if val, exists := request.Params.Arguments["state"]; exists && val != nil {
		payload["state"] = val
	} else {
		payload["state"] = "OPENED" // 默认值
	}

	if val, exists := request.Params.Arguments["resolved"]; exists && val != nil {
		payload["resolved"] = val
	} else {
		payload["resolved"] = false // 默认值
	}

	if val, exists := request.Params.Arguments["filePath"]; exists && val != nil {
		payload["filePath"] = val
	}

	// 创建客户端并发送请求
	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(payload))
	comments := &[]types.ChangeRequestComment{}
	return yunxiaoClient.HandleMCPResult(comments)
}
