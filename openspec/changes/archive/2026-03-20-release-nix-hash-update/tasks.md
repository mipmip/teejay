## 1. Nix vendorHash update function

- [x] 1.1 Add a `update_nix_vendor_hash` function to `scripts/release.sh` that: temporarily sets vendorHash to an empty string in `flake.nix`, runs `nix build` to get the expected hash from stderr, then updates `flake.nix` with the correct hash
- [x] 1.2 Add a `nix` availability check — if nix is not installed, warn and skip the hash update

## 2. Integrate into release flow

- [x] 2.1 Call the hash update function after VERSION/CHANGELOG updates but before the git commit
- [x] 2.2 Add `flake.nix` to the `git add` command in the release commit step

## 3. Verify

- [x] 3.1 Run the hash update function manually (outside of a release) to verify it correctly computes and updates the vendorHash
