package tmux

import (
	"os/exec"
	"strings"
)

// GetActivePaneID returns the pane ID of a currently focused tmux pane.
// With multiple attached sessions, returns the first match.
// Returns an empty string if not running inside tmux or if the query fails.
func GetActivePaneID() string {
	ids := GetActivePaneIDs()
	if len(ids) > 0 {
		for id := range ids {
			return id
		}
	}
	return ""
}

// GetActivePaneIDs returns the set of pane IDs that are currently focused
// (active pane in active window of any attached session).
// With multiple attached clients/sessions, multiple panes may be active.
func GetActivePaneIDs() map[string]bool {
	result := make(map[string]bool)
	if !IsInsideTmux() {
		return result
	}

	cmd := exec.Command("tmux", "list-panes", "-a", "-F", "#{pane_id}\t#{pane_active}\t#{window_active}\t#{session_attached}")
	output, err := cmd.Output()
	if err != nil {
		return result
	}

	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		parts := strings.Split(line, "\t")
		if len(parts) == 4 && parts[1] == "1" && parts[2] == "1" && parts[3] == "1" {
			result[parts[0]] = true
		}
	}
	return result
}
