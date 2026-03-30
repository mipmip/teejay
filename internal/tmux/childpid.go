package tmux

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// GetChildPID finds a child process of the given parent PID with the specified command name.
// Returns the child's PID, or an error if not found.
func GetChildPID(parentPID int, command string) (int, error) {
	cmd := exec.Command("pgrep", "-P", strconv.Itoa(parentPID))
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("no child processes for PID %d", parentPID)
	}

	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		childPID, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			continue
		}
		// Check the command name of this child
		commCmd := exec.Command("ps", "-p", strconv.Itoa(childPID), "-o", "comm=")
		commOutput, err := commCmd.Output()
		if err != nil {
			continue
		}
		comm := strings.TrimSpace(string(commOutput))
		if comm == command {
			return childPID, nil
		}
	}
	return 0, fmt.Errorf("no child process named %q for PID %d", command, parentPID)
}

// GetPanePID returns the shell PID for a tmux pane.
func GetPanePID(paneID string) (int, error) {
	cmd := exec.Command("tmux", "display-message", "-t", paneID, "-p", "#{pane_pid}")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to get pane PID: %w", err)
	}
	pid, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return 0, fmt.Errorf("invalid PID output: %s", string(output))
	}
	return pid, nil
}
