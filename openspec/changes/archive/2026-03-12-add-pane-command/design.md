## Context

tmon currently has a minimal TUI skeleton but no actual functionality. To become useful, users need a way to add panes to the watchlist. The `tmon add` command is the primary mechanism described in the README: "by running `[appname] add` from within the current pane which should be added."

The command must detect the current tmux pane ID and persist it to a watchlist file.

## Goals / Non-Goals

**Goals:**
- Implement `tmon add` subcommand
- Detect current tmux pane ID using `$TMUX_PANE` environment variable
- Persist watchlist to a JSON file in the user's config directory
- Provide clear error messages when not running in tmux

**Non-Goals:**
- Implementing the TUI watchlist display (separate change)
- Removing panes from the watchlist
- Browsing sessions/windows/panes interactively
- Any notification features

## Decisions

### Decision 1: Use $TMUX_PANE environment variable

Detect the current pane by reading the `$TMUX_PANE` environment variable, which tmux sets automatically in each pane (e.g., `%0`, `%1`).

**Rationale**: This is the simplest and most reliable method. No need to parse tmux commands or maintain socket connections.

**Alternative considered**: Running `tmux display-message -p '#{pane_id}'`. Rejected because it requires tmux to be in the PATH and adds subprocess overhead when the env var is already available.

### Decision 2: Store watchlist in ~/.config/tmon/watchlist.json

Use XDG-style config directory with a simple JSON format:

```json
{
  "panes": [
    {"id": "%0", "added_at": "2024-01-15T10:30:00Z"},
    {"id": "%1", "added_at": "2024-01-15T10:35:00Z"}
  ]
}
```

**Rationale**: JSON is human-readable, easy to debug, and Go has built-in support. XDG config dir is the standard location on Linux.

**Alternative considered**: SQLite database. Rejected as overkill for a simple list of pane IDs.

### Decision 3: Subcommand routing in main.go

Check `os.Args` for subcommands before launching the TUI:
- `tmon` (no args) → launch TUI
- `tmon add` → add current pane to watchlist
- Future: `tmon remove`, `tmon list`, etc.

**Rationale**: Keep it simple. No CLI framework needed yet—just check the first argument.

## Risks / Trade-offs

- [Risk] User runs `tmon add` outside tmux → Mitigated by checking `$TMUX_PANE` and showing clear error
- [Risk] Watchlist file corruption → Mitigated by atomic write (write to temp, then rename)
- [Trade-off] No duplicate detection yet → Acceptable for MVP; can add later
