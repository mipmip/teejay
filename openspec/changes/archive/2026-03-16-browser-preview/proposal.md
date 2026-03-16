## Why

When adding panes via the pane browser, users can only see pane IDs and running commands. This makes it difficult to identify which pane they want to add, especially when multiple panes run similar processes. A preview window showing the actual pane content would help users make better selections.

## What Changes

- Add a preview panel to the pane browser popup showing the currently selected pane's content
- Preview updates as the user navigates through the pane list
- Preview only shown when viewing panes (not during session selection)
- Browser popup layout changes from single list to split view (list + preview)

## Capabilities

### New Capabilities

None - this enhances an existing capability.

### Modified Capabilities

- `pane-browser`: Adding preview panel to show selected pane content while browsing

## Impact

- `internal/ui/app.go`: Modify `renderBrowserPopup()` to include preview panel, add preview state tracking during browsing, capture pane content on selection change
- Layout change: Browser popup becomes wider to accommodate preview panel
- No API changes
- No breaking changes
