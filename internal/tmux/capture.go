package tmux

import (
	"fmt"
	"os/exec"
	"strings"
)

// CapturePane captures the content of a tmux pane by its ID.
// Returns the pane content as a string, or an error if capture fails.
func CapturePane(paneID string) (string, error) {
	cmd := exec.Command("tmux", "capture-pane", "-p", "-t", paneID)
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("tmux capture-pane failed: %s", strings.TrimSpace(string(exitErr.Stderr)))
		}
		return "", fmt.Errorf("failed to run tmux: %w", err)
	}
	return string(output), nil
}
