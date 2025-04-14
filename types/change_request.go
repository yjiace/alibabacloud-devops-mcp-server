package types

type ChangeRequest struct {
	Ahead        int64  `json:"ahead"`
	Behind       int64  `json:"behind"`
	CreateTime   string `json:"createTime"`
	Description  string `json:"description"`
	ProjectId    int64  `json:"projectId"`
	SourceBranch string `json:"sourceBranch"`
	TargetBranch string `json:"targetBranch"`
	Title        string `json:"title"`
	UpdateTime   string `json:"updateTime"`
}
