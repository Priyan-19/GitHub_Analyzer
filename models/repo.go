package models

import "time"

type Repo struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	FullName    string     `json:"full_name"`
	Description string     `json:"description"`
	Stars       int        `json:"stargazers_count"`
	Forks       int        `json:"forks_count"`
	OpenIssues  int        `json:"open_issues_count"`
	Language    string     `json:"language"`
	UpdatedAt   time.Time  `json:"updated_at"`
	HTMLURL     string     `json:"html_url"`
	Fork        bool       `json:"fork"`
	Archived    bool       `json:"archived"`
	Topics      []string   `json:"topics"`
	Size        int        `json:"size"`
	Watchers    int        `json:"watchers_count"`
}
