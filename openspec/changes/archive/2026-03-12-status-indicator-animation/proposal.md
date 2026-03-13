## Why

The current status indicators are static symbols (●, ?, ○) that don't convey activity visually. A spinning/busy animation for the "Running" state would give users immediate visual feedback that something is happening, while a green circle for "Ready" would clearly indicate a pane is waiting for input.

## What Changes

- Replace static "●" Running indicator with an animated spinner (cycling through ⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏)
- Replace "?" Ready indicator with a green "●" circle
- Keep "○" for Idle state (no change needed)
- Animation cycles on the existing 100ms tick

## Capabilities

### New Capabilities

- `status-animation`: Animated spinner for running panes, green indicator for ready state

### Modified Capabilities

(none - this is implementation detail within the existing pane-list-view behavior)

## Impact

- `internal/monitor/status.go`: Change `Indicator()` to return frame-based spinner and colored indicators
- `internal/ui/app.go`: Pass tick/frame counter to indicator rendering
- Minimal change - uses existing tick mechanism
