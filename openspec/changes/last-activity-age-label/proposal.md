## Why

The activity sort order is not trustworthy at a glance — the recency color gradient is too subtle to distinguish "3 seconds ago" from "30 seconds ago." A compact textual age label (e.g., "3s", "14m", "2h") next to each waiting pane makes the sort order verifiable and gives precise timing information.

## What Changes

- Show a human-readable "time since last activity" label on the title row of each waiting pane, positioned just before the status indicator
- Format as compact duration: "3s", "14m", "2h", "1d"
- Only displayed for waiting panes (busy panes are actively changing, so it would flicker)
- Uses the existing `lastActivity` timestamp already tracked by the monitor

## Capabilities

### New Capabilities
- `activity-age-label`: Compact elapsed-time label shown per waiting pane in the browser list

### Modified Capabilities

(none)

## Impact

- Render changes in the browser item delegate (`internal/ui/app.go`)
- New duration formatting helper (compact "3s"/"14m" style)
- No new dependencies, data model changes, or config options
