## 1. Export config path

- [x] 1.1 Rename `configPath()` to `ConfigPath()` in watchlist.go (export it)

## 2. Add file monitoring state to Model

- [x] 2.1 Add `watchlistMtime time.Time` field to Model struct
- [x] 2.2 Initialize watchlistMtime in New() by stat-ing the watchlist file

## 3. Check for file changes on tick

- [x] 3.1 In Update() tick handler, stat the watchlist file to get current mtime
- [x] 3.2 Compare current mtime with stored watchlistMtime
- [x] 3.3 Skip check if in editing, deleting, or browsing mode

## 4. Reload watchlist on change

- [x] 4.1 If mtime changed, call watchlist.Load() to reload
- [x] 4.2 Update model's watchlist pointer with new data
- [x] 4.3 Update watchlistMtime to new value
- [x] 4.4 Call refreshList() to update the list view

## 5. Preserve selection

- [x] 5.1 Before refresh, store current selectedPaneID
- [x] 5.2 After refresh, check if selectedPaneID still exists in new watchlist
- [x] 5.3 If exists, keep selection; if not, select first pane (or clear if empty)
- [x] 5.4 Re-capture preview for the selected pane
