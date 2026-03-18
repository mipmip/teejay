## ADDED Requirements

### Requirement: vendorHash kept in sync during releases

The nix flake's vendorHash SHALL be updated during the release process to match the current Go module dependencies.

#### Scenario: Hash updated on release
- **WHEN** a release is created via `scripts/release.sh`
- **AND** Go dependencies have changed since the last release
- **THEN** the vendorHash in `flake.nix` reflects the new dependencies
- **AND** `nix build` succeeds after the release

#### Scenario: Hash unchanged when deps unchanged
- **WHEN** a release is created via `scripts/release.sh`
- **AND** Go dependencies have NOT changed since the last release
- **THEN** the vendorHash in `flake.nix` remains the same
- **AND** `nix build` continues to succeed
