## Why

When the user types in a monitored pane, the content changes and the monitor marks it as "busy" (spinning indicator). This is misleading — the pane is active because of user input, not because an agent is working. The focused pane's status should be frozen while the user is interacting with it, and resume monitoring after a short grace period when the user switches away. Relates to issue #24.

## What Changes

- Skip status updates for the currently focused tmux pane — keep its last known status while the user is interacting
- Track when focus leaves a pane and add a configurable grace period (default 2s) before resuming status monitoring for that pane
- The existing focus-aware alert suppression (already skips alerts for focused panes) continues to work as before

## Capabilities

### New Capabilities
- `focus-pause-monitoring`: Pause status monitoring for the focused pane and apply a grace period after defocus

### Modified Capabilities
_None — the existing focus-aware-alerts spec only covers alert suppression, not status monitoring_

## Impact

- `internal/ui/app.go`: The tick loop needs to skip `monitor.Update()` for the active pane and track defocus timestamps
- `internal/config/config.go`: Optional: configurable grace period duration (or hardcode 2s)
- No new dependencies
