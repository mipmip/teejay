## Why

Teejay is approaching its first release and needs a solid, repeatable release process. Currently there's no automation for building release binaries, creating GitHub releases, or documentation for maintainers on how to perform releases. A well-documented release process ensures consistent, reliable releases and reduces friction for the maintainer.

## What Changes

- Add goreleaser configuration for automated multi-platform binary builds
- Create GitHub Actions workflow for release automation (triggered by git tags)
- Add `RELEASING.md` maintainer documentation with step-by-step release instructions
- Add `CHANGELOG.md` to track release history
- Embed version information in the binary via ldflags

## Capabilities

### New Capabilities

- `release-automation`: Configuration for goreleaser and GitHub Actions to automate the release pipeline
- `version-embedding`: Embed version string in binary at build time via ldflags

### Modified Capabilities

None - this is infrastructure/tooling, not changing application behavior.

## Impact

- **New files**: `.goreleaser.yaml`, `.github/workflows/release.yml`, `RELEASING.md`, `CHANGELOG.md`
- **Modified files**: `flake.nix` (add goreleaser to devShell), `cmd/tj/main.go` (add version flag)
- **Build process**: Releases triggered by pushing `v*` tags to GitHub
- **Dependencies**: goreleaser (dev dependency only)
