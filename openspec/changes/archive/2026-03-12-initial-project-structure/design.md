## Context

This is a greenfield Go project for tmon, a tmux activity monitor TUI. No code exists yet. We need to establish the foundation: Go module, directory structure, and a minimal Bubbletea application that proves the stack works.

## Goals / Non-Goals

**Goals:**
- Initialize a working Go module with proper dependency management
- Create a directory structure that scales as the project grows
- Build a minimal Bubbletea TUI that compiles and runs
- Verify the Charm stack (bubbletea, lipgloss, bubbles) is properly integrated

**Non-Goals:**
- Implementing any actual tmon features (watchlist, pane preview, etc.)
- Creating a polished UI
- Adding tests (will come with feature work)

## Decisions

### Decision 1: Directory structure follows Go conventions

```
tmon/
├── cmd/tmon/          # CLI entry point
│   └── main.go
├── internal/          # Private packages
│   └── ui/            # TUI components (future)
│       └── app.go     # Bubbletea model
├── go.mod
└── go.sum
```

**Rationale**: Standard Go layout. `cmd/` for executables, `internal/` for private packages that can't be imported externally. Keeps the root clean.

**Alternative considered**: Flat structure with just `main.go` at root. Rejected because tmon will grow to have multiple packages and this structure scales better.

### Decision 2: Module path uses generic placeholder

Module path: `tmon` (simple, can be updated later if published)

**Rationale**: Keeps things simple for now. If the project gets published to GitHub, we can update the module path then.

### Decision 3: Minimal Bubbletea model

The initial TUI will be a single-file Bubbletea model that:
- Displays "tmon - press q to quit"
- Handles 'q' key to exit
- Uses basic lipgloss styling

**Rationale**: Proves the stack works without over-engineering. Features will be added incrementally.

## Risks / Trade-offs

- [Risk] Module path may need updating for distribution → Can run `go mod edit -module` later
- [Trade-off] Minimal UI provides no real functionality → Acceptable for scaffolding; features come next
