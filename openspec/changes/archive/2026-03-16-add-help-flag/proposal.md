## Why

Users have no way to discover available commands and flags when running `tj`. Without `--help`, users must read documentation or source code to understand how to use the CLI.

## What Changes

- Add `--help` and `-h` flags to display usage information
- Show available commands (`add`, `del`) with descriptions
- Show available global flags (`--config`, `--watchlist`, `--version`)
- Exit cleanly after displaying help

## Capabilities

### New Capabilities

- `cli-help`: Provides usage information via `--help`/`-h` flag, listing all commands and options

### Modified Capabilities

None - this is additive functionality that doesn't change existing behavior.

## Impact

- `cmd/tj/main.go`: Add help flag handling in `parseFlags` and display logic in `main`
- No breaking changes - purely additive
