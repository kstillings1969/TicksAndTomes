# UI Architecture - Management Screen

Mobile-first information architecture for empire management dashboard.

## Screen Overview

The management screen is a single-column, vertically scrolling mobile interface with collapsible sections. Web displays as 1-2 columns (TBD based on content testing).

## Information Hierarchy

### Primary Section: Empire Status (Always Visible)
Quick at-a-glance empire health. Sticky header on scroll.

```
┌─────────────────────────────┐
│ Empire Name                 │
│ Land: 98,723                │
│ Status: Alive (Avg Health)  │
└─────────────────────────────┘
```

### Section 1: Resources & Rates
Current resources with per-tick delta. Color-coded status (surplus/deficit/neutral).

**Data Layout:**
```
RESOURCES

Gold
├─ Current: 12,987
├─ Per Tick: +2,389
└─ Status: Surplus (green)

Grain
├─ Current: 324
├─ Per Tick: -2,328
└─ Status: Deficit (red)

Tomes
├─ Current: 23,421
├─ Per Tick: +34
├─ Meditation Towers: 187
├─ Ratio: 125.2:1 (optimal)
└─ Status: Optimal (blue)

Troops
├─ Current: 3,221 total
├─ Per Tick: +48
├─ Infantry: 2,100
├─ Archers: 321
├─ Cavalry: 32
├─ Navy: 98,723,223
└─ Breakdown: [collapsible sub-section]

Ticks (Regenerating Resource)
├─ Current: 287 / 500
├─ Regeneration: 1 per 5 min
└─ Next Tick In: 2m 34s
```

**Status Indicators:**
- Surplus: Light green background, small ✓ icon
- Deficit: Light red background, small ⚠ icon
- Neutral/Optimal: Blue or gray background

### Section 2: Buildings & Production
Expandable section showing building counts and their contribution.

**Data Layout:**
```
BUILDINGS & PRODUCTION (collapsed by default)

[Expand button]

─ Farms: 342
  └─ Grain Production: +142/tick (farm_level: 7)

─ Barracks: 156
  └─ Troop Production: +48/tick (drill_level: 5)
  └─ [CONTROL] Troop Distribution:
     - Infantry: 45%
     - Archers: 35%
     - Cavalry: 15%
     - Navy: 5%

─ Towers: 89
  └─ Defense: +8,900 points

─ Bazaars: 34
  └─ Gold Production: +2,389/tick

─ Meditation Towers: 187
  └─ Tome Production: +34/tick (meditation_level: 8)
```

### Section 3: Civilization Management
Economic engine controls and status.

**Data Layout:**
```
CIVILIZATION & ECONOMY

Civilians
├─ Current: 45,210,000
├─ Cap: 98,723,000 (land * 1000)
├─ Growth Rate: +523,000/tick
└─ Status: Healthy Growth

Tax Rate Control
├─ Current: 12%
├─ [SLIDER: 0-100%]
├─ Gold/Tick Impact: +1,489 per 1% increase
└─ Civilian Impact:
   - Below 9%: 🟢 +25% growth bonus
   - 10-22%: ⚪ Neutral
   - 23-34%: 🟡 -15% to -35% growth penalty
   - 35-99%: 🔴 Civilians fleeing
   - 100%: 🔴 -7% loss/tick
```

### Section 4: Action Levels & Experience
Skill progression for each action.

**Data Layout:**
```
ACTION PROGRESSION

Explore
├─ Level: 3 / 10
├─ Experience: ████████░░ 78% to next level
└─ Bonus Output: +60% per tick

Meditate
├─ Level: 8 / 10
├─ Experience: ██████░░░░ 52% to next level
└─ Bonus Output: +160% per tick

Drill
├─ Level: 5 / 10
├─ Experience: ███████░░░ 65% to next level
└─ Bonus Output: +100% per tick

Farm
├─ Level: 7 / 10
├─ Experience: █████░░░░░ 43% to next level
└─ Bonus Output: +140% per tick
```

### Section 5: Intelligence Report (Conditional)
After casting intelligence spell, show this instead of resource summary.

**Data Layout:**
```
INTELLIGENCE REPORT - [Opponent Name]

Gathered At: [timestamp]
Accuracy: [Based on caster's intelligence level]

Empire Status
├─ Land: 45,231
├─ Estimated Troops: 8,923
├─ Estimated Gold: 34,921
├─ Estimated Tomes: 12,345
└─ Estimated Defense: 4,500

Building Estimates
├─ Farms: ~145 (est.)
├─ Barracks: ~67 (est.)
├─ Towers: ~45 (est.)
└─ [etc]

[Close Report]
```

---

## Layout Details

### Mobile (< 768px)
- Single column, full width
- Sticky header: Empire name + critical status
- Sections stack vertically
- Controls are full-width inputs/sliders
- Collapsible sections to manage scroll depth

### Web (768px+)
- Two-column layout (TBD after prototype):
  - Left column: Resources, Buildings, Civilization
  - Right column: Action Levels, Intelligence Report
- Sections side-by-side where possible
- Wider input controls

### Color Scheme (Minimalist)
- Background: White or very light gray (#fafafa)
- Text: Dark gray (#333) for primary, medium gray (#666) for secondary
- Status Colors:
  - Surplus: #4caf50 (green)
  - Deficit: #f44336 (red)
  - Optimal: #2196f3 (blue)
  - Neutral: #9e9e9e (gray)
  - Warning: #ff9800 (orange)
- Borders: Light gray (#eee)
- Accents: Blue (#2196f3) for interactive elements

### Typography
- Headings: Sans-serif, bold, 18-24px
- Body: Sans-serif, regular, 14-16px
- Data values: Monospace for numbers, 16px
- Status labels: Sans-serif, 12px, uppercase

### Spacing
- Section padding: 16px mobile, 20px web
- Vertical gap between items: 12px
- Vertical gap between sections: 24px

---

## Interaction Patterns

### Collapsible Sections
- Click header to expand/collapse
- Persist state in browser localStorage
- Smooth 200ms animation

### Number Formatting
- Large numbers: Use comma separators (98,723,000)
- Per-tick values: Show as +/-X notation
- Percentages: Show decimal (12.5%) or whole (12%)
- Abbreviated form: Use K, M, B for thousands/millions/billions in tight spaces

### Controls
- **Sliders**: Tax rate (0-100%), troop distribution percentages
- **Dropdowns**: Filter/sort options (if added later)
- **Buttons**: Primary action (Apply, Save) in blue, secondary (Cancel) in gray
- **Input Fields**: Full-width on mobile, constrained on web

### Status Indicators
Use icons + background color + label:
```
Deficit    ⚠  Red background, left-aligned icon
Surplus    ✓  Green background, left-aligned icon
Neutral    -  Gray background
Optimal    ⚡ Blue background
Warning    !  Orange background
```

---

## Data Flow

### From Backend to UI
1. API returns empire state object every N seconds (configurable)
2. UI calculates derived values (per-tick deltas, ratios, status colors) locally
3. Intelligence report received as modal/section update after spell cast

### From UI to Backend
1. User adjusts tax rate → debounced API call (500ms)
2. User adjusts troop distribution → debounced API call (500ms)
3. User casts spell → modal opens, awaits response, displays result

---

## Accessibility

- All controls have associated labels
- Color is not the only indicator (use icons + labels)
- Sufficient contrast: WCAG AA minimum
- Touch targets: 44px minimum height
- Mobile: Landscape mode supported (media queries)
- Indicators: Screen-reader friendly alt text

