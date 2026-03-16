# version-embedding Specification

## Purpose
Embed version information in the binary from a single source of truth file.

## Requirements

### Requirement: Version flag

The CLI SHALL support a `--version` flag that displays the current version.

#### Scenario: Display version
- **WHEN** user runs `tj --version`
- **THEN** the version string is printed to stdout (e.g., `0.2.0`)

#### Scenario: Short version flag
- **WHEN** user runs `tj -v`
- **THEN** the version string is printed (same as `--version`)

### Requirement: Single source of truth VERSION file

The version SHALL be stored in `cmd/tj/VERSION` as the single source of truth, embedded at compile time via `go:embed`.

#### Scenario: Version embedded from file
- **WHEN** built with `go build ./cmd/tj`
- **THEN** the version matches the content of `cmd/tj/VERSION`

#### Scenario: Nix build reads version
- **WHEN** built via `nix build`
- **THEN** flake.nix reads version from `cmd/tj/VERSION`

