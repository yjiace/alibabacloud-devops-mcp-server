package types

type WorkItem struct {
	Id           int64        `json:"id"`
	Subject      string       `json:"subject"`
	Description  string       `json:"description"`
	GmtCreate    string       `json:"gmtCreate"`
	GmtModified  string       `json:"gmtModified"`
	WorkItemType WorkItemType `json:"workItemType"`
	Status       Status       `json:"status"`
	FormatType   string       `json:"formatType"`
}

type WorkItemType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	DisplayName string `json:"displayName"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	NameEn      string `json:"nameEn"`
}
