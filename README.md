<div align="center">

# 🔍 GitHub Developer Analyzer
### Intelligent Developer Insights & Profiling Platform ⚡📊

[![Go](https://img.shields.io/badge/Go_1.21-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![GitHub API](https://img.shields.io/badge/GitHub_API-181717?style=for-the-badge&logo=github)](https://docs.github.com/en/rest)

**GitHub Developer Analyzer** is a high-performance analytics platform that transforms raw GitHub profile data into actionable insights. It evaluates developer activity, repository impact, and language usage to generate a comprehensive developer persona.

</div>

---

## 📖 Project Overview

This platform provides a deep analysis of any GitHub user's public profile, offering structured insights into their coding behavior, expertise, and activity trends.

### Core Value Proposition
- **📊 Data-Driven Insights**: Transform raw GitHub API payloads into beautiful, clear analytics dashboards.
- **🧠 Developer Persona Engine**: Automatically classify developer behaviors and achievements.
- **⚡ High Performance**: Powered by a Go backend for lightning-fast parsing, computations, and rendering.
- **☁️ Serverless Ready**: Architected for native deployment as a Vercel Serverless Function.
- **📦 Portable Architecture**: Uses Go `embed` to package templates and static files inside a single binary.

---

## 🏗️ System Architecture

The application follows a **modular Go backend + embedded UI architecture**:

### 🐹 Backend: Go Application
- **HTTP Server**: High-performance routing using Go's standard library `http.NewServeMux`.
- **Modular Services**:
  - `github_service.go`: Fetches raw profile details, repository meta, and commits from the GitHub REST API.
  - `analytics_service.go`: Computes language percentage weights, repository scores, and spotlight tags.
- **Template Engine**: Compiles and parses Go `html/template` blocks into fully structured server-rendered pages.

### 🌐 Deployment Layer: Serverless (Vercel)
- **Function Entrypoint**: Vercel routes traffic through the entry function in `api/index.go`.
- **Serverless Integration**: Leverages Go `embed.FS` to dynamically bind and serve all UI templates and styles directly in serverless runtimes.

### 🎨 Presentation Layer
- **HTML Templates**: Fully responsive, clean layout engine designed to handle page swaps using modern HTMX.
- **Embedded Assets**: Serving custom CSS stylesheets and icons from the `static/` directory with edge-level routing.

---

## 🚀 Key Features

### 👤 Profile Intelligence
- Automatically aggregates account metrics:
  - Total followers & following
  - Count of public repositories
  - Personal profile avatar and bio details

### 📂 Repository Analysis
- Detailed repository scans measuring:
  - Star accumulations ⭐
  - Fork splits 🍴
  - Recency and date of last update

### 🧠 Developer Persona Engine
- Assigns personalized badges based on activity heuristics:
  - **High Impact Developer**: Substantial repository star count.
  - **GitHub Veteran**: Accounts created over 5 years ago.
  - **Consistent Contributor**: Solid activity across multiple public repositories.

### 💻 Language Breakdown
- Tallies exact programming language usage across all public repositories.
- Renders top-used languages using gorgeous, dynamic percentage bars.

### 📈 Activity Scoring System
- Custom ranking algorithm evaluating repository contributions based on:
  - Popularity metrics (Stars + Forks).
  - Recency weighting (recent updates have higher priority).
  - Score classification: Elite (green badge), Mid-tier (yellow badge), or Quiet (gray badge).

### 📱 Responsive Mobile UI
- Single-page non-scrolling desktop experience, optimized with a dynamic fluid layout for mobile viewports (`<= 768px`).
- Features a stacked 2x3 metrics grid, centered profiles, and touch-scrollable tables.

### 📦 Embedded Asset System
- All HTML structures and static styles are embedded during Go build time, removing filesystem dependencies and ensuring single-binary portability.

---

## 📂 Project Structure

```text
github-developer-analyzer/
├── api/
│   └── index.go          # Vercel Serverless Function adapter entry
├── cmd/
│   └── server/
│       └── main.go       # Local HTTP server bootstrap entry
├── handlers/             # HTTP controllers and template engines
├── models/               # Domain objects and GitHub API schemas
├── services/             # Core business engines (GitHub fetcher, scorer)
├── static/               # Shared stylesheets and favicon assets
├── templates/            # Go html/template structures (index, dashboard, error)
├── .gitignore            # Git exclusion rules
├── app.go                # Central package initializer & go:embed router
├── go.mod                # Dependency module definition
└── vercel.json           # Vercel cloud function configuration
```

---

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Internet connection (to fetch live data from the GitHub REST API)

---

### 1. Install Dependencies
```bash
go mod tidy
```

---

### 2. Run Locally
```bash
go run cmd/server/main.go
```

Once running, open your web browser and navigate to:
```
http://localhost:8080
```

---

### 3. Configure Port (Optional)

```bash
# Windows PowerShell
$env:PORT="9000"; go run cmd/server/main.go

# Windows CMD
set PORT=9000 && go run cmd/server/main.go

# Linux/macOS
PORT=9000 go run cmd/server/main.go
```

---

## ☁️ Deployment (Vercel)

This project is fully configured for zero-configuration deployments:

1. Push your repository to **GitHub**.
2. Log in to your **Vercel Dashboard**.
3. Import the `GitHub_Analyzer` repository.
4. Click **Deploy**.

Vercel will build the serverless package using Go, compile assets via `go:embed`, and host the application on global edge nodes.

---

## 🔒 API & Rate Limiting

The application leverages unauthenticated requests to public endpoints:

- Limit: **60 requests/hour**

### ⚠️ Recommendation
To prevent rate limitations under high usage, implement authorization headers using a Personal Access Token (PAT) within:
```
services/github_service.go
```

---

## 📦 Architecture Highlights

- Fully serverless-compatible design out of the box.
- Embedded resources for single-command builds and zero file-path dependencies.
- Modern, HTMX-backed client requests replacing complex JavaScript framework code.

---

## 📊 Use Cases

- Quick visual benchmarking of developer profiles.
- Recruitment assessment of technical contribution distributions.
- Developer portfolio visualization.

---

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to enhance analysis features, persona calculations, or styling rules.

---

<div align="center">
  <p>Built with 🔍 for Smarter Developer Insights</p>
  <p>Developed by <strong>Priyan</strong></p>
  <p>© 2026 GitHub Developer Analyzer. All Rights Reserved.</p>
</div>
