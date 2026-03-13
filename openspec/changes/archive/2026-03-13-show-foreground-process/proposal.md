## Why

When monitoring multiple tmux panes running AI agents or build processes, knowing what command is currently running in each pane provides valuable context at a glance. Currently the description line shows the pane ID and added date, which is less useful for quick status assessment.

## What Changes

- Replace current description line content (pane ID + added date) with the running foreground process
- Fetch current command from tmux for each watched pane
- Update description dynamically as the foreground process changes

## Capabilities

### New Capabilities

- `pane-foreground-process`: Display the current foreground process for each watched pane in the description line

### Modified Capabilities

(none - this enhances existing display, not changing spec-level behavior)

## Impact

- `internal/ui/app.go`: Update paneItem to include current command, modify Description() to show it
- `internal/tmux/list.go`: Already has `pane_current_command` - may need function to get command for single pane
- Periodic refresh already exists for preview - command info can be updated on same tick
