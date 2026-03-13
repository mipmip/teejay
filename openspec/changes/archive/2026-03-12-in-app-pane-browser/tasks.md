## 1. Tmux Pane Listing

- [x] 1.1 Create `PaneInfo` struct in `internal/tmux/list.go` with ID, Session, Window, Pane index, Command fields
- [x] 1.2 Add `ListAllPanes() ([]PaneInfo, error)` function using `tmux list-panes -a -F`
- [x] 1.3 Add tests for ListAllPanes in `internal/tmux/list_test.go`

## 2. Browser State and Data

- [x] 2.1 Add `browsing bool` and `browserList list.Model` fields to UI Model
- [x] 2.2 Create `browserItem` struct implementing list.Item for pane selection
- [x] 2.3 Add `loadBrowserPanes()` method that fetches panes and filters out already watched

## 3. Browser UI

- [x] 3.1 Add `a` key handler to enter browsing mode and load panes
- [x] 3.2 Add `updateBrowsing(msg)` method to handle browser key events
- [x] 3.3 Handle Enter to add selected pane, close browser, refresh list
- [x] 3.4 Handle Escape to close browser without changes
- [x] 3.5 Add `renderBrowserPopup()` method for centered overlay

## 4. View Integration

- [x] 4.1 Modify `View()` to render browser popup over main content when browsing
- [x] 4.2 Style popup with lipgloss (border, title, padding)
- [x] 4.3 Show "No panes available" message when browser list is empty

## 5. Verify

- [x] 5.1 Run `make test` and ensure all tests pass
- [x] 5.2 Run `make build` and test pane browser functionality
- [x] 5.3 Verify added panes appear in main list with correct status
