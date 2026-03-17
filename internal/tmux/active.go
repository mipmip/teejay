package tmux

import (
	"os/exec"
	"strings"
)

// GetActivePaneID returns the pane ID of the currently focused tmux pane.
// Returns an empty string if not running inside tmux or if the query fails.
func GetActivePaneID() string {
	if !IsInsideTmux() {
		return ""
	}

	cmd := exec.Command("tmux", "display-message", "-p", "#{pane_id}")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}
