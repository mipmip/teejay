## Why

The TUI currently shows a list of watched panes but doesn't show what's actually in them. Users need to see the pane content to know what's happening—this is the core value proposition of tmon as a monitoring tool.

## What Changes

- Add a split-panel layout: pane list on left, preview on right
- Capture tmux pane content using `tmux capture-pane`
- Display the selected pane's content in a readonly viewport
- Update preview when selection changes

## Capabilities

### New Capabilities
- `pane-preview`: Split-panel layout with readonly pane content preview that updates based on list selection

### Modified Capabilities
<!-- None - building on top of existing pane-list-view -->

## Impact

- `internal/ui/app.go`: Major refactor for split-panel layout
- New `internal/tmux/` package for tmux interaction
- Uses `tmux capture-pane -p -t <pane-id>` to get content
