## Context

The pane browser is triggered by pressing `a` in the main view. Currently it loads all tmux panes, filters out already-watched ones, and displays them in a flat `list.Model`. Users with many sessions see a long mixed list like "dev:0.0", "server:1.2", "dev:0.1" which requires scanning to find related panes.

## Goals / Non-Goals

**Goals:**
- Two-step selection: session list → pane list for selected session
- Escape navigates backward through the hierarchy
- Maintain current behavior for adding pane to watchlist

**Non-Goals:**
- Search/filter within session or pane lists
- Remembering last selected session between browser opens
- Showing pane preview in browser popup

## Decisions

### 1. State machine approach

Add a `browsingSession` boolean to track which level we're at:
- `true`: showing session list
- `false`: showing pane list for `selectedSession`

**Rationale**: Simple boolean is sufficient for two levels. Avoids over-engineering with enum or nested state.

### 2. Reuse single `browserList` for both views

Switch the contents of `browserList` between session items and pane items rather than maintaining two separate list models.

**Rationale**: Reduces state complexity. The popup only shows one list at a time anyway.

### 3. Cache all panes upfront

Store filtered panes in `allBrowserPanes` when browser opens, then filter by session as needed.

**Rationale**: Avoids repeated tmux calls. Pane list won't change while browser is open.

### 4. Session item displays pane count

Show "session-name" with description "N pane(s)" to help users know what to expect.

**Rationale**: Provides useful context without cluttering the UI.

### 5. Pane item displays window.pane and command

Format: "0.1 zsh" with description showing pane ID.

**Rationale**: Window/pane index is how tmux users think about panes. Command helps identify what's running.

## Risks / Trade-offs

- **Extra click required**: Users now need two selections instead of one. Acceptable trade-off for better organization with many panes.
- **Empty session edge case**: If all panes in a session are already watched, user sees empty pane list. Mitigated by showing helpful message.
