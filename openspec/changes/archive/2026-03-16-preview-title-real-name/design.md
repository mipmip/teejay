## Context

Currently, the preview panel title in `internal/ui/app.go` shows `"Preview: " + m.selectedPaneID`, which displays just the raw pane ID (e.g., "%0"). The `paneItem` struct already contains `name` and `command` fields that could provide more meaningful information.

The current code at line 1015:
```go
previewTitle := previewTitleStyle.Render("Preview: " + m.selectedPaneID)
```

## Goals / Non-Goals

**Goals:**
- Display the pane's display name (custom name or guessed name) in the preview title
- Optionally show the current foreground command as additional context
- Maintain consistent styling with the existing preview panel

**Non-Goals:**
- Changing the preview panel layout or styling
- Adding new configuration options for title format
- Modifying how pane names are stored or computed

## Decisions

### Decision 1: Use existing paneItem data

**Choice**: Get title information from the currently selected `paneItem` in the list

**Rationale**: The `paneItem` already contains `name` and `command` fields populated during `refreshList()`. This avoids additional lookups and keeps the code consistent with how the pane list displays items.

**Alternative considered**: Look up the pane directly from `m.watchlist.GetPane(m.selectedPaneID)` - rejected because `paneItem` already has the processed display information.

### Decision 2: Title format

**Choice**: Use format `"Preview: {name}"` where name falls back to pane ID if no custom name exists

**Rationale**: Simple, clean, and matches user expectations. The command is already visible in the pane list, so duplicating it in the preview title would be redundant.

**Alternative considered**: Including command in title like `"Preview: {name} ({command})"` - rejected to avoid clutter since command is already shown in the list.

## Risks / Trade-offs

**[Risk] paneItem not found for selected ID** → Find the matching item by iterating through list items, fallback to pane ID if not found.

**[Trade-off] Slight performance impact** → Iterating through list items to find the selected one. Acceptable since watchlists are typically small (<100 panes).
