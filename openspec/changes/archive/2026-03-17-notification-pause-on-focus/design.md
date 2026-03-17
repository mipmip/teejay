## Context

Teejay monitors tmux panes every 100ms, detecting Busy→Waiting transitions and firing alerts (sounds + desktop notifications). When a user is actively working in a monitored pane, their typing causes rapid content changes that trigger frequent status transitions, resulting in a flood of unwanted notifications. Notifications are meant to alert users who are away from a pane, not to "stalk" users who are actively working.

## Goals / Non-Goals

**Goals:**
- Suppress alerts for panes the user is currently focused on in tmux
- Detect the user's active tmux pane efficiently within the existing 100ms tick loop
- Make this behavior automatic — no configuration required

**Non-Goals:**
- Suppressing status indicator changes in the UI (the green dot should still appear)
- Queuing suppressed notifications for later delivery
- Detecting focus at the terminal emulator level (only tmux-level focus)

## Decisions

### Decision 1: Query active pane via `tmux display-message`

Use `tmux display-message -p '#{pane_id}'` to get the currently active pane ID in the user's tmux session.

**Why over alternatives:**
- `tmux list-panes -F '#{pane_id} #{pane_active}'` requires parsing multiple lines
- `display-message` returns a single pane ID string, minimal parsing needed
- The active pane is always from the user's current client perspective

### Decision 2: Check focus at alert trigger time, not continuously

Query the active pane only inside `triggerAlerts()` (or just before calling it), rather than polling it every tick.

**Why:** Busy→Waiting transitions are infrequent events. Querying tmux only when an alert would fire keeps overhead minimal — one extra tmux command per transition instead of 10/second.

### Decision 3: Add `GetActivePaneID()` to the tmux package

Add a new function in `internal/tmux/` that returns the currently focused pane ID. This keeps tmux interaction consolidated in one package.

### Decision 4: Guard in the update loop, not inside `triggerAlerts()`

Add the focus check at the call site in the update loop (around line 512 in app.go) rather than inside `triggerAlerts()`. This keeps `triggerAlerts()` as a pure "fire alerts" function and makes the suppression logic visible at the orchestration level.

## Risks / Trade-offs

- **[Risk] tmux command overhead**: One extra `tmux display-message` call per Busy→Waiting transition. → Mitigation: These transitions are rare (seconds/minutes apart), so overhead is negligible.
- **[Risk] Multi-client tmux sessions**: If multiple clients are attached, `display-message` returns the active pane of the most recently active client. → Mitigation: This is the correct behavior — if any client has the pane focused, suppress alerts.
- **[Trade-off] No queued notifications**: If a pane becomes ready while focused and the user later switches away, they won't get a retroactive notification. → Acceptable: The user saw the pane become ready since they were looking at it.
