## 1. Central VERSION File

- [x] 1.1 Create `VERSION` file with current version (0.1.0)
- [x] 1.2 Update `cmd/tj/main.go` to embed VERSION file using `//go:embed`
- [x] 1.3 Update `flake.nix` to read version from VERSION file
- [x] 1.4 Verify `go build` and `tj --version` work correctly

## 2. Release Script Setup

- [x] 2.1 Create `scripts/release.sh` with shebang and set -e
- [x] 2.2 Add gum dependency check with install instructions if missing
- [x] 2.3 Define color/style constants for output

## 3. Safety Checks

- [x] 3.1 Check for clean git working directory
- [x] 3.2 Check currently on main branch
- [x] 3.3 Check CHANGELOG.md exists and contains `[Unreleased]`

## 4. Version Calculation

- [x] 4.1 Read current version from VERSION file
- [x] 4.2 Display gum choose dropdown for bump type (major/minor/patch)
- [x] 4.3 Calculate new version based on selection
- [x] 4.4 Check if new version tag already exists

## 5. Release Execution

- [x] 5.1 Show confirmation prompt with version details
- [x] 5.2 Update VERSION file with new version
- [x] 5.3 Update CHANGELOG.md: insert new version header below `[Unreleased]`
- [x] 5.4 Commit VERSION and changelog changes with release message
- [x] 5.5 Create git tag with version
- [x] 5.6 Push commit and tag to remote

## 6. Testing

- [x] 6.1 Verify script is executable
- [x] 6.2 Test safety checks work (dirty git, wrong branch)
- [x] 6.3 Verify VERSION file is read correctly by Go and Nix
