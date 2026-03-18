## 1. Extend paneItem with session/window data

- [x] 1.1 Add `session` and `windowName` fields to the `paneItem` struct in `internal/ui/app.go`
- [x] 1.2 Update `Description()` method to return breadcrumb format `session > window : process` (omit `: process` when command is empty)

## 2. Populate breadcrumb data during refresh

- [x] 2.1 In `refreshListWithFrame()`, look up `tmux.PaneInfo` for each watched pane and populate the new `session` and `windowName` fields on `paneItem`

## 3. Testing

- [x] 3.1 Add unit test for `Description()` with full breadcrumb (session, window, and process)
- [x] 3.2 Add unit test for `Description()` when process is empty (should show `session > window`)
- [x] 3.3 Run existing tests to verify no regressions
