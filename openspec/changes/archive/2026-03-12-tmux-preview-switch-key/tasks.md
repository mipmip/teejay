## 1. Tmux Switch Function

- [x] 1.1 Add `SwitchToPane(paneID string) error` function in `internal/tmux/switch.go`
- [x] 1.2 Add `IsInsideTmux() bool` function to check TMUX environment variable
- [x] 1.3 Add tests for switch functions

## 2. UI Integration

- [x] 2.1 Add Enter key handler in `internal/ui/app.go` Update method
- [x] 2.2 Add state field for showing "not in tmux" message
- [x] 2.3 Update View to display the "not in tmux" message when triggered
- [x] 2.4 Update help footer to include Enter keybinding

## 3. Testing

- [x] 3.1 Add tests for Enter key behavior in app_test.go
