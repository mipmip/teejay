## Context

Bubbletea supports mouse events through `tea.MouseMsg`. The bubbles `list.Model` component already has built-in mouse support that is automatically enabled when the program has mouse mode enabled. Currently tmon doesn't enable mouse mode.

## Goals / Non-Goals

**Goals:**
- Enable mouse support in the Bubbletea program
- Clicking items in lists selects them
- Maintain all existing keyboard navigation

**Non-Goals:**
- Drag and drop functionality
- Mouse hover effects
- Custom click handling beyond list selection (bubbles handles this)

## Decisions

### Decision 1: Use tea.WithMouseCellMotion()

**Choice:** Enable mouse with `tea.WithMouseCellMotion()` option when creating the program.

**Rationale:**
- This is the standard Bubbletea way to enable mouse support
- `MouseCellMotion` provides click events which is what we need
- The bubbles list component will automatically handle mouse clicks for selection
- Alternative: `WithMouseAllMotion` - provides more events but unnecessary for our use case

### Decision 2: Let bubbles handle mouse events

**Choice:** Pass mouse events through to the list model via its Update() method without custom handling.

**Rationale:**
- The `list.Model` from bubbles already implements mouse click selection
- No custom code needed - just enable mouse mode and forward events
- Keeps implementation minimal and consistent with bubbles behavior

## Risks / Trade-offs

**[Risk] Terminal compatibility** → Some terminals may not support mouse events. Mitigation: This is optional UX enhancement; keyboard navigation remains fully functional.

**[Trade-off] Mouse capture** → Enabling mouse mode captures mouse events that would otherwise go to the terminal. Users who want to select text with mouse will need to hold Shift. This is standard behavior for TUI apps.
