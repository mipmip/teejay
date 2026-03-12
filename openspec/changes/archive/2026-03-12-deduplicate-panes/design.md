## Context

The watchlist currently allows duplicate pane entries. Users can run `tmon add` multiple times for the same pane, and the TUI will show the same pane multiple times. This is both a data integrity issue and a UX problem.

## Goals / Non-Goals

**Goals:**
- Add a `Contains(paneID)` method to check if pane exists
- Prevent adding duplicate panes via `tmon add`
- Deduplicate existing entries when loading watchlist
- Provide clear feedback when user tries to add existing pane

**Non-Goals:**
- Merging metadata from duplicate entries (just keep first)
- User confirmation for deduplication

## Decisions

### Decision 1: Add Contains() method to Watchlist

Add a simple `Contains(paneID string) bool` method that iterates through panes.

**Rationale**: Simple O(n) lookup is fine for small lists. No need for maps or indexes.

### Decision 2: Deduplicate on Load()

When loading the watchlist, deduplicate entries keeping the first occurrence (oldest `added_at`).

**Rationale**: Fixes existing data automatically. Users don't need to manually clean up. First occurrence preserves original add timestamp.

### Decision 3: Return specific error/message for duplicates

When `tmon add` detects a duplicate, show "Pane %s is already being watched" instead of an error.

**Rationale**: This is expected behavior, not an error. Don't exit with non-zero code.

## Risks / Trade-offs

- [Trade-off] O(n) lookup for Contains() → Acceptable for typical watchlist sizes (<100 panes)
- [Trade-off] Silent deduplication on load → Acceptable; no data is lost, just cleaned up
