## Why

When a tmux pane is closed while being watched, the app continues attempting to capture it every 100ms, resulting in persistent "can't find pane: %XX" errors displayed to the user. The stale pane remains in the watchlist until manually removed.

## What Changes

- Detect when a watched pane no longer exists in tmux (capture returns "can't find pane" error)
- Automatically remove stale panes from the watchlist when detected
- Provide visual feedback before removal so users understand what happened

## Capabilities

### New Capabilities

- `stale-pane-removal`: Automatic detection and removal of panes that no longer exist in tmux

### Modified Capabilities

- `watchlist-management`: Add ability to remove panes programmatically when they become stale

## Impact

- `internal/ui/app.go`: Add stale pane detection in the preview refresh loop
- `internal/watchlist/watchlist.go`: May need method to remove by ID with callback/notification
- `internal/tmux/capture.go`: Could add specific error type for missing panes (optional)
- User experience: Errors will auto-resolve instead of persisting until manual deletion
