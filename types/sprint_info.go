package types

// SprintInfo 迭代信息结构
type SprintInfo struct {
	CapacityHours int       `json:"capacityHours,omitempty"`
	Creator       UserInfo   `json:"creator,omitempty"`
	Description   string     `json:"description,omitempty"`
	EndDate       string     `json:"endDate,omitempty"`
	GmtCreate     string     `json:"gmtCreate,omitempty"`
	GmtModified   string     `json:"gmtModified,omitempty"`
	ID            string     `json:"id,omitempty"`
	Locked        bool       `json:"locked,omitempty"`
	Modifier      UserInfo   `json:"modifier,omitempty"`
	Name          string     `json:"name,omitempty"`
	Owners        []UserInfo `json:"owners,omitempty"`
	StartDate     string     `json:"startDate,omitempty"`
	Status        string     `json:"status,omitempty"`
} 