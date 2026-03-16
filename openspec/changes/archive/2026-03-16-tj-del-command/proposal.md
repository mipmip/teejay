## Why

Currently there's no CLI command to remove a pane from the watchlist - users must manually edit the watchlist.json file or use the TUI. A `tj del` command provides a symmetrical counterpart to `tj add` for quick workflow management. Additionally, the feedback messages from these commands should show the pane's human-readable name (using the naming system) for better UX.

## What Changes

- Add new CLI command `tj del` that removes the current pane from the watchlist
- Update `tj add` feedback message to show the resolved pane name
- Update `tj del` feedback message to show the resolved pane name
- Both commands use the `naming.GuessName()` system for consistent naming

## Capabilities

### New Capabilities

- `cli-del-command`: CLI command to remove current tmux pane from watchlist with named feedback

### Modified Capabilities

None

## Impact

- New file: `internal/cmd/del.go` - delete command implementation
- Modified: `cmd/tj/main.go` - add "del" case to CLI dispatch
- Modified: `internal/cmd/add.go` - improve feedback message (already shows name, verify consistency)
