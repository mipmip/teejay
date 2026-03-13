# Teejay

**Terminal Junky** - A tmux activity monitor for vibe coders

Teejay fuels your addiction to terminal-based development by letting you monitor multiple tmux panes at once. Perfect for vibe coding with parallel AI agent sessions.

## Why

We want to be free in our Tmux use, but it's obvious that in this vibecoding age
we want to use tmux for running parallel agent sessions. TeeJay does not want to force
you a coding workflow. Teejay serves as a watch list for panes which the user wants to monitor.


## Features

- All watched panes in a list in the left sidepanel
- Show status (busy/waiting for input)
- Detect process (agent) and show as extra info
- Open tmux pane from list for user input
- Switch to tmux session and window and focus pane
- Add a pane to the watch list:
  - by running `tj add` from within the current pane
  - by browsing through sessions->panes
- Remove a pane from the watch list
- Configure notifications per pane when activity stops

## Installation

### Nix

Run directly without installing:
```bash
nix run github:mipmip/teejay
```

Or enter a development shell:
```bash
nix develop github:mipmip/teejay
```

### From source

```bash
go build -o tj ./cmd/tj
```

## Usage

```bash
# Launch the TUI
tj

# Add current pane to watchlist (run from within tmux)
tj add
```

## Tech Stack

- Go + Bubbletea/Lipgloss for TUI
