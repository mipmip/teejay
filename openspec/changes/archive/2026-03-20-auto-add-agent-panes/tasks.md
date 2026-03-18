## 1. Core scan logic

- [x] 1.1 Create `internal/scan/scan.go` with `ScanResult` struct (Found, Added, Skipped int) and `ScanAndAdd(wl, cfg, allPanes)` function
- [x] 1.2 Implement detection: iterate panes, match `pane.Command` against `config.Detection.Apps` keys, skip already-watched panes, add with `naming.GuessName()`

## 2. TUI integration

- [x] 2.1 Add `s` keybinding in the main view key handler that calls `scan.ScanAndAdd()` and shows a status message
- [x] 2.2 Update the help footer to include `s: scan` in the keybinding list

## 3. CLI command

- [x] 3.1 Create `internal/cmd/scan.go` with `ScanPanes()` function that loads config/watchlist, calls `scan.ScanAndAdd()`, and prints results
- [x] 3.2 Register the `scan` subcommand in `cmd/tj/main.go`

## 4. Testing

- [x] 4.1 Add unit tests for `scan.ScanAndAdd()`: agent panes found, already-watched skipped, no agents found
- [x] 4.2 Run existing tests to verify no regressions
