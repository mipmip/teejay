## Context

The current pane switch behavior in `internal/ui/app.go` calls `tea.Quit` immediately after `tmux.SwitchToPane()`. This was likely implemented as a quick workflow: switch and exit. However, for a monitoring application, users need to stay in tmon to continue watching panes and switch between them multiple times.

## Goals / Non-Goals

**Goals:**
- App stays open after switching to a pane via Enter key
- User can continue navigating and switching to other panes
- Only explicit quit (q, ctrl+c) exits the application

**Non-Goals:**
- Changing tmux focus behavior (that stays as-is)
- Adding new keybindings or UI elements

## Decisions

### Decision 1: Remove tea.Quit after successful switch

**Choice:** Simply remove the `return m, tea.Quit` after `tmux.SwitchToPane()` and return `m, nil` instead.

**Rationale:**
- Minimal change - just one line modification
- The switch already happens before the return, so removing Quit is sufficient
- Alternative: Add a user setting for "quit after switch" - unnecessary complexity for now

## Risks / Trade-offs

**[Trade-off] User must explicitly quit** → Users who prefer the old behavior will need an extra keypress (q). This is acceptable as the monitoring use case benefits more from staying open.
