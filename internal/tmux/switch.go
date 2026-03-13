package tmux

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// IsInsideTmux returns true if the current process is running inside a tmux session.
func IsInsideTmux() bool {
	return os.Getenv("TMUX") != ""
}

// SwitchToPane switches the current tmux client to the specified pane.
// Returns an error if not running inside tmux or if the switch fails.
func SwitchToPane(paneID string) error {
	if !IsInsideTmux() {
		return fmt.Errorf("not running inside tmux")
	}

	cmd := exec.Command("tmux", "switch-client", "-t", paneID)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("tmux switch-client failed: %s", strings.TrimSpace(string(output)))
	}
	return nil
}
