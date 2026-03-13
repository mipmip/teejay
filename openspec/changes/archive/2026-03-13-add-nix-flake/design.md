## Context

tmon is a Go TUI application for monitoring tmux panes. It uses Go 1.25.5 with Bubbletea/Lipgloss. The main entry point is `cmd/tmon/main.go`. There are no separate Go modules or submodules to exclude.

## Goals / Non-Goals

**Goals:**
- Provide `nix build` to compile tmon
- Provide `nix run` to run tmon directly
- Provide `nix develop` for a development shell with Go tooling
- Support all common platforms (x86_64-linux, aarch64-linux, x86_64-darwin, aarch64-darwin)

**Non-Goals:**
- Version injection via ldflags (no existing version pattern in codebase)
- NixOS module or home-manager integration
- Publishing to nixpkgs

## Decisions

### 1. Use buildGoModule

Standard Nix function for Go projects. Handles vendorHash for reproducible dependency fetching.

### 2. Initial vendorHash = null

Start with `null`, then run `nix build` to get the actual hash from the error output. Update flake.nix with the real hash.

### 3. Minimal ldflags

Use only `-s -w` for smaller binary. No version injection needed since there's no existing version pattern.

### 4. MIT license in meta

Match the project's license.

### 5. Development shell with standard Go tools

Include: go, gopls, gotools, go-tools for a complete dev environment.

## Risks / Trade-offs

- **vendorHash maintenance**: Must be updated when go.mod dependencies change. Document this in README or CONTRIBUTING.
- **Go version**: Nix uses its own Go version, may differ from go.mod's 1.25.5. Generally not an issue for compatibility.
