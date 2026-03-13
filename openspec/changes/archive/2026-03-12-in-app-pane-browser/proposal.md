## Why

Currently, adding a pane to the watchlist requires leaving the TUI and running `tmon add` in the target pane. This breaks workflow and requires context switching. Users should be able to browse and add any tmux pane from within the application.

## What Changes

- Add `a` key binding to open a pane browser popup
- Display tmux session/window/pane tree structure in the popup
- Allow navigating and selecting a pane to add to the watchlist
- Show pane preview or identifier to help identify the right pane
- Close popup with Escape or after successful add

## Capabilities

### New Capabilities
- `pane-browser`: Browse and select tmux panes from within the TUI via popup overlay

### Modified Capabilities
<!-- None -->

## Impact

- `internal/tmux/`: Add functions to list sessions, windows, and panes
- `internal/ui/app.go`: Add popup state, overlay rendering, and key handling
- New `internal/ui/browser.go` or similar for browser component
