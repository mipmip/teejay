## Context

The tick loop (every 100ms) calls `monitor.Update()` for every watched pane, including the one the user is currently focused on. When the user types, the pane content changes and the monitor sees activity → marks it as Busy. The existing focus-aware logic only suppresses alerts but does NOT skip the status update itself.

## Goals / Non-Goals

**Goals:**
- Skip `monitor.Update()` entirely for the currently focused pane
- After the user switches focus away from a pane, wait a grace period before resuming monitoring
- Keep the pane's last known status visible while paused

**Non-Goals:**
- Making the grace period configurable in config.yaml (hardcode 2s for now, can be made configurable later)
- Changing how alerts are suppressed (existing logic is fine)

## Decisions

### 1. Track active pane ID and defocus timestamps in the Model
Add `lastActivePaneID string` and `paneFocusLostAt map[string]time.Time` to the Model. On each tick, compare current active pane with previous. When a pane loses focus, record the timestamp.

### 2. Skip monitoring for active pane + grace period panes
In the tick loop, before calling `monitor.Update()` for a pane:
- If `paneID == activePaneID` → skip (user is typing)
- If `paneID` is in `paneFocusLostAt` and less than 2s has elapsed → skip (grace period)
- Otherwise → monitor normally

### 3. Clean up expired grace periods
After the grace period expires, remove the entry from `paneFocusLostAt` so monitoring resumes and the map doesn't grow unbounded.

### 4. Still update preview content for focused pane
The preview panel should still show live content for the focused pane — only skip the status determination, not the content capture.

## Risks / Trade-offs

- **2s grace period may be too short or too long** → 2s matches the existing idle timeout default, feels natural. Can be made configurable later.
- **GetActivePaneID() called on every tick** → Already happening for alert suppression. No additional cost.
