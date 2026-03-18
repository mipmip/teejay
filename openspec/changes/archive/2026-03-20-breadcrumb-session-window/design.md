## Context

Teejay's watchlist currently shows each pane with a name (title) and the foreground process (description). The `paneItem` struct in `internal/ui/app.go` has `id`, `name`, `command` fields, but no session or window information. Meanwhile, `tmux.PaneInfo` already provides `Session` and `WindowName` from `tmux list-panes`.

Users monitoring panes across multiple sessions and windows need to see where each pane lives in the tmux hierarchy at a glance.

## Goals / Non-Goals

**Goals:**
- Show a breadcrumb in the pane list description: `session > window : process`
- Reuse existing tmux metadata — no new tmux commands needed
- Keep the display compact and readable within the watchlist panel width

**Non-Goals:**
- Changing the pane title/name display — only the description line changes
- Making the breadcrumb format configurable (can be added later)
- Adding breadcrumb to the preview panel header

## Decisions

### 1. Breadcrumb in the description line
The breadcrumb replaces the current `command` description with `session > window : process`. This keeps the 3-line item height (title, description, margin) unchanged.

**Alternative considered**: Adding a third line for the breadcrumb — rejected because it would change `itemHeight` from 3 to 4, breaking mouse click detection and increasing visual density.

### 2. Add session/windowName fields to paneItem
The `paneItem` struct gains `session` and `windowName` fields, populated during `refreshListWithFrame()` from the `tmux.PaneInfo` data already fetched during status updates.

**Alternative considered**: Looking up PaneInfo on each render — rejected because the data is already available during refresh.

### 3. Format: `session > window : process`
Use `>` as the hierarchy separator between session and window, and `:` before the process name. This mirrors the user's requested format (`technative-docs > proposals : claude`). When process is empty, omit the `: process` suffix. When session or window name is empty/generic, still show available parts.

## Risks / Trade-offs

- **Long breadcrumbs may truncate** → The watchlist panel is ~30% width. Long session/window names will be truncated by lipgloss rendering. This is acceptable — the most important info (session name) comes first.
- **Stale session/window info** → Session/window data refreshes on each tick (100ms), same as process info. No additional staleness risk.
