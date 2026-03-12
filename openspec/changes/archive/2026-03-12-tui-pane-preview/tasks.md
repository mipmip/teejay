## 1. Tmux Package

- [x] 1.1 Create `internal/tmux/` directory
- [x] 1.2 Create `internal/tmux/capture.go` with CapturePane function
- [x] 1.3 Create `internal/tmux/capture_test.go` with tests

## 2. Update UI Model

- [x] 2.1 Add bubbles/viewport to Model struct for preview panel
- [x] 2.2 Add preview content and selected pane ID fields to Model
- [x] 2.3 Update `New()` to initialize viewport component

## 3. Split-Panel Layout

- [x] 3.1 Update `View()` to create split-panel layout using lipgloss
- [x] 3.2 Set list width to ~30% and preview to ~70% of terminal width
- [x] 3.3 Add border/styling to separate panels

## 4. Preview Updates

- [x] 4.1 Detect selection change in `Update()` and capture new pane content
- [x] 4.2 Update viewport content when selection changes
- [x] 4.3 Handle capture errors gracefully (show error in preview)

## 5. Tests and Verify

- [x] 5.1 Update `internal/ui/app_test.go` for split-panel layout
- [x] 5.2 Run `make test` and ensure all tests pass
- [x] 5.3 Run `make build` and test TUI with multiple panes
