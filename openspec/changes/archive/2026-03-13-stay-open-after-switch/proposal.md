## Why

When pressing Enter to switch to a tmux pane, the application quits immediately after switching. This forces users to restart tmon every time they want to switch to another pane, which defeats the purpose of a persistent pane monitor. The app should stay open so users can switch between multiple panes in a session.

## What Changes

- Remove `tea.Quit` after successful pane switch
- App remains open after switching to a pane, allowing continued monitoring and further switches
- Only explicit quit actions (q, ctrl+c) should exit the application

## Capabilities

### New Capabilities

(none - this is a behavior fix within existing pane switching)

### Modified Capabilities

- `preview-switch`: Change switch behavior to not quit after switching

## Impact

- `internal/ui/app.go`: Remove `tea.Quit` from the Enter key handler after `tmux.SwitchToPane()`
