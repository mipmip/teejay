## 1. Monitor: Expose Last Activity Time

- [x] 1.1 Add `LastChangeTime(paneID string) time.Time` method to `Monitor` that returns `paneState.lastChangeTime` (zero time if pane unknown)
- [x] 1.2 Add tests for `LastChangeTime` — known pane returns correct time, unknown pane returns zero time

## 2. Config: Display Section

- [x] 2.1 Add `Display` struct with `RecencyColor bool` and `SortByActivity bool` to config, with defaults `true` and `false`
- [x] 2.2 Add `display` section to `configFile` YAML parsing
- [x] 2.3 Add test for default config values: `RecencyColor` is true, `SortByActivity` is false
- [x] 2.4 Add test for YAML parsing of display section with non-default values

## 3. UI: Track Last Activity Per Pane

- [x] 3.1 Add `lastActivity time.Time` field to `paneItem` struct
- [x] 3.2 In `refreshListWithFrame()`, populate `lastActivity` from `m.monitor.LastChangeTime(p.ID)`

## 4. UI: Recency Color Gradient

- [x] 4.1 Create `recencyColor(elapsed time.Duration) lipgloss.Color` helper that maps elapsed time to green intensity (5 tiers: #00FF00, #00DD00, #00BB00, #009900, #006600)
- [x] 4.2 Add tests for `recencyColor` — verify each tier boundary returns the expected color
- [x] 4.3 In `browserItemDelegate.Render()`, replace fixed `#00FF00` for waiting `●` indicator with `recencyColor(time.Since(p.lastActivity))` when config `RecencyColor` is true
- [x] 4.4 Apply same recency color in `renderMultiColumnItem()`

## 5. UI: Activity Sort

- [x] 5.1 Add `sortByActivity bool` field to `Model`, initialized from `config.Display.SortByActivity`
- [x] 5.2 Add `o` key handler in `Update()` to toggle `sortByActivity`
- [x] 5.3 In `refreshListWithFrame()`, when `sortByActivity` is true, sort items: busy panes first (by lastActivity desc), then waiting panes (by lastActivity desc)
- [x] 5.4 Update help footer text to include `o: order` keybinding
