## Why

Users need a centralized way to configure per-pane settings like notifications and sounds. Currently editing a name requires pressing 'e', but there's no way to configure alerts when a pane becomes ready. A configure popup (key 'c') provides a single place for all pane settings.

## What Changes

- Add 'c' key to open a configure popup for the selected pane
- Configure popup shows options:
  - Edit name (text input)
  - Play sound when ready (toggle)
  - Send notification when ready (toggle)
- Settings are stored per-pane in the watchlist
- When a pane transitions to Ready status, trigger configured alerts

## Capabilities

### New Capabilities

- `pane-configure`: Configure popup for per-pane settings (name, sound, notification toggles)
- `pane-alerts`: Sound and notification alerts when pane becomes ready

### Modified Capabilities

(none - new capabilities integrate with existing status monitoring)

## Impact

- `internal/watchlist/watchlist.go`: Add SoundOnReady and NotifyOnReady fields to Pane struct
- `internal/ui/app.go`: Add configure popup mode and 'c' key handler
- `internal/alerts/`: New package for sound playback and desktop notifications
- Help footer needs to include 'c' keybinding
