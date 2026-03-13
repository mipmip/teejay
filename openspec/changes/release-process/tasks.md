## 1. Version embedding

- [x] 1.1 Add `var version = "dev"` variable in `cmd/tj/main.go`
- [x] 1.2 Add `--version` and `-v` flag handling in main.go
- [x] 1.3 Print version and exit when flag is provided
- [x] 1.4 Test: `go run ./cmd/tj --version` shows "dev"
- [x] 1.5 Test: `go build -ldflags "-X main.version=test" -o tj ./cmd/tj && ./tj --version` shows "test"

## 2. goreleaser configuration

- [x] 2.1 Create `.goreleaser.yaml` with project metadata
- [x] 2.2 Configure builds for linux/amd64, linux/arm64, darwin/amd64, darwin/arm64
- [x] 2.3 Configure ldflags to inject version from git tag
- [x] 2.4 Configure archive format (tar.gz for unix)
- [x] 2.5 Configure checksum generation
- [x] 2.6 Add goreleaser to flake.nix devShell
- [x] 2.7 Test: `goreleaser check` passes
- [x] 2.8 Test: `goreleaser build --snapshot --clean` creates binaries

## 3. GitHub Actions workflow

- [x] 3.1 Create `.github/workflows/release.yml`
- [x] 3.2 Configure trigger on `v*` tag push
- [x] 3.3 Add checkout and Go setup steps
- [x] 3.4 Add goreleaser action with `GITHUB_TOKEN`
- [x] 3.5 Verify workflow syntax is valid

## 4. Changelog

- [x] 4.1 Create `CHANGELOG.md` with Keep a Changelog format
- [x] 4.2 Add Unreleased section header
- [x] 4.3 Document initial features for v0.1.0 release

## 5. Maintainer documentation

- [x] 5.1 Create `RELEASING.md` with overview
- [x] 5.2 Add pre-release checklist (tests pass, changelog updated, etc.)
- [x] 5.3 Document tagging process (`git tag -a v0.1.0 -m "Release v0.1.0"`)
- [x] 5.4 Document how to push tag and trigger release
- [x] 5.5 Document how to verify release on GitHub
- [x] 5.6 Add troubleshooting section (common issues)

## 6. Verification

- [x] 6.1 Run `go test ./...` to ensure existing tests pass
- [x] 6.2 Run `goreleaser check` to validate config
- [x] 6.3 Run `goreleaser build --snapshot --clean` to test local build
- [x] 6.4 Review all new files for completeness
