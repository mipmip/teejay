## 1. Initialize Go Module

- [x] 1.1 Run `go mod init tmon` to create go.mod
- [x] 1.2 Add Charm dependencies: bubbletea, lipgloss, bubbles

## 2. Create Directory Structure

- [x] 2.1 Create `cmd/tmon/` directory for CLI entry point
- [x] 2.2 Create `internal/ui/` directory for TUI components

## 3. Implement Minimal TUI

- [x] 3.1 Create `internal/ui/app.go` with Bubbletea model (Model, Init, Update, View)
- [x] 3.2 Create `cmd/tmon/main.go` entry point that runs the TUI
- [x] 3.3 Add basic lipgloss styling for the welcome message

## 4. Verify

- [x] 4.1 Run `go mod tidy` to ensure dependencies resolve
- [x] 4.2 Run `go run ./cmd/tmon` to verify TUI starts and 'q' quits
