# version-embedding Specification

## Purpose
TBD - created by archiving change release-process. Update Purpose after archive.
## Requirements
### Requirement: Version flag

The CLI SHALL support a `--version` flag that displays the current version.

#### Scenario: Display version
- **WHEN** user runs `tj --version`
- **THEN** the version string is printed to stdout (e.g., `tj version 0.1.0`)

#### Scenario: Short version flag
- **WHEN** user runs `tj -v`
- **THEN** the version string is printed (same as `--version`)

### Requirement: Version variable for build-time injection

The binary SHALL have a version variable that can be set at compile time via ldflags.

#### Scenario: Default version
- **WHEN** built without ldflags (e.g., `go build`)
- **THEN** the version displays as `dev`

#### Scenario: Injected version
- **WHEN** built with `-ldflags "-X main.version=0.1.0"`
- **THEN** the version displays as `0.1.0`

#### Scenario: goreleaser sets version
- **WHEN** built via goreleaser from a git tag `v0.1.0`
- **THEN** the version is automatically set to `0.1.0`

