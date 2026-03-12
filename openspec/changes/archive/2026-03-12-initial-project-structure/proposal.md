## Why

The tmon project needs a proper Go project foundation before any features can be built. Without a Go module, entry point, and organized directory structure, we cannot begin implementing the tmux activity monitor TUI.

## What Changes

- Initialize Go module (`github.com/user/tmon`)
- Create main entry point with basic Bubbletea TUI skeleton
- Set up directory structure following Go conventions
- Add initial dependencies (bubbletea, lipgloss, bubbles)

## Capabilities

### New Capabilities
- `project-scaffold`: Go module initialization, directory layout, and main entry point with minimal Bubbletea TUI that displays a placeholder message

### Modified Capabilities
<!-- None - this is a greenfield project -->

## Impact

- Creates `go.mod` and `go.sum` at project root
- Creates `main.go` entry point
- Creates `internal/` directory for internal packages
- Creates `cmd/tmon/` for the CLI application
