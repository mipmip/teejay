## Context

tmon is a TUI app that monitors tmux panes. Users can view a list of watched panes and preview their content in real-time. Currently, there's no way to switch to a previewed pane from within tmon - users must manually navigate tmux to find and switch to the pane they're viewing.

The pane IDs stored in the watchlist are tmux pane identifiers (e.g., `%5`, `session:window.pane`), which can be used with `tmux select-pane` and `tmux switch-client` commands.

## Goals / Non-Goals

**Goals:**
- Allow users to press Enter to switch to the currently previewed pane
- Handle both cases: running inside tmux (can switch) and outside tmux (show message)
- Exit tmon after switching so the user lands in the target pane

**Non-Goals:**
- Supporting switching without exiting tmon (would require complex state management)
- Keyboard customization for the switch key

## Decisions

### Decision 1: Use `tmux switch-client` for switching

**Choice:** Use `tmux switch-client -t <pane-id>` to switch to the target pane.

**Rationale:**
- `switch-client` changes the current client's active pane, which is what users expect
- Works with full pane IDs (session:window.pane format or %N format)
- Alternative `select-pane` only works within the same window

### Decision 2: Exit tmon after switching

**Choice:** After issuing the switch command, exit the tmon application.

**Rationale:**
- The user's intent is to interact with the target pane
- Keeping tmon running would show stale preview (target pane is now active)
- Simple UX: Enter = "go to this pane"

### Decision 3: Detect tmux environment via TMUX env var

**Choice:** Check for `TMUX` environment variable to determine if running inside tmux.

**Rationale:**
- Standard method used by tmux itself
- If empty/unset, we're not in tmux and cannot switch
- Alternative: try the command and handle error - but checking env var gives better UX

## Risks / Trade-offs

**[Risk] User accidentally switches** → The Enter key is a common key. Mitigation: Enter is intuitive for "select/go" actions, and the pane being switched to is clearly previewed. This matches user expectations.

**[Trade-off] Exit on switch** → User loses tmon state. Acceptable because switching indicates intent to work in that pane, and tmon can be restarted easily.
