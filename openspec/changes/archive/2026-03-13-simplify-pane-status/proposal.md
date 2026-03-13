## Why

The current three-state pane status (Idle, Running, Ready) is more complex than needed. Users primarily care about two things: is the pane busy doing something, or is it waiting for my input? The "Idle" state (content stable but no prompt) adds confusion without providing actionable information.

## What Changes

- Reduce pane status from three states to two:
  - **Busy**: Pane content is changing or stable without a prompt (animated spinner)
  - **Waiting**: Pane has a prompt and is waiting for user input (green indicator)
- Remove the "Idle" state entirely - if there's no prompt, the pane is considered busy
- Simplify the monitor logic by removing idle counter tracking
- Update status indicators: green dot for waiting, animated spinner for busy

## Capabilities

### New Capabilities

None - this is a simplification of existing behavior.

### Modified Capabilities

- `pane-status`: Simplify from three states (Idle/Running/Ready) to two states (Busy/Waiting)

## Impact

- `internal/monitor/status.go`: Remove `Idle` status, rename `Running` to `Busy` and `Ready` to `Waiting`
- `internal/monitor/monitor.go`: Remove idle counter logic, simplify state machine
- `internal/monitor/monitor_test.go`: Update tests for new two-state model
- `internal/ui/app.go`: Update any references to status indicators/colors
