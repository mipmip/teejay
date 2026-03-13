## 1. Activity Monitor Package

- [x] 1.1 Create `internal/monitor/status.go` with `PaneStatus` enum (Running, Ready, Idle)
- [x] 1.2 Create `internal/monitor/monitor.go` with `Monitor` struct holding hash and idle counter per pane
- [x] 1.3 Add `Update(paneID, content string) PaneStatus` method that computes hash, compares, and returns status
- [x] 1.4 Add `hasPrompt(content string) bool` helper with Claude/Aider prompt patterns
- [x] 1.5 Add tests for Monitor in `internal/monitor/monitor_test.go`

## 2. UI Integration

- [x] 2.1 Add `monitor *monitor.Monitor` field to UI Model
- [x] 2.2 Initialize Monitor in `New()` function
- [x] 2.3 Call `monitor.Update()` in tick handler for selected pane
- [x] 2.4 Store status in `paneStatuses map[string]monitor.PaneStatus` in Model

## 3. Status Display

- [x] 3.1 Update `paneItem` to include status field
- [x] 3.2 Modify `paneItem.Title()` to prepend status indicator (● Running, ? Ready, ○ Idle)
- [x] 3.3 Update `refreshList()` to populate status from `paneStatuses` map

## 4. Verify

- [x] 4.1 Run `make test` and ensure all tests pass
- [x] 4.2 Run `make build` and verify status indicators appear in TUI
- [x] 4.3 Test with an active pane to confirm Running/Idle transitions
