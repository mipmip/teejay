## Why

The TUI currently shows error messages that persist until manually dismissed with Esc. This creates unnecessary friction - temporary messages like "Cannot switch: not running inside tmux" should auto-dismiss after a few seconds. A general-purpose auto-dismiss mechanism will improve UX for this and all future temporary error/status messages.

## What Changes

- Add a reusable auto-dismiss mechanism for temporary messages in the TUI
- Messages auto-dismiss after a timeout (e.g., 3 seconds)
- Messages can still be dismissed early by pressing Esc
- Apply this mechanism to the "not in tmux" error as the first use case
- Future temporary messages can easily use the same pattern

## Capabilities

### New Capabilities
- `auto-dismiss-messages`: General-purpose automatic dismissal of temporary error/status messages after a configurable timeout

### Modified Capabilities
None - this is a new capability that doesn't modify existing spec requirements.

## Impact

- `internal/ui/app.go`: Add reusable timer-based auto-dismiss infrastructure
- Pattern established for all future temporary messages
- No API changes
- No breaking changes
- No new dependencies (uses existing Bubbletea tick mechanism)
