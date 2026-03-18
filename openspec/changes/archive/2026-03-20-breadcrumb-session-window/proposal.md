## Why

When monitoring multiple tmux panes across different sessions and windows, it's hard to tell at a glance where each pane lives. Adding a breadcrumb showing the session and window context helps users orient themselves, especially when running parallel AI agents across projects. Relates to issue #28.

## What Changes

- Add a breadcrumb display to each watched pane item showing the tmux hierarchy: `session > window : process`
- The breadcrumb replaces the current plain process-name description line with a structured path that includes session name, window name, and the active foreground process
- Example: `technative-docs > proposals : claude`

## Capabilities

### New Capabilities
- `pane-breadcrumb`: Display a breadcrumb trail showing session, window, and active process for each watched pane in the list view

### Modified Capabilities
- `watchlist-item-delegate`: The watchlist item description changes from showing just the process name to showing the full breadcrumb trail

## Impact

- `internal/ui/app.go`: paneItem struct and Description() method need session/window fields and breadcrumb formatting
- `internal/tmux/list.go`: PaneInfo already contains Session, WindowName — these need to flow through to the UI paneItem
- Watchlist item delegate rendering may need width adjustments for longer description text
