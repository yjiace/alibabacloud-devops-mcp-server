package types

import "time"

type ChangeRequest struct {
	Ahead                  int64  `json:"ahead"`
	Behind                 int64  `json:"behind"`
	CreateTime             string `json:"createTime"`
	Description            string `json:"description"`
	ProjectId              int64  `json:"projectId"`
	SourceBranch           string `json:"sourceBranch"`
	TargetBranch           string `json:"targetBranch"`
	Title                  string `json:"title"`
	UpdateTime             string `json:"updateTime"`
	LocalId                int64  `json:"localId"`
	AllRequirementsPassed  bool   `json:"allRequirementsPassed"`
	CreateFrom             string `json:"createFrom"`
	DetailUrl              string `json:"detailUrl"`
	MergedRevision         string `json:"mergedRevision"`
	Status                 string `json:"status"`
	TotalCommentCount      int64  `json:"totalCommentCount"`
	UnResolvedCommentCount int64  `json:"unResolvedCommentCount"`
	WebUrl                 string `json:"WebUrl"`
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

type PatchSet struct {
	CommitId             string    `json:"commitId"`
	CreateTime           time.Time `json:"createTime"`
	PatchSetBizId        string    `json:"patchSetBizId"`
	PatchSetName         string    `json:"patchSetName"`
	Ref                  string    `json:"ref"`
	RelatedMergeItemType string    `json:"relatedMergeItemType"`
	ShortId              string    `json:"shortId"`
	VersionNo            int       `json:"versionNo"`
}
