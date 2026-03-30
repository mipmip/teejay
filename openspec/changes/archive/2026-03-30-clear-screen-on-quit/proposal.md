## Why

When quitting teejay, the TUI output remains on screen, cluttering the shell. The terminal should return to a clean state after exit.

## What Changes

- Enable Bubbletea's alternate screen buffer (`tea.WithAltScreen()`), so the TUI renders on a separate screen buffer and the terminal restores to its previous state on quit

## Capabilities

### New Capabilities

_None_

### Modified Capabilities

_None — this is a program option change, not a spec-level behavior change_

## Impact

- `cmd/tj/main.go` — add `tea.WithAltScreen()` to `tea.NewProgram()` options
