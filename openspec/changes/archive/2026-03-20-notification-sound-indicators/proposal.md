## Why

Users have no visual feedback about which notification and sound settings are active — neither globally nor per-pane. When managing many parallel AI agent sessions, it's unclear which panes will alert you and which won't. Showing modest colored symbols makes the alert configuration visible at a glance. Relates to issue #23.

## What Changes

- Add a global alert status indicator in the branding footer area (bottom right, next to "Terminal Junkie") showing the current global sound and notification settings — this acts as the application "systray"
- Add per-pane alert indicators on watchlist items when a pane has overridden the global defaults
- Use small colored symbols: one for sound, one for notifications
  - Enabled: colored symbol (e.g., 🔔 or ♪ for sound, 📢 or ✉ for notification)
  - Disabled/off: dimmed or absent

## Capabilities

### New Capabilities
- `alert-status-indicators`: Visual indicators showing sound and notification configuration state, both globally in the footer and per-pane when overridden

### Modified Capabilities
- `branding-footer`: The branding footer area gains a global alert status display next to "Terminal Junkie"
- `watchlist-item-delegate`: Pane items show override indicators when per-pane alert settings differ from global defaults

## Impact

- `internal/ui/app.go`: renderBrandingFooter() gains alert status symbols; paneItem rendering gains per-pane override indicators
- `internal/watchlist/watchlist.go`: Need to query effective alert state per pane to determine if override is active
- No new dependencies — purely UI/rendering changes using existing config and watchlist data
