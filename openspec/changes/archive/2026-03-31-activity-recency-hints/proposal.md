## Why

When monitoring many agent panes, it's hard to tell at a glance which ones recently became active or finished. All waiting panes look the same green `●` regardless of whether they finished 2 seconds or 20 minutes ago. Users need a visual signal of recency to prioritize attention, and optionally the ability to sort panes by last activity for a natural triage order.

## What Changes

- Track a `lastActivityTime` per pane (when the pane last transitioned state or had content changes)
- Add a recency color gradient on the waiting indicator: bright green for recently finished, fading to dim green over time
- Add a keybind (`o`) to toggle between watchlist order and activity-sorted order (busy panes first, then waiting panes sorted by most-recently-finished)
- Add config options: `display.recency_color` (bool, default true) and `display.sort_by_activity` (bool, default false) to set defaults

## Capabilities

### New Capabilities

- `activity-recency`: Tracks per-pane last-activity timestamps and provides recency-based color gradient and activity-sorted ordering

### Modified Capabilities

- `activity-detection`: Exposes last-activity timestamp from the monitor for use by the UI

## Impact

- `internal/monitor/monitor.go` — expose `LastChangeTime(paneID)` method
- `internal/ui/app.go` — track `lastActivityTime` per pane, recency color in delegate rendering, sort toggle with `o` keybind, help text update
- `internal/config/config.go` — add `Display` section with `RecencyColor` and `SortByActivity` fields
