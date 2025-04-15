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
	LocalId      int64  `json:"localId"`
}

type ChangeRequestComment struct {
	ProjectId          string `json:"project_id"`
	CommentBizId       string `json:"comment_biz_id"`
	CommentType        string `json:"comment_type"`
	Content            string `json:"content"`
	FilePath           string `json:"filePath"`
	LineNumber         int64  `json:"line_number"`
	ParentCommentBizId string `json:"parent_comment_biz_id"`
}
