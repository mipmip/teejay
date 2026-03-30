## 1. Disable list quit keybindings

- [x] 1.1 Disable `KeyMap.Quit` on the main pane list after each `list.New()` call in `internal/ui/app.go` (both the empty-state and normal initialization paths)
- [x] 1.2 Disable `KeyMap.Quit` on the browser list after each `list.New()` call in `loadSessionList()` and `loadPaneList()`

## 2. Verification

- [x] 2.1 Build and run the app, verify pressing Escape in the main view does not quit
- [x] 2.2 Verify Escape still closes popups (edit, delete, browse, configure, quick-answer)
- [x] 2.3 Verify `q` and `ctrl+c` still quit the app from the main view
