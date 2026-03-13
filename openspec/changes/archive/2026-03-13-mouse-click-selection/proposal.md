## Why

Currently the TUI only supports keyboard navigation. Users accustomed to mouse interaction cannot click to select items in the pane list, session list, or other UI elements. Adding mouse support improves accessibility and matches user expectations from modern terminal applications.

## What Changes

- Enable mouse support in the Bubbletea program
- Allow mouse clicks to select items in the main pane list
- Allow mouse clicks to select items in the browser popup (sessions and panes)
- Allow mouse clicks to select items in the configure popup menu

## Capabilities

### New Capabilities

- `mouse-support`: Enable mouse click selection across all list-based UI components

### Modified Capabilities

(none - this adds a new input method without changing existing keyboard behavior)

## Impact

- `cmd/tmon/main.go`: Add `tea.WithMouseCellMotion()` option to enable mouse support
- `internal/ui/app.go`: Handle `tea.MouseMsg` events in Update() to select items on click
