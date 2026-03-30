## Context

The app currently creates `tea.NewProgram()` without `tea.WithAltScreen()`. This means the TUI renders in the main terminal buffer and its output persists after exit.

## Goals / Non-Goals

**Goals:**
- Clean terminal on quit using the standard alternate screen buffer approach

**Non-Goals:**
- Custom screen clearing logic

## Decisions

**Use `tea.WithAltScreen()`** — The standard Bubbletea approach. The terminal switches to an alternate buffer on startup and restores the original buffer on exit. No manual clearing needed.
