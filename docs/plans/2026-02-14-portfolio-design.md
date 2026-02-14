# Brutalist DevOps Portfolio - Design

## Project Overview
- **Type**: Personal portfolio website
- **Purpose**: Showcase DevOps engineering work and skills
- **Style**: Brutalist, monochrome
- **Tech**: Plain HTML/CSS/JS with HTMX

## Visual Design

### Color Palette
| Role | Color | Hex |
|------|-------|-----|
| Background | Dark gray | #1a1a1a |
| Card background | Charcoal | #252525 |
| Borders | Medium gray | #444 |
| Primary text | Light gray | #e5e5e5 |
| Headings | White | #ffffff |
| Accent | Gray | #888888 |

### Typography
- **Headings**: JetBrains Mono (Google Fonts)
- **Body**: IBM Sans (Google Fonts)
- **Code/Details**: JetBrains Mono

### Layout
- Full-width sections
- Exposed 1px borders on all elements
- No border-radius (sharp brutalist corners)
- 2rem section gaps
- 1.5rem card padding

## Page Structure

### Hero Section
- Full viewport height
- Large brutalist typography: Your name
- Subtitle: "DevOps Engineer"
- 1-2 sentence intro
- Raw, high contrast

### Projects Section
- 2x2 grid on desktop
- Single column on mobile
- 4 project cards:

1. **Kubernetes**
   - Container orchestration
   - K8s clusters, Helm charts

2. **CI/CD**
   - Pipeline automation
   - GitHub Actions, Jenkins

3. **Go**
   - Custom tooling
   - CLI applications

4. **Observability/SRE**
   - Monitoring, logging
   - SLIs, SLOs

### Card Design
- Dark card background (#252525)
- 1px border (#444)
- Project title in JetBrains Mono
- Brief description in IBM Sans
- Hover: invert colors (light bg, dark text)
- Subtle transition

## Interactions
- HTMX for potential content loading
- Hover states with color inversion
- Raw underline animations on links

## Acceptance Criteria
- [ ] Hero displays name, title, intro
- [ ] 4 project cards in grid layout
- [ ] Brutalist styling (borders, no radius, monochrome)
- [ ] JetBrains Mono and IBM Sans fonts loaded
- [ ] Hover effects work on cards
- [ ] Responsive on mobile
- [ ] Page loads without errors
