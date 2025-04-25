package projex

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibabacloud-devops-mcp-server/types"
	"github.com/aliyun/alibabacloud-devops-mcp-server/utils"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

const (
	SearchWorkitems = "search_workitems"
)

// SearchWorkitemsResponse 定义工作项列表的返回结构
type SearchWorkitemsResponse []types.WorkItem

var SearchWorkitemsOptions = []mcp.ToolOption{
	mcp.WithDescription("搜索工作项"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),
	mcp.WithString(
		"category", mcp.Description("搜索的工作项类型，例如 Req(需求)、Task(任务)、Bug(缺陷)等，多值用逗号隔开"),
		mcp.Required()),
	mcp.WithString(
		"spaceId", mcp.Description("项目ID, 项目唯一标识"),
		mcp.Required()),

	// 简化的搜索参数
	mcp.WithString(
		"subject", mcp.Description("标题包含的文本")),
	mcp.WithString(
		"status", mcp.Description("状态ID，多个用逗号分隔。状态名称及其ID：待确认（28），待处理（100005），再次打开（30），推迟修复（34），已确认（32），已选择（625489），分析中（154395），分析完成（165115），处理中（100010），设计中（156603），设计完成（307012），开发中（142838），开发完成（100011），测试中（100012）")),
	mcp.WithString(
		"createdAfter", mcp.Description("创建时间不早于，格式：YYYY-MM-DD")),
	mcp.WithString(
		"createdBefore", mcp.Description("创建时间不晚于，格式：YYYY-MM-DD")),
	mcp.WithString(
		"creator", mcp.Description("创建者ID，多个用逗号分隔")),
	mcp.WithString(
		"assignedTo", mcp.Description("负责人ID，多个用逗号分隔")),

	// 高级参数
	mcp.WithString(
		"advancedConditions", mcp.Description("高级过滤条件，JSON格式")),
	mcp.WithString(
		"orderBy", mcp.Description("排序字段，默认为gmtCreate"), mcp.DefaultString("gmtCreate")),
}

var SearchWorkitemsTool = func() mcp.Tool {
	return mcp.NewTool(SearchWorkitems, SearchWorkitemsOptions...)
}()

// WorkitemFilterCondition 工作项过滤条件
type WorkitemFilterCondition struct {
	ClassName       string        `json:"className"`
	FieldIdentifier string        `json:"fieldIdentifier"`
	Format          string        `json:"format"`
	Operator        string        `json:"operator"`
	ToValue         interface{}   `json:"toValue"`
	Value           []interface{} `json:"value"`
}

// WorkitemConditions 完整条件对象
type WorkitemConditions struct {
	ConditionGroups [][]WorkitemFilterCondition `json:"conditionGroups"`
}

func SearchWorkitemsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)

	// 构建API URL
	apiUrl := fmt.Sprintf("/oapi/v1/projex/organizations/%s/workitems:search", organizationId)

	// 创建请求体
	payload := make(map[string]interface{})

	// 添加必填参数
	payload["category"] = request.Params.Arguments["category"]
	payload["spaceId"] = request.Params.Arguments["spaceId"]

	// 处理条件参数
	conditions := buildWorkitemConditions(request.Params.Arguments)
	if conditions != "" {
		payload["conditions"] = conditions
	}

	// 添加其他可选参数
	if orderBy, ok := request.Params.Arguments["orderBy"].(string); ok && orderBy != "" {
		payload["orderBy"] = orderBy
	} else {
		payload["orderBy"] = "gmtCreate"
	}

	// 创建客户端
	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(payload))

	// 创建响应对象
	response := SearchWorkitemsResponse{}

	return yunxiaoClient.HandleMCPResult(&response)
}

// buildWorkitemConditions 构建工作项搜索条件
func buildWorkitemConditions(args map[string]interface{}) string {
	// 如果直接提供了高级条件，优先使用
	if advancedConditions, ok := args["advancedConditions"].(string); ok && advancedConditions != "" {
		return advancedConditions
	}

	// 构建条件组
	var filterConditions []WorkitemFilterCondition

	// 处理标题
	if subject, ok := args["subject"].(string); ok && subject != "" {
		filterConditions = append(filterConditions, WorkitemFilterCondition{
			ClassName:       "string",
			FieldIdentifier: "subject",
			Format:          "input",
			Operator:        "CONTAINS",
			ToValue:         nil,
			Value:           []interface{}{subject},
		})
	}

	// 处理状态
	if status, ok := args["status"].(string); ok && status != "" {
		statusValues := strings.Split(status, ",")
		values := make([]interface{}, len(statusValues))
		for i, v := range statusValues {
			values[i] = strings.TrimSpace(v)
		}

		filterConditions = append(filterConditions, WorkitemFilterCondition{
			ClassName:       "status",
			FieldIdentifier: "status",
			Format:          "list",
			Operator:        "CONTAINS",
			ToValue:         nil,
			Value:           values,
		})
	}

	// 处理创建时间范围
	if createdAfter, ok := args["createdAfter"].(string); ok && createdAfter != "" {
		createdBefore := ""
		if beforeVal, ok := args["createdBefore"].(string); ok && beforeVal != "" {
			createdBefore = beforeVal + " 23:59:59"
		}

		filterConditions = append(filterConditions, WorkitemFilterCondition{
			ClassName:       "date",
			FieldIdentifier: "gmtCreate",
			Format:          "input",
			Operator:        "BETWEEN",
			ToValue:         createdBefore,
			Value:           []interface{}{createdAfter + " 00:00:00"},
		})
	}

	// 处理创建者
	if creator, ok := args["creator"].(string); ok && creator != "" {
		creatorValues := strings.Split(creator, ",")
		values := make([]interface{}, len(creatorValues))
		for i, v := range creatorValues {
			values[i] = strings.TrimSpace(v)
		}

		filterConditions = append(filterConditions, WorkitemFilterCondition{
			ClassName:       "user",
			FieldIdentifier: "creator",
			Format:          "list",
			Operator:        "CONTAINS",
			ToValue:         nil,
			Value:           values,
		})
	}

	// 处理负责人
	if assignedTo, ok := args["assignedTo"].(string); ok && assignedTo != "" {
		assignedToValues := strings.Split(assignedTo, ",")
		values := make([]interface{}, len(assignedToValues))
		for i, v := range assignedToValues {
			values[i] = strings.TrimSpace(v)
		}

		filterConditions = append(filterConditions, WorkitemFilterCondition{
			ClassName:       "user",
			FieldIdentifier: "assignedTo",
			Format:          "list",
			Operator:        "CONTAINS",
			ToValue:         nil,
			Value:           values,
		})
	}

	// 如果没有任何条件，返回空字符串
	if len(filterConditions) == 0 {
		return ""
	}

	// 构建完整条件对象
	conditions := WorkitemConditions{
		ConditionGroups: [][]WorkitemFilterCondition{filterConditions},
	}

	// 序列化为JSON
	conditionsJson, err := json.Marshal(conditions)
	if err != nil {
		return ""
	}

	return string(conditionsJson)
}
