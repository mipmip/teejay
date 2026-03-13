## 1. Stale Pane Detection

- [x] 1.1 Add helper function `isStalePaneError(err error) bool` in `internal/ui/app.go` that checks if error contains "can't find pane"
- [x] 1.2 Add unit test for `isStalePaneError` with both matching and non-matching error cases

## 2. Auto-Removal Logic

- [x] 2.1 Add `statusMessage string` field to Model in `internal/ui/app.go` for displaying notifications
- [x] 2.2 In `captureSelectedPane()`, after detecting stale pane error, call `watchlist.Remove()` and save
- [x] 2.3 Set status message when pane is auto-removed (e.g., "Removed stale pane %65")
- [x] 2.4 Refresh the pane list after removal by calling `m.refreshList()`

## 3. UI Feedback

- [x] 3.1 Add status message display to the UI view (bottom of screen or in preview area)
- [x] 3.2 Clear status message on next user input or after a few seconds
- [x] 3.3 Handle selection adjustment when the removed pane was selected

## 4. Testing

- [x] 4.1 Add integration test that simulates stale pane scenario and verifies removal
- [x] 4.2 Verify watchlist is saved after auto-removal
