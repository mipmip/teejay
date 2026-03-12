## Why

Users need a way to add tmux panes to the watchlist. The `tmon add` command allows users to run a simple command from within any tmux pane to add that pane to monitoring. This is the first step toward making tmon functional as a pane monitor.

## What Changes

- Add `tmon add` subcommand that captures the current tmux pane ID
- Create a watchlist storage mechanism (JSON file in user config directory)
- Add watchlist data types and persistence logic

## Capabilities

### New Capabilities
- `add-pane`: CLI subcommand to add the current tmux pane to the watchlist, including tmux pane ID detection and watchlist persistence

### Modified Capabilities
<!-- None -->

## Impact

- `cmd/tmon/main.go`: Add subcommand routing (TUI vs CLI commands)
- New `internal/watchlist/` package for watchlist management
- New config file at `~/.config/tmon/watchlist.json`
