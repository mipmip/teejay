## Why

The preview panel title currently shows just the pane ID (e.g., "Preview: %0"), which is not user-friendly. Users would benefit from seeing the actual pane name and description (like the running command) to better identify which pane they're previewing.

## What Changes

- Update the preview panel title to display the pane's display name instead of just the pane ID
- Include the current foreground command as a subtitle/description when available
- Format: "Preview: {name}" with optional command info

## Capabilities

### New Capabilities

(none - this is an enhancement to existing preview functionality)

### Modified Capabilities

- `pane-preview`: The preview title format is changing from showing pane ID to showing pane name and description

## Impact

- `internal/ui/app.go`: The `View()` method where `previewTitle` is rendered needs modification
- No API changes
- No breaking changes - purely UI improvement
