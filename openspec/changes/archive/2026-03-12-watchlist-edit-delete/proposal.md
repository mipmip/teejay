## Why

The TUI currently displays watched panes but offers no way to manage them. Users cannot rename entries for clarity or remove panes they no longer want to watch without manually editing the JSON file.

## What Changes

- Add keyboard shortcut `e` to edit the selected pane's display name
- Add keyboard shortcut `d` to delete the selected pane from the watchlist
- Add name field to pane entries for custom display names
- Show confirmation before deleting a pane

## Capabilities

### New Capabilities
- `watchlist-management`: Edit and delete pane entries from the TUI via keyboard shortcuts

### Modified Capabilities
<!-- None - watchlist data structure will get a new optional field but no spec-level behavior changes -->

## Impact

- `internal/watchlist/watchlist.go`: Add `Name` field to Pane struct, add `Remove()` and `Rename()` methods
- `internal/ui/app.go`: Handle `e` and `d` key bindings, add edit mode and delete confirmation
