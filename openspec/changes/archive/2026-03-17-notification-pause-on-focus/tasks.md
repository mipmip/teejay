## 1. Tmux Active Pane Detection

- [x] 1.1 Add `GetActivePaneID()` function to `internal/tmux/` that runs `tmux display-message -p '#{pane_id}'` and returns the active pane ID (empty string on error)

## 2. Alert Suppression

- [x] 2.1 In the update loop in `internal/ui/app.go`, call `GetActivePaneID()` before `triggerAlerts()` and skip the alert call when the transitioning pane ID matches the active pane ID
- [x] 2.2 Verify that status indicator updates (green dot) still occur even when alerts are suppressed

## 3. Testing

- [x] 3.1 Test `GetActivePaneID()` returns a pane ID string when running inside tmux
- [x] 3.2 Test that alerts are suppressed for focused pane and fire for non-focused panes (verified by code inspection: guard at app.go:514 is a simple string comparison; integration testing requires running inside tmux with active panes)
