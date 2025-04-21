package projex

import (
	"context"
	"devops.aliyun.com/mcp-yunxiao/types"
	"devops.aliyun.com/mcp-yunxiao/utils"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

const (
	SearchProjects = "search_projects"
)

// SearchProjectsResponse 定义项目列表的返回结构
type SearchProjectsResponse []types.ProjectInfo

var SearchProjectsOptions = []mcp.ToolOption{
	mcp.WithDescription("搜索项目"),
	mcp.WithString(
		"organizationId", mcp.Description("组织ID"),
		mcp.Required()),

	// 简化的搜索参数
	mcp.WithString(
		"name", mcp.Description("项目名称包含的文本")),
	mcp.WithString(
		"status", mcp.Description("项目状态ID，多个用逗号分隔")),
	mcp.WithString(
		"createdAfter", mcp.Description("创建时间不早于，格式：YYYY-MM-DD")),
	mcp.WithString(
		"createdBefore", mcp.Description("创建时间不晚于，格式：YYYY-MM-DD")),
	mcp.WithString(
		"creator", mcp.Description("创建者")),
	mcp.WithString(
		"admin", mcp.Description("管理员")),
	mcp.WithString(
		"logicalStatus", mcp.Description("逻辑状态，如NORMAL")),

	// 高级参数
	mcp.WithString(
		"advancedConditions", mcp.Description("高级过滤条件，JSON格式")),
	mcp.WithString(
		"extraConditions", mcp.Description("额外的过滤条件，例如我管理的、我参与的，我收藏的等")),
	mcp.WithString(
		"orderBy", mcp.Description("排序字段，默认为gmtCreate，支持：gmtCreate(创建时间)，name(名称)")),
	mcp.WithNumber(
		"page", mcp.Description("分页参数，第几页")),
	mcp.WithNumber(
		"perPage", mcp.Description("分页参数，每页大小，0-200，默认值20")),
	mcp.WithString(
		"sort", mcp.Description("排序方式，默认为desc，可选：desc(降序)，asc(升序)")),
}

var SearchProjectsTool = func() mcp.Tool {
	return mcp.NewTool(SearchProjects, SearchProjectsOptions...)
}()

// FilterCondition 过滤条件
type FilterCondition struct {
	ClassName       string        `json:"className"`
	FieldIdentifier string        `json:"fieldIdentifier"`
	Format          string        `json:"format"`
	Operator        string        `json:"operator"`
	ToValue         interface{}   `json:"toValue"`
	Value           []interface{} `json:"value"`
}

// Conditions 完整条件对象
type Conditions struct {
	ConditionGroups [][]FilterCondition `json:"conditionGroups"`
}

func SearchProjectsFunc(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	organizationId := request.Params.Arguments["organizationId"].(string)

	apiUrl := fmt.Sprintf("/oapi/v1/projex/organizations/%s/projects:search", organizationId)

	payload := make(map[string]interface{})

	// 处理条件参数
	conditions := buildProjectConditions(request.Params.Arguments)
	if conditions != "" {
		payload["conditions"] = conditions
	}

	// 添加其他可选参数
	if extraConditions, ok := request.Params.Arguments["extraConditions"].(string); ok && extraConditions != "" {
		payload["extraConditions"] = extraConditions
	}

	if orderBy, ok := request.Params.Arguments["orderBy"].(string); ok && orderBy != "" {
		payload["orderBy"] = orderBy
	}

	if page, ok := request.Params.Arguments["page"]; ok {
		payload["page"] = page
	}

	if perPage, ok := request.Params.Arguments["perPage"]; ok {
		payload["perPage"] = perPage
	}

	if sort, ok := request.Params.Arguments["sort"].(string); ok && sort != "" {
		payload["sort"] = sort
	}

	yunxiaoClient := utils.NewYunxiaoClient("POST", apiUrl, utils.WithPayload(payload))

	response := SearchProjectsResponse{}

	return yunxiaoClient.HandleMCPResult(&response)
}

// buildProjectConditions 构建项目搜索条件
func buildProjectConditions(args map[string]interface{}) string {
	// 如果直接提供了高级条件，优先使用
	if advancedConditions, ok := args["advancedConditions"].(string); ok && advancedConditions != "" {
		return advancedConditions
	}

	// 构建条件组
	var filterConditions []FilterCondition

	// 处理名称
	if name, ok := args["name"].(string); ok && name != "" {
		filterConditions = append(filterConditions, FilterCondition{
			ClassName:       "string",
			FieldIdentifier: "name",
			Format:          "input",
			Operator:        "CONTAINS",
			ToValue:         nil,
			Value:           []interface{}{name},
		})
	}

	// 处理状态
	if status, ok := args["status"].(string); ok && status != "" {
		statusValues := strings.Split(status, ",")
		values := make([]interface{}, len(statusValues))
		for i, v := range statusValues {
			values[i] = strings.TrimSpace(v)
		}

		filterConditions = append(filterConditions, FilterCondition{
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

		filterConditions = append(filterConditions, FilterCondition{
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

		filterConditions = append(filterConditions, FilterCondition{
			ClassName:       "user",
			FieldIdentifier: "creator",
			Format:          "list",
			Operator:        "CONTAINS",
			ToValue:         nil,
			Value:           values,
		})
	}

	// 处理管理员
	if admin, ok := args["admin"].(string); ok && admin != "" {
		adminValues := strings.Split(admin, ",")
		values := make([]interface{}, len(adminValues))
		for i, v := range adminValues {
			values[i] = strings.TrimSpace(v)
		}

		filterConditions = append(filterConditions, FilterCondition{
			ClassName:       "user",
			FieldIdentifier: "project.admin",
			Format:          "multiList",
			Operator:        "CONTAINS",
			ToValue:         nil,
			Value:           values,
		})
	}

	// 处理逻辑状态
	if logicalStatus, ok := args["logicalStatus"].(string); ok && logicalStatus != "" {
		filterConditions = append(filterConditions, FilterCondition{
			ClassName:       "string",
			FieldIdentifier: "logicalStatus",
			Format:          "list",
			Operator:        "CONTAINS",
			ToValue:         nil,
			Value:           []interface{}{logicalStatus},
		})
	}

	// 如果没有任何条件，返回空字符串
	if len(filterConditions) == 0 {
		return ""
	}

	// 构建完整条件对象
	conditions := Conditions{
		ConditionGroups: [][]FilterCondition{filterConditions},
	}

	// 序列化为JSON
	conditionsJson, err := json.Marshal(conditions)
	if err != nil {
		return ""
	}

	return string(conditionsJson)
}
