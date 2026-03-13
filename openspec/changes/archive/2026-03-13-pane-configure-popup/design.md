## Context

The TUI currently has modal states for editing (e key) and deleting (d key). A configure popup would be another modal that provides multiple settings at once. The watchlist stores panes with ID, Name, and AddedAt - we need to extend this with alert settings. There's already a browser popup pattern (`browsing` state) that can be followed.

## Goals / Non-Goals

**Goals:**
- Centralized configuration popup for each pane
- Per-pane toggles for sound and notification alerts
- Ability to edit name from within the popup
- Trigger alerts when pane transitions to Ready status

**Non-Goals:**
- Global alert settings (this is per-pane only)
- Custom sound file selection (use system bell)
- Custom notification text (use default message)

## Decisions

### Decision 1: Use system bell for sound alerts

**Choice:** Use terminal bell character (`\a` / `\007`) for sound alerts.

**Rationale:**
- No external dependencies needed
- Works in any terminal that supports bell
- Alternative: External sound library - adds complexity and dependencies

### Decision 2: Use os/exec with notify-send for notifications

**Choice:** Shell out to `notify-send` on Linux for desktop notifications.

**Rationale:**
- Simple, no CGO required
- Standard on most Linux systems
- Alternative: D-Bus library - more complex, CGO dependency
- Future: Can add macOS `osascript` support if needed

### Decision 3: Store alert settings in watchlist.json

**Choice:** Add `sound_on_ready` and `notify_on_ready` boolean fields to Pane struct.

**Rationale:**
- Keeps all pane config in one place
- Persists across sessions
- No additional config file needed

### Decision 4: Configure popup with selectable menu items

**Choice:** Use a simple menu in the popup with items: Name, Sound, Notify. Arrow keys to navigate, Enter to toggle/edit.

**Rationale:**
- Consistent with existing TUI patterns
- Clear visual indication of current settings
- Alternative: Inline toggles - harder to implement, less clear

## Risks / Trade-offs

**[Risk] notify-send not available** → Mitigation: Fail silently with log message if command not found. Consider adding a check and showing message in configure popup.

**[Trade-off] Terminal bell may be annoying** → Users can disable in their terminal settings. This is acceptable as it's opt-in per pane.

**[Risk] Rapid status changes cause notification spam** → Mitigation: Only trigger on transition TO Ready, not while staying Ready.
