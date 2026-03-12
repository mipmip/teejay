## 1. Watchlist Data Layer

- [x] 1.1 Add `Name` field to `Pane` struct in `watchlist.go`
- [x] 1.2 Add `Remove(paneID string)` method to Watchlist
- [x] 1.3 Add `Rename(paneID, name string)` method to Watchlist
- [x] 1.4 Add `DisplayName()` method to Pane that returns Name or ID
- [x] 1.5 Add tests for Remove, Rename, and DisplayName in `watchlist_test.go`

## 2. TUI Edit Mode

- [x] 2.1 Add `textinput` component to app model
- [x] 2.2 Add `editing` state field to track edit mode
- [x] 2.3 Handle `e` key to enter edit mode with current pane's display name
- [x] 2.4 Handle `Enter` in edit mode to save and exit
- [x] 2.5 Handle `Escape` in edit mode to cancel
- [x] 2.6 Update list item to show updated name after edit

## 3. TUI Delete Confirmation

- [x] 3.1 Add `deleting` state field to track delete confirmation mode
- [x] 3.2 Handle `d` key to enter delete confirmation mode
- [x] 3.3 Display "Delete [name]? (y/n)" prompt
- [x] 3.4 Handle `y` to confirm delete, remove pane, save watchlist
- [x] 3.5 Handle `n` or `Escape` to cancel delete

## 4. Verify

- [x] 4.1 Run `make test` and ensure all tests pass
- [x] 4.2 Run `make build` and manually test edit functionality
- [x] 4.3 Manually test delete functionality with confirmation
