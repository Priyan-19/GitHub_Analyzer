package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github-developer-analyzer/models"
)

const githubAPIBase = "https://api.github.com"

type cacheEntry struct {
	repos     []models.Repo
	user      models.GitHubUser
	fetchedAt time.Time
}

type GitHubService struct {
	client    *http.Client
	cache     map[string]cacheEntry
	cacheMu   sync.RWMutex
	cacheTTL  time.Duration
}

func NewGitHubService() *GitHubService {
	return &GitHubService{
		client:   &http.Client{Timeout: 15 * time.Second},
		cache:    make(map[string]cacheEntry),
		cacheTTL: 5 * time.Minute,
	}
}

func (s *GitHubService) doRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "github-developer-analyzer/1.0")
	return s.client.Do(req)
}

func (s *GitHubService) FetchUser(username string) (models.GitHubUser, error) {
	url := fmt.Sprintf("%s/users/%s", githubAPIBase, username)
	resp, err := s.doRequest(url)
	if err != nil {
		return models.GitHubUser{}, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return models.GitHubUser{}, fmt.Errorf("user not found")
	}
	if resp.StatusCode == 403 {
		return models.GitHubUser{}, fmt.Errorf("github api rate limit exceeded, please try again later")
	}
	if resp.StatusCode != 200 {
		return models.GitHubUser{}, fmt.Errorf("github api error: status %d", resp.StatusCode)
	}

	var user models.GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return models.GitHubUser{}, fmt.Errorf("failed to parse user data")
	}
	return user, nil
}

func (s *GitHubService) FetchAllRepos(username string) ([]models.Repo, error) {
	s.cacheMu.RLock()
	if entry, ok := s.cache[username]; ok && time.Since(entry.fetchedAt) < s.cacheTTL {
		s.cacheMu.RUnlock()
		return entry.repos, nil
	}
	s.cacheMu.RUnlock()

	var allRepos []models.Repo
	page := 1
	for {
		url := fmt.Sprintf("%s/users/%s/repos?per_page=100&page=%d&sort=updated", githubAPIBase, username, page)
		resp, err := s.doRequest(url)
		if err != nil {
			return nil, fmt.Errorf("network error fetching repos: %w", err)
		}

		if resp.StatusCode == 404 {
			resp.Body.Close()
			return nil, fmt.Errorf("user not found")
		}
		if resp.StatusCode == 403 {
			resp.Body.Close()
			return nil, fmt.Errorf("github api rate limit exceeded, please try again later")
		}
		if resp.StatusCode != 200 {
			resp.Body.Close()
			return nil, fmt.Errorf("github api error: status %d", resp.StatusCode)
		}

		var repos []models.Repo
		if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("failed to parse repo data")
		}
		resp.Body.Close()

		allRepos = append(allRepos, repos...)
		if len(repos) < 100 {
			break
		}
		page++
		if page > 10 { // safety cap at 1000 repos
			break
		}
	}

	s.cacheMu.Lock()
	s.cache[username] = cacheEntry{repos: allRepos, fetchedAt: time.Now()}
	s.cacheMu.Unlock()

	return allRepos, nil
}
