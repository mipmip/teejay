## Why

Users watching panes running long processes (builds, AI assistants, log streams) cannot tell at a glance whether a pane is actively outputting content or idle/waiting for input. This requires constantly watching the preview, which defeats the purpose of a monitoring tool.

## What Changes

- Add content hash tracking to detect when pane output changes
- Add prompt pattern detection to identify when tools (Claude, Aider, etc.) are waiting for input
- Track pane status: `Running` (content changing), `Ready` (waiting for input), `Idle` (no recent changes)
- Display status indicator next to each pane in the list

## Capabilities

### New Capabilities
- `activity-detection`: Detect pane activity state using content hashing and prompt pattern matching

### Modified Capabilities
<!-- None -->

## Impact

- `internal/tmux/`: New monitor package for hash tracking and status detection
- `internal/ui/app.go`: Store and display pane status, update status on each tick
- `internal/watchlist/`: Add Status field to track per-pane state
