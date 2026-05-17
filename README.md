# GitHub Developer Analyzer

A Go-based web application that analyzes a GitHub user's profile and repositories to provide insights into their activity, most used languages, and developer "persona".

This project is fully structured for modern production hosting on platforms like **Vercel** as a serverless function, while remaining easy to run and test locally.

## 🚀 Features

- **Profile Overview**: Fetches user metadata (followers, following, public repos, etc.).
- **Repository Analysis**: Analyzes all public repositories for stars, forks, and activity.
- **Language Breakdown**: Calculates the top 5 programming languages used by the user.
- **Activity Scoring**: Custom algorithm to rank repositories by recency and popularity.
- **Developer Persona**: Assigns "Insight Tags" based on developer stats (e.g., "High Impact Developer", "GitHub Veteran").
- **Responsive UI**: Clean, modern dashboard built with HTML templates and CSS with optimized, precise spacings.
- **Asset Embedding**: Uses Go `embed` to package static assets and templates into a single distributable binary, enabling seamless serverless hosting.

## 🛠️ Project Structure

```text
github-developer-analyzer/
├── api/
│   └── index.go          # Vercel serverless function entrypoint
├── cmd/
│   └── server/
│       └── main.go       # Local execution entrypoint (main package)
├── handlers/              # HTTP request handlers
├── models/                # Data structures and domain models
├── services/              # Business logic (GitHub API, Analytics)
├── static/                # Static assets (CSS, JS, Images)
├── templates/             # HTML templates
├── .gitignore             # Standard Go git ignore rules
├── app.go                 # Core app package (analyzer) using go:embed
├── go.mod                 # Go module definition
└── vercel.json            # Vercel deployment routing configuration
```

## 📋 Prerequisites

- **Go 1.21** or higher installed.
- Internet connection (to fetch data from GitHub API).

## 📥 Setup & Installation

1. **Clone or Download** the project.
2. **Download Dependencies**:
   Open your terminal in the project root and run:
   ```bash
   go mod tidy
   ```

## 🏃 Running Locally

To start the local development server, run:
```bash
go run cmd/server/main.go
```
The application will be available at `http://localhost:8080`.

To change the default port, set the `PORT` environment variable:
```bash
# Windows (PowerShell)
$env:PORT="9000"; go run cmd/server/main.go

# Windows (Command Prompt)
set PORT=9000 && go run cmd/server/main.go

# Linux/macOS
PORT=9000 go run cmd/server/main.go
```

## ⚡ Vercel Deployment

This project is pre-configured for zero-config serverless deployment to **Vercel**:

1. Push this repository to **GitHub**.
2. Go to your [Vercel Dashboard](https://vercel.com).
3. Click **Add New...** and select **Project**.
4. Import your repository.
5. Click **Deploy**. Vercel will automatically read the `vercel.json` configuration, build the serverless Go function in `api/index.go`, package your static assets via `go:embed`, and host the app globally.

## 🛑 Rate Limiting Note

This application currently uses unauthenticated requests to the GitHub API. 
- **Unauthenticated requests**: 60 requests per hour.
- If you encounter a "rate limit exceeded" error, please consider adding authentication support in `services/github_service.go`.

## 🤝 Contributing

Feel free to fork the project and submit pull requests!
