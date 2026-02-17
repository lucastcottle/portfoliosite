---
date: 2026-02-14
topic: portfolio-brutalist-enhancements
status: approved
---

# Portfolio: Brutalist Enhancements Design

## Context
Current portfolio uses a clean dark brutalist aesthetic with monospace typography (JetBrains Mono/IBM Sans), dark theme (1a1a1a background, 252525 cards), and minimalist card-based projects section.

## Design Goals
Maintain strict minimalist/brutalist aesthetic while adding:
- Bold ASCII-style separators
- Navigation bar for section jumping
- Timeline section for work experience

## Visual Design Decisions

### Navigation Bar
- Font: JetBrains Mono (700 weight)
- Position: Fixed/sticky at top
- Style: Inline text links matching footer style
- Content: [About] | [Experience] | [Projects] | [Contact]
- Interaction: Hover underline effect using --text-accent color

### Bold Separators
- Type: ASCII character bars (`========` or `----------`)
- Colors: --text-accent or --border
- Positions: 
  1. After hero content (below intro text)
  2. After timeline content (before projects section)
- Size: Full container width (1200px)

### Timeline Section
- Layout: Single column vertical list
- Font: JetBrains Mono monospace throughout
- Structure:
  - Company Name: 700 weight
  - Role: 500 weight
  - Period/Location: 400 weight, --text-accent color
  - Description: Optional, minimal tech keywords
- Spacing: Mirroring project card rhythm (1.5rem gap, 2rem padding sections)
- Interaction: Simple hover - subtle background shift from 1a1a1a to darker or lighter

## Technical Implementation
- All CSS uses existing custom properties (--bg-dark, --text-primary, etc.)
- No new fonts or external dependencies
- ASCII separators as text content with monospace spacing
- Smooth scroll behavior for nav links
- Mobile responsive: vertical nav/menu on narrow screens

## Success Criteria
- Visual language consistent with existing brutalist minimalism
- No visual clutter or unnecessary decorations
- Typography remains hero element
- Dark theme preserved
- All interactive elements have minimal subtle feedback (not flashy)
