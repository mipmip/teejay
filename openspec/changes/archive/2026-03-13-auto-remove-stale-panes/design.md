## Context

The TUI app monitors tmux panes by capturing their content every 100ms. When a pane is closed in tmux, the capture fails with "can't find pane: %XX" error. Currently this error displays persistently until the user manually deletes the pane from the watchlist.

The watchlist is stored in `~/.config/teejay/watchlist.json` and managed by `internal/watchlist/watchlist.go`. The capture logic is in `internal/tmux/capture.go` and the UI refresh loop is in `internal/ui/app.go`.

## Goals / Non-Goals

**Goals:**
- Automatically detect when a watched pane no longer exists
- Remove stale panes from the watchlist without user intervention
- Provide brief visual feedback before removal so users understand what happened

**Non-Goals:**
- Proactively validating panes on startup (could be added later)
- Periodic background validation independent of capture attempts
- Confirmation dialogs before auto-removal (would defeat the purpose)

## Decisions

### 1. Detection point: In the UI capture loop

**Decision**: Detect stale panes in `app.go` when `captureSelectedPane()` returns an error containing "can't find pane".

**Rationale**: This is the simplest approach with minimal code changes. The error already surfaces here, we just need to act on it.

**Alternatives considered**:
- Create a custom error type in `capture.go` - adds complexity for little benefit
- Validate with `ListAllPanes()` before capture - extra tmux call, slower
- Background goroutine for validation - overkill for this use case

### 2. Removal timing: Immediate with brief notification

**Decision**: Remove the pane immediately when detected as stale. Show a brief status message (e.g., "Removed stale pane %65") that clears on next action.

**Rationale**: Delaying removal provides no benefit since the pane is already gone. Immediate cleanup prevents error spam.

**Alternatives considered**:
- Confirmation dialog - defeats the purpose of auto-removal
- Delay/retry - pane won't magically reappear
- Queue for batch removal - unnecessary complexity

### 3. Error detection method: String matching

**Decision**: Check if error message contains "can't find pane" substring.

**Rationale**: This is the exact error tmux returns. While somewhat brittle, it's simple and tmux error messages are stable.

**Alternatives considered**:
- Custom error type wrapping - more code for same outcome
- Regex parsing - overkill for a single known error pattern

## Risks / Trade-offs

**[Risk] False positive removal if error message changes** → Low risk. Tmux error messages are stable. If it changes, we'd simply stop auto-removing until updated.

**[Risk] Race condition if pane recreated with same ID** → Tmux assigns incrementing pane IDs, so a new pane won't reuse the old ID. Not a real concern.

**[Trade-off] No undo for auto-removal** → Acceptable. User can re-add the pane if needed. The pane was already gone from tmux anyway.
