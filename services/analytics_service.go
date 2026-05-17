package services

import (
	"math"
	"sort"
	"strings"
	"time"

	"github-developer-analyzer/models"
)

type AnalyticsService struct{}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{}
}

func (a *AnalyticsService) Analyze(user models.GitHubUser, repos []models.Repo) models.UserAnalysis {
	analysis := models.UserAnalysis{User: user}

	if len(repos) == 0 {
		return analysis
	}

	langCount := make(map[string]int)
	var repoSummaries []models.RepoSummary
	now := time.Now()

	for _, r := range repos {
		if r.Fork {
			analysis.ForkedRepos++
		}
		if r.Archived {
			analysis.ArchivedRepos++
		}
		if !r.Fork {
			analysis.OriginalRepos++
		}

		analysis.TotalStars += r.Stars
		analysis.TotalForks += r.Forks
		analysis.TotalOpenIssues += r.OpenIssues

		if r.Language != "" {
			langCount[r.Language]++
		}

		score := computeActivityScore(r, now)
		repoSummaries = append(repoSummaries, models.RepoSummary{
			Name:          r.Name,
			HTMLURL:       r.HTMLURL,
			Stars:         r.Stars,
			Forks:         r.Forks,
			Language:      r.Language,
			OpenIssues:    r.OpenIssues,
			UpdatedAt:     r.UpdatedAt,
			Description:   r.Description,
			ActivityScore: score,
		})
	}

	analysis.TotalRepos = len(repos)

	// Sort repos by activity score descending for table
	sort.Slice(repoSummaries, func(i, j int) bool {
		return repoSummaries[i].ActivityScore > repoSummaries[j].ActivityScore
	})
	analysis.Repos = repoSummaries

	// Most popular by stars
	sort.Slice(repoSummaries, func(i, j int) bool {
		return repoSummaries[i].Stars > repoSummaries[j].Stars
	})
	if len(repoSummaries) > 0 {
		analysis.MostPopularRepo = repoSummaries[0]
	}

	// Most active by updated_at
	mostActive := repoSummaries[0]
	for _, rs := range repoSummaries {
		if rs.UpdatedAt.After(mostActive.UpdatedAt) {
			mostActive = rs
		}
	}
	analysis.MostActiveRepo = mostActive

	// Average activity score
	var totalScore float64
	for _, rs := range repoSummaries {
		totalScore += rs.ActivityScore
	}
	analysis.AvgActivityScore = totalScore / float64(len(repoSummaries))

	// Top languages
	type ls struct {
		name  string
		count int
	}
	var langSlice []ls
	for k, v := range langCount {
		langSlice = append(langSlice, ls{k, v})
	}
	sort.Slice(langSlice, func(i, j int) bool {
		return langSlice[i].count > langSlice[j].count
	})
	total := 0
	for _, l := range langSlice {
		total += l.count
	}
	for i, l := range langSlice {
		if i >= 5 {
			break
		}
		pct := 0.0
		if total > 0 {
			pct = float64(l.count) / float64(total) * 100
		}
		analysis.TopLanguages = append(analysis.TopLanguages, models.LanguageStat{
			Name:  l.name,
			Count: l.count,
			Pct:   math.Round(pct*10) / 10,
		})
	}

	// Account age
	analysis.AccountAgeDays = int(time.Since(user.CreatedAt).Hours() / 24)

	// Insight tags
	analysis.InsightTags = generateInsights(analysis, langSlice[0].name)

	// Re-sort repos by activity score for display
	sort.Slice(analysis.Repos, func(i, j int) bool {
		return analysis.Repos[i].ActivityScore > analysis.Repos[j].ActivityScore
	})

	return analysis
}

func computeActivityScore(r models.Repo, now time.Time) float64 {
	daysSinceUpdate := now.Sub(r.UpdatedAt).Hours() / 24
	recencyScore := math.Max(0, 100-daysSinceUpdate*0.5)
	starScore := math.Log1p(float64(r.Stars)) * 10
	forkScore := math.Log1p(float64(r.Forks)) * 5
	return math.Round((recencyScore+starScore+forkScore)*10) / 10
}

func generateInsights(a models.UserAnalysis, topLang string) []string {
	var tags []string

	if a.TotalStars > 1000 {
		tags = append(tags, "⭐ High Impact Developer")
	} else if a.TotalStars > 100 {
		tags = append(tags, "🌟 Rising Star")
	}

	if a.AccountAgeDays > 365*5 {
		tags = append(tags, "🏛️ GitHub Veteran")
	} else if a.AccountAgeDays > 365*2 {
		tags = append(tags, "📅 Experienced Developer")
	}

	if a.OriginalRepos > 50 {
		tags = append(tags, "🚀 Prolific Creator")
	} else if a.OriginalRepos > 20 {
		tags = append(tags, "💡 Active Builder")
	}

	if topLang != "" {
		tags = append(tags, "🔧 "+topLang+" Specialist")
	}

	if a.User.Followers > 500 {
		tags = append(tags, "👥 Community Leader")
	}

	forkRatio := 0.0
	if a.TotalRepos > 0 {
		forkRatio = float64(a.ForkedRepos) / float64(a.TotalRepos)
	}
	if forkRatio > 0.5 {
		tags = append(tags, "🤝 Open Source Contributor")
	} else if forkRatio < 0.2 && a.TotalRepos > 5 {
		tags = append(tags, "🎨 Original Creator")
	}

	if len(tags) == 0 {
		tags = append(tags, "💻 GitHub Developer")
	}

	// Clamp to relevant set
	if len(tags) > 5 {
		tags = tags[:5]
	}
	_ = strings.Join(tags, "") // suppress unused import warning
	return tags
}
