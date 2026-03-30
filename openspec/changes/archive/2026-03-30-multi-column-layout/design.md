## Context

The current `View()` method in `app.go` renders a fixed two-panel layout: 30% list + 70% preview. On narrow terminals (<25 char sidebar), it falls back to a single full-width list. The `browserItemDelegate` renders each pane item at a height of 5 lines with a minimum usable width of ~30 characters. The main watchlist uses the same delegate and list model from `bubbles/list`.

## Goals / Non-Goals

**Goals:**
- Add an alternative layout that hides the preview and shows pane items in multiple columns
- Calculate column count dynamically from terminal width and minimum column width (30 chars)
- Toggle between layouts with a single keybind

**Non-Goals:**
- Persisting layout preference across app restarts
- Changing pane-item delegate rendering or sizing
- Custom column widths or user-configurable minimum widths

## Decisions

### Layout mode as Model state

Add a `layoutMode` field to `Model` (e.g., an int/enum: `layoutDefault = 0`, `layoutMultiColumn = 1`). This keeps the toggle simple and avoids config file changes.

**Alternative considered**: Storing in config file — rejected because this is a view preference that doesn't need persistence yet.

### Multi-column rendering approach

Rather than using multiple `list.Model` instances side-by-side (which would complicate keyboard navigation), render the multi-column view manually:

1. Calculate `numColumns = max(1, (availableWidth) / minColumnWidth)` where `minColumnWidth = 30`
2. Distribute the actual width evenly: `colWidth = availableWidth / numColumns`
3. Slice the full item list into column-sized chunks (items fill top-to-bottom per column, then next column)
4. Render each column of items using the existing `browserItemDelegate.Render()` pattern
5. Join columns horizontally with `lipgloss.JoinHorizontal`

The selected item highlight needs to work across columns — maintain the single `m.list` for state/navigation, but render its items across columns in the View.

**Alternative considered**: Multiple `list.Model` instances — rejected because it would require complex focus management and synchronized selection state.

### Keybinding: `v` to toggle layout

Use `v` as the toggle key — intuitive for "switch **v**iew", single-press, and unbound.

**Alternative considered**: `tab`/`shift+tab` — reserved for future use (field focus, etc.). `l` for "layout" — conflicts with potential vim-style navigation.

### Navigation in multi-column mode

Keep using the existing single `list.Model` for item state and selection index. In multi-column mode, intercept arrow keys to provide spatial navigation: up/down moves within a column, left/right jumps between columns (same row). This requires translating between the flat list index and a (column, row) position based on the current column count and items-per-column.

## Risks / Trade-offs

- **[Low] Delegate rendering coupling** — We reuse the delegate's render logic outside the list widget. If the delegate changes, the multi-column renderer needs updating too. → Mitigated by keeping the rendering in the same file.
- **[Low] Selection visibility** — The selected item might be in a non-obvious column position. → Mitigated by the highlight styling making selection clear regardless of position.
