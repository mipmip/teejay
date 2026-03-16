## Context

The current `PlayBell()` function writes `\a` to stdout, relying on terminal bell support. This is unreliable - many terminals have bell disabled or muted. Users need reliable audio alerts when panes become ready.

Current state:
- `internal/alerts/alerts.go` has `PlayBell()` using terminal escape code
- Config has `sound_on_ready` boolean (global and per-pane)
- No audio library dependency currently

## Goals / Non-Goals

**Goals:**
- Replace terminal bell with native Go audio playback
- Provide 5 distinct, recognizable notification sounds
- Allow users to select preferred sound (global default + per-pane override)
- Embed sounds in binary (no external files)
- Cross-platform support (Linux, macOS)

**Non-Goals:**
- Custom user-provided sound files
- Volume control
- Windows support (not a primary target platform)
- Sound preview in the TUI

## Decisions

### Decision 1: Use `github.com/gopxl/beep` for audio

Use the beep library for cross-platform audio playback. It supports WAV playback and works on Linux/macOS.

**Alternatives considered:**
- `github.com/faiface/beep`: Original, but unmaintained
- `github.com/hajimehoshi/oto`: Lower-level, more setup required
- System commands (`aplay`, `afplay`): Not native Go, requires external tools

**Rationale:** beep/v2 is actively maintained (gopxl fork), has simple API, supports embedded audio.

### Decision 2: Embed WAV files with `//go:embed`

Use Go's embed directive to include WAV files directly in the binary.

**Rationale:** No external file dependencies; sounds always available; single binary distribution.

### Decision 3: Sound type as string enum

Use string type for sound selection: "chime", "bell", "ping", "pop", "ding".

**Alternatives considered:**
- Integer constants: Less readable in config files
- iota enum: Harder to serialize/deserialize from YAML

**Rationale:** String values are human-readable in config.yaml and watchlist.json.

### Decision 4: Tri-state sound type per pane

Per-pane `sound_type` can be:
- `nil` (omitted): Use global default
- `""` (empty string): Explicitly use default
- `"chime"`, `"bell"`, etc.: Override with specific sound

**Rationale:** Consistent with existing tri-state pattern for `sound_on_ready` and `notify_on_ready`.

### Decision 5: Initialize audio speaker once

Initialize the beep speaker on first sound play, then reuse. Keep it simple - no background audio context.

**Rationale:** Avoid startup overhead; most users may not enable sounds.

## Risks / Trade-offs

- **[Risk] Audio library may fail on some systems** → Fallback to terminal bell if audio init fails; log warning
- **[Risk] Binary size increase** → Accept ~50-100KB increase; sounds are small WAV files
- **[Trade-off] No Windows support** → Accept for now; Linux/macOS are primary targets
- **[Trade-off] 5 fixed sounds vs custom** → Keep it simple; can expand later if requested
