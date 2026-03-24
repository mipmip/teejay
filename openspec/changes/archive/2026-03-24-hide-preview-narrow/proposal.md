## Why

On narrow terminals the 30/70 split between the watchlist sidebar and preview panel squeezes the sidebar items to an unreadable width. When the sidebar would be narrower than ~25 characters, the preview panel should be hidden entirely so the sidebar can use the full terminal width.

## What Changes

- Add a responsive breakpoint to the main View() layout: when the sidebar width at 30% would be less than ~25 chars, hide the preview panel and give the sidebar the full width
- Apply the same logic to the WindowSizeMsg handler and mouse click width calculation so all layout paths are consistent

## Capabilities

### New Capabilities
_None_

### Modified Capabilities
- `pane-preview`: The preview panel is hidden on narrow terminals, giving full width to the watchlist sidebar

## Impact

- `internal/ui/app.go`: View() layout calculation, WindowSizeMsg handler, mouse click width calculation
