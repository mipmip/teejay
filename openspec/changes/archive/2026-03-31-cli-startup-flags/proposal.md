## Why

Users need to control the TUI's initial state from the command line — for scripting, tmux automation, and personal workflow shortcuts (e.g., `alias tj-monitor='tj --columns --sort-activity --sound'`). Currently only `--config` and `--watchlist` are supported as flags. All config file options and runtime-toggleable states should be settable via CLI flags, with flags overriding config file values.

## What Changes

- Add CLI flags that map to every config option and runtime-toggleable state
- Flags override config file values (config file → flag override → runtime toggle)
- Boolean flags support `--flag` (enable) and `--no-flag` (disable) convention
- Update `--help` output with all new flags

### Flags to add

| Flag | Short | Type | Overrides | Description |
|---|---|---|---|---|
| `--sound` | | bool | `alerts.sound_on_ready` | Enable sound alerts |
| `--no-sound` | | bool | `alerts.sound_on_ready` | Disable sound alerts |
| `--notify` | | bool | `alerts.notify_on_ready` | Enable desktop notifications |
| `--no-notify` | | bool | `alerts.notify_on_ready` | Disable desktop notifications |
| `--sort-activity` | | bool | `display.sort_by_activity` | Start with activity sort |
| `--sort-watchlist` | | bool | `display.sort_by_activity` | Start with watchlist order (default) |
| `--columns` | | bool | layout mode | Start in multi-column layout |
| `--recency-color` | | bool | `display.recency_color` | Enable recency color gradient |
| `--no-recency-color` | | bool | `display.recency_color` | Disable recency color gradient |
| `--picker` | | bool | `display.picker_mode` | Picker mode: Enter switches to pane and quits |

## Capabilities

### New Capabilities

- `cli-startup-flags`: CLI flags for controlling initial TUI state, overriding config file values

### Modified Capabilities

_None — existing behavior unchanged, flags are additive_

## Impact

- `cmd/tj/main.go` — extend `parseFlags()` with new boolean flags, apply overrides to config before passing to `ui.New()`, update help text
- `internal/ui/app.go` — `New()` may need a layout mode parameter (currently derived internally)
