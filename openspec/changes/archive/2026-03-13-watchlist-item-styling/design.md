## Context

The main watchlist uses `list.NewDefaultDelegate()` which renders items with simple text styling. We recently implemented custom styling with background colors for the browser popup items using `browserItemDelegate`. The same pattern should be applied to the main watchlist panel. Currently, mouse click detection uses hardcoded item height (2 lines).

## Goals / Non-Goals

**Goals:**
- Dark grey background (#333333) for unselected items
- Lighter grey background (#555555) for selected items
- 1-line visual gap between items
- Full-width clickable area per item
- Consistent styling with the browser popup items

**Non-Goals:**
- Themeable colors (will be added later)
- Different styling for different pane statuses (e.g., ready vs running)
- Hover effects (not supported without bubblezone)

## Decisions

### Decision 1: Reuse browserItemDelegate

**Choice:** Reuse the existing `browserItemDelegate` struct that was just created for browser popup styling.

**Rationale:**
- Same visual requirements (background colors, 3-line height)
- Avoids code duplication
- Ensures visual consistency across the app
- Alternative: Create separate `watchlistItemDelegate` - rejected as unnecessary duplication

### Decision 2: Item height = 3 lines (title, desc, margin)

**Choice:** Each item takes 3 lines: 1 for title, 1 for description, 1 for margin/gap.

**Rationale:**
- Matches browser popup styling
- Creates clear visual separation between items
- Mouse click detection needs to account for this new height

### Decision 3: Update main list mouse click detection

**Choice:** Update the mouse click detection in the main `Update()` method to use itemHeight=3.

**Rationale:**
- Must match the delegate's Height() return value
- Ensures clicks select the correct item

## Risks / Trade-offs

**[Trade-off] Reduced visible items** → With 3-line items instead of 2, fewer items fit in view. Acceptable for clarity.

**[Risk] Mouse hit detection mismatch** → If itemHeight doesn't match delegate Height(), clicks will select wrong items. Mitigation: Update in same change, test thoroughly.
