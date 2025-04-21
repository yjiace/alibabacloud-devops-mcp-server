package types

// Branch 分支信息结构体
type Branch struct {
	Commit        Commit `json:"commit"`
	DefaultBranch bool   `json:"defaultBranch"`
	Name          string `json:"name"`
	Protected     bool   `json:"protected"`
	WebUrl        string `json:"webUrl"`
}
