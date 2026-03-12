## 1. Update Capture Command

- [x] 1.1 Update `CapturePane` in `internal/tmux/capture.go` to add `-e` flag for ANSI preservation
- [x] 1.2 Add `-J` flag to join wrapped lines

## 2. Verify

- [x] 2.1 Run `make test` and ensure all tests pass
- [x] 2.2 Run `make build` and test preview with colored terminal content
