## Context

Teejay currently requires users to manually add panes one at a time — either via `tj add` in a pane or through the browser popup ('a' key). The app already has all the building blocks: `tmux.ListAllPanes()` returns every pane with its foreground command, `config.Detection.Apps` maps known agent names, and `naming.GuessName()` auto-generates display names. The missing piece is a function that ties these together to batch-add agent panes.

## Goals / Non-Goals

**Goals:**
- One-action scan that finds and adds all panes running known agents
- Available both in TUI (keybinding) and CLI (`tj scan`)
- Reuse existing config app names as the agent detection list
- Show clear feedback: how many found, how many added (vs already watched)

**Non-Goals:**
- Continuous/automatic background scanning (explicit user action only)
- Detecting agents by content analysis — just match foreground process command name
- Adding a config option to control which apps are considered "agents" (the existing `detection.apps` map is sufficient)

## Decisions

### 1. Detection via foreground command name matching
Match `paneInfo.Command` against the keys in `config.Detection.Apps`. This is the same data `tmux list-panes` already provides and is how the monitor already determines which app patterns to use. Simple string equality check.

**Alternative considered**: Scanning pane content for agent signatures — rejected because it's slow (requires capture per pane) and the foreground command is sufficient and reliable.

### 2. Shared scan logic in a new `internal/scan` package
Create a `scan.ScanAndAdd()` function that both the TUI keybinding and CLI command call. Returns a result struct with counts (found, added, skipped). This avoids duplicating logic between the two entry points.

**Alternative considered**: Inline in each caller — rejected because the logic is non-trivial (list panes, filter by app, skip watched, add with names) and would be duplicated.

### 3. TUI keybinding: `s` for scan
The `s` key triggers scan from the main watchlist view. Shows a status message like "Scan: added 3 panes (2 already watched)". The key is mnemonic and not currently bound.

### 4. CLI command: `tj scan`
New subcommand that runs the scan non-interactively, prints results to stdout, and exits. Follows the pattern of existing `tj add` and `tj del` commands.

### 5. Auto-naming via GuessName()
Use `naming.GuessName()` for each added pane, same as the browser popup flow. Accept the guessed name without prompting — the user can rename later via the configure popup.

## Risks / Trade-offs

- **Panes may be running an agent indirectly** (e.g., a shell that spawned claude) → The foreground command check handles this correctly since tmux reports the foreground process, not the shell.
- **Many panes added at once may clutter watchlist** → Acceptable — the user explicitly requested the scan. They can remove unwanted panes individually.
- **New agents not in config won't be detected** → Expected — users can add custom apps to config. The scan uses whatever apps are configured.
