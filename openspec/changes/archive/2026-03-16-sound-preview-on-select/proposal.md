## Why

When selecting a sound type in the pane configuration menu, users have no way to hear what the sound actually sounds like without waiting for a pane to become ready. Playing a preview when cycling through sound types provides immediate feedback and makes selection intuitive.

## What Changes

- Play the selected sound when cycling through sound types in the pane configuration menu
- The preview plays immediately after each selection change

## Capabilities

### New Capabilities

- `sound-preview`: Play sound preview when selecting sound type in pane configuration

### Modified Capabilities

None

## Impact

- Modified: `internal/ui/app.go` - add `alerts.PlaySound()` call when cycling sound type in config menu
- No new dependencies (uses existing alerts package)
