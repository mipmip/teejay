package monitor

import (
	"crypto/sha256"
	"strings"
)

// paneState holds the tracking state for a single pane.
type paneState struct {
	hash [32]byte
}

// Monitor tracks activity state for multiple panes.
type Monitor struct {
	panes map[string]*paneState
}

// New creates a new Monitor instance.
func New() *Monitor {
	return &Monitor{
		panes: make(map[string]*paneState),
	}
}

// Update checks the pane content and returns the current status.
// It computes a hash of the content and compares with the previous hash.
// Returns Waiting if a prompt is detected, Busy otherwise.
func (m *Monitor) Update(paneID, content string) PaneStatus {
	hash := sha256.Sum256([]byte(content))

	state, exists := m.panes[paneID]
	if !exists {
		// First time seeing this pane
		m.panes[paneID] = &paneState{
			hash: hash,
		}
	} else {
		// Update hash if content changed
		if hash != state.hash {
			state.hash = hash
		}
	}

	// Simple two-state logic: prompt detected → Waiting, else → Busy
	if hasPrompt(content) {
		return Waiting
	}
	return Busy
}

// promptPatterns contains strings that indicate a pane is waiting for input.
var promptPatterns = []string{
	// Claude Code prompts
	"No, and tell Claude what to do differently",
	"Do you want to proceed?",
	"Would you like me to",
	// Aider prompts
	"(Y)es/(N)o",
	"(y)es/(n)o",
	// Generic shell prompts (at end of content)
	// These are checked separately in hasPrompt
}

// hasPrompt checks if the content contains any known prompt patterns.
func hasPrompt(content string) bool {
	// Check for known tool prompts
	for _, pattern := range promptPatterns {
		if strings.Contains(content, pattern) {
			return true
		}
	}

	// Check for shell prompt at end of last non-empty line
	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	if len(lines) > 0 {
		lastLine := strings.TrimSpace(lines[len(lines)-1])
		if len(lastLine) > 0 {
			lastChar := lastLine[len(lastLine)-1]
			// Common shell prompts end with $, >, #, or %
			if lastChar == '$' || lastChar == '>' || lastChar == '#' || lastChar == '%' {
				return true
			}
		}
	}

	return false
}
