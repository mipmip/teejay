## 1. Update Model Structure

- [x] 1.1 Add bubbles/list to Model struct in `internal/ui/app.go`
- [x] 1.2 Create paneItem type that implements list.Item interface
- [x] 1.3 Update `New()` to load watchlist and initialize list component

## 2. Implement List Display

- [x] 2.1 Update `View()` to render the bubbles/list
- [x] 2.2 Add empty state message when no panes are watched
- [x] 2.3 Style list items to show pane ID and added timestamp

## 3. Keyboard Navigation

- [x] 3.1 Update `Update()` to delegate key events to bubbles/list
- [x] 3.2 Ensure q/Ctrl+C still quits the application

## 4. Tests

- [x] 4.1 Update `internal/ui/app_test.go` with tests for pane list display
- [x] 4.2 Ensure all tests pass with `make test`

## 5. Verify

- [x] 5.1 Run `make build` and test TUI with empty watchlist
- [x] 5.2 Add a pane with `tmon add` and verify it appears in the list
