package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"tj/internal/cmd"
	"tj/internal/ui"
)

// version is set at build time via ldflags
var version = "dev"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v":
			fmt.Println(version)
			return
		case "add":
			if err := cmd.AddPane(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			return
		default:
			fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
			os.Exit(1)
		}
	}

	// No args: launch TUI
	p := tea.NewProgram(ui.New(version), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
