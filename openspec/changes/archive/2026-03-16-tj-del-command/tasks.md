## 1. Create Delete Command

- [x] 1.1 Create `internal/cmd/del.go` with `DelPane()` function
- [x] 1.2 Implement pane ID lookup via `TMUX_PANE` env var
- [x] 1.3 Load watchlist and check if pane exists
- [x] 1.4 Get pane name before removal (from watchlist or via naming.GuessName)
- [x] 1.5 Remove pane and save watchlist
- [x] 1.6 Display named feedback message

## 2. Wire Up CLI

- [x] 2.1 Add "del" case to `cmd/tj/main.go` switch statement
- [x] 2.2 Call `cmd.DelPane()` and handle errors

## 3. Testing

- [x] 3.1 Test `tj del` removes pane and shows correct name
- [x] 3.2 Test `tj del` on unwatched pane shows info message
- [x] 3.3 Test `tj del` outside tmux shows error
