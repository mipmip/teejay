## 1. Create flake.nix

- [x] 1.1 Create flake.nix with description "tmon - tmux activity monitor"
- [x] 1.2 Add nixpkgs input (github:NixOS/nixpkgs/nixos-unstable)
- [x] 1.3 Add packages output with buildGoModule (pname="tmon", vendorHash=null initially)
- [x] 1.4 Add ldflags for -s -w
- [x] 1.5 Add meta with description, homepage, MIT license, mainProgram="tmon"
- [x] 1.6 Add apps output pointing to tmon binary
- [x] 1.7 Add devShells output with go, gopls, gotools, go-tools

## 2. Generate lock file and resolve vendorHash

- [x] 2.1 Run `git add flake.nix` to register with git
- [x] 2.2 Run `nix flake update` to generate flake.lock
- [x] 2.3 Run `nix build` to get vendorHash from error output
- [x] 2.4 Update flake.nix with correct vendorHash

## 3. Verify build and functionality

- [x] 3.1 Run `nix build --no-link` to verify successful build
- [x] 3.2 Run `nix run . -- --help` or similar to verify binary works
- [x] 3.3 Run `nix flake check` to validate flake
- [x] 3.4 Run `nix develop -c go version` to verify dev shell

## 4. Update documentation

- [x] 4.1 Add Nix/NixOS section to README.md with installation and usage instructions
