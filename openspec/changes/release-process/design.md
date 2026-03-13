## Context

Teejay is a Go TUI application using Nix flakes for development. Currently:
- `flake.nix` defines the build with `pname = "tj"` and `version = "0.1.0"`
- No release automation exists
- Version is hardcoded in flake.nix only, not embedded in binary
- No maintainer documentation for releases

The project is hosted at `github.com/mipmip/teejay` and targets Linux and macOS (amd64/arm64).

## Goals / Non-Goals

**Goals:**
- Automate multi-platform binary releases via goreleaser
- Trigger releases automatically when pushing `v*` git tags
- Embed version in binary (queryable via `tj --version`)
- Document the release process for maintainers
- Maintain changelog for user visibility

**Non-Goals:**
- Package manager distribution (homebrew, apt, etc.) - future work
- Signing binaries or releases
- Automated changelog generation from commits
- Docker image releases

## Decisions

### Decision 1: Use goreleaser for release automation

**Choice:** goreleaser
**Rationale:** De facto standard for Go projects. Handles cross-compilation, checksums, GitHub release creation, and changelog inclusion. Single config file.

**Alternatives considered:**
- Manual release scripts: Error-prone, tedious
- GitHub Actions matrix build: More complex, less features
- ko: Focused on containers, not general binaries

### Decision 2: Version embedding via ldflags

**Choice:** Use `-ldflags "-X main.version=..."` pattern
**Rationale:** Standard Go practice. goreleaser sets this automatically. Works with Nix build too.

**Implementation:**
```go
// cmd/tj/main.go
var version = "dev"  // overridden at build time
```

### Decision 3: Semantic versioning with v prefix

**Choice:** Tags like `v0.1.0`, `v1.0.0`
**Rationale:** Go ecosystem convention. goreleaser expects this format.

### Decision 4: Manual changelog maintenance

**Choice:** Hand-maintained `CHANGELOG.md` using Keep a Changelog format
**Rationale:** Better quality than auto-generated. Maintainer writes meaningful release notes.

**Alternatives considered:**
- git-cliff / conventional commits: Requires commit discipline, noisy output
- GitHub release notes: Less discoverable in repo

### Decision 5: Target platforms

Build for:
- `linux/amd64`
- `linux/arm64`
- `darwin/amd64`
- `darwin/arm64`

**Rationale:** Covers common developer machines. Windows excluded (tmux doesn't run on Windows).

## Risks / Trade-offs

**[Risk]** goreleaser needs GitHub token for release creation → Document in RELEASING.md, use `GITHUB_TOKEN` secret

**[Risk]** flake.nix version may diverge from git tags → Document that flake version is for nix builds only; goreleaser uses git tag

**[Trade-off]** Manual changelog is more work → Worth it for quality; releases are infrequent
