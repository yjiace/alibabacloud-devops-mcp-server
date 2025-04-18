package types

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
	ProjectId                    int64                  `json:"project_id"`
	CommentBizId                 string                 `json:"comment_biz_id"`
	CommentType                  string                 `json:"comment_type"`
	Content                      string                 `json:"content"`
	FilePath                     string                 `json:"filePath"`
	LineNumber                   int64                  `json:"line_number"`
	ParentCommentBizId           string                 `json:"parent_comment_biz_id"`
	RootCommentBizId             string                 `json:"root_comment_biz_id"`
	MrBizId                      string                 `json:"mr_biz_id"`
	Author                       CodeupUser             `json:"author"`
	CommentTime                  string                 `json:"comment_time"`
	LastEditTime                 string                 `json:"last_edit_time"`
	LastEditUser                 CodeupUser             `json:"last_edit_user"`
	Resolved                     bool                   `json:"resolved"`
	State                        string                 `json:"state"`
	IsDeleted                    bool                   `json:"is_deleted"`
	OutDated                     bool                   `json:"out_dated"`
	ChildCommentsList            []ChangeRequestComment `json:"child_comments_list"`
	ExpressionReplyList          []interface{}          `json:"expression_reply_list"`
	FromPatchsetBizId            string                 `json:"from_patchset_biz_id"`
	ToPatchsetBizId              string                 `json:"to_patchset_biz_id"`
	LastResolvedStatusChangeTime string                 `json:"last_resolved_status_change_time"`
	LastResolvedStatusChangeUser CodeupUser             `json:"last_resolved_status_change_user"`
	Location                     CommentLocation        `json:"location"`
	RelatedPatchset              PatchSet               `json:"related_patchset"`
}

type CodeupUser struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	State    string `json:"state"`
}

type CommentLocation struct {
	CanLocated           bool   `json:"can_located"`
	LocatedFilePath      string `json:"located_file_path"`
	LocatedLineNumber    int64  `json:"located_line_number"`
	LocatedPatchSetBizId string `json:"located_patch_set_biz_id"`
}

type PatchSet struct {
	CommitId             string `json:"commitId"`
	CreateTime           string `json:"createTime"`
	PatchSetBizId        string `json:"patchSetBizId"`
	PatchSetName         string `json:"patchSetName"`
	Ref                  string `json:"ref"`
	RelatedMergeItemType string `json:"relatedMergeItemType"`
	ShortId              string `json:"shortId"`
	VersionNo            int    `json:"versionNo"`
}
