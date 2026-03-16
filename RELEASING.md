# Releasing Teejay

This document describes the release process for Teejay.

## Overview

Releases are automated via GitHub Actions and goreleaser. When a `v*` tag is pushed to GitHub, the release workflow builds binaries for all supported platforms and creates a GitHub Release with the artifacts.

## Supported Platforms

- Linux (amd64, arm64)
- macOS (amd64, arm64)

## Pre-Release Checklist

Before creating a release, ensure:

- [ ] All tests pass: `go test ./...`
- [ ] The application builds: `go build ./cmd/tj`
- [ ] `CHANGELOG.md` is updated with the new version's changes
- [ ] The `[Unreleased]` section has been moved to a versioned section
- [ ] Version number follows semantic versioning (e.g., `0.1.0`, `1.0.0`)
- [ ] in cmd/tj/main.go update the version number
- [ ] in flake.nix update the version number

## Creating a Release

### 1. Update the Changelog

Edit `CHANGELOG.md`:
- Rename the `[Unreleased]` section to `[X.Y.Z] - YYYY-MM-DD`
- Add a new empty `[Unreleased]` section at the top
- Commit the change: `git commit -am "Prepare release vX.Y.Z"`

### 2. Create and Push the Tag

```bash
# Create an annotated tag
git tag -a v0.1.0 -m "Release v0.1.0"

# Push the tag to GitHub
git push origin v0.1.0
```

### 3. Wait for the Release

The GitHub Action will automatically:
1. Build binaries for all platforms
2. Create checksums
3. Create a GitHub Release with all artifacts

This typically takes 2-3 minutes.

### 4. Verify the Release

1. Go to the [Releases page](https://github.com/mipmip/teejay/releases)
2. Verify the new release appears with the correct tag
3. Check that all platform binaries are attached:
   - `tj_X.Y.Z_linux_amd64.tar.gz`
   - `tj_X.Y.Z_linux_arm64.tar.gz`
   - `tj_X.Y.Z_darwin_amd64.tar.gz`
   - `tj_X.Y.Z_darwin_arm64.tar.gz`
   - `checksums.txt`
4. Optionally download and test a binary

## Local Testing

Before pushing a tag, you can test the release locally:

```bash
# Validate goreleaser config
goreleaser check

# Build a snapshot (doesn't create a release)
goreleaser build --snapshot --clean

# Check the built binaries
ls -la dist/
```

## Troubleshooting

### Release workflow failed

1. Check the Actions tab on GitHub for error details
2. Common issues:
   - Missing `GITHUB_TOKEN` permissions (should be automatic)
   - goreleaser config errors (run `goreleaser check` locally)
   - Build failures (run `go build ./cmd/tj` locally)

### Tag already exists

If you need to redo a release:
```bash
# Delete the local tag
git tag -d v0.1.0

# Delete the remote tag
git push origin :refs/tags/v0.1.0

# Delete the GitHub Release manually via the web UI

# Create and push the tag again
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

### Version not showing in binary

Ensure the tag follows the `vX.Y.Z` format. goreleaser extracts the version from the git tag and injects it via ldflags.

Test locally:
```bash
go build -ldflags "-X main.version=test" -o tj ./cmd/tj
./tj --version
```
