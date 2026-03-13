## Why

Enable Nix users to install and run tmon through the Nix package manager. This provides reproducible builds, easy installation via `nix run`, and a consistent development environment for contributors using Nix.

## What Changes

- Add `flake.nix` with buildGoModule derivation, development shell, and multi-system support (x86_64/aarch64 Linux and Darwin)
- Add `flake.lock` for pinned dependencies
- Update README with Nix installation instructions

## Capabilities

### New Capabilities

- `nix-flake`: Nix flake configuration for building, running, and developing tmon

### Modified Capabilities

(none)

## Impact

- Root directory: new `flake.nix` and `flake.lock` files
- `README.md`: add Nix/NixOS installation section
