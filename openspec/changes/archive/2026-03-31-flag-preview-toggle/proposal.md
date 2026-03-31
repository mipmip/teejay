## Why

Users may want to disable the preview panel entirely — to save resources (preview captures pane content every 100ms), to reduce visual noise, or when using teejay purely as a status monitor. Currently the preview is always shown when space allows, with no way to disable it.

## What Changes

- Add `--preview` / `--no-preview` CLI flags
- Add `display.show_preview` config option (default `true`)
- When disabled, hide the preview panel in both default layout (no right panel) and multi-column layout (no bottom panel)
- Update documentation: printHelp(), README CLI Flags, README Configuration Options, config.example.yaml

## Capabilities

### New Capabilities

_None_

### Modified Capabilities

- `pane-preview`: Preview can be disabled globally via config/flag

## Impact

- `internal/config/config.go` — add `ShowPreview bool` to Display struct
- `cmd/tj/main.go` — add `--preview`/`--no-preview` flags, parse and apply override
- `internal/ui/app.go` — check `showPreview` config before rendering preview in both layouts
- `README.md` — add flag and config docs
- `config.example.yaml` — add `show_preview` option
- `cmd/tj/main_test.go` — add flag parsing and override tests
