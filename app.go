package analyzer

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github-developer-analyzer/handlers"
	"github-developer-analyzer/services"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed static/*
var staticFS embed.FS

// App exposes the main handler for Vercel or local server
func App() http.Handler {
	tmpl, err := template.New("").Funcs(templateFuncs()).ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}

	githubSvc := services.NewGitHubService()
	analyticsSvc := services.NewAnalyticsService()
	analyzeHandler := handlers.NewAnalyzeHandler(githubSvc, analyticsSvc, tmpl)

	mux := http.NewServeMux()
	
	// Serve static files
	staticSubFS, _ := fs.Sub(staticFS, "static")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticSubFS))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			log.Printf("index template error: %v", err)
			http.Error(w, "Internal server error", 500)
		}
	})

	mux.Handle("/analyze-user", analyzeHandler)

	return mux
}

func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("Jan 2, 2006")
		},
		"truncate": func(s string, n int) string {
			if len(s) <= n {
				return s
			}
			return s[:n] + "…"
		},
		"pctWidth": func(pct float64) string {
			if pct > 100 {
				pct = 100
			}
			return fmt.Sprintf("%.1f%%", pct)
		},
		"scoreColor": func(score float64) string {
			switch {
			case score >= 80:
				return "#22c55e"
			case score >= 50:
				return "#f59e0b"
			default:
				return "#94a3b8"
			}
		},
		"langColor": langColor,
		"add":       func(a, b int) int { return a + b },
		"gt0":       func(n int) bool { return n > 0 },
		"floatFmt": func(f float64) string {
			return fmt.Sprintf("%.1f", f)
		},
	}
}

func langColor(lang string) string {
	colors := map[string]string{
		"Go":         "#00ADD8",
		"Python":     "#3776AB",
		"JavaScript": "#F7DF1E",
		"TypeScript": "#3178C6",
		"Rust":       "#CE422B",
		"Java":       "#ED8B00",
		"C":          "#A8B9CC",
		"C++":        "#F34B7D",
		"C#":         "#178600",
		"Ruby":       "#CC342D",
		"PHP":        "#777BB4",
		"Swift":      "#FA7343",
		"Kotlin":     "#7F52FF",
		"Scala":      "#DC322F",
		"Shell":      "#89E051",
		"HTML":       "#E34C26",
		"CSS":        "#563D7C",
		"Vue":        "#4FC08D",
		"Svelte":     "#FF3E00",
		"Dart":       "#00B4AB",
		"R":          "#276DC3",
		"Lua":        "#000080",
		"Haskell":    "#5D4F85",
		"Elixir":     "#6E4A7E",
		"Clojure":    "#5881D8",
		"Zig":        "#EC915C",
	}
	if c, ok := colors[lang]; ok {
		return c
	}
	return "#64748b"
}
