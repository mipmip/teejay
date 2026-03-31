## Context

The monitor already tracks `lastChangeTime` per pane in `paneState` but doesn't expose it. The UI tracks `paneStatuses` (Busy/Waiting) but not when transitions happened. The existing `browserItemDelegate.Render()` colors the `●` indicator with a fixed `#00FF00` for all waiting panes.

## Goals / Non-Goals

**Goals:**
- Expose last-activity time from the monitor
- Color the waiting indicator with a recency gradient
- Toggle between watchlist order and activity-sorted order with `o`
- Configurable defaults via config file

**Non-Goals:**
- Showing exact timestamps or "2m ago" text in the UI
- Persistent sort preference across restarts (runtime toggle only, config sets default)

## Decisions

### Expose `LastChangeTime` from monitor

Add a `LastChangeTime(paneID string) time.Time` method to `Monitor`. This returns `paneState.lastChangeTime` — the last time content changed for this pane. This is cheap (map lookup) and doesn't change the monitoring logic.

### Recency color gradient

Map elapsed time since last activity to a green intensity:

| Elapsed | Color | Description |
|---------|-------|-------------|
| 0-10s | `#00FF00` | Bright neon green — just finished |
| 10-30s | `#00DD00` | Bright green |
| 30s-2min | `#00BB00` | Medium green |
| 2-5min | `#009900` | Dimmer green |
| 5min+ | `#006600` | Dim green — been waiting a while |

This applies to the `●` indicator only (not `?` which stays yellow, not spinners which stay default). The gradient is calculated at render time from `time.Since(lastActivity)`.

**Alternative considered**: Applying the gradient to the entire row background — rejected as too visually noisy. The indicator color is subtle enough.

### Activity sort order

When activity sort is enabled:
1. **Busy panes first** — sorted by last content change (most recent first)
2. **Waiting panes second** — sorted by last content change (most recent first)

This puts actively working panes at the top (they're the "hottest"), then recently finished panes next (they might need input), then long-idle panes at the bottom.

The sort is applied in `refreshListWithFrame()` before calling `m.list.SetItems()`. A `sortByActivity` bool on the Model controls this, toggled by `o`.

### Config structure

Add a `Display` section to the config:

```yaml
display:
  recency_color: true
  sort_by_activity: false
```

Both default to their indicated values. The runtime `o` toggle overrides `sort_by_activity` for the current session.

## Risks / Trade-offs

- **[Low] Color perception** — The gradient relies on green intensity which may not be distinguishable on all terminals or for color-blind users. Mitigated: it's supplementary information, not the only signal.
- **[Low] Sort instability** — When multiple panes are busy, their order might shuffle as content changes. Mitigated: busy panes change content frequently but the sort only runs on list refresh (~100ms), and the visual difference between adjacent busy panes is minimal.
