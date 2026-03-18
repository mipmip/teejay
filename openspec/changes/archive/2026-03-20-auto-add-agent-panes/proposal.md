## Why

Manually adding panes one by one is tedious when running multiple AI agents across tmux sessions. Teejay already knows which apps are "agents" (claude, aider, codex, opencode) from the config. A scan-and-auto-add function lets users populate their watchlist instantly with all panes running known agents. Relates to issue #27.

## What Changes

- Add an auto-scan function that scans all tmux panes, detects those running configured agent apps, and adds them to the watchlist automatically
- Expose this as a keybinding (`s` for scan) in the main TUI view
- Also expose as a CLI command (`tj scan`) for non-interactive use
- Use the app names from `config.detection.apps` as the list of known agents to detect
- Skip panes already in the watchlist (no duplicates)
- Auto-name added panes using the existing `naming.GuessName()` logic

## Capabilities

### New Capabilities
- `auto-scan-agents`: Scan all tmux panes for known agent processes and auto-add them to the watchlist

### Modified Capabilities
_None — this is additive. The watchlist, naming, and config systems are used as-is._

## Impact

- `internal/ui/app.go`: New keybinding handler for scan, status message showing results
- `internal/cmd/`: New `scan.go` command for CLI `tj scan`
- `cmd/tj/main.go`: Register the scan subcommand
- Uses existing: `tmux.ListAllPanes()`, `config.Detection.Apps`, `naming.GuessName()`, `watchlist.AddWithName()`
