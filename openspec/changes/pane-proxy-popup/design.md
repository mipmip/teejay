## Context

TJ currently switches tmux to the selected pane when Enter is pressed. Users lose their place in TJ and must manually navigate back. We want to enable quick pane interaction via a tmux popup overlay.

The proxy approach mirrors the pane's display and forwards input, creating an interactive session without actually switching away.

## Goals / Non-Goals

**Goals:**
- Enable interactive pane access without leaving TJ
- Support full terminal input including special keys (Ctrl-C, arrows, etc.)
- Preserve ANSI colors and styling from the target pane
- Provide a clear exit mechanism (double-Escape)
- Feel responsive enough for typing and navigation

**Non-Goals:**
- Perfect latency parity with native tmux (some delay is acceptable)
- Supporting tmux versions below 3.2 (no `display-popup`)
- Resizing the target pane to match popup dimensions
- Mouse support within the proxy

## Decisions

### 1. Proxy Architecture: Capture-Display-Forward Loop

**Decision:** Build a standalone Go program that runs in the popup and proxies the pane via tmux commands.

**Rationale:**
- `capture-pane -p -e` gives us content with ANSI escapes (~12ms)
- `send-keys -l` forwards literal input to the pane (~17ms)
- Cursor position available via `display-message -p '#{cursor_x} #{cursor_y}'`
- No PTY stealing or complex IPC needed

**Alternatives considered:**
- PTY attachment (reptyr-style): Too complex, potential conflicts with tmux's PTY ownership
- pipe-pane bidirectional: Doesn't handle display, only I/O piping

### 2. Input Handling: Raw Terminal Mode

**Decision:** Use Go's `term` package to put stdin in raw mode and read bytes directly.

**Rationale:**
- Need to capture all keystrokes including Ctrl sequences and escape codes
- Must handle multi-byte sequences (arrow keys, function keys)
- `term.MakeRaw()` + restore on exit is well-supported in Go

**Alternatives considered:**
- Bubbletea for proxy TUI: Overkill, adds unnecessary abstraction for simple I/O loop

### 3. Refresh Rate: 30fps with Adaptive Backoff

**Decision:** Target ~33ms refresh interval, skip refresh if no pane output change detected.

**Rationale:**
- 30fps feels smooth for most terminal interactions
- Comparing content hashes avoids unnecessary redraws
- Balance between responsiveness and CPU usage

### 4. Exit Mechanism: Double-Escape

**Decision:** Two Escape presses within 500ms exits the proxy.

**Rationale:**
- Single Escape is used by vim, less, and many TUI apps
- Double-Escape is uncommon enough to be safe
- Provides clear, memorable exit gesture

**Alternatives considered:**
- Ctrl-Q: Could conflict with XON/XOFF flow control
- Prefix key (like tmux prefix): Adds complexity, learning curve

### 5. Popup Invocation: TJ Spawns via tmux display-popup

**Decision:** When Enter is pressed in TJ, execute `tmux display-popup -E -w 80% -h 80% "tj proxy <pane-id>"`.

**Rationale:**
- `-E` closes popup when proxy exits
- 80% size leaves visible border showing you're in a popup
- TJ stays running underneath, ready when popup closes

## Risks / Trade-offs

**[Latency]** 20-50ms round-trip may feel sluggish for fast typing
→ Mitigation: Show input immediately in local buffer, reconcile on next refresh

**[Terminal Size Mismatch]** Popup may be different size than target pane
→ Mitigation: Display what fits, indicate truncation. User can resize popup.

**[Complex TUI Apps]** vim, htop etc. may have rendering quirks
→ Mitigation: Capture with `-e` flag preserves escapes. Test with common apps.

**[tmux Version]** Requires tmux 3.2+ for display-popup
→ Mitigation: Check version at startup, fall back to regular switch if unavailable

**[Escape Sequence Parsing]** Multi-byte input sequences need correct handling
→ Mitigation: Use established patterns from terminal libraries, buffer partial sequences
