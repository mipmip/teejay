## Context

The TUI currently displays watched panes in a list but provides no management capabilities. Users must manually edit `~/.config/tmon/watchlist.json` to rename or remove entries. The TUI uses Bubbletea with a split layout (list on left, preview on right).

## Goals / Non-Goals

**Goals:**
- Add `e` key binding to enter edit mode for renaming the selected pane
- Add `d` key binding to delete selected pane with confirmation
- Add optional `Name` field to Pane struct for custom display names
- Provide clear visual feedback during edit and delete operations

**Non-Goals:**
- Bulk operations (multi-select delete)
- Undo functionality
- Reordering panes

## Decisions

### Decision 1: Use bubbles/textinput for edit mode

Use the `textinput` component from bubbles library for inline name editing.

**Rationale**: Consistent with Bubbletea patterns, handles cursor, backspace, and input validation out of the box.

**Alternatives considered**: Raw key handling - rejected due to complexity of cursor management.

### Decision 2: Simple confirmation for delete

Show "Delete [pane]? (y/n)" message at the bottom of the screen.

**Rationale**: Simple and familiar pattern. No need for a modal dialog for single-item delete.

**Alternatives considered**: Modal dialog - overkill for this use case.

### Decision 3: Display name falls back to pane ID

If `Name` is empty, display the pane ID (e.g., `%65`). The name is optional.

**Rationale**: Backwards compatible with existing watchlists. Users can rename if they want friendlier names.

## Risks / Trade-offs

- [Trade-off] No undo for delete → Acceptable; users can re-add panes easily with `tmon add`
- [Risk] Edit mode captures all keys → Mitigation: Escape key always exits edit mode
