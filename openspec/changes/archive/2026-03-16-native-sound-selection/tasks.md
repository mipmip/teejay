## 1. Dependencies and Sound Files

- [x] 1.1 Add `github.com/gopxl/beep/v2` dependency to go.mod
- [x] 1.2 Create `internal/alerts/sounds/` directory for embedded sound files
- [x] 1.3 Create or obtain 5 short WAV sound files (chime, bell, ping, pop, ding)
- [x] 1.4 Add embed.go with `//go:embed` directive for sound files

## 2. Sound Playback Implementation

- [x] 2.1 Create `internal/alerts/sounds/sounds.go` with sound type constants and `GetSound(soundType string)` function
- [x] 2.2 Implement `PlaySound(soundType string)` function using beep library
- [x] 2.3 Add speaker initialization with lazy init pattern (init on first play)
- [x] 2.4 Implement fallback to terminal bell if audio initialization fails
- [x] 2.5 Update `PlayBell()` to call `PlaySound("chime")` for backwards compatibility

## 3. Config Integration

- [x] 3.1 Add `SoundType` field to `Alerts` struct in config.go
- [x] 3.2 Update config loading to parse `sound_type` from YAML
- [x] 3.3 Add default sound type "chime" in `Default()` config

## 4. Watchlist Integration

- [x] 4.1 Add `SoundType *string` field to `WatchedPane` struct (pointer for tri-state)
- [x] 4.2 Add `SetSoundType(paneID string, soundType *string)` method to Watchlist
- [x] 4.3 Add `GetEffectiveSoundType(config *config.Config) string` method to WatchedPane

## 5. Alert Triggering

- [x] 5.1 Update `triggerAlerts()` in app.go to use `GetEffectiveSoundType()` and `PlaySound()`

## 6. UI Integration

- [x] 6.1 Add sound type menu item to configure popup
- [x] 6.2 Implement sound type cycling in configure popup (chime → bell → ping → pop → ding → chime)

## 7. Testing

- [x] 7.1 Add unit tests for sound type selection logic
- [x] 7.2 Verify existing tests pass
- [x] 7.3 Manual test: verify each sound plays correctly
