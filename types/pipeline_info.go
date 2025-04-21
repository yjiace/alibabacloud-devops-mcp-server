package types

// PipelineBaseInfo 流水线基本信息
type PipelineBaseInfo struct {
	CreateAccountId string `json:"createAccountId,omitempty"`
	CreateTime      int64  `json:"createTime,omitempty"`
	PipelineId      int64  `json:"pipelineId,omitempty"`
	PipelineName    string `json:"pipelineName,omitempty"`
}

// PipelineInfo 流水线详情信息
type PipelineInfo struct {
	CreateTime        int64          `json:"createTime,omitempty"`
	CreatorAccountId  string         `json:"creatorAccountId,omitempty"`
	EnvId             int64          `json:"envId,omitempty"`
	EnvName           string         `json:"envName,omitempty"`
	GroupId           int64          `json:"groupId,omitempty"`
	ModifierAccountId string         `json:"modifierAccountId,omitempty"`
	Name              string         `json:"name,omitempty"`
	PipelineConfig    PipelineConfig `json:"pipelineConfig,omitempty"`
	TagList           []PipelineTag  `json:"tagList,omitempty"`
	UpdateTime        int64          `json:"updateTime,omitempty"`
}

type PipelineConfig struct {
	Flow     string           `json:"flow,omitempty"`
	Settings string           `json:"settings,omitempty"`
	Sources  []PipelineSource `json:"sources,omitempty"`
}

type PipelineSource struct {
	Data PipelineSourceData `json:"data,omitempty"`
	Sign string             `json:"sign,omitempty"`
	Type string             `json:"type,omitempty"`
}

type PipelineSourceData struct {
	Branch              string   `json:"branch,omitempty"`
	CloneDepth          int      `json:"cloneDepth,omitempty"`
	CredentialId        int64    `json:"credentialId,omitempty"`
	CredentialLabel     string   `json:"credentialLabel,omitempty"`
	CredentialType      string   `json:"credentialType,omitempty"`
	Events              []string `json:"events,omitempty"`
	IsBranchMode        bool     `json:"isBranchMode,omitempty"`
	IsCloneDepth        bool     `json:"isCloneDepth,omitempty"`
	IsSubmodule         bool     `json:"isSubmodule,omitempty"`
	IsTrigger           bool     `json:"isTrigger,omitempty"`
	Label               string   `json:"label,omitempty"`
	Namespace           string   `json:"namespace,omitempty"`
	Repo                string   `json:"repo,omitempty"`
	ServiceConnectionId int64    `json:"serviceConnectionId,omitempty"`
	TriggerFilter       string   `json:"triggerFilter,omitempty"`
	Webhook             string   `json:"webhook,omitempty"`
}

type PipelineTag struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
