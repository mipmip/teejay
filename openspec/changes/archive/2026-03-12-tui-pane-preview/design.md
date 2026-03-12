## Context

The TUI currently shows a full-width pane list. We need to add a preview panel that shows the actual content of the selected tmux pane. This requires:
1. Splitting the layout into left (list) and right (preview) panels
2. Capturing pane content from tmux
3. Displaying it in a scrollable viewport

## Goals / Non-Goals

**Goals:**
- Split-panel layout with list on left (~30% width), preview on right (~70%)
- Capture pane content via `tmux capture-pane -p -t <pane-id>`
- Display content in bubbles/viewport for scrolling
- Update preview when selection changes

**Non-Goals:**
- Live/auto-refresh of pane content (manual refresh or future change)
- Sending input to the pane (readonly only)
- ANSI color rendering (display raw text for now)

## Decisions

### Decision 1: Use `tmux capture-pane -p` for content capture

Run `tmux capture-pane -p -t <pane-id>` to get pane content as text.

**Rationale**: Simple, reliable, built-in tmux feature. The `-p` flag outputs to stdout instead of a buffer.

**Alternative considered**: Using tmux control mode. Rejected as overkill for readonly capture.

### Decision 2: Use bubbles/viewport for preview display

Use `github.com/charmbracelet/bubbles/viewport` for the scrollable preview panel.

**Rationale**: Handles scrolling, line wrapping, and keyboard navigation. Consistent with our use of bubbles/list.

### Decision 3: Lipgloss for split-panel layout

Use `lipgloss.JoinHorizontal` to create the split layout.

**Rationale**: lipgloss provides layout primitives. No need for a separate layout library.

### Decision 4: Create internal/tmux package

Put tmux interaction in `internal/tmux/` to keep it separate from UI code.

**Rationale**: Clean separation of concerns. Easy to test tmux functions independently.

### Decision 5: Fetch content on selection change only

Don't auto-refresh. Capture content when user changes selection.

**Rationale**: Simplicity. Auto-refresh would require timers/goroutines. Can add later.

## Risks / Trade-offs

- [Risk] Pane no longer exists when capturing → Show error message in preview
- [Risk] Large pane content is slow → `capture-pane` is fast; viewport handles large content
- [Trade-off] No ANSI colors → Acceptable for MVP; can add terminal rendering later
- [Trade-off] No auto-refresh → Users must navigate away and back to refresh
