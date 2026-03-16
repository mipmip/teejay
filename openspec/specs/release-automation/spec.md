# release-automation Specification

## Purpose
Automate the release process with scripts, goreleaser, and GitHub Actions.

## Requirements
### Requirement: goreleaser configuration

The project SHALL have a `.goreleaser.yaml` configuration file that builds binaries for multiple platforms.

#### Scenario: Build targets
- **WHEN** goreleaser runs
- **THEN** it builds binaries for linux/amd64, linux/arm64, darwin/amd64, and darwin/arm64

#### Scenario: Binary naming
- **WHEN** goreleaser creates release artifacts
- **THEN** binaries are named `tj` with platform-specific archive names (e.g., `tj_0.1.0_linux_amd64.tar.gz`)

#### Scenario: Checksum file
- **WHEN** goreleaser completes
- **THEN** a `checksums.txt` file is generated containing SHA256 hashes of all artifacts

### Requirement: GitHub Actions release workflow

The project SHALL have a GitHub Actions workflow that triggers releases on version tags.

#### Scenario: Trigger on version tag
- **WHEN** a tag matching `v*` pattern is pushed (e.g., `v0.1.0`)
- **THEN** the release workflow runs automatically

#### Scenario: Create GitHub release
- **WHEN** the workflow completes successfully
- **THEN** a GitHub release is created with the tag name as the release title
- **AND** release artifacts are uploaded

### Requirement: Maintainer release documentation

The project SHALL have a `RELEASING.md` document with step-by-step release instructions.

#### Scenario: Document content
- **WHEN** a maintainer reads RELEASING.md
- **THEN** they find instructions for: updating changelog, creating tag, verifying release

#### Scenario: Pre-release checklist
- **WHEN** preparing a release
- **THEN** RELEASING.md includes a checklist of required steps before tagging

### Requirement: Changelog file

The project SHALL maintain a `CHANGELOG.md` file documenting changes in each release.

#### Scenario: Changelog format
- **WHEN** viewing CHANGELOG.md
- **THEN** it follows Keep a Changelog format with sections for Added, Changed, Fixed, Removed

#### Scenario: Unreleased section
- **WHEN** changes are made between releases
- **THEN** they are documented under an "Unreleased" section at the top

### Requirement: Release script

The project SHALL have a `scripts/release.sh` script for interactive release creation.

#### Scenario: Version bump selection
- **WHEN** maintainer runs `scripts/release.sh`
- **THEN** an interactive prompt allows selecting major, minor, or patch bump

#### Scenario: Safety checks
- **WHEN** release script runs
- **THEN** it verifies: clean git working directory, on main branch, changelog has [Unreleased] section

#### Scenario: Automated updates
- **WHEN** release is confirmed
- **THEN** script updates `cmd/tj/VERSION`, updates CHANGELOG.md with version and date, commits, tags, and pushes

