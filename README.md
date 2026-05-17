# GitHub Developer Analyzer

A Go-based web application that analyzes a GitHub user's profile and repositories to provide insights into their activity, most used languages, and developer "persona".

## 🚀 Features

- **Profile Overview**: Fetches user metadata (followers, following, public repos, etc.).
- **Repository Analysis**: Analyzes all public repositories for stars, forks, and activity.
- **Language Breakdown**: Calculates the top 5 programming languages used by the user.
- **Activity Scoring**: Custom algorithm to rank repositories by recency and popularity.
- **Developer Persona**: Assigns "Insight Tags" based on developer stats (e.g., "High Impact Developer", "GitHub Veteran").
- **Responsive UI**: Clean, modern dashboard built with HTML templates and CSS.

## 🛠️ Project Structure

```text
github-developer-analyzer/
├── handlers/          # HTTP request handlers
├── models/            # Data structures and domain models
├── services/          # Business logic (GitHub API, Analytics)
├── static/            # Static assets (CSS, JS, Images)
├── templates/         # HTML templates
├── .gitignore         # Standard Go git ignore rules
├── go.mod             # Go module definition
└── main.go            # Application entry point
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

## 🏃 Running the Application

To start the server, run:
```bash
go run main.go
```
The application will be available at `http://localhost:8080`.

## ⚙️ Configuration

By default, the server runs on port `8080`. You can change this by setting the `PORT` environment variable:
```bash
# Windows (PowerShell)
$env:PORT="9000"; go run main.go

# Windows (Command Prompt)
set PORT=9000 && go run main.go

# Linux/macOS
PORT=9000 go run main.go
```

## 🛑 Rate Limiting Note

This application currently uses unauthenticated requests to the GitHub API. 
- **Unauthenticated requests**: 60 requests per hour.
- If you encounter a "rate limit exceeded" error, please wait or consider adding authentication support in `services/github_service.go`.

## 🤝 Contributing

Feel free to fork the project and submit pull requests!
