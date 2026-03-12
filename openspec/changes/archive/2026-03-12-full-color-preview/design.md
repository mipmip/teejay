## Context

The current `CapturePane` function uses `tmux capture-pane -p` which strips ANSI escape sequences. The claude-squad project (github.com/smtg-ai/claude-squad) demonstrates that adding `-e` flag preserves these sequences, and Bubbletea's viewport natively renders ANSI codes.

## Goals / Non-Goals

**Goals:**
- Preserve ANSI escape sequences when capturing pane content
- Display full-color terminal output in the preview pane
- Join wrapped lines for cleaner display

**Non-Goals:**
- Custom ANSI parsing or color manipulation
- Theme-based color remapping
- Capturing pane scrollback history (just visible content)

## Decisions

### Decision 1: Use `-e` flag for ANSI preservation

Add `-e` flag to `tmux capture-pane` command.

**Rationale**: This is the standard tmux approach. The flag tells tmux to include escape sequences in output. Bubbletea/viewport handles rendering these natively.

**Alternatives considered**: Manual ANSI stripping and re-colorization - rejected as unnecessarily complex.

### Decision 2: Add `-J` flag for line joining

Add `-J` flag to join wrapped lines.

**Rationale**: Prevents awkward line breaks in the preview when content wraps in the source pane.

## Risks / Trade-offs

- [Trade-off] More data transferred per capture → Acceptable; ANSI sequences add minimal overhead
- [Risk] Some terminals may not render all ANSI codes → Mitigation: Bubbletea handles most common codes; edge cases are acceptable
