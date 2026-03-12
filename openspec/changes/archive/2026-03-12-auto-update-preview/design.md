## Context

The TUI currently updates the preview panel only when the user navigates to a different pane. Users monitoring active processes (builds, logs, long-running commands) need to see live output without manual intervention. Claude-squad implements this using a ticker-based approach with 100ms intervals.

Current state:
- `captureSelectedPane()` is called only on selection change
- `Init()` returns nil (no initial commands)
- No tick message type exists

## Goals / Non-Goals

**Goals:**
- Auto-refresh preview content at regular intervals (~100ms)
- Maintain smooth UX without visual stuttering
- Keep existing selection-based update behavior

**Non-Goals:**
- Configurable refresh rate (can be added later if needed)
- Pause/resume auto-refresh functionality
- Differential updates (always re-capture full content)

## Decisions

### Decision 1: Ticker interval of 100ms
Using 100ms refresh interval (same as claude-squad) provides responsive updates without excessive CPU usage. `tmux capture-pane` is lightweight (~1-2ms execution).

**Alternatives considered:**
- 50ms: More responsive but potentially wasteful for terminal content that rarely changes that fast
- 500ms: Noticeable lag for fast-scrolling output
- 100ms: Good balance, proven in production (claude-squad)

### Decision 2: Self-scheduling tick command
Use Bubbletea's `tea.Cmd` pattern where each tick schedules the next tick:

```go
func tickCmd() tea.Cmd {
    return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
        return previewTickMsg{}
    })
}
```

**Alternatives considered:**
- External ticker goroutine: More complex, harder to integrate with Bubbletea's message loop
- `time.Ticker`: Requires cleanup, channel management
- `tea.Tick`: Built-in, idiomatic, auto-cleanup

### Decision 3: Skip refresh when in edit/delete mode
Don't refresh preview during modal operations (editing pane name, confirming delete) to avoid visual disruption.

## Risks / Trade-offs

- [Unnecessary refreshes when pane content hasn't changed] → Acceptable; capture-pane is cheap and detecting changes adds complexity
- [Preview flicker if content changes length significantly] → Viewport component handles this smoothly
- [Higher CPU usage when idle] → Minimal impact; 10 calls/second with <2ms each
