## 1. Indicator styles

- [x] 1.1 Add lipgloss styles for alert indicator symbols: sound enabled (green), notification enabled (yellow), disabled (dim gray)
- [x] 1.2 Add a helper function `renderAlertIndicators(soundEnabled, notifyEnabled bool) string` that returns the styled `♪ ✉` string

## 2. Global systray in branding footer

- [x] 2.1 Update `renderBrandingFooter()` to append global alert status indicators after the version, reading `m.config.Alerts.SoundOnReady` and `m.config.Alerts.NotifyOnReady`

## 3. Per-pane override indicators

- [x] 3.1 Add `soundOverride *bool` and `notifyOverride *bool` fields to `paneItem` struct
- [x] 3.2 Update `Description()` on `paneItem` to append alert indicators when overrides are present
- [x] 3.3 Populate the override fields in `refreshListWithFrame()` from the watchlist pane data

## 4. Testing

- [x] 4.1 Add unit test for `renderAlertIndicators` with various enabled/disabled combinations
- [x] 4.2 Add unit test for `paneItem.Description()` with overrides present (shows indicators)
- [x] 4.3 Add unit test for `paneItem.Description()` with no overrides (no indicators)
- [x] 4.4 Run existing tests to verify no regressions
