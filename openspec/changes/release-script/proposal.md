## Why

Currently there's no automated way to create releases. The process requires manually updating the changelog, creating git tags, and pushing - which is error-prone and tedious. Additionally, version is hardcoded in multiple places (`cmd/tj/main.go`, `flake.nix`) making it easy to have inconsistent versions (addresses #21).

A release script will standardize the process, prevent mistakes, and maintain a single source of truth for versioning.

## What Changes

- Add a `VERSION` file as the single source of truth for version
- Update `cmd/tj/main.go` to read version from VERSION file (or embed at build)
- Update `flake.nix` to read version from VERSION file
- Add a `scripts/release.sh` script that automates the release process:
  - Interactive version bump selection (major, minor, patch)
  - Updates VERSION file with new version
  - Updates CHANGELOG.md: replace `[Unreleased]` with version and date
  - Git tag creation and push with safety checks

## Capabilities

### New Capabilities

- `release-automation`: Shell script for automated semantic versioning releases with central version file management

### Modified Capabilities

None

## Impact

- New file: `VERSION` - single source of truth for version number
- New file: `scripts/release.sh` - interactive release automation script
- Modified: `cmd/tj/main.go` - read version from VERSION file
- Modified: `flake.nix` - read version from VERSION file
- No new dependencies
