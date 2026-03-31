## 1. Config: Layout Mode and Picker Mode

- [x] 1.1 Add `LayoutMode string` field to `Display` struct (values: `"default"`, `"columns"`, default `"default"`)
- [x] 1.2 Add `PickerMode bool` field to `Display` struct (default `false`)
- [x] 1.3 Add YAML parsing for `layout_mode` and `picker_mode` in configFile, wire into `Load()`

## 2. CLI: Flag Parsing

- [x] 2.1 Define `CLIOverrides` struct with `*bool` pointer fields: Sound, Notify, SortActivity, Columns, RecencyColor, PickerMode
- [x] 2.2 Extend `parseFlags()` to parse `--sound`, `--no-sound`, `--notify`, `--no-notify`, `--sort-activity`, `--sort-watchlist`, `--columns`, `--recency-color`, `--no-recency-color`, `--picker` — returning a `CLIOverrides`
- [x] 2.3 Implement `applyOverrides(cfg *config.Config, overrides CLIOverrides)` that sets config values for non-nil override pointers (including `cfg.Display.LayoutMode` for `--columns`)

## 3. Wire Up

- [x] 3.1 In `main()`, call `applyOverrides(cfg, overrides)` after `config.Load()` and before `ui.New()`
- [x] 3.2 In `ui.New()`, initialize `m.layoutMode` from `cfg.Display.LayoutMode` instead of always defaulting to 0
- [x] 3.3 In `ui.New()`, store `pickerMode` from `cfg.Display.PickerMode` on Model
- [x] 3.4 In the `"enter"` key handler, when `pickerMode` is true, switch to pane and return `tea.Quit` instead of just switching

## 4. Help Text

- [x] 4.1 Update `printHelp()` to document all new flags grouped by category (Alerts, Display, Mode)

## 5. Tests

- [x] 5.1 Add tests for `parseFlags` with new flags — verify CLIOverrides are correctly populated
- [x] 5.2 Add tests for `applyOverrides` — verify config values are only modified when override is non-nil
- [x] 5.3 Add test for config `layout_mode` and `picker_mode` YAML parsing
