## Context

The preview panel is rendered in two places: `View()` for default layout (right-side panel) and `renderMultiColumnLayout()` for multi-column layout (bottom panel). Both check terminal dimensions before showing the preview. A `ShowPreview` config flag adds a third gate: if false, skip preview rendering entirely.

## Goals / Non-Goals

**Goals:**
- Global preview on/off via config and CLI flag
- Follow established patterns for config/flag/docs

**Non-Goals:**
- Per-layout preview control (same setting applies to both layouts)
- Runtime keybind toggle for preview (could be added later)

## Decisions

### Config field: `display.show_preview`

Add `ShowPreview bool` to the `Display` struct, default `true`. Uses the same `*bool` pointer pattern in configFile for YAML parsing (to distinguish unset from false).

### CLI flags: `--preview` / `--no-preview`

Add `Preview *bool` to `CLIOverrides`. `--preview` sets true, `--no-preview` sets false. Applied in `applyOverrides`.

### Render check

In both render paths, gate preview on `m.config.Display.ShowPreview` before the existing width/height checks. When false, the default layout shows full-width list (same as narrow-terminal behavior) and multi-column layout shows columns only (no bottom preview).
