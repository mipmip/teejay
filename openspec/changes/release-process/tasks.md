## 1. Version embedding

- [ ] 1.1 Add `var version = "dev"` variable in `cmd/tj/main.go`
- [ ] 1.2 Add `--version` and `-v` flag handling in main.go
- [ ] 1.3 Print version and exit when flag is provided
- [ ] 1.4 Test: `go run ./cmd/tj --version` shows "dev"
- [ ] 1.5 Test: `go build -ldflags "-X main.version=test" -o tj ./cmd/tj && ./tj --version` shows "test"

## 2. goreleaser configuration

- [ ] 2.1 Create `.goreleaser.yaml` with project metadata
- [ ] 2.2 Configure builds for linux/amd64, linux/arm64, darwin/amd64, darwin/arm64
- [ ] 2.3 Configure ldflags to inject version from git tag
- [ ] 2.4 Configure archive format (tar.gz for unix)
- [ ] 2.5 Configure checksum generation
- [ ] 2.6 Add goreleaser to flake.nix devShell
- [ ] 2.7 Test: `goreleaser check` passes
- [ ] 2.8 Test: `goreleaser build --snapshot --clean` creates binaries

## 3. GitHub Actions workflow

- [ ] 3.1 Create `.github/workflows/release.yml`
- [ ] 3.2 Configure trigger on `v*` tag push
- [ ] 3.3 Add checkout and Go setup steps
- [ ] 3.4 Add goreleaser action with `GITHUB_TOKEN`
- [ ] 3.5 Verify workflow syntax is valid

## 4. Changelog

- [ ] 4.1 Create `CHANGELOG.md` with Keep a Changelog format
- [ ] 4.2 Add Unreleased section header
- [ ] 4.3 Document initial features for v0.1.0 release

## 5. Maintainer documentation

- [ ] 5.1 Create `RELEASING.md` with overview
- [ ] 5.2 Add pre-release checklist (tests pass, changelog updated, etc.)
- [ ] 5.3 Document tagging process (`git tag -a v0.1.0 -m "Release v0.1.0"`)
- [ ] 5.4 Document how to push tag and trigger release
- [ ] 5.5 Document how to verify release on GitHub
- [ ] 5.6 Add troubleshooting section (common issues)

## 6. Verification

- [ ] 6.1 Run `go test ./...` to ensure existing tests pass
- [ ] 6.2 Run `goreleaser check` to validate config
- [ ] 6.3 Run `goreleaser build --snapshot --clean` to test local build
- [ ] 6.4 Review all new files for completeness
