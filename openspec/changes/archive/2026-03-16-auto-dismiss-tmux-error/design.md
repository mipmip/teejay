## Context

The TUI displays various temporary messages (errors, status updates) that currently require manual dismissal. A general-purpose auto-dismiss mechanism will improve UX and establish a consistent pattern for all temporary messages. The app already uses Bubbletea's `tea.Tick` mechanism for preview refresh timing, establishing a pattern we can extend.

## Goals / Non-Goals

**Goals:**
- Create a reusable auto-dismiss mechanism for temporary messages
- Auto-dismiss messages after a configurable timeout
- Maintain the ability to dismiss early with Esc
- Apply to the "not in tmux" error as first use case
- Make it easy for future messages to use the same pattern

**Non-Goals:**
- Per-message configurable timeouts (single default timeout is sufficient for now)
- Visual countdown/timer indicators
- Queueing multiple messages (one at a time is fine)

## Decisions

### Decision 1: Generic temporary message state

Replace specific boolean flags (like `notInTmuxMsg`) with a generic `temporaryMessage` string field. When non-empty, it displays as the footer message and triggers auto-dismiss.

**Alternatives considered:**
- Keep separate booleans per message type: Doesn't scale, duplicates dismiss logic
- Message queue with priorities: Over-engineered for current needs

**Rationale:** Single field is simple, reusable, and handles the common case well.

### Decision 2: Use tea.Tick for timeout

Use Bubbletea's `tea.Tick` command to schedule auto-dismissal, consistent with existing `tickCmd()` pattern for preview refresh.

**Alternatives considered:**
- Custom goroutine + channel: More complex, requires cleanup handling
- Shared tick with preview: Would couple unrelated concerns

**Rationale:** `tea.Tick` is the idiomatic Bubbletea approach, already proven in codebase.

### Decision 3: Dedicated message type for dismiss

Create a `dismissTemporaryMsg` type rather than reusing existing tick messages.

**Rationale:** Keeps timeout handling isolated; prevents accidental dismissal from unrelated ticks.

### Decision 4: 3-second default timeout

Use a 3-second delay before auto-dismiss.

**Rationale:** Long enough to read the message, short enough to not be annoying. Consistent with common UI patterns.

## Risks / Trade-offs

- **[Risk] Message dismissed before user reads it** → 3 seconds is sufficient for short messages; user can trigger action again if needed
- **[Trade-off] Single message vs queue** → Keeping it simple; new message replaces existing one, which is acceptable UX
- **[Trade-off] Fixed timeout vs configurable** → Starting simple; can add config later if users request it
