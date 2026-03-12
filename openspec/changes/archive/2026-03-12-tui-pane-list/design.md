## Context

The TUI currently shows a static placeholder ("tmon - press q to quit"). The watchlist package exists and can load panes from `~/.config/tmon/watchlist.json`. We need to connect these: load the watchlist on startup and display it in a navigable list.

## Goals / Non-Goals

**Goals:**
- Display watched panes in a vertical list using bubbles/list
- Support keyboard navigation (up/down, j/k)
- Show pane ID and "added at" timestamp for each entry
- Handle empty state with a helpful message

**Non-Goals:**
- Pane content preview (separate change)
- Removing panes from the list
- Real-time updates/refreshing
- Left panel / right panel split layout (keep it simple for now)

## Decisions

### Decision 1: Use bubbles/list component

Use the `github.com/charmbracelet/bubbles/list` component for the pane list.

**Rationale**: It handles keyboard navigation, scrolling, and selection out of the box. Well-tested and consistent with Charm ecosystem patterns.

**Alternative considered**: Custom list implementation. Rejected because bubbles/list provides everything needed and is maintained upstream.

### Decision 2: Load watchlist once at startup

Load the watchlist in `New()` when creating the model. No live refresh.

**Rationale**: Simple and sufficient for MVP. Users can restart tmon to see new panes. Live refresh adds complexity (file watching, goroutines) that can come later.

### Decision 3: Simple single-panel layout

Display the list full-width for now, no split panels.

**Rationale**: The README mentions a left sidepanel with preview in main body, but that's a larger UI change. This change focuses on getting a functional list first.

## Risks / Trade-offs

- [Trade-off] No live refresh → Users must restart to see new panes → Acceptable for MVP
- [Trade-off] No panel layout → Doesn't match final vision → Incremental progress is fine
- [Risk] List doesn't fit terminal → bubbles/list handles scrolling automatically
