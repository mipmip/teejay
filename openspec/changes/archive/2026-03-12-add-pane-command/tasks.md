## 1. Watchlist Package

- [x] 1.1 Create `internal/watchlist/` directory
- [x] 1.2 Create `internal/watchlist/watchlist.go` with Pane struct and Watchlist type
- [x] 1.3 Implement `Load()` function to read watchlist from JSON file
- [x] 1.4 Implement `Save()` function to write watchlist to JSON file (atomic write)
- [x] 1.5 Implement `Add(paneID string)` function to add a pane to the watchlist

## 2. Add Command

- [x] 2.1 Create `internal/cmd/` directory for CLI commands
- [x] 2.2 Create `internal/cmd/add.go` with AddPane function
- [x] 2.3 Implement tmux pane detection (read $TMUX_PANE env var)
- [x] 2.4 Handle error case when not running in tmux

## 3. Command Routing

- [x] 3.1 Update `cmd/tmon/main.go` to check os.Args for subcommands
- [x] 3.2 Route "add" subcommand to AddPane function
- [x] 3.3 Keep default behavior (no args) launching TUI

## 4. Tests

- [x] 4.1 Create `internal/watchlist/watchlist_test.go` with tests for Load, Save, Add
- [x] 4.2 Create `internal/cmd/add_test.go` with tests for pane detection
- [x] 4.3 Ensure all tests pass with `make test`

## 5. Verify

- [x] 5.1 Build and test `tmon add` outside tmux (should show error)
- [x] 5.2 Build and verify the binary compiles without errors
