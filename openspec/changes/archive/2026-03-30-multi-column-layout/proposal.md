## Why

When monitoring many tmux panes, the current layout dedicates 70% of the screen to the preview panel, showing only a single column of pane items in the remaining 30%. Users who want an overview of all their panes at a glance have no way to trade preview space for more visible pane items.

## What Changes

- Add a new "multi-column" layout mode that hides the preview panel and fills the full width with multiple columns of pane items
- Columns are calculated dynamically based on terminal width and the existing minimum pane-item width (30 characters)
- Add a keybinding to toggle between the default (list + preview) layout and the multi-column layout
- Persist layout preference in the Model state (not across sessions)

## Capabilities

### New Capabilities

- `multi-column-layout`: A layout mode that replaces the preview panel with additional columns of pane items, with a keybind to toggle between layouts

### Modified Capabilities

- `pane-preview`: The preview panel can now be hidden via layout toggle (not just narrow terminals)

## Impact

- `internal/ui/app.go` — View() rendering, Model state, Update() keybinding handler, help text
- The existing `browserItemDelegate` and pane-item rendering remain unchanged
