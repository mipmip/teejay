## Why

Users need the ability to run multiple teejay instances with different configurations, or use project-specific watchlists. Currently, config and watchlist paths are hardcoded to `~/.config/teejay/`, making it impossible to override without modifying the home directory files.

## What Changes

- Add `--config` / `-c` flag to specify an alternative config.yaml path
- Add `--watchlist` / `-w` flag to specify an alternative watchlist.json path
- Thread custom paths through config.Load() and watchlist.Load() functions
- Update TUI and CLI subcommands to respect custom paths

## Capabilities

### New Capabilities

- `cli-path-flags`: CLI flag parsing for --config and --watchlist path overrides

### Modified Capabilities

- `config-file`: Load() must accept optional custom path parameter
- `watchlist-management`: Load() and Save() must accept optional custom path parameter

## Impact

- `cmd/tj/main.go`: Parse new flags, pass paths to config/watchlist loaders
- `internal/config/config.go`: Modify Load() signature to accept optional path
- `internal/watchlist/watchlist.go`: Modify Load()/Save() to accept optional path
- `internal/ui/app.go`: Store and use custom watchlist path for saves
- `internal/cmd/add.go` and `internal/cmd/del.go`: Accept custom watchlist path
