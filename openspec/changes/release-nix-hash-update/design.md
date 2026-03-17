## Context

The release script (`scripts/release.sh`) automates version bumping, changelog updates, git commits, and tag pushing. The nix flake (`flake.nix`) uses a `vendorHash` for Go module dependency verification. When dependencies change, this hash becomes stale, breaking `nix build` after a release.

Currently the release commit includes only `cmd/tj/VERSION` and `CHANGELOG.md`. The `flake.nix` vendorHash must be updated manually.

## Goals / Non-Goals

**Goals:**
- Automatically compute and update `vendorHash` in `flake.nix` during the release script
- Include `flake.nix` in the release commit

**Non-Goals:**
- Updating `flake.lock` (nixpkgs pin) during releases — that's a separate maintenance task
- CI-based hash updates — this is a local release script concern

## Decisions

### Decision 1: Use `nix-prefetch` to compute vendorHash

Run `nix build .#default.goModules --no-link 2>&1` with an intentionally wrong hash, then parse the expected hash from the error output. This is the standard nix approach for getting the correct hash.

**Alternative considered:** Using `nix-prefetch-url` or `go mod vendor` + manual hashing. These are more complex and less reliable than letting nix itself compute the hash.

**Simpler alternative:** Use `nix build .#default.goModules` with `vendorHash = lib.fakeHash` temporarily, capture the expected hash from stderr. However, modifying flake.nix twice is fragile.

**Chosen approach:** Use `nix-prefetch` approach — set vendorHash to empty string, attempt build, parse correct hash from error. Actually the simplest: use `nix hash path` on the vendor directory after `go mod vendor`, or just use `nix build` with a known-bad hash. But the cleanest is: temporarily set the hash to `lib.fakeHash`, run `nix build`, capture the correct hash.

**Final approach:** Use `sed` to temporarily replace the vendorHash with an empty string (`""`), run `nix build` which will fail and print the expected hash, parse it, then `sed` the correct hash back in. This is the standard pattern used by nix maintainers.

### Decision 2: Place the hash update after VERSION update but before git commit

The hash update step goes between "Update VERSION/CHANGELOG" and "git commit" in the release flow. This way the release commit includes the correct vendorHash.

## Risks / Trade-offs

- **[Risk] `nix` not available on release machine**: Not all maintainers may have nix installed. → Mitigation: Check for `nix` binary; if not found, warn and skip the hash update (don't block the release).
- **[Risk] Hash computation takes time**: `nix build` downloads and hashes all Go deps. → Acceptable: This is a one-time cost per release (~10-30s).
- **[Trade-off] Requires network access**: `nix build` may need to fetch Go modules. → Acceptable: Releases already require network for git push.
