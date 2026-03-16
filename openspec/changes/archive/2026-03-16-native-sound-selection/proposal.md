## Why

The current terminal bell (`\a`) sound is inconsistent across systems and often silent. Users need reliable, audible alerts when panes become ready. A native Go audio implementation with 5 distinct selectable sounds provides consistent, cross-platform audio without external dependencies.

## What Changes

- Replace terminal bell with native Go audio playback using embedded WAV sounds
- Add 5 built-in sound options: chime, bell, ping, pop, ding
- Add `sound_type` config option (global in config.yaml, per-pane in watchlist)
- Default sound type: "chime"
- Sounds are embedded in the binary (no external files needed)

## Capabilities

### New Capabilities

- `native-sounds`: Native Go audio playback with 5 selectable notification sounds

### Modified Capabilities

- `pane-alerts`: Add sound_type selection to alert configuration

## Impact

- `internal/alerts/alerts.go`: Replace `PlayBell()` with `PlaySound(soundType string)`
- `internal/alerts/sounds/`: New package with embedded WAV files
- `internal/config/config.go`: Add `SoundType` field to Alerts struct
- `internal/watchlist/watchlist.go`: Add `SoundType` field to WatchedPane
- `internal/ui/app.go`: Update configure popup to allow sound type selection
- New dependency: Go audio library (e.g., `github.com/gopxl/beep` or similar)
- Binary size increase: ~50-100KB for embedded sounds
