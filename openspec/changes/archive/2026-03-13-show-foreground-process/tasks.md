## 1. Model Updates

- [x] 1.1 Add `command` field to `paneItem` struct in `internal/ui/app.go`
- [x] 1.2 Update `paneItem.Description()` to return `command • paneID` format

## 2. Refresh Logic

- [x] 2.1 In tick handler, call `tmux.GetPaneByID()` to get current command for selected pane
- [x] 2.2 Update `refreshListWithFrame()` to fetch and set command for all visible panes
- [x] 2.3 Handle errors gracefully (keep last known command on failure)

## 3. Testing

- [x] 3.1 Verify description shows current foreground process
- [x] 3.2 Verify process updates when changing commands in pane
- [x] 3.3 Verify pane ID is still visible in description
