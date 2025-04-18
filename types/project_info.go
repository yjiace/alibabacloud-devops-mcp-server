package types

// ProjectInfo 项目信息结构
type ProjectInfo struct {
	Creator           UserInfo            `json:"creator,omitempty"`
	CustomCode        string              `json:"customCode,omitempty"`
	CustomFieldValues []CustomFieldValues `json:"customFieldValues,omitempty"`
	Description       string              `json:"description,omitempty"`
	GmtCreate         string              `json:"gmtCreate,omitempty"`
	GmtModified       string              `json:"gmtModified,omitempty"`
	Icon              string              `json:"icon,omitempty"`
	ID                string              `json:"id,omitempty"`
	LogicalStatus     string              `json:"logicalStatus,omitempty"`
	Modifier          UserInfo            `json:"modifier,omitempty"`
	Name              string              `json:"name,omitempty"`
	Scope             string              `json:"scope,omitempty"`
	Status            ProjectStatusInfo   `json:"status,omitempty"`
}

type UserInfo struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CustomFieldValues struct {
	FieldID   string      `json:"fieldId,omitempty"`
	FieldName string      `json:"fieldName,omitempty"`
	Values    []FieldItem `json:"values,omitempty"`
}

type FieldItem struct {
	DisplayValue string `json:"displayValue,omitempty"`
	Identifier   string `json:"identifier,omitempty"`
}

type ProjectStatusInfo struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
