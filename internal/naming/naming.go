package naming

import (
	"strings"

	"tj/internal/tmux"
)

// genericNames is a set of names that are considered too generic to be useful
// as pane identifiers. These are common shell names, tools, default tmux names, etc.
var genericNames = map[string]bool{
	// Shells
	"bash":  true,
	"zsh":   true,
	"fish":  true,
	"sh":    true,
	"dash":  true,
	"ksh":   true,
	"tcsh":  true,
	"csh":   true,
	"ash":   true,
	"shell": true,

	// Multiplexers and terminals
	"tmux":   true,
	"screen": true,

	// AI coding tools
	"claude":   true,
	"opencode": true,
	"aider":    true,
	"cursor":   true,

	// Default tmux names and numbers
	"0":       true,
	"1":       true,
	"2":       true,
	"3":       true,
	"4":       true,
	"5":       true,
	"6":       true,
	"7":       true,
	"8":       true,
	"9":       true,
	"main":    true,
	"default": true,
	"new":     true,
	"window":  true,
}

// IsGeneric returns true if the name is considered too generic to be useful.
func IsGeneric(name string) bool {
	if name == "" {
		return true
	}
	return genericNames[strings.ToLower(name)]
}

// GuessName attempts to determine a meaningful name for a pane based on tmux metadata.
// It tries sources in priority order: session > window name > process (command).
// Returns the guessed name and a boolean indicating whether the name is generic.
func GuessName(paneInfo tmux.PaneInfo) (string, bool) {
	// Try session name first (user's chosen context)
	if paneInfo.Session != "" && !IsGeneric(paneInfo.Session) {
		return paneInfo.Session, false
	}

	// Try window name next
	if paneInfo.WindowName != "" && !IsGeneric(paneInfo.WindowName) {
		return paneInfo.WindowName, false
	}

	// Try command last
	if paneInfo.Command != "" && !IsGeneric(paneInfo.Command) {
		return paneInfo.Command, false
	}

	// All sources are generic, return the best available (session if set, else window, else command)
	if paneInfo.Session != "" {
		return paneInfo.Session, true
	}
	if paneInfo.WindowName != "" {
		return paneInfo.WindowName, true
	}
	if paneInfo.Command != "" {
		return paneInfo.Command, true
	}

	return "", true
}
