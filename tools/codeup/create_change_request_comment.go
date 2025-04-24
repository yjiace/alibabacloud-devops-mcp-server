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
	CreateChangeRequestComment = "create_change_request_comment"
)

var CreateChangeRequestCommentOptions = []mcp.ToolOption{
	mcp.WithDescription("创建合并请求评论"),
	mcp.WithString("organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"), mcp.Required()),
	mcp.WithString("repositoryId", mcp.Description("代码库ID或者组织ID与仓库名称的组合，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F）"), mcp.Required()),
	mcp.WithString("localId", mcp.Description("局部ID，表示代码库中第几个合并请求"), mcp.Required()),
	mcp.WithString("comment_type", mcp.Description("评论类型"), mcp.Required(), mcp.Enum("GLOBAL_COMMENT", "INLINE_COMMENT")),
	mcp.WithString("content", mcp.Description("评论内容，长度必须在1到65535之间"), mcp.Required()),
	mcp.WithBoolean("draft", mcp.Description("是否草稿评论"), mcp.Required()),
	mcp.WithBoolean("resolved", mcp.Description("是否标记已解决"), mcp.Required()),
	mcp.WithString("patchset_biz_id", mcp.Description("关联版本ID，如果是INLINE_COMMENT，则选择from_patchset_biz_id或to_patchset_biz_id中的一个"), mcp.Required()),
	// 可选参数
	mcp.WithString("file_path", mcp.Description("文件名称，只有行内评论才有")),
	mcp.WithNumber("line_number", mcp.Description("行号，只有行内评论才有")),
	mcp.WithString("from_patchset_biz_id", mcp.Description("比较的起始版本ID，INLINE_COMMENT类型的评论必传")),
	mcp.WithString("to_patchset_biz_id", mcp.Description("比较的目标版本ID，INLINE_COMMENT类型的评论必传")),
	mcp.WithString("parent_comment_biz_id", mcp.Description("父评论ID")),
}

var CreateChangeRequestCommentTool = func() mcp.Tool {
	return mcp.NewTool(CreateChangeRequestComment, CreateChangeRequestCommentOptions...)
}()

func CreateChangeRequestCommentFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	localId := request.Params.Arguments["localId"].(string)
	commentType := request.Params.Arguments["comment_type"].(string)

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
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/changeRequests/%s/comments",
		organizationId, repositoryId, localId)

	// 准备payload
	payload := map[string]interface{}{
		"comment_type":    commentType,
		"content":         request.Params.Arguments["content"],
		"patchset_biz_id": request.Params.Arguments["patchset_biz_id"],
	}

	// 处理布尔值参数，如果未提供则使用默认值false
	if draft, ok := request.Params.Arguments["draft"]; ok && draft != nil {
		payload["draft"] = draft
	} else {
		payload["draft"] = false
	}

	if resolved, ok := request.Params.Arguments["resolved"]; ok && resolved != nil {
		payload["resolved"] = resolved
	} else {
		payload["resolved"] = false
	}

	// 根据评论类型添加必要参数
	if commentType == "INLINE_COMMENT" {
		// 检查INLINE_COMMENT必需的参数
		requiredParams := []string{"file_path", "line_number", "from_patchset_biz_id", "to_patchset_biz_id"}
		for _, param := range requiredParams {
			if _, exists := request.Params.Arguments[param]; !exists {
				return mcp.NewToolResultError(fmt.Sprintf("Parameter '%s' is required for INLINE_COMMENT", param)),
					fmt.Errorf("parameter '%s' is required for INLINE_COMMENT", param)
			}

			// 添加必需参数到payload（确保这些参数以正确类型添加到payload中）
			if param == "line_number" {
				payload[param] = request.Params.Arguments[param]
			} else {
				payload[param] = request.Params.Arguments[param].(string)
			}
		}
	}

	// 添加可选参数
	if val, exists := request.Params.Arguments["parent_comment_biz_id"]; exists && val != nil {
		payload["parent_comment_biz_id"] = val.(string)
	}

	// 创建客户端并发送请求
	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(payload))
	comment := &types.ChangeRequestComment{}
	return yunxiaoClient.HandleMCPResult(comment)
}
