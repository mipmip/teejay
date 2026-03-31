## Context

The current `parseFlags()` in `main.go` hand-parses `--config`, `--watchlist`, `--help`, and `--version` flags. Config is loaded, then passed to `ui.New()`. The UI reads initial state from the config (e.g., `cfg.Display.SortByActivity`). Layout mode is not configurable — it always starts in default mode.

## Goals / Non-Goals

**Goals:**
- Add boolean flags for all config options and runtime-toggleable states
- Flags override config values using a clear precedence: config file → CLI flags
- Clean `--help` output documenting all flags
- Extensible pattern for future flags

**Non-Goals:**
- Switching to a flag parsing library (keep the hand-rolled parser for now — it's simple and has no dependencies)
- Adding flags for detection patterns (these are complex and config-file-only)

## Decisions

### Flag parsing approach: extend existing `parseFlags`

Add a `CLIOverrides` struct that captures all flag values as `*bool` pointers (nil = not specified, non-nil = override). After config loading, apply overrides.

```go
type CLIOverrides struct {
    Sound         *bool
    Notify        *bool
    SortActivity  *bool
    Columns       *bool
    RecencyColor  *bool
    PickerMode    *bool
}
```

Pointer bools let us distinguish "not specified" from "explicitly set to false". This is important because `--no-sound` and "no flag" are different — the first overrides config, the second doesn't.

**Alternative considered**: Using Go's `flag` package — rejected because it doesn't support `--no-*` negation patterns well, and the hand-rolled parser is already the established pattern.

### Layout mode and picker mode as config

Add `layout_mode` to the `Display` config section (values: `"default"`, `"columns"`). The `--columns` flag sets this.

Add `picker_mode` to the `Display` config section (bool, default false). The `--picker` flag sets this. When enabled, pressing Enter on a pane calls `tmux switch-client` and then returns `tea.Quit` instead of just switching. This makes teejay behave as a pane picker/selector rather than a persistent monitor.

### Override application order

```
config.yaml → Load() → cfg
                          ↓
CLI flags  → CLIOverrides → applyOverrides(cfg, overrides)
                          ↓
                     final cfg → ui.New()
```

The `applyOverrides` function only sets values where the override pointer is non-nil.

## Risks / Trade-offs

- **[Low] Flag proliferation** — As more features are added, the flag list grows. Mitigated by grouping in help output and using consistent naming conventions.
