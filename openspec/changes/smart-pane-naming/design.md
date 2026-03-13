## Context

Currently, when a pane is added to the watchlist (via `tj add` or the in-app browser), it stores only the pane ID with no name. Users see cryptic IDs like `%5` until they manually rename. Tmux provides rich metadata (session name, window name, running command) that can be leveraged to auto-generate meaningful names.

## Goals / Non-Goals

**Goals:**
- Automatically suggest a meaningful name when adding a pane
- Use tmux metadata: running command, window name, session name (in priority order)
- Detect and filter out generic names that aren't useful
- Prompt the user when the name is too generic (CLI only; in-app uses the name and allows later rename)

**Non-Goals:**
- Advanced NLP or fuzzy matching for name generation
- Renaming existing panes retroactively
- Fetching process tree or deep command inspection

## Decisions

### Decision 1: Name guessing priority order

The system will try to extract a name in this order:
1. **Pane command** - the current command running in the pane (e.g., `nvim`, `cargo`, `npm`)
2. **Window name** - if the user has named the window distinctively
3. **Session name** - if the session has a meaningful name

**Rationale:** The running command is most specific to what the pane is doing. Window and session names are fallbacks but may be generic (e.g., "0", "main").

**Alternatives considered:**
- Use only command: Too narrow, misses user-assigned tmux names
- Use session:window:pane format: Too verbose, not human-friendly

### Decision 2: Generic name list

Maintain a hardcoded list of generic names that trigger prompting:
```
bash, zsh, fish, sh, tmux, screen,
claude, opencode, aider, cursor,
0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
main, default, new, window
```

**Rationale:** These are common shell/tool names or tmux default names that don't identify what work is happening in the pane.

**Alternatives considered:**
- Regex patterns: More complex, harder to maintain
- User-configurable list: Over-engineered for this use case

### Decision 3: Prompt behavior differs by context

- **CLI (`tj add`)**: If name is generic, use stdin/stdout to prompt user for a name
- **TUI (pane browser)**: Always use guessed name; user can rename via configure popup later

**Rationale:** CLI has blocking I/O available. TUI should not interrupt the selection flow with modal prompts - the existing rename flow handles this.

### Decision 4: New package for name guessing

Create `internal/naming` package with:
- `GuessName(paneID string) (string, bool)` - returns name and whether it's generic
- `IsGeneric(name string) bool` - checks if name is in generic list

**Rationale:** Keeps naming logic isolated and testable. Both CLI and TUI can use the same logic.

## Risks / Trade-offs

**[Risk]** Window name may require additional tmux query → Extend `PaneInfo` struct and `ListAllPanes` to include window name

**[Risk]** User may find guessed name unhelpful → They can always rename; the default is better than `%5`

**[Trade-off]** Hardcoded generic list may miss edge cases → Start conservative, can expand based on feedback
