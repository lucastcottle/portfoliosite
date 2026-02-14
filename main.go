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
