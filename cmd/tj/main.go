package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"tj/internal/cmd"
	"tj/internal/config"
	"tj/internal/ui"
)

// Version is embedded from the VERSION file (single source of truth).
//
//go:embed VERSION
var version string

func init() {
	version = strings.TrimSpace(version)
}

func printHelp() {
	fmt.Print(`tj - tmux pane watchlist manager

Usage:
  tj [flags]
  tj <command> [flags]

Commands:
  add         Add the current tmux pane to the watchlist
  del         Remove the current tmux pane from the watchlist

Flags:
  -h, --help              Show this help message
  -v, --version           Show version
  -c, --config <path>     Path to config file
  -w, --watchlist <path>  Path to watchlist file

Run 'tj' without arguments to launch the TUI.
`)
}

// parseFlags extracts --config/-c and --watchlist/-w flags from args.
// Returns the remaining args (with flags removed), configPath, watchlistPath, and whether help was requested.
func parseFlags(args []string) (remaining []string, configPath, watchlistPath string, helpRequested bool) {
	i := 0
	for i < len(args) {
		arg := args[i]

		// Check for --help or -h anywhere in args
		if arg == "--help" || arg == "-h" {
			helpRequested = true
			i++
			continue
		}

		// Check for --config or -c
		if arg == "--config" || arg == "-c" {
			if i+1 < len(args) {
				configPath = args[i+1]
				i += 2
				continue
			}
		}
		// Check for --config=value
		if strings.HasPrefix(arg, "--config=") {
			configPath = strings.TrimPrefix(arg, "--config=")
			i++
			continue
		}
		if strings.HasPrefix(arg, "-c=") {
			configPath = strings.TrimPrefix(arg, "-c=")
			i++
			continue
		}

		// Check for --watchlist or -w
		if arg == "--watchlist" || arg == "-w" {
			if i+1 < len(args) {
				watchlistPath = args[i+1]
				i += 2
				continue
			}
		}
		// Check for --watchlist=value
		if strings.HasPrefix(arg, "--watchlist=") {
			watchlistPath = strings.TrimPrefix(arg, "--watchlist=")
			i++
			continue
		}
		if strings.HasPrefix(arg, "-w=") {
			watchlistPath = strings.TrimPrefix(arg, "-w=")
			i++
			continue
		}

		remaining = append(remaining, arg)
		i++
	}
	return
}

func main() {
	// Parse global flags before subcommand dispatch
	args, configPath, watchlistPath, helpRequested := parseFlags(os.Args[1:])

	// Handle help flag before anything else
	if helpRequested {
		printHelp()
		return
	}

	// Load config with optional custom path
	cfg := config.Load(configPath)

	if len(args) > 0 {
		switch args[0] {
		case "--version", "-v":
			fmt.Println(version)
			return
		case "add":
			if err := cmd.AddPane(watchlistPath); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			return
		case "del":
			if err := cmd.DelPane(watchlistPath); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			return
		default:
			fmt.Fprintf(os.Stderr, "Unknown command: %s\n", args[0])
			os.Exit(1)
		}
	}

	// No args: launch TUI
	p := tea.NewProgram(ui.New(version, cfg, watchlistPath), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
