## Context

VHS by Charmbracelet generates terminal recordings from declarative `.tape` files. The demo needs to run inside tmux to showcase Teejay's actual functionality. VHS supports running commands that spawn tmux sessions.

## Goals / Non-Goals

**Goals:**
- Create a compelling 15-30 second showcase GIF
- Demonstrate core workflow: view panes → add pane → see status → configure
- Reproducible via `vhs demo.tape`
- Suitable for GitHub README (reasonable file size)

**Non-Goals:**
- Full feature documentation (that's for docs)
- Multiple format outputs (just GIF for now)
- CI/CD integration for auto-regeneration

## Decisions

### Decision 1: GIF output format

**Choice:** Generate `.gif` for maximum compatibility.

**Rationale:**
- Works everywhere (GitHub, social, browsers)
- No video player needed
- Alternative: MP4 - better quality but less compatible in READMEs

### Decision 2: Demo script structure

**Choice:** Script a realistic workflow:
1. Start with tj running, showing watched panes
2. Press 'a' to open browser, navigate to add a pane
3. Show status change animation
4. Open configure popup, toggle notification
5. End with clean view

**Rationale:**
- Shows real user value in under 30 seconds
- Highlights unique features (status animation, pane browser)

### Decision 3: Terminal dimensions

**Choice:** 1200x700 pixels, font size 18-20.

**Rationale:**
- Wide enough for two-panel layout
- Large enough text to be readable on GitHub
- Fits well in README without scrolling

### Decision 4: Pre-seeded tmux state

**Choice:** The tape file will set up tmux sessions/panes before launching tj.

**Rationale:**
- VHS needs actual tmux panes to demonstrate
- Can create fake "agent" sessions with running commands

## Risks / Trade-offs

**[Risk] tmux dependency** → VHS must run in environment with tmux. Mitigation: Document requirement in tape file comments.

**[Risk] GIF file size** → Long recordings get large. Mitigation: Keep demo under 20 seconds, optimize with reasonable framerate.

**[Trade-off] Scripted vs real** → Demo is staged, not real agent activity. Acceptable for showcase purposes.
