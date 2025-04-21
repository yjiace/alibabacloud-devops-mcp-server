package types

// CreateFileRequest 创建文件的请求
type CreateFileRequest struct {
	Branch        string `json:"branch"`
	CommitMessage string `json:"commitMessage"`
	Content       string `json:"content"`
	Encoding      string `json:"encoding"`
	FilePath      string `json:"filePath"`
}

// CreateFileResponse 创建文件的响应
type CreateFileResponse struct {
	Branch   string `json:"branch"`
	FilePath string `json:"filePath"`
	NewOid   string `json:"newOid"`
}

// DeleteFileResponse 删除文件的响应
type DeleteFileResponse struct {
	Result bool `json:"result"`
}

// FileInfo 文件信息
type FileInfo struct {
	ID    string `json:"id"`
	IsLFS bool   `json:"isLFS"`
	Mode  string `json:"mode"`
	Name  string `json:"name"`
	Path  string `json:"path"`
	Type  string `json:"type"`
}

// FileContent 文件内容
type FileContent struct {
	BlobId       string `json:"blobId"`
	CommitId     string `json:"commitId"`
	Content      string `json:"content"`
	Encoding     string `json:"encoding"`
	FileName     string `json:"fileName"`
	FilePath     string `json:"filePath"`
	LastCommitId string `json:"lastCommitId"`
	Ref          string `json:"ref"`
	Size         int    `json:"size"`
}

// UpdateFileRequest 更新文件内容的请求
type UpdateFileRequest struct {
	Branch        string `json:"branch"`
	CommitMessage string `json:"commitMessage"`
	Content       string `json:"content"`
	Encoding      string `json:"encoding,omitempty"`
} 