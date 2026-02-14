# Go HTMX Server - Design

## Project Overview
- **Type**: Go HTTP server with HTMX
- **Purpose**: Serve portfolio with dynamic project details
- **Architecture**: Go net/http with html/template, HTMX for interactivity

## Functionality

### Routes
| Path | Description |
|------|-------------|
| `/` | Main page with hero + project cards |
| `/projects/{id}` | HTMX fragment with project details |

### Data Model
```go
type Project struct {
    ID          string
    Title       string
    ShortDesc   string
    FullDesc    string
    Tech        string
    Features    []string
}
```

### Projects
1. **Kubernetes** - Container orchestration, K8s clusters, Helm charts
2. **CI/CD** - Pipeline automation, GitHub Actions, Jenkins
3. **Go** - Custom tooling, CLI applications
4. **Observability/SRE** - Monitoring, logging, SLIs/SLOs

### UI Flow
1. User loads `/` - sees hero + 4 project cards (collapsed)
2. Click project card → HTMX GET `/projects/{id}`
3. Server returns HTML fragment
4. Fragment replaces card content with expanded details

## Styling
- Match existing brutalist CSS from index.html
- JetBrains Mono + IBM Sans fonts
- Dark theme (#1a1a1a, #252525, #444, #e5e5e5)

## Acceptance Criteria
- [ ] Server runs on localhost:8080
- [ ] Main page serves at `/`
- [ ] 4 project cards displayed
- [ ] Click card loads details via HTMX
- [ ] Styling matches brutalist design
