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
  scan        Scan all panes and add those running known agents

General:
  -h, --help              Show this help message
  -v, --version           Show version
  -c, --config <path>     Path to config file
  -w, --watchlist <path>  Path to watchlist file

Alerts:
  --sound                 Enable sound alerts
  --no-sound              Disable sound alerts
  --notify                Enable desktop notifications
  --no-notify             Disable desktop notifications

Display:
  --columns               Start in multi-column layout
  --sort-activity         Start with activity sort (busy first, then recently finished)
  --sort-watchlist        Start with watchlist order (default)
  --recency-color         Enable recency color gradient on indicators
  --no-recency-color      Disable recency color gradient
  --preview               Show pane preview panel (default)
  --no-preview            Hide pane preview panel

Mode:
  --picker                Picker mode: Enter switches to pane and quits

Run 'tj' without arguments to launch the TUI.
`)
}

// CLIOverrides holds flag values that override config settings.
// Nil pointers mean "not specified" (preserve config value).
type CLIOverrides struct {
	Sound        *bool
	Notify       *bool
	SortActivity *bool
	Columns      *bool
	RecencyColor *bool
	PickerMode   *bool
	Preview      *bool
}

func boolPtr(v bool) *bool { return &v }

// parseFlags extracts all flags from args.
// Returns remaining args, configPath, watchlistPath, helpRequested, and CLI overrides.
func parseFlags(args []string) (remaining []string, configPath, watchlistPath string, helpRequested bool, overrides CLIOverrides) {
	i := 0
	for i < len(args) {
		arg := args[i]

		switch {
		// Help
		case arg == "--help" || arg == "-h":
			helpRequested = true
			i++
			continue

		// Config path
		case arg == "--config" || arg == "-c":
			if i+1 < len(args) {
				configPath = args[i+1]
				i += 2
				continue
			}
		case strings.HasPrefix(arg, "--config="):
			configPath = strings.TrimPrefix(arg, "--config=")
			i++
			continue
		case strings.HasPrefix(arg, "-c="):
			configPath = strings.TrimPrefix(arg, "-c=")
			i++
			continue

		// Watchlist path
		case arg == "--watchlist" || arg == "-w":
			if i+1 < len(args) {
				watchlistPath = args[i+1]
				i += 2
				continue
			}
		case strings.HasPrefix(arg, "--watchlist="):
			watchlistPath = strings.TrimPrefix(arg, "--watchlist=")
			i++
			continue
		case strings.HasPrefix(arg, "-w="):
			watchlistPath = strings.TrimPrefix(arg, "-w=")
			i++
			continue

		// Alert flags
		case arg == "--sound":
			overrides.Sound = boolPtr(true)
			i++
			continue
		case arg == "--no-sound":
			overrides.Sound = boolPtr(false)
			i++
			continue
		case arg == "--notify":
			overrides.Notify = boolPtr(true)
			i++
			continue
		case arg == "--no-notify":
			overrides.Notify = boolPtr(false)
			i++
			continue

		// Display flags
		case arg == "--columns":
			overrides.Columns = boolPtr(true)
			i++
			continue
		case arg == "--sort-activity":
			overrides.SortActivity = boolPtr(true)
			i++
			continue
		case arg == "--sort-watchlist":
			overrides.SortActivity = boolPtr(false)
			i++
			continue
		case arg == "--recency-color":
			overrides.RecencyColor = boolPtr(true)
			i++
			continue
		case arg == "--no-recency-color":
			overrides.RecencyColor = boolPtr(false)
			i++
			continue

		// Mode flags
		case arg == "--picker":
			overrides.PickerMode = boolPtr(true)
			i++
			continue
		case arg == "--preview":
			overrides.Preview = boolPtr(true)
			i++
			continue
		case arg == "--no-preview":
			overrides.Preview = boolPtr(false)
			i++
			continue
		}

		remaining = append(remaining, arg)
		i++
	}
	return
}

// applyOverrides applies CLI flag overrides to the config.
// Only non-nil overrides modify the config.
func applyOverrides(cfg *config.Config, overrides CLIOverrides) {
	if overrides.Sound != nil {
		cfg.Alerts.SoundOnReady = *overrides.Sound
		if !*overrides.Sound {
			cfg.Alerts.MuteSound = true // --no-sound overrules per-pane settings
		}
	}
	if overrides.Notify != nil {
		cfg.Alerts.NotifyOnReady = *overrides.Notify
		if !*overrides.Notify {
			cfg.Alerts.MuteNotify = true // --no-notify overrules per-pane settings
		}
	}
	if overrides.SortActivity != nil {
		cfg.Display.SortByActivity = *overrides.SortActivity
	}
	if overrides.Columns != nil && *overrides.Columns {
		cfg.Display.LayoutMode = "columns"
	}
	if overrides.RecencyColor != nil {
		cfg.Display.RecencyColor = *overrides.RecencyColor
	}
	if overrides.PickerMode != nil {
		cfg.Display.PickerMode = *overrides.PickerMode
	}
	if overrides.Preview != nil {
		cfg.Display.ShowPreview = *overrides.Preview
	}
}

func main() {
	// Parse global flags before subcommand dispatch
	args, configPath, watchlistPath, helpRequested, overrides := parseFlags(os.Args[1:])

	// Handle help flag before anything else
	if helpRequested {
		printHelp()
		return
	}

	// Load config with optional custom path
	cfg := config.Load(configPath)

	// Apply CLI overrides
	applyOverrides(cfg, overrides)

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
		case "scan":
			if err := cmd.ScanPanes(cfg, watchlistPath); err != nil {
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
	p := tea.NewProgram(ui.New(version, cfg, watchlistPath), tea.WithMouseCellMotion(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
