## Why

The TUI currently shows only a placeholder message. Users have no way to see which panes are being watched. This change transforms the TUI into a functional pane list view, making tmon actually useful as a monitoring tool.

## What Changes

- Load and display the watchlist in the TUI
- Show panes in a vertical list in the left panel
- Use bubbles/list component for navigation and selection
- Display pane ID and when it was added
- Handle empty watchlist state gracefully

## Capabilities

### New Capabilities
- `pane-list-view`: TUI component that displays watched panes in a navigable list with selection support

### Modified Capabilities
<!-- None - the current TUI is a placeholder with no real functionality -->

## Impact

- `internal/ui/app.go`: Replace placeholder with real pane list UI
- `internal/ui/`: May add additional files for list styling/helpers
- Imports `internal/watchlist` package to load pane data
