## ADDED Requirements

### Requirement: Central VERSION file

The project SHALL use a single VERSION file as the source of truth for version numbers.

#### Scenario: VERSION file exists
- **WHEN** the project is built
- **THEN** a VERSION file exists in the repository root
- **AND** it contains only the semantic version number (e.g., `0.2.0`)

#### Scenario: Go binary reads VERSION
- **WHEN** the Go binary is built
- **THEN** the version is embedded from the VERSION file
- **AND** `tj --version` outputs the correct version

#### Scenario: Nix flake reads VERSION
- **WHEN** the Nix flake builds the package
- **THEN** the version is read from the VERSION file

### Requirement: Interactive version selection

The release script SHALL provide an interactive dropdown to select the version bump type.

#### Scenario: User selects version bump type
- **WHEN** user runs the release script
- **THEN** a dropdown is displayed with options: major (1.x.x), minor (x.1.x), patch (x.x.1)
- **AND** the new version number is calculated and shown based on selection

### Requirement: Safety checks before release

The release script SHALL verify preconditions before proceeding with a release.

#### Scenario: Dirty working directory
- **WHEN** user runs release script with uncommitted changes
- **THEN** the script exits with an error message about uncommitted changes

#### Scenario: Not on main branch
- **WHEN** user runs release script from a non-main branch
- **THEN** the script exits with an error message about being on wrong branch

#### Scenario: Version tag already exists
- **WHEN** user selects a version that already has a git tag
- **THEN** the script exits with an error message about duplicate tag

#### Scenario: Missing Unreleased section
- **WHEN** CHANGELOG.md does not contain `[Unreleased]` section
- **THEN** the script exits with an error message

### Requirement: VERSION file update

The release script SHALL update the VERSION file with the new version.

#### Scenario: Successful VERSION update
- **WHEN** user confirms the release
- **THEN** the VERSION file is updated with the new version number

### Requirement: Changelog update

The release script SHALL update CHANGELOG.md with the new version and date.

#### Scenario: Successful changelog update
- **WHEN** user confirms the release
- **THEN** `## [Unreleased]` is followed by a new `## [X.Y.Z] - YYYY-MM-DD` section
- **AND** the date is the current date in ISO format

### Requirement: Git tag creation and push

The release script SHALL create a git tag and push changes.

#### Scenario: Successful release
- **WHEN** user confirms the release
- **THEN** the VERSION and changelog changes are committed together
- **AND** a git tag `vX.Y.Z` is created
- **AND** changes and tag are pushed to remote

### Requirement: Confirmation before release

The release script SHALL require user confirmation before making changes.

#### Scenario: User confirms release
- **WHEN** user is shown the version to be released
- **THEN** user must confirm before any file modifications, commit, tag, and push occur

#### Scenario: User cancels release
- **WHEN** user declines confirmation
- **THEN** no changes are made and script exits cleanly
