package types

// PipelineRunInfo 流水线运行实例信息
type PipelineRunInfo struct {
	CreateTime        int64            `json:"createTime,omitempty"`
	CreatorAccountId  string           `json:"creatorAccountId,omitempty"`
	GlobalParams      []PipelineParam  `json:"globalParams,omitempty"`
	Groups            []PipelineGroup  `json:"groups,omitempty"`
	ModifierAccountId string           `json:"modifierAccountId,omitempty"`
	PipelineId        int64            `json:"pipelineId,omitempty"`
	PipelineRunId     int64            `json:"pipelineRunId,omitempty"`
	Sources           []PipelineSource `json:"sources,omitempty"`
	StageGroup        []string         `json:"stageGroup,omitempty"`
	Stages            []PipelineStage  `json:"stages,omitempty"`
	Status            string           `json:"status,omitempty"`
	TriggerMode       int              `json:"triggerMode,omitempty"`
	UpdateTime        int64            `json:"updateTime,omitempty"`
}

type PipelineParam struct {
	Encrypted bool   `json:"encrypted,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
}

type PipelineGroup struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PipelineStage struct {
	Index     string    `json:"index,omitempty"`
	Name      string    `json:"name,omitempty"`
	StageInfo StageInfo `json:"stageInfo,omitempty"`
}

type StageInfo struct {
	EndTime   int64     `json:"endTime,omitempty"`
	Jobs      []JobInfo `json:"jobs,omitempty"`
	Name      string    `json:"name,omitempty"`
	StartTime int64     `json:"startTime,omitempty"`
	Status    string    `json:"status,omitempty"`
}

type JobInfo struct {
	Actions   []JobAction `json:"actions,omitempty"`
	EndTime   int64       `json:"endTime,omitempty"`
	ID        int64       `json:"id,omitempty"`
	JobSign   string      `json:"jobSign,omitempty"`
	Result    string      `json:"result,omitempty"`
	Name      string      `json:"name,omitempty"`
	Params    string      `json:"params,omitempty"`
	StartTime int64       `json:"startTime,omitempty"`
	Status    string      `json:"status,omitempty"`
}

type JobAction struct {
	Data        string      `json:"data,omitempty"`
	Disable     bool        `json:"disable,omitempty"`
	DisplayType string      `json:"displayType,omitempty"`
	Name        string      `json:"name,omitempty"`
	Order       int         `json:"order,omitempty"`
	Params      interface{} `json:"params,omitempty"`
	Title       string      `json:"title,omitempty"`
	Type        string      `json:"type,omitempty"`
}
