package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github-developer-analyzer/services"
)

type AnalyzeHandler struct {
	githubSvc    *services.GitHubService
	analyticsSvc *services.AnalyticsService
	tmpl         *template.Template
}

func NewAnalyzeHandler(gs *services.GitHubService, as *services.AnalyticsService, tmpl *template.Template) *AnalyzeHandler {
	return &AnalyzeHandler{githubSvc: gs, analyticsSvc: as, tmpl: tmpl}
}

func (h *AnalyzeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimSpace(r.URL.Query().Get("username"))
	if username == "" {
		h.renderError(w, "Please enter a GitHub username.")
		return
	}
	// Sanitize
	for _, c := range username {
		if !isValidUsernameChar(c) {
			h.renderError(w, "Invalid username. GitHub usernames may only contain alphanumeric characters and hyphens.")
			return
		}
	}

	user, err := h.githubSvc.FetchUser(username)
	if err != nil {
		h.renderError(w, friendlyError(err))
		return
	}

	repos, err := h.githubSvc.FetchAllRepos(username)
	if err != nil {
		h.renderError(w, friendlyError(err))
		return
	}

	if len(repos) == 0 {
		h.renderError(w, "This user has no public repositories.")
		return
	}

	analysis := h.analyticsSvc.Analyze(user, repos)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.ExecuteTemplate(w, "analysis.html", analysis); err != nil {
		log.Printf("template error: %v", err)
		http.Error(w, "Internal server error", 500)
	}
}

func (h *AnalyzeHandler) renderError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.ExecuteTemplate(w, "error.html", msg); err != nil {
		http.Error(w, msg, http.StatusBadRequest)
	}
}

func isValidUsernameChar(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') || c == '-' || c == '_'
}

func friendlyError(err error) string {
	msg := err.Error()
	if strings.Contains(msg, "user not found") {
		return "GitHub user not found. Please check the username and try again."
	}
	if strings.Contains(msg, "rate limit") {
		return "GitHub API rate limit reached. Please wait a few minutes and try again."
	}
	if strings.Contains(msg, "network") {
		return "Unable to reach GitHub API. Please check your internet connection."
	}
	return "An unexpected error occurred: " + msg
}
