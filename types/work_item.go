package types

type WorkItem struct {
	Id              string           `json:"id"`
	Subject         string           `json:"subject"`
	Description     string           `json:"description"`
	GmtCreate       string           `json:"gmtCreate"`
	GmtModified     string           `json:"gmtModified"`
	WorkItemType    WorkItemType     `json:"workitemType"`
	Status          Status           `json:"status"`
	FormatType      string           `json:"formatType"`
	CategoryId      string           `json:"categoryId"`
	LogicalStatus   string           `json:"logicalStatus"`
	ParentId        string           `json:"parentId"`
	SerialNumber    string           `json:"serialNumber"`
	StatusStageId   string           `json:"statusStageId"`
	UpdateStatusAt  string           `json:"updateStatusAt"`
	IdPath          string           `json:"idPath"`
	AssignedTo      User             `json:"assignedTo"`
	Creator         User             `json:"creator"`
	Modifier        User             `json:"modifier"`
	Verifier        User             `json:"verifier"`
	Space           Space            `json:"space"`
	Sprint          Sprint           `json:"sprint"`
	Labels          []Label          `json:"labels"`
	Participants    []User           `json:"participants"`
	Trackers        []User           `json:"trackers"`
	Versions        []Version        `json:"versions"`
	CustomFieldValues []CustomFieldValue `json:"customFieldValues"`
}

type WorkItemType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	DisplayName string `json:"displayName"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	NameEn      string `json:"nameEn"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Space struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Sprint struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Label struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Version struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CustomFieldValue struct {
	FieldFormat string   `json:"fieldFormat"`
	FieldId     string   `json:"fieldId"`
	FieldName   string   `json:"fieldName"`
	Values      []Value  `json:"values"`
}

type Value struct {
	DisplayValue string `json:"displayValue"`
	Identifier   string `json:"identifier"`
}
