package codeup

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-devops-mcp-server/types"
	"github.com/aliyun/alibaba-devops-mcp-server/utils"
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
	mcp.WithNumber("sourceProjectId", mcp.Description("源库ID（如不提供，会尝试自动获取）")),
	mcp.WithString("targetBranch", mcp.Description("目标分支名称"), mcp.Required()),
	mcp.WithNumber("targetProjectId", mcp.Description("目标库ID（如不提供，会尝试自动获取）")),
	mcp.WithArray("reviewerUserIds", mcp.Description("评审人用户ID列表")),
	mcp.WithArray("workItemIds", mcp.Description("关联工作项ID列表")),
	mcp.WithString("createFrom", mcp.Description("创建来源：WEB - 页面创建；COMMAND_LINE - 命令行创建；默认为WEB"), mcp.Enum("WEB", "COMMAND_LINE")),
}

// 提示：sourceProjectId和targetProjectId会自动获取。如果repositoryId是数字ID，则直接使用该数字；
// 如果是organizationId%2Frepo-name格式，会自动调用API获取对应的数字ID。

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

// 通过API获取仓库的数字ID
func getRepositoryNumericId(organizationId, repositoryId string) (string, error) {
	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s", organizationId, repositoryId)
	yunxiaoClient := utils.NewYunxiaoClient("GET", apiUrl)

	// 执行请求
	client, err := yunxiaoClient.Do()
	if err != nil {
		return "", fmt.Errorf("调用GetRepository API失败: %v", err)
	}

	// 读取响应体
	body, err := client.GetRespBody()
	if err != nil {
		return "", fmt.Errorf("读取GetRepository API响应失败: %v", err)
	}

	// 解析响应
	var repoInfo struct {
		Id int `json:"id"`
	}

	if err := json.Unmarshal(body, &repoInfo); err != nil {
		return "", fmt.Errorf("解析GetRepository API响应失败: %v", err)
	}

	if repoInfo.Id == 0 {
		return "", fmt.Errorf("无法获取代码库ID")
	}

	return strconv.Itoa(repoInfo.Id), nil
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
	var sourceProjectId, targetProjectId string
	var sourceIdExists, targetIdExists bool
	var tempValue interface{}

	// 检查并获取sourceProjectId
	if tempValue, sourceIdExists = request.Params.Arguments["sourceProjectId"]; sourceIdExists && tempValue != nil {
		sourceProjectId = floatToIntString(tempValue)
	}

	// 检查并获取targetProjectId
	if tempValue, targetIdExists = request.Params.Arguments["targetProjectId"]; targetIdExists && tempValue != nil {
		targetProjectId = floatToIntString(tempValue)
	}

	// 如果repositoryId是纯数字，且sourceProjectId或targetProjectId未提供，直接使用repositoryId的值
	if numericId, err := strconv.Atoi(repositoryId); err == nil {
		// 是数字ID，可以直接使用
		if !sourceIdExists {
			sourceProjectId = strconv.Itoa(numericId)
		}
		if !targetIdExists {
			targetProjectId = strconv.Itoa(numericId)
		}
	} else if strings.Contains(repositoryId, "%2F") {
		// 如果是组织ID与仓库名称的组合，调用API获取数字ID
		if !sourceIdExists || !targetIdExists {
			numericId, err := getRepositoryNumericId(organizationId, repositoryId)
			if err != nil {
				return mcp.NewToolResultError("当使用'组织ID%2F仓库名称'格式时，必须先获取该仓库的数字ID并用于sourceProjectId和targetProjectId参数"),
					fmt.Errorf("请先使用get_repository工具获取'%s'的数字ID，然后使用该ID作为sourceProjectId和targetProjectId的值", repositoryId)
			}

			if !sourceIdExists {
				sourceProjectId = numericId
			}
			if !targetIdExists {
				targetProjectId = numericId
			}
		}
	}

	// 确保sourceProjectId和targetProjectId已设置
	if sourceProjectId == "" {
		return mcp.NewToolResultError("无法获取sourceProjectId"),
			fmt.Errorf("无法获取sourceProjectId，请手动提供此参数")
	}
	if targetProjectId == "" {
		return mcp.NewToolResultError("无法获取targetProjectId"),
			fmt.Errorf("无法获取targetProjectId，请手动提供此参数")
	}

	apiUrl := fmt.Sprintf("/oapi/v1/codeup/organizations/%s/repositories/%s/changeRequests", organizationId, repositoryId)

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
