## 1. Configuration

- [x] 1.1 Add `Alerts` struct to config.go with `SoundOnReady` and `NotifyOnReady` bools
- [x] 1.2 Add `Alerts` field to `Config` struct
- [x] 1.3 Update `Default()` to include alerts defaults (both false)
- [x] 1.4 Update `configFile` struct for YAML parsing
- [x] 1.5 Update `Load()` to parse alerts section

## 2. Watchlist

- [x] 2.1 Change `Pane.SoundOnReady` from `bool` to `*bool`
- [x] 2.2 Change `Pane.NotifyOnReady` from `bool` to `*bool`
- [x] 2.3 Update `SetSound()` to accept `*bool`
- [x] 2.4 Update `SetNotify()` to accept `*bool`
- [x] 2.5 Add helper methods `GetEffectiveSound(cfg)` and `GetEffectiveNotify(cfg)` that apply defaults

## 3. UI Logic

- [x] 3.1 Update alert triggering in app.go to use effective values with config defaults
- [x] 3.2 Update toggle handlers to cycle: default → enabled → disabled → default
- [x] 3.3 Update pane list rendering to show override state indicators

## 4. Documentation

- [x] 4.1 Add `alerts` section to config.example.yaml
- [x] 4.2 Add alerts options to README.md configuration table

## 5. Testing

- [x] 5.1 Test config loads alerts section correctly
- [x] 5.2 Test effective value resolution with various combinations
- [x] 5.3 Test backward compatibility with existing watchlist.json files
