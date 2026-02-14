# Go HTMX Server - Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Go HTTP server that serves portfolio with HTMX dynamic project details

**Architecture:** Go net/http with html/template, serving main page and HTMX fragments

**Tech Stack:** Go 1.21+, standard library net/http, html/template

---

### Task 1: Create project data and handlers

**Files:**
- Create: `main.go`

**Step 1: Write main.go**

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

type Project struct {
	ID        string
	Title     string
	ShortDesc string
	FullDesc  string
	Tech      string
	Features  []string
}

var projects = []Project{
	{
		ID:        "kubernetes",
		Title:     "Kubernetes",
		ShortDesc: "Container orchestration at scale. K8s clusters, Helm charts, GitOps workflows.",
		FullDesc:  "Production-grade Kubernetes infrastructure with automated deployments, GitOps workflows using ArgoCD, and comprehensive Helm chart library for reusable deployments.",
		Tech:      "KUBERNETES",
		Features:  []string{"EKS/GKE clusters", "Helm charts", "ArgoCD GitOps", "Terraform IaC"},
	},
	{
		ID:        "cicd",
		Title:     "CI/CD",
		ShortDesc: "Pipeline automation for rapid, reliable deployments. GitHub Actions, Jenkins.",
		FullDesc:  "End-to-end CI/CD pipelines with automated testing, security scanning, and multi-environment deployments with rollback capabilities.",
		Tech:      "GITHUB ACTIONS",
		Features:  []string{"GitHub Actions", "Jenkins pipelines", "Security scanning", "Blue-green deploys"},
	},
	{
		ID:        "go",
		Title:     "Go",
		ShortDesc: "Custom tooling and CLI applications built for DevOps workflows.",
		FullDesc:  "Custom Go applications including CLI tools for infrastructure management, automation scripts, and internal developer platforms.",
		Tech:      "GOLANG",
		Features:  []string{"CLI tools", "Infrastructure APIs", "Automation", "Custom operators"},
	},
	{
		ID:        "observability",
		Title:     "Observability/SRE",
		ShortDesc: "Monitoring, logging, and reliability engineering. SLIs, SLOs, and beyond.",
		FullDesc:  "Comprehensive observability stack with Prometheus metrics, Loki logging, Grafana dashboards, and SLO-based alerting for production systems.",
		Tech:      "PROMETHEUS",
		Features:  []string{"Prometheus/Grafana", "Loki logging", "SLO tracking", "Incident response"},
	},
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/projects/", handleProject)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", projects)
}

func handleProject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/projects/"):]
	
	for _, p := range projects {
		if p.ID == id {
			tmpl.ExecuteTemplate(w, "project-detail.html", p)
			return
		}
	}
	http.NotFound(w, r)
}
```

**Step 2: Verify Go is available**

Run: `go version`

---

### Task 2: Create templates directory and index template

**Files:**
- Create: `templates/index.html`

**Step 1: Create templates directory**

Run: `mkdir -p templates`

**Step 2: Write index.html template**

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>DevOps Portfolio</title>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=IBM+Sans:wght@400;500;700&family=JetBrains+Mono:wght@400;700&display=swap" rel="stylesheet">
    <script src="https://unpkg.com/htmx.org@2.0.0"></script>
    <style>
        :root {
            --bg-dark: #1a1a1a;
            --bg-card: #252525;
            --border: #444;
            --text-primary: #e5e5e5;
            --text-heading: #ffffff;
            --text-accent: #888;
        }

        * { margin: 0; padding: 0; box-sizing: border-box; }

        body {
            font-family: 'IBM Sans', sans-serif;
            background-color: var(--bg-dark);
            color: var(--text-primary);
            line-height: 1.6;
        }

        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 0 2rem;
        }

        .hero {
            min-height: 60vh;
            display: flex;
            flex-direction: column;
            justify-content: center;
            border-bottom: 1px solid var(--border);
            padding: 4rem 0;
        }

        .hero h1 {
            font-family: 'JetBrains Mono', monospace;
            font-size: clamp(3rem, 8vw, 6rem);
            font-weight: 700;
            color: var(--text-heading);
            letter-spacing: -0.02em;
            line-height: 1.1;
        }

        .hero .subtitle {
            font-family: 'JetBrains Mono', monospace;
            font-size: 1.25rem;
            color: var(--text-accent);
            margin-top: 1rem;
        }

        .hero .intro {
            font-size: 1.125rem;
            max-width: 600px;
            margin-top: 2rem;
            color: var(--text-primary);
        }

        .projects {
            padding: 4rem 0;
        }

        .projects h2 {
            font-family: 'JetBrains Mono', monospace;
            font-size: 1.5rem;
            color: var(--text-heading);
            margin-bottom: 2rem;
            padding-bottom: 1rem;
            border-bottom: 1px solid var(--border);
        }

        .projects-grid {
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 1.5rem;
        }

        @media (max-width: 768px) {
            .projects-grid { grid-template-columns: 1fr; }
        }

        .project-card {
            background: var(--bg-card);
            border: 1px solid var(--border);
            padding: 1.5rem;
            transition: all 0.2s ease;
            cursor: pointer;
        }

        .project-card:hover {
            background: var(--text-primary);
            color: var(--bg-dark);
        }

        .project-card:hover .project-title,
        .project-card:hover .project-desc,
        .project-card:hover .project-tech {
            color: var(--bg-dark);
        }

        .project-card:hover .project-tech {
            border-color: var(--bg-dark);
        }

        .project-title {
            font-family: 'JetBrains Mono', monospace;
            font-size: 1.25rem;
            font-weight: 700;
            color: var(--text-heading);
            margin-bottom: 0.75rem;
        }

        .project-desc {
            font-size: 0.95rem;
            color: var(--text-primary);
            margin-bottom: 1rem;
        }

        .project-tech {
            font-family: 'JetBrains Mono', monospace;
            font-size: 0.75rem;
            color: var(--text-accent);
            border: 1px solid var(--border);
            padding: 0.25rem 0.5rem;
            display: inline-block;
        }

        .project-features {
            margin-top: 1rem;
            display: flex;
            flex-wrap: wrap;
            gap: 0.5rem;
        }

        .project-features span {
            font-family: 'JetBrains Mono', monospace;
            font-size: 0.7rem;
            color: var(--text-accent);
            border: 1px solid var(--border);
            padding: 0.2rem 0.4rem;
        }

        .project-card:hover .project-features span {
            color: var(--bg-dark);
            border-color: var(--bg-dark);
        }

        footer {
            border-top: 1px solid var(--border);
            padding: 2rem 0;
            text-align: center;
        }

        footer a {
            font-family: 'JetBrains Mono', monospace;
            color: var(--text-accent);
            text-decoration: none;
            font-size: 0.875rem;
        }

        footer a:hover {
            color: var(--text-primary);
            text-decoration: underline;
        }
    </style>
</head>
<body hx-boost="true">
    <main class="container">
        <section class="hero">
            <h1>LUCAS COTTLE</h1>
            <p class="subtitle">DevOps Engineer</p>
            <p class="intro">Building reliable infrastructure and automating everything in sight.</p>
        </section>

        <section class="projects">
            <h2>PROJECTS</h2>
            <div class="projects-grid">
                {{range .}}
                <article 
                    class="project-card"
                    hx-get="/projects/{{.ID}}"
                    hx-swap="innerHTML"
                    hx-target="this"
                >
                    <h3 class="project-title">{{.Title}}</h3>
                    <p class="project-desc">{{.ShortDesc}}</p>
                    <span class="project-tech">{{.Tech}}</span>
                </article>
                {{end}}
            </div>
        </section>
    </main>

    <footer>
        <div class="container">
            <a href="mailto:lucas.t.cottle@gmail.com">lucas.t.cottle@gmail.com</a>
        </div>
    </footer>
</body>
</html>
```

---

### Task 3: Create project detail template

**Files:**
- Create: `templates/project-detail.html`

**Step 1: Write project-detail.html**

```html
<h3 class="project-title">{{.Title}}</h3>
<p class="project-desc">{{.FullDesc}}</p>
<span class="project-tech">{{.Tech}}</span>
<div class="project-features">
    {{range .Features}}
    <span>{{.}}</span>
    {{end}}
</div>
```

---

### Task 4: Test the server

**Step 1: Run the server**

Run: `go run main.go`

**Step 2: Verify main page**

Open: http://localhost:8080

**Step 3: Test HTMX**

Click a project card - should expand with full description and features.

---

### Task 5: Commit

Run:
```bash
git add main.go templates/
git commit -m "feat: add Go HTMX server for dynamic portfolio"
```
