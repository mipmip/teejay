# Teejay

**Terminal Junky** - A tmux activity monitor for vibe coders

Teejay monitors multiple tmux panes you select yourself at once. Perfect for
vibe coding with parallel AI agent sessions.

Here's a demo, though it's diffucult to tape a multisession vibe marathon.

![Teejay Demo](demo.gif)

## Why

We want to be free in our Tmux use, but it's obvious that in this vibecoding age
we want to use tmux for running parallel agent sessions. Teejay does not force
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

# Show help
tj --help

# Add current pane to watchlist (run from within tmux)
tj add

# Remove current pane from watchlist
tj del

# Use custom config and watchlist paths
tj --config /path/to/config.yaml --watchlist /path/to/watchlist.json
```

## Configuration

Teejay can be configured via `~/.config/teejay/config.yaml`. Copy the example
file to get started:

```bash
cp config.example.yaml ~/.config/teejay/config.yaml
```

### Configuration Options

| Option                                  | Type       | Default    | Description                                                                                                       |
|:----------------------------------------|:-----------|:-----------|:------------------------------------------------------------------------------------------------------------------|
| `detection.idle_timeout`                | duration   | `2s`       | How long content must be unchanged before marking a pane as ''waiting''. Set to `0s` to disable idle detection.   |
| `detection.prompt_endings`              | string[]   | `[]`       | Global prompt endings - characters that indicate a shell prompt when found at the end of the last non-empty line. |
| `detection.waiting_strings`             | string[]   | `[]`       | Global waiting strings - substrings that indicate the pane is waiting for user input.                             |
| `detection.apps.<name>.prompt_endings`  | string[]   |            | App-specific prompt endings for a named application.                                                              |
| `detection.apps.<name>.waiting_strings` | string[]   |            | App-specific waiting strings for a named application.                                                             |
| `alerts.sound_on_ready`                 | bool       | `false`    | Global default: play terminal bell when a pane becomes ready. Per-pane settings override this.                    |
| `alerts.notify_on_ready`                | bool       | `false`    | Global default: send desktop notification when a pane becomes ready. Per-pane settings override this.             |

**Note:** When a pane is running a configured application (matched by process
name), only the app-specific patterns are used - global patterns are ignored.
This allows precise detection for each tool without interference.

### Built-in App Defaults

Teejay includes default patterns for common AI coding tools:

- **claude**: Detects `? for shortcuts`
- **aider**: Detects `(Y)es/(N)o`
- **codex**: Detects `[Y/n]`
- **opencode**: Detects `Continue?`

### Adding Custom App Patterns

```yaml
detection:
  apps:
    myapp:
      prompt_endings:
        - ">"
      waiting_strings:
        - "Press Enter to continue"
```

## Tech Stack

- Go + Bubbletea/Lipgloss for TUI
