## 1. Implementation

- [x] 1.1 Remove `tea.Quit` after `tmux.SwitchToPane()` in Enter key handler
- [x] 1.2 Return `m, nil` instead to keep app running

## 2. Testing

- [x] 2.1 Verify Enter switches pane but app stays open
- [x] 2.2 Verify q still quits the application
- [x] 2.3 Verify ctrl+c still quits the application
