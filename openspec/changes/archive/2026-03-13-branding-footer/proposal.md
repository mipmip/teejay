## Why

The app lacks visual identity. Adding a branded footer with "Terminal Junkie" in neon-style ASCII art plus the version number gives the app personality and helps users quickly identify the app and version they're running.

## What Changes

- Add a branded footer section in the bottom-right corner of the TUI
- Display "Terminal Junkie" in stylized neon/ASCII art characters
- Include the current version number below or alongside the branding
- Footer should be subtle and not interfere with main content

## Capabilities

### New Capabilities

- `branding-footer`: Display branded footer with app name in neon-style text and version number

### Modified Capabilities

None - this is additive UI enhancement with no spec-level behavior changes.

## Impact

- `internal/ui/app.go`: Add footer rendering to the View() function
- `cmd/tj/main.go`: Already has version variable from goreleaser
- May need to pass version into the UI model
