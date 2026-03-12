## Why

The pane preview currently displays plain text without ANSI colors. Terminal applications use colors extensively (syntax highlighting, status indicators, progress bars), and losing these colors makes the preview much less useful for understanding pane content at a glance.

## What Changes

- Update `tmux capture-pane` command to include `-e` flag to preserve ANSI escape sequences
- Add `-J` flag to join wrapped lines for cleaner display
- The viewport will render the raw ANSI content, which Bubbletea/Lipgloss handles natively

## Capabilities

### New Capabilities
- `ansi-color-capture`: Capture tmux pane content with preserved ANSI escape sequences for full-color rendering

### Modified Capabilities
<!-- None - this enhances existing capture without changing spec-level behavior -->

## Impact

- `internal/tmux/capture.go`: Update CapturePane to use `-e -J` flags
