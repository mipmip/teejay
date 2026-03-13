package monitor

import (
	"crypto/sha256"
	"strings"
)

// idleThreshold is the number of ticks (at 100ms each) before a pane is considered idle.
// 20 ticks = 2 seconds.
const idleThreshold = 20

// paneState holds the tracking state for a single pane.
type paneState struct {
	hash        [32]byte
	idleCounter int
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
func (m *Monitor) Update(paneID, content string) PaneStatus {
	hash := sha256.Sum256([]byte(content))

	state, exists := m.panes[paneID]
	if !exists {
		// First time seeing this pane
		m.panes[paneID] = &paneState{
			hash:        hash,
			idleCounter: 0,
		}
		// Check for prompt on first view
		if hasPrompt(content) {
			return Ready
		}
		return Running
	}

	// Check if content changed
	contentChanged := hash != state.hash
	if contentChanged {
		state.hash = hash
		state.idleCounter = 0
	}

	// Check for prompt - takes precedence over running state
	if hasPrompt(content) {
		state.idleCounter = 0
		return Ready
	}

	// Content changed but no prompt - running
	if contentChanged {
		return Running
	}

	// Content stable, no prompt - increment idle counter
	state.idleCounter++
	if state.idleCounter >= idleThreshold {
		return Idle
	}

	// Still in running state but content hasn't changed recently
	return Running
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
