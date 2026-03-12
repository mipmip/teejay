## 1. Watchlist Package Updates

- [x] 1.1 Add `Contains(paneID string) bool` method to Watchlist
- [x] 1.2 Add `Deduplicate()` method to remove duplicate pane entries
- [x] 1.3 Call `Deduplicate()` in `Load()` after unmarshaling JSON
- [x] 1.4 Add tests for `Contains()` in `watchlist_test.go`
- [x] 1.5 Add tests for `Deduplicate()` in `watchlist_test.go`

## 2. Add Command Updates

- [x] 2.1 Update `AddPane()` to check if pane already exists before adding
- [x] 2.2 Show "already being watched" message instead of adding duplicate
- [x] 2.3 Add test for duplicate pane handling in `add_test.go`

## 3. Verify

- [x] 3.1 Run `make test` and ensure all tests pass
- [x] 3.2 Run `make build` and test `tmon add` with duplicate pane
- [x] 3.3 Verify TUI shows deduplicated list
