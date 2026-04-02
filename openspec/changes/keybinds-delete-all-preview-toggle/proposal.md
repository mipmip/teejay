## Why

Two common operations lack keyboard shortcuts:
1. Clearing the entire watchlist — currently requires deleting panes one by one with `d`
2. Toggling the preview panel at runtime — currently only controllable via `--no-preview` flag or config, with no way to toggle during a session

## What Changes

- Add `D` (shift+d) keybinding: delete all panes from watchlist with a confirmation popup ("Delete all N panes? y/n")
- Add `p` keybinding: toggle preview panel visibility at runtime (toggles `m.config.Display.ShowPreview`)
- Update help footer, README keybindings table, and CHANGELOG

## Capabilities

### New Capabilities

_None_

### Modified Capabilities

- `pane-list-view`: Add "delete all" action with confirmation
- `pane-preview`: Preview can be toggled at runtime with `p` key

## Impact

- `internal/ui/app.go` — add `D` key handler with confirmation popup, add `p` key handler toggling ShowPreview, update help footer
- README.md — add to keybindings table
- CHANGELOG.md
