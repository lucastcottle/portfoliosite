# Portfolio: Brutalist Enhancements Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Add navigation bar, bold ASCII separators, and timeline section to portfolio while preserving brutalist minimalism

**Architecture:** Pure HTML/CSS additions using existing CSS variables; ASCII text separators; minimalist navigation with smooth scroll and monospace typography

**Tech Stack:** HTML5, CSS3 with CSS custom properties, no new dependencies

**Prerequisites:** Current palette `--bg-dark: #1a1a1a`, `--border: #444`, `--text-accent: #888`, monospace fonts already loaded

---

### Task 1: Add Top Navigation Bar HTML

**Files:**
- Modify: `/var/home/lucas/repos/portfolio2/docs/plans/2026-02-14-portfolio-brutalist-enhancements-design.md`
- Modify: `/var/home/lucas/repos/portfolio2/index.html:165` (before main container)

**Step 1: Add nav HTML structure**

```html
<!-- Add after <body> opening tag, before <main> -->
<nav class="nav-bar">
    <div class="container">
        <a href="#hero" class="nav-link">About</a>
        <span class="nav-separator">|</span>
        <a href="#experience" class="nav-link">Experience</a>
        <span class="nav-separator">|</span>
        <a href="#projects" class="nav-link">Projects</a>
        <span class="nav-separator">|</span>
        <a href="#contact" class="nav-link">Contact</a>
    </div>
</nav>
```

**Step 2: Add section IDs for navigation targets**

```html
<!-- Update these existing sections -->
<section class="hero" id="hero">
<!-- Update projects section title element -->
<h2 id="projects-title">PROJECTS</h2>
```

### Task 2: Add Navigation Bar CSS Styles

**Files:**
- Modify: `/var/home/lucas/repos/portfolio2/index.html:11-162` (style section)

**Step 1: Add nav component styles**

```css
/* Add after :root variables and basic reset */
.nav-bar {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    background-color: var(--bg-dark);
    border-bottom: 1px solid var(--border);
    padding: 1rem 0;
    z-index: 100;
    font-family: 'JetBrains Mono', monospace;
    font-weight: 700;
    font-size: 0.875rem;
}

.nav-bar .container {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 2rem;
}

.nav-link {
    color: var(--text-accent);
    text-decoration: none;
    transition: color 0.2s ease;
}

.nav-link:hover {
    color: var(--text-primary);
    text-decoration: underline;
}

.nav-separator {
    color: var(--border);
}

/* Add body padding to account for fixed nav */
body {
    padding-top: 60px; /* Height of nav bar */
}

/* Smooth scroll */
html {
    scroll-behavior: smooth;
}

@media (max-width: 768px) {
    .nav-bar .container {
        gap: 1rem;
        font-size: 0.75rem;
    }
    
    .nav-separator {
        display: none;
    }
}
```

### Task 3: Add Bold ASCII Separator After Hero

**Files:**
- Modify: `/var/home/lucas/repos/portfolio2/index.html:171` (after hero section closing)

**Step 1: Add HTML separator element**

```html
<!-- After hero section closing tag -->
<div class="separator">
    <div class="container">
        <div class="ascii-line">================================================================================</div>
    </div>
</div>
```

**Step 2: Add CSS for ASCII separator**

```css
/* Add to styles section */
.separator {
    padding: 2rem 0;
    margin: 2rem 0;
}

.ascii-line {
    font-family: 'JetBrains Mono', monospace;
    font-size: 1.2rem;
    font-weight: 400;
    color: var(--text-accent);
    letter-spacing: -0.05em;
    text-align: center;
    overflow: hidden;
    white-space: nowrap;
}

@media (max-width: 768px) {
    .ascii-line {
        font-size: 0.9rem;
    }
}
```

### Task 4: Create Timeline Section HTML

**Files:**
- Modify: `/var/home/lucas/repos/portfolio2/index.html:172-173` (between hero separator and projects)

**Step 1: Add experience anchor and section**

```html
<!-- Add between separators -->
<section class="experience" id="experience">
    <div class="container">
        <h2 class="section-title">EXPERIENCE</h2>
        
        <div class="timeline">
            <div class="timeline-entry">
                <h3 class="timeline-company">Tech Corp</h3>
                <p class="timeline-role">Senior DevOps Engineer</p>
                <p class="timeline-period">2022 - 2024 | San Francisco, CA</p>
                <p class="timeline-desc">Kubernetes orchestration, CI/CD pipelines, infrastructure as code</p>
            </div>
            
            <div class="timeline-entry">
                <h3 class="timeline-company">Dev Solutions Inc</h3>
                <p class="timeline-role">DevOps Lead</p>
                <p class="timeline-period">2019 - 2022 | Remote</p>
                <p class="timeline-desc">AWS infrastructure, Docker containerization, monitoring systems</p>
            </div>
            
            <div class="timeline-entry">
                <h3 class="timeline-company">StartupXYZ</h3>
                <p class="timeline-role">Systems Administrator</p>
                <p class="timeline-period">2017 - 2019 | Austin, TX</p>
                <p class="timeline-desc">Linux server management, shell scripting, network configuration</p>
            </div>
        </div>
    </div>
</section>
```

### Task 5: Add Timeline Section Styles

**Files:**
- Modify: `/var/home/lucas/repos/portfolio2/index.html:11-162` (styles section)

**Step 1: Add timeline component styles**

```css
/* Add to styles after projects section */
/* Section title */
.section-title {
    font-family: 'JetBrains Mono', monospace;
    font-size: 1.5rem;
    font-weight: 700;
    color: var(--text-heading);
    margin-bottom: 2rem;
    padding-bottom: 1rem;
    border-bottom: 1px solid var(--border);
}

/* Timeline */
.experience {
    padding: 4rem 0;
}

.timeline {
    display: flex;
    flex-direction: column;
    gap: 2rem;
}

.timeline-entry {
    transition: background-color 0.2s ease;
    padding: 1rem 0;
}

.timeline-entry:hover {
    background-color: #222; /* Slightly lighter bg-dark variant */
}

.timeline-company {
    font-family: 'JetBrains Mono', monospace;
    font-size: 1.25rem;
    font-weight: 700;
    color: var(--text-heading);
    margin-bottom: 0.25rem;
}

.timeline-role {
    font-family: 'JetBrains Mono', monospace;
    font-size: 1rem;
    font-weight: 500;
    color: var(--text-primary);
    margin-bottom: 0.5rem;
}

.timeline-period {
    font-family: 'JetBrains Mono', monospace;
    font-size: 0.875rem;
    font-weight: 400;
    color: var(--text-accent);
    margin-bottom: 0.75rem;
}

.timeline-desc {
    font-size: 0.95rem;
    color: var(--text-primary);
    max-width: 600px;
    line-height: 1.5;
}
```

### Task 6: Add Second Separator Before Projects

**Files:**
- Modify: `/var/home/lucas/repos/portfolio2/index.html:175` (after experience section)

**Step 1: Add second ASCII separator**

```html
<!-- Add after experience section closing -->
<div class="separator">
    <div class="container">
        <div class="ascii-line">================================================================================</div>
    </div>
</div>

<!-- Projects section follows -->
<section class="projects">
    <h2 id="projects-title">PROJECTS</h2>
```

### Task 7: Add Contact Anchor for Navigation

**Files:**
- Modify: `/var/home/lucas/repos/portfolio2/index.html:203` (footer)

**Step 1: Add contact ID to footer**

```html
<footer id="contact">
    <div class="container">
        <a href="mailto:lucas.t.cottle@gmail.com">lucas.t.cottle@gmail.com</a>
    </div>
</footer>
```

### Task 8: Verify File Structure and Test Navigation

**Files:**
- Read: `/var/home/lucas/repos/portfolio2/index.html` (check all changes)

**Step 1: Verify HTML sections in order**

```html
<!-- Should be: -->
<body>
    <nav>...</nav>
    <main class="container">
        <section class="hero" id="hero">...</section>
        <div class="separator">...</div>
        <section class="experience" id="experience">...</section>
        <div class="separator">...</div>
        <section class="projects">...</section>
    </main>
    <footer id="contact">...</footer>
</body>
```

**Step 2: Open file in browser to test**

Command to serve local file:
```bash
cd /var/home/lucas/repos/portfolio2 && python3 -m http.server 8000
```
Expected: File should open at `http://localhost:8000/index.html`

Test interactions:
- Click nav links - should scroll to corresponding sections
- Hover nav links - should show underline
- Hover timeline entries - subtle background color change
- Verify ASCII separators are visible and monospace-sized
- Check mobile responsiveness (nav separators hide, minimal spacing)

### Task 9: Commit Changes

**Files:**
- Staged: `index.html` (all modifications)

**Step 1: Stage and commit**

```bash
git add index.html
git commit -m "feat: add brutalist navigation, experience timeline, and ASCII separators

- Add fixed top navigation with smooth-scroll anchors
- Create minimal experience timeline using monospace typography
- Add ASCII separators (========) after hero and before projects
- Maintain dark brutalist aesthetic with existing CSS variables
- Mobile responsive navigation (hides separators on narrow screens)"
```

**Step 2: Verify commit**

Command: `git log -1 --stat`
Expected: Shows `index.html` with many changed lines

