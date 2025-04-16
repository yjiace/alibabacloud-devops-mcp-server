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
	mcp.WithString(
		"organizationId", mcp.Description("organization id"),
		mcp.Required()),
	mcp.WithString("repositoryId", mcp.Description("repository id"), mcp.Required()),
	mcp.WithString("localId", mcp.Description("local id"), mcp.Required()),
	mcp.WithString("comment_type", mcp.Description("comment type, GLOBAL_COMMENT,INLINE_COMMENT"), mcp.Required(), mcp.Enum("GLOBAL_COMMENT", "INLINE_COMMENT")),
	mcp.WithString("content", mcp.Description("content"), mcp.Required()),
	mcp.WithBoolean("draft", mcp.Description("draft"), mcp.Required()),
	mcp.WithNumber("line_number", mcp.Description("line number")),
	mcp.WithBoolean("resolved", mcp.Description("resolved"), mcp.Required()),
	mcp.WithString("patchset_biz_id", mcp.Description("patchset_biz_id"), mcp.Required()),
}

var CreateChangeRequestCommentTool = func() mcp.Tool {
	return mcp.NewTool(CreateChangeRequestComment, CreateChangeRequestCommentOptions...)
}()

func CreateChangeRequestCommentFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)
	localId := request.Params.Arguments["localId"].(string)
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/changeRequests/%s/comments", organizationId, repositoryId, localId)

	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(request.Params.Arguments))
	comment := &types.ChangeRequestComment{}
	return yunxiaoClient.HandleMCPResult(comment)
}
