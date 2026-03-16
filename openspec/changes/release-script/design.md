## Context

The project uses semantic versioning with git tags (e.g., `v0.1.0`) and maintains a CHANGELOG.md following Keep a Changelog format. Currently version is hardcoded in `cmd/tj/main.go` and `flake.nix`, leading to potential inconsistencies (#21).

## Goals / Non-Goals

**Goals:**
- Single source of truth for version via `VERSION` file
- Provide interactive version bump selection (major/minor/patch)
- Automatically update VERSION, CHANGELOG.md with version and release date
- Create and push git tags safely
- Prevent releases from dirty working directories
- Show clear progress and confirmation

**Non-Goals:**
- GitHub release creation (goreleaser handles this)
- Building binaries (goreleaser handles this)
- Complex branching strategies

## Decisions

### Decision 1: VERSION file as single source of truth

Create a plain text `VERSION` file containing just the version number (e.g., `0.2.0`). Other files read from this.

**Rationale:** Simple, universally readable, easy to update programmatically.

### Decision 2: Go embed for version in main.go

Use `//go:embed VERSION` to embed the version file content at compile time.

**Alternatives considered:**
- Read file at runtime: Requires VERSION file to exist at runtime
- Build-time ldflags: Works but VERSION file is cleaner

**Rationale:** Embeds version at build time, no runtime file dependency, stays in sync automatically.

### Decision 3: Nix reads VERSION file

Use `builtins.readFile ./VERSION` in flake.nix to read version.

**Rationale:** Native Nix approach, no external tools needed.

### Decision 4: Bash script using gum for UI

Use a bash script with [gum](https://github.com/charmbracelet/gum) for the interactive dropdown selection.

**Rationale:** Gum is simple, looks great, and aligns with the project's Charm ecosystem usage.

### Decision 5: Safety checks before release

Perform these checks before proceeding:
1. Git working directory is clean (no uncommitted changes)
2. Currently on main branch
3. CHANGELOG.md contains `[Unreleased]` section
4. Selected version tag doesn't already exist

**Rationale:** Prevents accidental releases and ensures consistency.

### Decision 6: Release script updates VERSION file

The release script updates VERSION file first, then CHANGELOG, then commits all changes together.

**Rationale:** Single commit contains all version-related changes.

### Decision 7: Changelog update format

Replace `## [Unreleased]` with:
```
## [Unreleased]

## [X.Y.Z] - YYYY-MM-DD
```

**Rationale:** Follows Keep a Changelog convention, preserves Unreleased section.

## Risks / Trade-offs

- **[Risk] gum not installed** → Check for gum and show install instructions if missing
- **[Risk] VERSION file deleted** → Build fails clearly with missing embed file
- **[Trade-off] Bash vs Go** → Bash is simpler but less portable; acceptable for dev tooling
