## 1. Extend tmux metadata

- [ ] 1.1 Add `WindowName` field to `PaneInfo` struct in `internal/tmux/list.go`
- [ ] 1.2 Update `ListAllPanes` format string to include `#{window_name}`
- [ ] 1.3 Parse and populate `WindowName` in `ListAllPanes`
- [ ] 1.4 Add test for window name in `list_test.go`

## 2. Create naming package

- [ ] 2.1 Create `internal/naming/naming.go` with `IsGeneric(name string) bool`
- [ ] 2.2 Define generic names list (shells, tools, numbers, defaults)
- [ ] 2.3 Implement `GuessName(paneInfo PaneInfo) (string, bool)` that tries command, window, session
- [ ] 2.4 Create `internal/naming/naming_test.go` with tests for generic detection
- [ ] 2.5 Add tests for name guessing priority (command > window > session)

## 3. Update watchlist

- [ ] 3.1 Add `AddWithName(paneID, name string)` method to `Watchlist`
- [ ] 3.2 Add test for `AddWithName`

## 4. Update CLI add command

- [ ] 4.1 Import naming package in `internal/cmd/add.go`
- [ ] 4.2 Fetch pane info for current pane (need new helper to get info by ID)
- [ ] 4.3 Call `GuessName` to get suggested name
- [ ] 4.4 If name is generic, prompt user for input via stdin
- [ ] 4.5 Use `AddWithName` instead of `Add` to save with name
- [ ] 4.6 Update success message to show assigned name
- [ ] 4.7 Add test for CLI with distinctive command
- [ ] 4.8 Add test for CLI prompting behavior (mocked stdin)

## 5. Update pane browser

- [ ] 5.1 Import naming package in `internal/ui/app.go`
- [ ] 5.2 When adding pane from browser, call `GuessName` on selected `browserItem.paneInfo`
- [ ] 5.3 Use `AddWithName` to save pane with guessed name
- [ ] 5.4 Verify pane appears with guessed name in list

## 6. Verification

- [ ] 6.1 Run `go test ./...` to verify all tests pass
- [ ] 6.2 Manual test: `tj add` in a pane running nvim (should auto-name)
- [ ] 6.3 Manual test: `tj add` in a plain zsh pane (should prompt)
- [ ] 6.4 Manual test: Add pane from browser (should show guessed name)
