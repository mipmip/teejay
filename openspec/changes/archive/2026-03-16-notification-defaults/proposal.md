## Why

Currently, notification and sound alerts must be enabled individually for each pane. Users who want alerts for all their monitored panes must toggle settings for every pane manually. There's no way to set sensible defaults that apply to new panes.

## What Changes

- Add global default settings for `sound_on_ready` and `notify_on_ready` in config.yaml
- Per-pane settings in watchlist.json override the global defaults
- New panes inherit the global defaults when added to the watchlist

## Capabilities

### New Capabilities

- `notification-defaults`: Global default settings for notification and sound alerts, configurable via config.yaml with per-pane overrides

### Modified Capabilities

None

## Impact

- Modified: `internal/config/config.go` - add Alerts section with default settings
- Modified: `internal/ui/app.go` - use global defaults when pane has no explicit setting
- Modified: `internal/watchlist/watchlist.go` - may need tri-state for override logic
- Modified: `config.example.yaml` - document new settings
- Modified: `README.md` - document new config options
