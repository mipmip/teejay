## Why

A good showcase video/GIF is essential for demonstrating Teejay's value proposition on GitHub and social media. Using VHS (charmbracelet/vhs) allows us to create reproducible, version-controlled demo recordings that can be regenerated as the UI evolves.

## What Changes

- Add a VHS tape file to generate a showcase GIF/video
- Demo should highlight key features:
  - Main TUI with watched panes list and live preview
  - Status indicators (busy/ready animations)
  - Adding a pane via browser (session → pane selection)
  - Switching to a pane
  - Configure popup for notifications
- Output suitable for README and social sharing

## Capabilities

### New Capabilities

- `vhs-showcase`: VHS tape file and generated assets for project showcase

### Modified Capabilities

(none - this is documentation/assets, not code behavior)

## Impact

- New file: `demo.tape` (VHS script)
- New file: `demo.gif` or `assets/demo.gif` (generated output)
- README.md update to embed the showcase GIF
