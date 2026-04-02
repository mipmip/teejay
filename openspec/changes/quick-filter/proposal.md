## Why

When monitoring many panes, finding a specific one requires scanning the entire list visually. A quick filter (vim-style `/` search) lets users type a query to instantly narrow the list to matching panes. This is especially useful in multi-column mode with many agents.

Addresses GitHub issue #36.

## What Changes

- Press `/` to enter filter mode: a text input appears in the footer
- Typing filters the pane list in real-time, showing only panes whose name, session, window name, or command contain the query (case-insensitive)
- `Enter` confirms the filter and returns to normal navigation (filter stays active)
- `Esc` clears the filter and returns to showing all panes
- When a filter is active, the footer shows the filter query as a reminder
- Works in both default and multi-column layouts

## Capabilities

### New Capabilities

- `quick-filter`: `/`-triggered text filter on the pane list, matching against pane name, session, window, and command

### Modified Capabilities

_None_

## Impact

- `internal/ui/app.go` — filter mode state, `/` keybinding, filter logic in `refreshListWithFrame()`, footer rendering
- README.md — keybindings table
- CHANGELOG.md
