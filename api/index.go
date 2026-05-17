package api

import (
	"net/http"

	analyzer "github-developer-analyzer"
)

// Handler is the Vercel serverless function entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	analyzer.App().ServeHTTP(w, r)
}
