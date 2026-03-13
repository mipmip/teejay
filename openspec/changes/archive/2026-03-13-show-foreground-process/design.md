## Context

The tmux package already fetches `pane_current_command` via `ListAllPanes()`. The UI refreshes pane content every 100ms via `previewTickMsg`. Currently `paneItem.Description()` returns static info (pane ID + added date).

## Goals / Non-Goals

**Goals:**
- Show current foreground process in description line (e.g., "python", "npm", "claude")
- Update dynamically as process changes
- Keep pane ID visible for reference

**Non-Goals:**
- Full process tree / child processes
- Process arguments (just the command name)
- Historical process tracking

## Decisions

### Decision 1: Add command field to paneItem

**Choice:** Add a `command` field to `paneItem` struct, updated during refresh.

**Rationale:**
- Mirrors the existing pattern for `status` field
- Updated in the same tick loop that refreshes preview
- Alternative: Fetch on-demand in Description() - rejected, would cause tmux calls on every render

### Decision 2: Use GetPaneByID for single pane lookup

**Choice:** Call `tmux.GetPaneByID()` during tick refresh to get current command.

**Rationale:**
- Already exists and returns PaneInfo with Command field
- One tmux call per selected pane per tick (100ms) - acceptable overhead
- Alternative: Batch fetch all panes - overkill for refreshing selected pane

### Decision 3: Description format

**Choice:** Show command prominently, keep pane ID as secondary info: `"python • %42"`

**Rationale:**
- Command is the most useful info for quick scanning
- Pane ID still needed for identification
- Removed "added" date - less useful for active monitoring

## Risks / Trade-offs

**[Trade-off] Removed added date** → Less historical context, but command info is more actionable.

**[Risk] Stale command** → If tmux call fails, show last known command. Don't break the UI.
