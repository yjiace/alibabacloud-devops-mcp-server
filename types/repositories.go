package types

type Repository struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	WebUrl      string `json:"WebUrl"`
	Description string `json:"description"`
	Path        string `json:"path"`
}
