package models

import "time"

type GitHubUser struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	Blog      string `json:"blog"`
	HTMLURL   string `json:"html_url"`
	PublicRepos int  `json:"public_repos"`
	Followers   int  `json:"followers"`
	Following   int  `json:"following"`
	CreatedAt   time.Time `json:"created_at"`
}

type LanguageStat struct {
	Name  string
	Count int
	Pct   float64
}

type RepoSummary struct {
	Name        string
	HTMLURL     string
	Stars       int
	Forks       int
	Language    string
	OpenIssues  int
	UpdatedAt   time.Time
	Description string
	ActivityScore float64
}

type UserAnalysis struct {
	User              GitHubUser
	TotalRepos        int
	OriginalRepos     int
	TotalStars        int
	TotalForks        int
	TotalOpenIssues   int
	TopLanguages      []LanguageStat
	MostPopularRepo   RepoSummary
	MostActiveRepo    RepoSummary
	AvgActivityScore  float64
	Repos             []RepoSummary
	InsightTags       []string
	AccountAgeDays    int
	ForkedRepos       int
	ArchivedRepos     int
}
