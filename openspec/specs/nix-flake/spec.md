# nix-flake Specification

## Purpose
TBD - created by archiving change add-nix-flake. Update Purpose after archive.
## Requirements
### Requirement: Build tmon with Nix

The system SHALL provide a Nix flake that builds tmon using buildGoModule.

#### Scenario: Build succeeds
- **WHEN** user runs `nix build`
- **THEN** tmon binary is built successfully in `result/bin/tmon`

#### Scenario: Build is reproducible
- **WHEN** user runs `nix build` on any supported system
- **THEN** the build uses pinned dependencies from flake.lock

### Requirement: Run tmon with Nix

The system SHALL provide an app output for running tmon directly.

#### Scenario: Run without install
- **WHEN** user runs `nix run . -- --help`
- **THEN** tmon executes and shows help output

### Requirement: Development shell

The system SHALL provide a development shell with Go tooling.

#### Scenario: Enter dev shell
- **WHEN** user runs `nix develop`
- **THEN** shell includes go, gopls, gotools, and go-tools

#### Scenario: Go version available
- **WHEN** user runs `nix develop -c go version`
- **THEN** Go version is displayed

### Requirement: Multi-platform support

The system SHALL support common platforms.

#### Scenario: Supported systems
- **WHEN** flake is evaluated
- **THEN** packages are available for x86_64-linux, aarch64-linux, x86_64-darwin, aarch64-darwin

### Requirement: Flake validation

The system SHALL pass Nix flake checks.

#### Scenario: Flake check passes
- **WHEN** user runs `nix flake check`
- **THEN** no errors are reported

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

### Requirement: README documentation

The system SHALL document Nix installation in README.

#### Scenario: Nix section present
- **WHEN** user reads README.md
- **THEN** there is a section explaining `nix run` and `nix develop` usage

