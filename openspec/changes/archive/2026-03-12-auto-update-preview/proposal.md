## Why

The preview panel only updates when the user navigates to a different pane. Users watching active tmux sessions (running builds, logs, etc.) need to see live content updates without manual intervention.

## What Changes

- Add ticker-based auto-refresh for the preview panel (100ms interval, inspired by claude-squad)
- Introduce a custom tea.Msg type for preview refresh ticks
- Modify `Init()` to start the refresh ticker
- Update `Update()` to handle tick messages and re-capture pane content

## Capabilities

### New Capabilities
- `preview-auto-refresh`: Automatic periodic refresh of the preview panel content using Bubbletea's command pattern

### Modified Capabilities
- `pane-preview`: Add requirement for automatic content refresh at regular intervals

## Impact

- `internal/ui/app.go`: Add tick message type, ticker command, and tick handler
- No breaking changes - existing behavior preserved, auto-refresh is additive
- Minimal performance impact: `tmux capture-pane` is lightweight
