## Context

The multi-column layout (`renderMultiColumnLayout()`) currently fills the full terminal width with pane item columns but ignores vertical space. Each pane item is 4 lines tall (blank + title + desc + blank). With few panes or many columns, the grid may only use 20-30% of the terminal height, leaving the rest empty.

The existing preview infrastructure (`m.viewport`, `m.previewContent`, `m.previewErr`) is already maintained regardless of layout mode — content is captured every 100ms tick. Only the rendering is skipped in multi-column mode.

## Goals / Non-Goals

**Goals:**
- Show the preview panel below the multi-column grid when vertical space allows
- Reuse the existing viewport and preview content — no new data fetching
- Automatically adapt when the terminal is resized

**Non-Goals:**
- Making the preview height user-configurable
- Adding a keybind to toggle the below-preview independently
- Changing the default layout's side-by-side preview

## Decisions

### Height calculation approach

Calculate available space after rendering the column grid:

```
totalHeight = m.height
footerHeight = 2 (help + branding line)
gridHeight = itemsPerCol * 4 (each item is 4 lines)
remainingHeight = totalHeight - gridHeight - footerHeight
```

Show the preview below if `remainingHeight >= 8` (enough for border + title + a few content lines). This threshold keeps the preview useful — anything shorter would just be noise.

**Alternative considered**: Fixed split (e.g., 60% grid, 40% preview) — rejected because it would shrink the grid even when there are many panes, defeating the purpose of multi-column mode.

### Rendering approach

`renderMultiColumnLayout()` currently returns just the column grid string. Extend it to:

1. Calculate grid height and remaining space
2. If enough space, render the preview panel below at full width with the remaining height
3. Join vertically: columns on top, preview below
4. Set viewport dimensions to match the available preview area

The preview panel uses the same `previewPanelStyle` (rounded border, gray) and `previewTitleStyle` as the default layout, maintaining visual consistency.

### Viewport height adjustment

The viewport height needs to be set correctly for the below-preview scenario. This is done in the render path since the available height depends on the number of items and columns — which can change dynamically. The viewport's `Height` is set before calling `m.viewport.View()`.

## Risks / Trade-offs

- **[Low] Preview flicker on resize** — When the terminal is resized across the threshold, the preview appears/disappears. This is consistent with how the default layout handles narrow terminals.
- **[Low] Viewport height mismatch** — Setting viewport height in the render path (View) rather than Update is slightly unusual for Bubbletea, but since we're only reading the viewport (not mutating Model state), this is safe with a value receiver.
