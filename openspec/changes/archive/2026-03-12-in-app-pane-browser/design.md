## Context

The TUI currently requires users to run `tmon add` from within the target pane to add it to the watchlist. This is inconvenient when monitoring multiple panes. Users want to browse available tmux panes and add them without leaving the TUI.

Tmux provides commands to list sessions, windows, and panes:
- `tmux list-sessions` - lists all sessions
- `tmux list-windows -t <session>` - lists windows in a session
- `tmux list-panes -t <window>` - lists panes in a window
- `tmux list-panes -a` - lists all panes across all sessions

## Goals / Non-Goals

**Goals:**
- Show popup overlay when pressing `a` key
- Display flat list of all panes with session/window context
- Navigate with arrow keys, select with Enter
- Add selected pane to watchlist and close popup
- Cancel with Escape

**Non-Goals:**
- Hierarchical tree navigation (flat list is simpler)
- Pane content preview in browser (adds complexity)
- Multi-select (add one pane at a time)
- Filtering/search (can add later)

## Decisions

### Decision 1: Flat list of all panes

Use `tmux list-panes -a -F "#{pane_id} #{session_name}:#{window_index}.#{pane_index} #{pane_current_command}"` to get all panes in one call.

**Rationale**: Simpler than hierarchical tree navigation. Most users have few enough panes that a flat list works well. Format string provides context (session:window.pane) and current command.

**Alternatives considered**:
- Hierarchical tree: More complex UI, unnecessary for typical use
- Multiple tmux calls: Slower, more complex

### Decision 2: Modal popup using overlay rendering

Add `browsing bool` state to Model. When true, render a centered popup over the main UI. Use lipgloss for styling and positioning.

**Rationale**: Bubbletea doesn't have built-in modals, but overlay rendering with lipgloss.Place() works well. Similar pattern to delete confirmation but larger.

### Decision 3: Reuse bubbles/list for pane selection

Use the same list component as the main pane list, but with pane browser items.

**Rationale**: Consistent UX, keyboard navigation already handled.

### Decision 4: Skip already-watched panes

Filter out panes that are already in the watchlist from the browser list.

**Rationale**: Prevents confusion and duplicate add attempts.

## Risks / Trade-offs

- [Stale pane list] → Re-fetch on each popup open; panes may change during browsing but acceptable
- [Large number of panes] → Flat list may be unwieldy with 50+ panes; add filtering later if needed
- [Pane identification] → Show command helps identify panes; could add preview later
