## Context

The TUI app loads the watchlist once at startup and holds it in memory. The `tmon add` command modifies `~/.config/tmon/watchlist.json` directly. Currently there's no mechanism for the running TUI to detect this change.

The app already has a tick mechanism (100ms interval for preview refresh) that could be leveraged for periodic file checks.

## Goals / Non-Goals

**Goals:**
- Detect when watchlist.json is modified externally
- Reload watchlist and refresh UI automatically
- Preserve user's current selection if the pane still exists

**Non-Goals:**
- Real-time filesystem events (fsnotify) - adds dependency complexity
- Detecting which specific panes were added/removed
- Notification when new pane is added

## Decisions

### 1. Poll-based detection using file modification time

Check the file's mtime on each tick (100ms). If changed, reload the watchlist.

**Rationale**: Simple, no new dependencies. 100ms latency is imperceptible to users. The tick loop already exists.

**Alternative considered**: fsnotify - more complex, adds dependency, cross-platform concerns.

### 2. Store last known mtime in Model

Add `watchlistMtime time.Time` field to track when we last loaded the file.

### 3. Export ConfigPath from watchlist package

Make the config path accessible so the UI can stat the file.

### 4. Preserve selection after refresh

After reloading, if the previously selected pane ID still exists, keep it selected. Otherwise select the first pane.

### 5. Check file only when not in modal modes

Skip file check during editing, deleting, or browsing to avoid disruptive refreshes.

## Risks / Trade-offs

- **Polling overhead**: Stat call every 100ms is negligible. File is small, reload is fast.
- **Race condition**: If user is editing while file changes, edit completes first, then refresh happens. Acceptable behavior.
