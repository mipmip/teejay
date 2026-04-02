## Context

The app already has a single-pane delete flow: `d` → confirmation "Delete X? (y/n)" → `y` to confirm. The "delete all" follows the same pattern but operates on all panes. The preview toggle is a simple boolean flip on `m.config.Display.ShowPreview`, which is already checked in both render paths.

## Goals / Non-Goals

**Goals:**
- `D` opens a confirmation popup, `y` clears the watchlist
- `p` toggles preview on/off at runtime

**Non-Goals:**
- Persisting the preview toggle to config file (runtime only, like `v` and `o`)

## Decisions

### Delete all: reuse existing confirmation pattern

Add a `deletingAll` bool to Model (distinct from `deleting` which is single-pane). When `D` is pressed and watchlist is not empty, set `deletingAll = true`. The `updateDeletingAll` handler waits for `y`/`n`. On `y`, call `m.watchlist.Clear()` (or remove all panes), save, refresh.

### Preview toggle: flip config at runtime

`p` toggles `m.config.Display.ShowPreview`. The View() method already checks this value, so the next render immediately reflects the change. No new state field needed.
