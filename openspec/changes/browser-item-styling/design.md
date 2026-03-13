## Context

The browser popup uses `list.NewDefaultDelegate()` which renders items with a simple text style. We need custom styling with background colors. The bubbles list component supports custom delegates via the `list.ItemDelegate` interface. Currently, mouse click detection uses hardcoded item height (2 lines).

## Goals / Non-Goals

**Goals:**
- Dark grey background (#333333) for unselected items
- Lighter grey background (#555555) for selected items
- 1-line visual gap between items
- Full-width clickable area per item
- Works for both session list and pane list

**Non-Goals:**
- Themeable colors (will be added later)
- Different styling for different item types
- Hover effects (not supported without bubblezone)

## Decisions

### Decision 1: Create custom ItemDelegate

**Choice:** Implement `list.ItemDelegate` interface with custom `Render()` method that applies background styles.

**Rationale:**
- Standard bubbles pattern for custom list rendering
- Full control over item appearance
- Alternative: Modify item Title()/Description() - but can't control background color that way

### Decision 2: Use lipgloss Background() for item backgrounds

**Choice:** Use `lipgloss.NewStyle().Background(color).Width(width)` to create full-width background.

**Rationale:**
- Lipgloss handles terminal color rendering correctly
- Width ensures background extends to full item width
- Padding adds breathing room inside items

### Decision 3: Item height = 3 lines (title, desc, margin)

**Choice:** Each item takes 3 lines: 1 for title, 1 for description, 1 for margin/gap.

**Rationale:**
- Creates clear visual separation between items
- Margin is just empty space, not rendered content
- Mouse click detection needs to account for this new height

### Decision 4: Colors hardcoded for now

**Choice:** Use hardcoded grey values (#333333 normal, #555555 selected).

**Rationale:**
- Theming is explicitly a non-goal for now
- Simple to change later when theming is added
- Matches the dark theme aesthetic

## Risks / Trade-offs

**[Trade-off] Reduced visible items** → With 3-line items instead of 2, fewer items fit in view. Acceptable for clarity.

**[Risk] Mouse hit detection needs update** → Must update `updateBrowsing()` to use itemHeight=3. Mitigation: Update in same change.
