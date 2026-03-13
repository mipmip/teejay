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
            version = "0.1.0";

            src = ./.;
            vendorHash = "sha256-yAmydoJZXlipqhZsjojoPA3uoI8BhaU4sPzs9OZ1+3w=";

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
            ];

            shellHook = ''
              echo "Teejay development environment"
              echo "Go version: $(go version)"
            '';
          };
        }
      );
    };
}
