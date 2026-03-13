## 1. Monitor Package

- [x] 1.1 Add spinner frames constant array in `internal/monitor/status.go`
- [x] 1.2 Add `IndicatorAnimated(frame int) string` method to PaneStatus
- [x] 1.3 Add tests for animated indicator

## 2. UI Integration

- [x] 2.1 Add `tickFrame` counter to Model struct in `internal/ui/app.go`
- [x] 2.2 Increment tickFrame on each previewTickMsg
- [x] 2.3 Update paneItem.Title() to use animated indicator with frame
- [x] 2.4 Apply green color styling to Ready indicator

## 3. Testing

- [x] 3.1 Verify spinner cycles through frames correctly
- [x] 3.2 Verify Ready shows green indicator
