## MODIFIED Requirements

### Requirement: Safety checks before release

The release script SHALL verify preconditions before proceeding with a release. The checks SHALL work in both pure-git and jj-colocated repositories.

#### Scenario: Dirty working directory
- **WHEN** user runs release script with uncommitted changes (staged or unstaged)
- **THEN** the script SHALL exit with an error message, detected via `git diff --quiet` and `git diff --cached --quiet`

#### Scenario: Version tag already exists
- **WHEN** user selects a version that already has a git tag
- **THEN** the script exits with an error message about duplicate tag

#### Scenario: Missing Unreleased section
- **WHEN** CHANGELOG.md does not contain `[Unreleased]` section
- **THEN** the script exits with an error message

### Requirement: Git tag creation and push

The release script SHALL create a git tag and push changes using commands that work in both pure-git and jj-colocated repos.

#### Scenario: Successful release
- **WHEN** user confirms the release
- **THEN** the VERSION, changelog, and flake.nix changes are committed together
- **AND** a git tag `vX.Y.Z` is created
- **AND** changes are pushed to remote main via `git push origin HEAD:main`
- **AND** the tag is pushed via `git push origin vX.Y.Z`

## REMOVED Requirements

### Requirement: Not on main branch check
**Reason**: The branch check (`git branch --show-current`) fails under jj-colocated repos (detached HEAD). The check is unnecessary — the push to `HEAD:main` ensures changes land on main regardless.
**Migration**: The branch check is removed. The `git push origin HEAD:main` command serves as the implicit guarantee that changes target main.
