## Why

The browser popup currently uses the default bubbles list styling which doesn't provide clear visual separation between items. Adding background colors to each item (dark grey normally, lighter grey when selected) with margins between them improves visual clarity and makes it obvious what area is clickable.

## What Changes

- Replace default list delegate with custom delegate for browser items
- Each item rendered with a dark grey background
- Selected item rendered with a lighter grey background
- 1-line margin between items for visual separation
- Full item area (including background) is clickable for mouse selection
- Apply to both session list and pane list in browser popup

## Capabilities

### New Capabilities

- `browser-item-delegate`: Custom list delegate with background styling for browser popup items

### Modified Capabilities

(none - this is a visual enhancement, not a behavior change)

## Impact

- `internal/ui/app.go`: Create custom delegate type with styled rendering
- `internal/ui/app.go`: Update loadSessionList() and loadPaneListForSession() to use custom delegate
- Mouse click hit detection may need adjustment for new item heights
