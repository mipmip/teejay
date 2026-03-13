## Why

When the TUI app is running and the user runs `tmon add` from another terminal, the watchlist file changes but the app doesn't reflect the new pane. Users expect the list to update automatically without restarting the app.

## What Changes

- Detect changes to the watchlist.json file while the TUI is running
- Automatically reload the watchlist and refresh the UI when changes are detected
- Preserve current selection when possible after refresh

## Capabilities

### New Capabilities

- `watchlist-file-watch`: Monitor watchlist.json for external changes and trigger UI refresh

### Modified Capabilities

(none)

## Impact

- `internal/ui/app.go`: Add file watching, handle refresh messages, reload watchlist on change
- `internal/watchlist/watchlist.go`: May need to expose config path for file watching
