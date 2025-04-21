package types

type Compare struct {
	Commits []Commit `json:"commits"`
	Diffs   []Diff   `json:"diffs"`
}

type Commit struct {
	AuthorEmail    string   `json:"authorEmail"`
	AuthorName     string   `json:"authorName"`
	AuthoredDate   string   `json:"authoredDate"`
	CommittedDate  string   `json:"committedDate"`
	CommitterEmail string   `json:"committerEmail"`
	CommitterName  string   `json:"committerName"`
	ID             string   `json:"id"`
	Message        string   `json:"message"`
	ParentIds      []string `json:"parentIds"`
	ShortId        string   `json:"shortId"`
	Stats          Stats    `json:"stats"`
	Title          string   `json:"title"`
	WebUrl         string   `json:"webUrl"`
}

type Stats struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Total     int `json:"total"`
}

type Diff struct {
	AMode       string `json:"aMode"`
	BMode       string `json:"bMode"`
	DeletedFile bool   `json:"deletedFile"`
	Diff        string `json:"diff"`
	IsBinary    bool   `json:"isBinary"`
	NewFile     bool   `json:"newFile"`
	NewId       string `json:"newId"`
	NewPath     string `json:"newPath"`
	OldId       string `json:"oldId"`
	OldPath     string `json:"oldPath"`
	RenamedFile bool   `json:"renamedFile"`
}
