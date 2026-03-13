## Context

Currently, the monitor package tracks three pane states:
- `Idle` (○): Content stable for 2+ seconds without a prompt
- `Running` (animated): Content actively changing
- `Ready` (●): Prompt detected, waiting for input

The state machine uses a hash comparison with an idle counter to transition between Running → Idle after 20 ticks (2 seconds) of no changes. The `hasPrompt()` function checks for known prompt patterns to detect the Ready state.

## Goals / Non-Goals

**Goals:**
- Simplify to two states: Busy (animated) and Waiting (green)
- Remove the idle counter complexity
- Keep prompt detection logic intact
- Maintain visual feedback clarity

**Non-Goals:**
- Changing the prompt detection patterns
- Modifying the tick rate or refresh behavior
- Adding new visual indicators

## Decisions

### Decision 1: Rename states for clarity

**Choice:** Rename `Running` to `Busy` and `Ready` to `Waiting`

**Rationale:** These names better describe what the user cares about:
- "Busy" = something is happening, don't interrupt
- "Waiting" = ready for your input

### Decision 2: Merge Idle into Busy

**Choice:** Remove `Idle` state entirely. Any pane without a detected prompt is `Busy`.

**Rationale:** The Idle state was meant to indicate "probably done but no prompt" - but this is ambiguous. If there's no prompt, we can't confirm the pane is ready for input, so treating it as busy is safer. Users will see the animated spinner until a prompt appears.

**Trade-off:** Long-running processes with stable output will show as "busy" forever. This is acceptable because without a prompt, we genuinely don't know if input is expected.

### Decision 3: Remove idle counter from Monitor

**Choice:** Remove `idleCounter` field and `idleThreshold` constant.

**Rationale:** With no Idle state, there's no need to track how long content has been stable. The state machine simplifies to:
- Prompt detected → Waiting
- No prompt → Busy

### Decision 4: Visual indicators

**Choice:**
- Waiting: Green filled dot `●` (styled green)
- Busy: Animated spinner (braille dots)

**Rationale:** Green universally signals "go" / "ready". Animation signals activity. This is intuitive and requires no learning.

## Risks / Trade-offs

**[Trade-off]** Panes with stable output but no prompt will spin forever → Acceptable. We can't know if input is needed without a prompt.

**[Risk]** Breaking change for any code depending on `Idle` status → Low risk, internal API only. Update all references in same change.
