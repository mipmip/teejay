package tmux

import (
	"fmt"
	"os/exec"
	"strings"
)

// PaneInfo holds metadata about a tmux pane.
type PaneInfo struct {
	ID        string // e.g., "%0"
	Session   string // session name
	Window    int    // window index
	Pane      int    // pane index within window
	Command   string // current command running in pane
	SessionID string // full session:window.pane identifier
}

// ListAllPanes returns information about all tmux panes across all sessions.
func ListAllPanes() ([]PaneInfo, error) {
	// Format: pane_id, session_name, window_index, pane_index, pane_current_command
	format := "#{pane_id}\t#{session_name}\t#{window_index}\t#{pane_index}\t#{pane_current_command}"
	cmd := exec.Command("tmux", "list-panes", "-a", "-F", format)
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("tmux list-panes failed: %s", strings.TrimSpace(string(exitErr.Stderr)))
		}
		return nil, fmt.Errorf("failed to run tmux: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return []PaneInfo{}, nil
	}

	panes := make([]PaneInfo, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if len(parts) < 5 {
			continue
		}

		var windowIdx, paneIdx int
		fmt.Sscanf(parts[2], "%d", &windowIdx)
		fmt.Sscanf(parts[3], "%d", &paneIdx)

		panes = append(panes, PaneInfo{
			ID:        parts[0],
			Session:   parts[1],
			Window:    windowIdx,
			Pane:      paneIdx,
			Command:   parts[4],
			SessionID: fmt.Sprintf("%s:%d.%d", parts[1], windowIdx, paneIdx),
		})
	}

	return panes, nil
}
