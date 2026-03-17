## Why

When a user is actively interacting with a tmux pane (typing, editing), the pane content changes rapidly. Each change can trigger Busy→Waiting status transitions, flooding the user with sound alerts and desktop notifications. This defeats the purpose of notifications — they should inform the user when they're away, not stalk them while they're working in the pane.

## What Changes

- Add a "notification pause" state that suppresses alerts for a pane when the user has that pane focused in tmux
- Detect which tmux pane is currently active using `tmux display-message -p '#{pane_id}'` (or similar)
- Skip `triggerAlerts()` when the transitioning pane is the one the user is currently focused on in tmux
- Resume normal notification behavior when the user switches away from the pane

## Capabilities

### New Capabilities
- `focus-aware-alerts`: Suppress notifications for tmux panes that are currently focused by the user, detecting active pane via tmux and pausing alerts accordingly

### Modified Capabilities
- `pane-alerts`: Add a focus-check guard before triggering alerts — alerts are suppressed when the pane is the user's active tmux pane

## Impact

- `internal/ui/app.go`: Modify `triggerAlerts()` or its call site to check focus state before firing
- `internal/tmux/`: Add function to query the currently active tmux pane ID
- No config changes needed — this is a sensible default behavior (user shouldn't need to opt in)
- No breaking changes
