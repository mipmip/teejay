{
  description = "Teejay - Terminal Junky - tmux activity monitor for vibe coders";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.buildGoModule rec {
            pname = "tj";
            version = builtins.replaceStrings ["\n"] [""] (builtins.readFile ./VERSION);

            src = ./.;
            vendorHash = "sha256-yAmydoJZXlipqhZsjojoPA3uoI8BhaU4sPzs9OZ1+3w=";

            nativeBuildInputs = with pkgs; [ pkg-config ];
            buildInputs = with pkgs; pkgs.lib.optionals pkgs.stdenv.isLinux [ alsa-lib ];

            subPackages = [ "cmd/tj" ];

            ldflags = [
              "-s"
              "-w"
            ];

            meta = with pkgs.lib; {
              description = "Teejay - TUI app for monitoring tmux panes";
              homepage = "https://github.com/mipmip/teejay";
              license = licenses.mit;
              mainProgram = "tj";
            };
          };
        }
      );

      apps = forAllSystems (system: {
        default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/tj";
        };
      });

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              gopls
              gotools
              go-tools
              goreleaser
              # Audio dependencies for native sound support (Linux only)
              pkg-config
            ] ++ pkgs.lib.optionals pkgs.stdenv.isLinux [ alsa-lib ];

            shellHook = ''
              echo "Teejay development environment"
              echo "Go version: $(go version)"
            '';
          };
        }
      );
    };
}
