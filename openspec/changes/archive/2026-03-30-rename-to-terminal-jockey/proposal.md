## Why

The app's full name "Terminal Junkie" doesn't quite fit the intended vibe. "Terminal Jockey" better captures the idea of someone who skillfully rides/manages their terminal sessions — a tmux jockey.

## What Changes

- Rename all instances of "Terminal Junkie" to "Terminal Jockey" in the branding footer display
- Update the branding footer spec to reflect the new name

## Capabilities

### New Capabilities

_None_

### Modified Capabilities

- `branding-footer`: The displayed name changes from "Terminal Junkie" to "Terminal Jockey"

## Impact

- `internal/ui/app.go` — the `renderBrandingFooter()` function hardcodes the name
- `openspec/specs/branding-footer/spec.md` — spec references the old name
