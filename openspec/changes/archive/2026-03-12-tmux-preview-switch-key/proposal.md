## Why

When previewing a tmux pane in tmon, users often want to switch to that pane to interact with it. Currently there's no way to jump directly to the previewed pane from within tmon.

## What Changes

- Add a keybinding (Enter) that switches tmux to the currently selected/previewed pane
- When running inside tmux, pressing Enter will switch to the session/window/pane being previewed
- When not running inside tmux, show an informative message that switching requires tmux

## Capabilities

### New Capabilities

- `preview-switch`: Keybinding to switch tmux to the currently previewed pane

### Modified Capabilities

- `pane-preview`: Add switch keybinding to preview interaction

## Impact

- `internal/ui/app.go`: Add Enter key handler to switch to previewed pane
- `internal/tmux/`: Add function to switch to a specific pane
- Help text in footer needs to include the new keybinding
