## Context

The project already has `tj add` for adding panes to the watchlist, along with a naming system (`internal/naming`) that intelligently guesses pane names from tmux metadata (session > window > command). The watchlist package already has a `Remove()` method. This change adds the symmetric `tj del` command and ensures both add/del use consistent named feedback.

## Goals / Non-Goals

**Goals:**
- Provide a `tj del` CLI command that removes current pane from watchlist
- Show human-readable pane name in feedback messages for both add and del
- Mirror the structure and error handling of `tj add`

**Non-Goals:**
- Bulk delete operations
- Delete by name (only current pane via TMUX_PANE)
- Interactive confirmation prompts

## Decisions

### Decision 1: Mirror add.go structure for del.go

Create `internal/cmd/del.go` following the same pattern as `add.go`:
1. Get current pane ID from `TMUX_PANE` env var
2. Load watchlist
3. Check if pane exists (error if not)
4. Look up pane name for feedback before removing
5. Remove and save

**Rationale:** Consistency with existing code, reuse established patterns.

### Decision 2: Get name before removal for feedback

When deleting, fetch the pane's stored name or guess it from tmux metadata BEFORE removing, so we can show "Removed 'project-x' from watchlist" rather than just the pane ID.

**Rationale:** Better UX - users recognize the name they assigned or that was guessed.

### Decision 3: Use existing watchlist.GetPane() for name lookup

Look up the pane in watchlist first to get its stored `Name`. If not stored, fall back to `naming.GuessName()` using tmux metadata.

**Rationale:** Prefer the user-assigned or previously-guessed name that's already saved.

## Risks / Trade-offs

- **[Risk] Pane not in watchlist** → Return clear error message "Pane %s is not being watched"
- **[Risk] Pane no longer exists in tmux** → Still works - we just remove from watchlist, tmux state is irrelevant
- **[Trade-off] No undo** → Acceptable for simple CLI tool; users can re-add with `tj add`
