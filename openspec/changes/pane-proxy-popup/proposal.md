## Why

Currently, pressing Enter on a pane switches tmux to that pane, which takes the user away from TJ. To return, they must manually navigate back. This breaks the monitoring workflow when users want to quickly interact with a pane and return.

A popup-based proxy would let users interact with a pane without leaving TJ, maintaining context and enabling quick return.

## What Changes

- Add a new "proxy" subcommand (`tj proxy <pane-id>`) that creates an interactive terminal proxy to a tmux pane
- Modify the Enter key behavior in the TUI to open the pane in a tmux popup running the proxy instead of switching
- The proxy captures pane content with ANSI escapes and displays it, while forwarding all keyboard input via `send-keys`
- Double-Escape exits the proxy and closes the popup, returning to TJ

## Capabilities

### New Capabilities

- `pane-proxy`: Interactive terminal proxy that mirrors a tmux pane's display and forwards input, enabling interaction without switching away from the current context

### Modified Capabilities

- `preview-switch`: Change Enter key to open pane in popup proxy instead of switching (when running in tmux)

## Impact

- New file: `internal/proxy/proxy.go` - proxy implementation
- New file: `internal/cmd/proxy.go` - CLI subcommand
- Modified: `internal/ui/app.go` - Enter key handler to spawn popup
- Modified: `internal/tmux/` - may need additional capture/send-keys helpers
- Dependency: Requires tmux 3.2+ for `display-popup` support
