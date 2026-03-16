## Context

The `tj` CLI currently handles flags via a manual `parseFlags` function in `cmd/tj/main.go`. It supports `--version`/`-v`, `--config`/`-c`, and `--watchlist`/`-w`, plus subcommands `add` and `del`. There is no help flag, so users must read source code or external docs to discover usage.

## Goals / Non-Goals

**Goals:**
- Add `--help` and `-h` flags to display usage information
- Show all available commands and global flags in a readable format
- Maintain consistency with existing flag parsing patterns

**Non-Goals:**
- Adding subcommand-specific help (e.g., `tj add --help`)
- Using an external flag-parsing library (keep manual parsing for simplicity)
- Colored/styled help output

## Decisions

### 1. Help flag handling location

**Decision**: Handle `--help`/`-h` in the same switch statement as `--version`

**Rationale**: Follows the existing pattern where early-exit flags are handled as pseudo-commands in the `switch args[0]` block. This keeps all command/flag dispatch in one place.

**Alternative considered**: Handle in `parseFlags()` - rejected because `parseFlags` is for extracting values, not triggering actions.

### 2. Help text format

**Decision**: Use a simple multi-line string with manual formatting

**Rationale**: The CLI is small (3 commands, 4 flags). A hardcoded string is straightforward and easy to maintain. No need for auto-generation complexity.

**Alternative considered**: Using `text/template` or generating from a data structure - rejected as over-engineering for this scope.

## Risks / Trade-offs

**[Help text maintenance]** → Manual help text can drift from actual behavior. Mitigation: Keep help text close to command handling code for easy updates.

**[No subcommand help]** → Users can't get help for specific commands. Mitigation: Document this as a future enhancement if needed.
