## 1. Proxy Core

- [ ] 1.1 Create `internal/proxy/proxy.go` with main proxy loop structure
- [ ] 1.2 Implement raw terminal mode setup using `golang.org/x/term`
- [ ] 1.3 Implement pane content capture with ANSI escapes (`capture-pane -p -e`)
- [ ] 1.4 Implement cursor position retrieval (`display-message -p '#{cursor_x} #{cursor_y}'`)
- [ ] 1.5 Implement display rendering with cursor positioning
- [ ] 1.6 Implement keyboard input reading in raw mode
- [ ] 1.7 Implement input forwarding via `send-keys -l`
- [ ] 1.8 Implement special key handling (Ctrl sequences, arrow keys, function keys)
- [ ] 1.9 Implement double-Escape exit detection with 500ms timeout

## 2. CLI Subcommand

- [ ] 2.1 Create `internal/cmd/proxy.go` with proxy subcommand
- [ ] 2.2 Add pane ID argument parsing and validation
- [ ] 2.3 Wire up proxy subcommand in `cmd/tj/main.go`

## 3. Tmux Integration

- [ ] 3.1 Add `TmuxVersion()` function to detect tmux version
- [ ] 3.2 Add `OpenPopup(paneID string)` function to spawn popup with proxy
- [ ] 3.3 Update Enter key handler in `internal/ui/app.go` to use popup when tmux >= 3.2
- [ ] 3.4 Keep fallback to `SwitchToPane` for tmux < 3.2

## 4. Testing

- [ ] 4.1 Add unit tests for double-Escape detection logic
- [ ] 4.2 Add unit tests for tmux version parsing
- [ ] 4.3 Manual testing: verify proxy works with shell prompts
- [ ] 4.4 Manual testing: verify proxy works with vim
- [ ] 4.5 Manual testing: verify proxy works with htop/btop
