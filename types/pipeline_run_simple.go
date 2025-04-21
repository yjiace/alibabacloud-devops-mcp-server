package types

// PipelineRunSimple 流水线运行实例简要信息
type PipelineRunSimple struct {
	CreatorAccountId string `json:"creatorAccountId,omitempty"`
	EndTime          int64  `json:"endTime,omitempty"`
	PipelineId       int64  `json:"pipelineId,omitempty"`
	PipelineRunId    int64  `json:"pipelineRunId,omitempty"`
	StartTime        int64  `json:"startTime,omitempty"`
	TriggerMode      int    `json:"triggerMode,omitempty"`
} 