## Context

The scan logic (`scan.ScanAndAdd`) already exists and is triggered by pressing `s` at runtime. It calls `tmux.ListAllPanes()`, matches against configured app patterns, and adds discovered panes to the watchlist. We need to trigger this same logic at startup.

## Goals / Non-Goals

**Goals:**
- Run scan at TUI startup when flag/config is set
- Reuse existing scan logic, no duplication

**Non-Goals:**
- Continuous background scanning (that's a separate feature)
- Changing what the scan detects

## Decisions

### Trigger point: Model Init()

Run the scan in the Bubbletea `Init()` method by returning a `tea.Cmd` that performs the scan. This runs after the model is constructed but before the first render, and doesn't block the UI since `tea.Cmd` runs in a goroutine.

The result arrives as a message, which the `Update()` handler processes to update the watchlist and refresh the list — same pattern as the async prompt check.

**Alternative considered**: Running scan synchronously in `New()` — rejected because it would block during model construction and delay first render.

### Config field: `display.scan_on_start`

Add `ScanOnStart bool` to Display, default `false`. The `--scan` flag sets it to true.
