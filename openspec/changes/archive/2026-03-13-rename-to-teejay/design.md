## Context

The project is currently named "tmon" with:
- Module: `tmon` in go.mod
- Binary: `cmd/tmon/main.go`
- Import paths: `tmon/internal/...`

We're renaming to "Teejay" (CLI: `tj`) - "Terminal Junky" - to give the project a distinctive identity that reflects the addictive nature of terminal-based development.

## Goals / Non-Goals

**Goals:**
- Rename Go module from `tmon` to `tj`
- Move binary from `cmd/tmon/` to `cmd/tj/`
- Update all import paths across the codebase
- Preserve all existing functionality unchanged

**Non-Goals:**
- No functional changes to any component
- No new features or capabilities
- No changes to external dependencies
- No user data migration (none exists)

## Decisions

### Decision 1: Module name `tj` (not `teejay`)
CLI commands should be short and easy to type. Users will invoke `tj add`, `tj` etc. The module name matches the binary name for consistency.

**Alternatives considered:**
- `teejay`: Full name is longer to type, less convenient for frequent CLI use
- `tjunky`: Longer, harder to type

### Decision 2: Single atomic rename
Perform all renames in one change rather than phased approach.

**Rationale:** The project is small (< 20 Go files) and not yet distributed. A single coordinated rename is simpler and avoids temporary inconsistent states.

### Decision 3: Preserve directory structure
Keep `cmd/tj/main.go` pattern rather than moving to root.

**Rationale:** Standard Go project layout supports multiple binaries if needed later.

## Risks / Trade-offs

**[Risk]** IDE/editor caches may show stale import errors → Run `go mod tidy` after rename; restart language server if needed

**[Risk]** Existing `./tmon` build artifacts cause confusion → Clean build directory before rebuilding as `tj`

**[Trade-off]** Short name `tj` is less descriptive → Acceptable because memorability and typing speed matter more for CLI tools
