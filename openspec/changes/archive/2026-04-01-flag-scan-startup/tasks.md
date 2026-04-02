## 1. Config

- [x] 1.1 Add `ScanOnStart bool` to `Display` struct (default `false`)
- [x] 1.2 Add `*bool` YAML parsing for `scan_on_start` in configFile, wire into `Load()`

## 2. CLI Flag

- [x] 2.1 Add `Scan *bool` to `CLIOverrides`
- [x] 2.2 Parse `--scan` in `parseFlags()`
- [x] 2.3 Apply scan override in `applyOverrides()`

## 3. Startup Scan

- [x] 3.1 Add `scanResultMsg` type to carry scan results
- [x] 3.2 In `Init()`, if `m.config.Display.ScanOnStart`, return a `tea.Cmd` that runs the scan and returns `scanResultMsg`
- [x] 3.3 Handle `scanResultMsg` in `Update()` — update watchlist, refresh list, set status message

## 4. Documentation

- [x] 4.1 Add `--scan` to `printHelp()` in Mode section
- [x] 4.2 Add `--scan` to README CLI Flags table
- [x] 4.3 Add `display.scan_on_start` to README Configuration Options table
- [x] 4.4 Add `scan_on_start` to config.example.yaml with comment
- [x] 4.5 Update CHANGELOG.md under [Unreleased]

## 5. Tests

- [x] 5.1 Add test for `parseFlags` with `--scan`
- [x] 5.2 Add test for `applyOverrides` with scan override
- [x] 5.3 Add test for config default (`ScanOnStart` is `false`) and YAML parsing
