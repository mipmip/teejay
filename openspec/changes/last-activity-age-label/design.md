## Context

Each pane already tracks `lastActivity` (a `time.Time` from `monitor.LastChangeTime()`). The browser item delegate renders a title row with pane name on the left and a status indicator on the right. The recency color gradient provides visual feedback but lacks precision.

## Goals / Non-Goals

**Goals:**
- Show compact elapsed-time text ("3s", "14m", "2h") on the title row, right-aligned before the status indicator
- Only for waiting panes — busy panes are actively changing

**Non-Goals:**
- Replacing the recency color gradient
- Adding a config toggle for this feature (always shown)
- Showing age in the multi-column layout (too tight on space)

## Decisions

### 1. Compact duration format without libraries

Use a simple helper function that picks the single largest unit:
- `<60s` → "3s"
- `<60m` → "14m"
- `<24h` → "2h"
- `≥24h` → "1d"

No need for a library — this is ~15 lines of Go. Similar approach to how `recencyColor` already uses elapsed duration thresholds.

### 2. Positioned on title row, right-aligned before indicator

The label sits between the pane name and the status dot. Styled dim so it doesn't compete with the pane name or indicator for attention.

```
  claude-frontend       3s ●
  api-refactor         14m ●
  data-pipeline            ◍   ← busy, no label
```

### 3. Only in list delegate, not column delegate

The multi-column layout has minimal horizontal space per cell. Adding the age label there would crowd the content. List view has plenty of room.

## Risks / Trade-offs

- [Visual noise] Adding more text to each row → Mitigated by dim styling and only showing for waiting panes.
- [Stale display] The label updates each render tick, which is already driven by the monitor poll interval (~1s). Good enough granularity.
