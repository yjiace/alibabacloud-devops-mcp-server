package codeup

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"math"
	"net/url"
	"strconv"
	"strings"
)

const (
	CreateChangeRequest = "create_change_request"
)

var CreateChangeRequestOptions = []mcp.ToolOption{
	mcp.WithDescription("创建合并请求"),
	mcp.WithString("organizationId", mcp.Description("组织ID，前往组织管理后台的基本信息页面获取"), mcp.Required()),
	mcp.WithString("repositoryId", mcp.Description("代码库ID或者组织ID与仓库名称的组合，例如: 2835387 或 organizationId%2Frepo-name（注意：斜杠需URL编码为%2F）"), mcp.Required()),
	mcp.WithString("title", mcp.Description("标题，不超过256个字符"), mcp.Required()),
	mcp.WithString("description", mcp.Description("描述，不超过10000个字符")),
	mcp.WithString("sourceBranch", mcp.Description("源分支名称"), mcp.Required()),
	mcp.WithNumber("sourceProjectId", mcp.Description("源库ID（必须是数字ID，不能使用组织ID/仓库名称组合）"), mcp.Required()),
	mcp.WithString("targetBranch", mcp.Description("目标分支名称"), mcp.Required()),
	mcp.WithNumber("targetProjectId", mcp.Description("目标库ID（必须是数字ID，不能使用组织ID/仓库名称组合）"), mcp.Required()),
	mcp.WithArray("reviewerUserIds", mcp.Description("评审人用户ID列表")),
	mcp.WithArray("workItemIds", mcp.Description("关联工作项ID列表")),
	mcp.WithString("createFrom", mcp.Description("创建来源：WEB - 页面创建；COMMAND_LINE - 命令行创建；默认为WEB"), mcp.Enum("WEB", "COMMAND_LINE")),
}

// 提示：sourceProjectId和targetProjectId通常需要与repositoryId相同，所以在调用此工具时，
// 如果是使用数字ID作为repositoryId，也请将相同的值用于sourceProjectId和targetProjectId。
// 如果使用的是organizationId%2Frepo-name格式，请先调用get_repository工具获取该仓库的数字ID。

var CreateChangeRequestTool = func() mcp.Tool {
	return mcp.NewTool(CreateChangeRequest, CreateChangeRequestOptions...)
}()

// 将浮点数转换为整数字符串（去除小数点和小数部分）
func floatToIntString(value interface{}) string {
	// 如果传入的是字符串，先尝试转为浮点数
	if strValue, ok := value.(string); ok {
		if floatValue, err := strconv.ParseFloat(strValue, 64); err == nil {
			value = floatValue
		} else {
			return strValue // 如果转换失败，返回原字符串
		}
	}
	
	// 处理浮点数
	if floatValue, ok := value.(float64); ok {
		intValue := int(math.Floor(floatValue + 0.5)) // 四舍五入转整数
		return strconv.Itoa(intValue)
	}
	
	// 处理其他情况，直接转字符串
	return fmt.Sprintf("%v", value)
}

func CreateChangeRequestFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)
	repositoryId := request.Params.Arguments["repositoryId"].(string)

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

	// 检查参数中是否已包含sourceProjectId和targetProjectId
	_, sourceIdExists := request.Params.Arguments["sourceProjectId"]
	_, targetIdExists := request.Params.Arguments["targetProjectId"]

	// 如果repositoryId是纯数字，且sourceProjectId或targetProjectId未提供，尝试使用repositoryId的值
	if _, err := strconv.Atoi(repositoryId); err == nil {
		// 如果是数字ID，检查是否需要自动填充sourceProjectId或targetProjectId
		if !sourceIdExists {
			return mcp.NewToolResultError("必须提供参数sourceProjectId"),
				fmt.Errorf("必须提供参数sourceProjectId，可使用与repositoryId相同的值: %s", repositoryId)
		}
		if !targetIdExists {
			return mcp.NewToolResultError("必须提供参数targetProjectId"),
				fmt.Errorf("必须提供参数targetProjectId，可使用与repositoryId相同的值: %s", repositoryId)
		}
	} else if strings.Contains(repositoryId, "%2F") {
		// 如果是组织ID与仓库名称的组合，提示用户需要获取对应的数字ID
		return mcp.NewToolResultError("当使用'组织ID%2F仓库名称'格式时，必须先获取该仓库的数字ID并用于sourceProjectId和targetProjectId参数"),
			fmt.Errorf("请先使用get_repository工具获取'%s'的数字ID，然后使用该ID作为sourceProjectId和targetProjectId的值", repositoryId)
	}

	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/changeRequests", organizationId, repositoryId)

	// 获取并处理sourceProjectId和targetProjectId，确保它们是整数字符串（没有小数点）
	sourceProjectId := floatToIntString(request.Params.Arguments["sourceProjectId"])
	targetProjectId := floatToIntString(request.Params.Arguments["targetProjectId"])

	// 准备payload，确保类型正确
	payload := map[string]interface{}{
		"title":           request.Params.Arguments["title"],
		"sourceBranch":    request.Params.Arguments["sourceBranch"],
		"targetBranch":    request.Params.Arguments["targetBranch"],
		"sourceProjectId": sourceProjectId,
		"targetProjectId": targetProjectId,
	}

	// 添加可选参数
	if val, ok := request.Params.Arguments["description"]; ok && val != nil {
		payload["description"] = val
	}

	if val, ok := request.Params.Arguments["reviewerUserIds"]; ok && val != nil {
		payload["reviewerUserIds"] = val
	}

	if val, ok := request.Params.Arguments["workItemIds"]; ok && val != nil {
		payload["workItemIds"] = val
	}

	if val, ok := request.Params.Arguments["createFrom"]; ok && val != nil {
		payload["createFrom"] = val
	} else {
		payload["createFrom"] = "WEB" // 默认为WEB
	}

	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(payload))
	changeRequest := &types.ChangeRequest{}
	return yunxiaoClient.HandleMCPResult(changeRequest)
}
