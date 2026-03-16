## Context

The pane configuration menu allows cycling through sound types (chime → bell → ping → pop → ding → chime) using Enter key. Currently the selection updates but no preview is played, requiring users to wait for a pane to become "ready" to hear the sound.

The `alerts.PlaySound(soundType)` function is already available and handles audio playback with fallback to terminal bell.

## Goals / Non-Goals

**Goals:**
- Play preview of selected sound immediately when cycling sound types
- Use existing audio infrastructure (no new dependencies)

**Non-Goals:**
- Volume control for previews
- Muting preview sounds separately from alert sounds
- Preview for other configuration options

## Decisions

### Decision 1: Play sound after updating selection

Call `alerts.PlaySound(nextType)` immediately after cycling to the next sound type, using the newly selected sound type.

**Rationale:** Simple, direct approach using existing infrastructure. The sound plays non-blocking (handled by beep library).

## Risks / Trade-offs

- **[Risk] Rapid cycling plays many sounds** → Acceptable - user can stop pressing Enter, and sounds are short
- **[Trade-off] No way to disable preview** → Keeps implementation simple; if users want silence they can stop cycling
