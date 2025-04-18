package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateChangeRequestComment = "create_change_request_comment"
)

var CreateChangeRequestCommentOptions = []mcp.ToolOption{
	mcp.WithDescription("create change request comments"),
	mcp.WithString("organizationId", mcp.Description("organization id"), mcp.Required()),
	mcp.WithString("repositoryId", mcp.Description("repository id"), mcp.Required()),
	mcp.WithString("localId", mcp.Description("local id"), mcp.Required()),
	mcp.WithString("comment_type", mcp.Description("comment type, GLOBAL_COMMENT or INLINE_COMMENT"), mcp.Required(), mcp.Enum("GLOBAL_COMMENT", "INLINE_COMMENT")),
	mcp.WithString("content", mcp.Description("comment content (1-65535 chars)"), mcp.Required()),
	mcp.WithBoolean("draft", mcp.Description("whether this is a draft comment (defaults to false)"), mcp.Required()),
	mcp.WithBoolean("resolved", mcp.Description("whether this comment is resolved (defaults to false)"), mcp.Required()),
	mcp.WithString("patchset_biz_id", mcp.Description("associated version id"), mcp.Required()),
	// 可选参数
	mcp.WithString("file_path", mcp.Description("file path (required for INLINE_COMMENT)")),
	mcp.WithNumber("line_number", mcp.Description("line number (required for INLINE_COMMENT)")),
	mcp.WithString("from_patchset_biz_id", mcp.Description("from patchset biz id (required for INLINE_COMMENT)")),
	mcp.WithString("to_patchset_biz_id", mcp.Description("to patchset biz id (required for INLINE_COMMENT)")),
	mcp.WithString("parent_comment_biz_id", mcp.Description("parent comment biz id")),
}

var CreateChangeRequestCommentTool = func() mcp.Tool {
	return mcp.NewTool(CreateChangeRequestComment, CreateChangeRequestCommentOptions...)
}()

func CreateChangeRequestCommentFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	localId := request.Params.Arguments["localId"].(string)
	commentType := request.Params.Arguments["comment_type"].(string)

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
