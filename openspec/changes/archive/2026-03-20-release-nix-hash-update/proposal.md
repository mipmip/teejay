## Why

The release script (`scripts/release.sh`) does not update the `vendorHash` in `flake.nix`. When Go dependencies change between releases, `nix build` fails because the hash no longer matches. This requires manual intervention after every release that includes dependency changes.

## What Changes

- Add a step to `scripts/release.sh` that computes the correct `vendorHash` and updates it in `flake.nix` before committing the release
- Include `flake.nix` in the release commit so the nix flake is always in sync

## Capabilities

### New Capabilities

_(none)_

### Modified Capabilities
- `release-automation`: Add nix vendorHash update step to the release script
- `nix-flake`: Require that vendorHash is kept in sync during releases

## Impact

- `scripts/release.sh`: New step to compute and update vendorHash
- `flake.nix`: vendorHash value updated automatically during releases
- Release commit will include `flake.nix` in addition to VERSION and CHANGELOG.md
