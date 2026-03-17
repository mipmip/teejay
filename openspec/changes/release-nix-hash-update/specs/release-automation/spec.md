## MODIFIED Requirements

### Requirement: Automated updates

- **WHEN** release is confirmed
- **THEN** script updates `cmd/tj/VERSION`, updates CHANGELOG.md with version and date, updates `vendorHash` in `flake.nix`, commits, tags, and pushes

#### Scenario: Successful release with nix hash update
- **WHEN** user confirms the release
- **AND** `nix` is available on the system
- **THEN** the vendorHash in `flake.nix` is updated to the correct value
- **AND** `flake.nix` is included in the release commit

#### Scenario: Release without nix available
- **WHEN** user confirms the release
- **AND** `nix` is NOT available on the system
- **THEN** a warning is shown that vendorHash was not updated
- **AND** the release proceeds without updating `flake.nix`

### Requirement: Git tag creation and push

The release script SHALL create a git tag and push changes.

#### Scenario: Successful release
- **WHEN** user confirms the release
- **THEN** the VERSION, changelog, and flake.nix changes are committed together
- **AND** a git tag `vX.Y.Z` is created
- **AND** changes and tag are pushed to remote
