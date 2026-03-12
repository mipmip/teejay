## Why

Currently, `tmon add` allows adding the same pane multiple times, resulting in duplicates in the watchlist. Additionally, when the TUI loads, these duplicates appear as separate entries. This clutters the UI and wastes user effort navigating duplicate entries.

## What Changes

- Prevent `tmon add` from adding a pane that's already in the watchlist
- Remove duplicate pane entries when loading the watchlist (cleanup existing data)
- Show appropriate message when user tries to add an existing pane

## Capabilities

### New Capabilities
- `pane-deduplication`: Prevent duplicate panes in watchlist, both when adding and when loading

### Modified Capabilities
<!-- None - this is a new capability for data integrity -->

## Impact

- `internal/watchlist/watchlist.go`: Add `Contains()` method, dedupe on `Load()`
- `internal/cmd/add.go`: Check for existing pane before adding
