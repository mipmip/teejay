## Why

In multi-column mode, pane items are cramped at 30 chars wide. Names and breadcrumbs truncate aggressively, making it hard to distinguish panes. When only a few panes are watched, the columns are still narrow even though there's plenty of screen space.

## What Changes

- Increase `minColWidth` from 30 to 45 for more readable columns
- Cap the number of columns to the number of items — no empty columns wasting space. With 4 panes on a 200-char terminal, get 4 × 50 instead of 6 × 33.

## Capabilities

### New Capabilities

_None_

### Modified Capabilities

- `multi-column-layout`: Wider minimum columns and adaptive column count

## Impact

- `internal/ui/app.go` — change `minColWidth` constant in `calcMultiColumn()` and cap `numColumns` to `totalItems`
- CHANGELOG.md
