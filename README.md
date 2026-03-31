# Teejay

**Terminal Jockey** - A tmux activity monitor for vibe coders

Teejay monitors multiple tmux panes you select yourself at once. Perfect for
vibe coding with parallel AI agent sessions.

Here's a demo, though it's diffucult to tape a multisession vibe marathon.

![Teejay Demo](demo.gif)

## Why

We want to be free in our Tmux use, but it's obvious that in this vibecoding age
we want to use tmux for running parallel agent sessions. Teejay does not force
you a coding workflow. Teejay serves as a watch list for panes which the user wants to monitor.

## Features

- All watched panes in a list with live status (busy/waiting for input)
- Multi-column layout mode for overview of many panes at once
- Preview panel (side or below columns when space allows)
- Quick-answer popup: respond to agent prompts without switching panes
- Claude Code transcript-based prompt recognition (structured, version-resistant)
- `?` indicator for panes waiting with an actionable question
- Recency color gradient on waiting indicators (bright→dim green over time)
- Activity sort: order panes by most recently active
- Picker mode: use as a pane selector that quits after switching
- Detect process (agent) and show as extra info
- Switch to tmux session and window and focus pane
- Add panes by running `tj add` or by browsing sessions→panes
- Remove panes from the watch list
- Configure sound and notification alerts per pane
- Screen-scraping fallback for non-Claude agents

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

# Launch in multi-column mode with activity sort
tj --columns --sort-activity

# Launch as a pane picker (Enter switches and quits)
tj --picker --columns --sort-activity

# Launch with sound alerts enabled, notifications disabled
tj --sound --no-notify

# Show help
tj --help

# Add current pane to watchlist (run from within tmux)
tj add

# Remove current pane from watchlist
tj del

# Use custom config and watchlist paths
tj --config /path/to/config.yaml --watchlist /path/to/watchlist.json
```

## CLI Flags

### General

| Flag | Description |
|:-----|:------------|
| `-h`, `--help` | Show help message |
| `-v`, `--version` | Show version |
| `-c`, `--config <path>` | Path to config file |
| `-w`, `--watchlist <path>` | Path to watchlist file |

### Alerts

| Flag | Description |
|:-----|:------------|
| `--sound` | Enable sound alerts |
| `--no-sound` | Disable sound alerts (overrules per-pane settings) |
| `--notify` | Enable desktop notifications |
| `--no-notify` | Disable desktop notifications (overrules per-pane settings) |

### Display

| Flag | Description |
|:-----|:------------|
| `--columns` | Start in multi-column layout |
| `--sort-activity` | Start with activity sort (busy first, then recently finished) |
| `--sort-watchlist` | Start with watchlist order (default) |
| `--recency-color` | Enable recency color gradient on indicators |
| `--no-recency-color` | Disable recency color gradient |
| `--preview` | Show pane preview panel (default) |
| `--no-preview` | Hide pane preview panel |

### Mode

| Flag | Description |
|:-----|:------------|
| `--picker` | Picker mode: Enter switches to pane and quits |

## Keybindings

| Key | Action |
|:----|:-------|
| `↑`/`k`, `↓`/`j` | Navigate up/down in pane list |
| `←`/`h`, `→`/`l` | Navigate between columns (multi-column mode) |
| `Enter` | Switch to selected pane (quits in picker mode) |
| `Space` | Open quick-answer popup for waiting panes |
| `v` | Toggle layout: default (list+preview) ↔ multi-column |
| `o` | Toggle sort: watchlist order ↔ activity order |
| `a` | Browse and add panes |
| `s` | Scan for agent panes and auto-add |
| `c` | Configure selected pane (name, alerts) |
| `e` | Edit/rename selected pane |
| `d` | Delete selected pane from watchlist |
| `Esc` | Dismiss messages |
| `q`, `Ctrl+C` | Quit |

## Configuration

Teejay can be configured via `~/.config/teejay/config.yaml`. Copy the example
file to get started:

```bash
cp config.example.yaml ~/.config/teejay/config.yaml
```

### Configuration Options

| Option                                  | Type       | Default     | Description                                                                                                       |
|:----------------------------------------|:-----------|:------------|:------------------------------------------------------------------------------------------------------------------|
| `detection.idle_timeout`                | duration   | `2s`        | How long content must be unchanged before marking a pane as "waiting". Set to `0s` to disable idle detection.    |
| `detection.prompt_endings`              | string[]   | `[]`        | Global prompt endings - characters that indicate a shell prompt when found at the end of the last non-empty line. |
| `detection.waiting_strings`             | string[]   | `[]`        | Global waiting strings - substrings that indicate the pane is waiting for user input.                             |
| `detection.apps.<name>.prompt_endings`  | string[]   |             | App-specific prompt endings for a named application.                                                              |
| `detection.apps.<name>.waiting_strings` | string[]   |             | App-specific waiting strings for a named application.                                                             |
| `alerts.sound_on_ready`                 | bool       | `false`     | Global default: play sound when a pane becomes ready. Per-pane settings override this.                            |
| `alerts.notify_on_ready`                | bool       | `false`     | Global default: send desktop notification when a pane becomes ready. Per-pane settings override this.             |
| `display.recency_color`                 | bool       | `true`      | Color gradient on waiting indicators based on how recently the pane was active.                                    |
| `display.sort_by_activity`              | bool       | `false`     | Default sort order: `true` for activity sort, `false` for watchlist order.                                        |
| `display.layout_mode`                   | string     | `"default"` | Initial layout: `"default"` (list+preview) or `"columns"` (multi-column).                                        |
| `display.picker_mode`                   | bool       | `false`     | Picker mode: Enter switches to pane and quits.                                                                    |
| `display.show_preview`                  | bool       | `true`      | Show pane preview panel. Set to `false` to hide preview in both layouts.                                          |

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
