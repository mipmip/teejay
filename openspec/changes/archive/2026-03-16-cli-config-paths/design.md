## Context

Teejay currently hardcodes config and watchlist paths to `~/.config/teejay/`. The config.Load() and watchlist.Load() functions internally call ConfigPath() which constructs the path from the user's home directory.

Users want to:
- Run multiple teejay instances with different watchlists (e.g., per-project)
- Test configurations without modifying their main config
- Use teejay in CI/test environments with custom paths

## Goals / Non-Goals

**Goals:**
- Add --config/-c and --watchlist/-w flags to the CLI
- Modify Load functions to accept optional custom paths
- Ensure custom watchlist path is used for saves (not just loads)
- Maintain backward compatibility (no flags = current behavior)

**Non-Goals:**
- Environment variable support (can be added later)
- Config file discovery (searching parent directories)
- Merging multiple config files

## Decisions

### 1. Flag Parsing Approach: Manual parsing before subcommand dispatch

**Decision:** Parse --config and --watchlist flags manually in main.go before the switch on subcommands, storing in package-level or passed variables.

**Rationale:**
- Current CLI uses simple os.Args switching, not a flag library
- Adding flag library (cobra, urfave/cli) is overkill for two flags
- Manual parsing keeps consistency with existing code style

**Alternatives considered:**
- Use `flag` stdlib: Conflicts with subcommand structure, would need FlagSet per command
- Use cobra: Heavy dependency for simple needs

### 2. API Change: Optional path parameter via variadic

**Decision:** Change `config.Load()` to `config.Load(customPath ...string)` and same for watchlist.

**Rationale:**
- Variadic allows zero or one argument, backward compatible
- Cleaner than LoadFromPath() separate function
- All existing Load() calls continue to work unchanged

**Alternatives considered:**
- LoadFromPath(path string) + Load(): Two entry points, more code
- Load(path *string): Requires pointer, awkward call sites

### 3. Watchlist Path Threading: Store in Watchlist struct

**Decision:** Add a `path` field to `Watchlist` struct, set on Load, used in Save.

**Rationale:**
- Watchlist.Save() needs to know where to save
- Storing path in struct avoids passing it through UI layer
- Natural place for the state

**Alternatives considered:**
- Pass path to Save(path string): Requires threading path through all callers
- Global variable: Implicit state, harder to test

## Risks / Trade-offs

**[API Change]** Modifying Load() signatures could break external consumers
→ Mitigation: Using variadic makes it backward compatible; no external consumers known

**[Save Path]** If user loads default watchlist, modifies, but TUI was started with custom path, confusion could occur
→ Mitigation: Path is set once at load time and consistently used; this is expected behavior

**[Flag Position]** Flags must come before subcommand (e.g., `tj -c foo.yaml add` not `tj add -c foo.yaml`)
→ Mitigation: Document this; consistent with many CLIs (docker, kubectl)
