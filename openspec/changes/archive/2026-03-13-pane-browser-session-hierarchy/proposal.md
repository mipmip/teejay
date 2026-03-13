## Why

The current pane browser shows a flat list of all tmux panes across all sessions. When users have many sessions with multiple panes, this becomes unwieldy to navigate. A hierarchical session-first selection improves UX by reducing cognitive load and allowing users to focus on one session at a time.

## What Changes

- Replace flat pane list with two-step selection: first choose session, then choose pane within that session
- Escape key navigates back: from pane list to session list, from session list closes the browser
- Session list shows session name and pane count
- Pane list shows window.pane index and running command

## Capabilities

### New Capabilities

(none - this modifies existing pane-browser capability)

### Modified Capabilities

- `pane-browser`: Change from flat pane list to hierarchical session→pane selection with back navigation

## Impact

- `internal/ui/app.go`: Update browser state machine, add session list view, modify navigation logic
