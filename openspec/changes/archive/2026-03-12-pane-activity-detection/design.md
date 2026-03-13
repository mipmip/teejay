## Context

The TUI currently refreshes pane content every 100ms but provides no indication of whether content is actively changing. Users monitoring AI assistants (Claude Code, Aider) or builds need to know when a process is waiting for input vs actively working.

Claude-squad solves this with SHA256 hash comparison + prompt pattern matching. We'll adopt the same proven approach.

## Goals / Non-Goals

**Goals:**
- Detect content changes using hash comparison
- Identify "waiting for input" state via prompt patterns (Claude, Aider, shell prompts)
- Show status indicator in pane list (icon or color)
- Minimal performance overhead

**Non-Goals:**
- Auto-responding to prompts (like claude-squad's daemon)
- Notifications/alerts for state changes
- Custom/user-configurable prompt patterns (can add later)

## Decisions

### Decision 1: SHA256 hash for change detection

Store SHA256 hash of pane content, compare on each tick. If hash differs → Running state.

**Rationale**: Same approach as claude-squad. Hash is fixed size (32 bytes), fast to compute, and avoids storing full content history.

**Alternatives considered**:
- String equality: Memory-intensive for large panes
- Line count comparison: Misses content changes that don't add/remove lines

### Decision 2: In-memory status tracking in UI model

Store `map[string]PaneStatus` in the UI Model, keyed by pane ID. Update status on each tick.

**Rationale**: Status is transient display state, not persistent data. No need to save to watchlist.json.

**Alternatives considered**:
- Store in watchlist.Pane struct: Pollutes data model with display concerns
- Separate status file: Over-engineered for volatile data

### Decision 3: Three states with idle timeout

States:
- `Running`: Hash changed since last tick
- `Ready`: Hash stable AND prompt pattern detected
- `Idle`: Hash stable for N ticks, no prompt detected

Use 2-second idle timeout (20 ticks at 100ms).

**Rationale**: Distinguishes between "actively working" and "finished but no prompt" vs "waiting for input".

### Decision 4: Built-in prompt patterns for common tools

Detect prompts for:
- Claude Code: `"No, and tell Claude what to do differently"`, `"Do you want to proceed?"`
- Aider: `"(Y)es/(N)o"`, `">"` at line start
- Generic shell: `"$"`, `">"`, `"#"` at line end

**Rationale**: Cover the most common use cases. Can extend later.

## Risks / Trade-offs

- [False positives in prompt detection] → Use specific patterns, not generic ones
- [Hash computation overhead] → SHA256 is fast (~100μs for typical pane), negligible at 100ms intervals
- [Idle timeout feels arbitrary] → 2 seconds is reasonable; can make configurable later
