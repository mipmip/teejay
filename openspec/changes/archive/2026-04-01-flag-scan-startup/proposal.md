## Why

Users who frequently spin up new agent sessions want teejay to automatically discover and add them on startup, without having to press `s` every time. A `--scan` flag lets users include it in their launch alias (e.g., `alias tj='tj --scan --columns'`) so new agent panes are picked up immediately.

## What Changes

- Add `--scan` CLI flag that runs the agent scan (same as pressing `s`) at TUI startup
- Add `display.scan_on_start` config option (default `false`)
- Update documentation: printHelp(), README, config.example.yaml, CHANGELOG

## Capabilities

### New Capabilities

_None — reuses existing scan logic_

### Modified Capabilities

- `auto-scan-agents`: Scan can now be triggered at startup via flag/config

## Impact

- `internal/config/config.go` — add `ScanOnStart bool` to Display struct
- `cmd/tj/main.go` — add `--scan` flag, parse and apply override
- `internal/ui/app.go` — run scan in `Init()` or early in first tick when `ScanOnStart` is true
- Docs: README, config.example.yaml, printHelp(), CHANGELOG
