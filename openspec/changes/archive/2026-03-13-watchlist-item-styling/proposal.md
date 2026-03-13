## Why

The main watchlist panel currently uses the default bubbles list styling which doesn't provide clear visual separation between items. Adding background colors to each item (dark grey normally, lighter grey when selected) with margins between them improves visual clarity and makes it obvious which item is selected and what area is clickable.

## What Changes

- Replace default list delegate with custom delegate for watchlist pane items
- Each item rendered with a dark grey background
- Selected item rendered with a lighter grey background
- 1-line margin between items for visual separation
- Full item area (including background) is clickable for mouse selection

## Capabilities

### New Capabilities

- `watchlist-item-delegate`: Custom list delegate with background styling for watched pane items in the main list

### Modified Capabilities

(none - this is a visual enhancement, not a behavior change)

## Impact

- `internal/ui/app.go`: Create custom delegate type with styled rendering (can reuse `browserItemDelegate` pattern)
- `internal/ui/app.go`: Update `New()` to use custom delegate instead of `list.NewDefaultDelegate()`
- Mouse click hit detection in main list may need adjustment for new item heights
