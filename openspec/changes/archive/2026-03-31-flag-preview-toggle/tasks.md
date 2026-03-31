## 1. Config

- [x] 1.1 Add `ShowPreview bool` to `Display` struct (default `true`)
- [x] 1.2 Add `*bool` YAML parsing for `show_preview` in configFile, wire into `Load()`

## 2. CLI Flags

- [x] 2.1 Add `Preview *bool` to `CLIOverrides`
- [x] 2.2 Parse `--preview` and `--no-preview` in `parseFlags()`
- [x] 2.3 Apply preview override in `applyOverrides()`

## 3. UI Rendering

- [x] 3.1 In `View()` default layout: gate `showPreview` on `m.config.Display.ShowPreview` — when false, render full-width list
- [x] 3.2 In `renderMultiColumnLayout()`: gate bottom preview on `m.config.Display.ShowPreview`

## 4. Documentation

- [x] 4.1 Add `--preview`/`--no-preview` to `printHelp()` in Display section
- [x] 4.2 Add `--preview`/`--no-preview` to README CLI Flags table
- [x] 4.3 Add `display.show_preview` to README Configuration Options table
- [x] 4.4 Add `show_preview` to config.example.yaml with comment

## 5. Tests

- [x] 5.1 Add test for `parseFlags` with `--preview` and `--no-preview`
- [x] 5.2 Add test for `applyOverrides` with preview override (set, unset, nil)
- [x] 5.3 Add test for config default (`ShowPreview` is `true`)
- [x] 5.4 Add test for YAML parsing of `show_preview: false`
