package githubapi

import "time"

const ApiURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	CreatedAt time.Time `json:"created_at"`
	Body      string
	User      *User
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Flags struct {
	Repo   string
	Status bool
	Count  int
	Asc    bool
}
