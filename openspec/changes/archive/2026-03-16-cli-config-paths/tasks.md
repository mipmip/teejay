## 1. Config Module

- [x] 1.1 Modify `config.Load()` to accept variadic `customPath ...string` parameter
- [x] 1.2 Use custom path if provided, else call `ConfigPath()` for default
- [x] 1.3 Add unit test for `Load()` with custom path

## 2. Watchlist Module

- [x] 2.1 Add `path` field to `Watchlist` struct to store the load path
- [x] 2.2 Modify `watchlist.Load()` to accept variadic `customPath ...string` parameter
- [x] 2.3 Store the resolved path in the `Watchlist.path` field on load
- [x] 2.4 Update `Save()` to use `wl.path` instead of calling `ConfigPath()`
- [x] 2.5 Add unit test for `Load()` with custom path
- [x] 2.6 Add unit test for `Save()` writing to custom path

## 3. CLI Flag Parsing

- [x] 3.1 Add flag parsing in `main.go` to extract `--config`/`-c` and `--watchlist`/`-w` before subcommand
- [x] 3.2 Remove consumed flags from `os.Args` before subcommand dispatch
- [x] 3.3 Pass custom config path to `config.Load()`
- [x] 3.4 Pass custom watchlist path to `watchlist.Load()` and subcommands

## 4. TUI Integration

- [x] 4.1 Update `ui.New()` to accept config and watchlist paths
- [x] 4.2 Pass watchlist path through to Model for save operations
- [x] 4.3 Ensure all watchlist saves in UI use the stored path

## 5. Subcommand Integration

- [x] 5.1 Update `cmd.AddPane()` to accept optional watchlist path
- [x] 5.2 Update `cmd.DelPane()` to accept optional watchlist path
- [x] 5.3 Thread watchlist path from main.go to subcommands
