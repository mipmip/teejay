package monitor

import (
	"crypto/sha256"
	"regexp"
	"strings"
	"time"

	"tj/internal/config"
)

// paneState holds the tracking state for a single pane.
type paneState struct {
	hash           [32]byte
	lastChangeTime time.Time
}

// ansiRegex matches ANSI escape sequences (CSI, OSC, etc.)
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]|\x1b\][^\x1b]*\x1b\\|\x1b\[[\?]?[0-9;]*[a-zA-Z]`)

// Monitor tracks activity state for multiple panes.
type Monitor struct {
	panes       map[string]*paneState
	config      *config.Config
	idleTimeout time.Duration
}

// New creates a new Monitor instance with the given configuration.
func New(cfg *config.Config) *Monitor {
	if cfg == nil {
		cfg = config.Default()
	}
	return &Monitor{
		panes:       make(map[string]*paneState),
		config:      cfg,
		idleTimeout: cfg.Detection.IdleTimeout,
	}
}

// Update checks the pane content and returns the current status.
// It uses configurable patterns and idle timeout detection.
// appName is the current foreground process (e.g., "claude", "fish").
func (m *Monitor) Update(paneID, content, appName string) PaneStatus {
	// Strip ANSI escape sequences before hashing and pattern matching
	// so that cursor/focus changes don't trigger false activity detection
	content = ansiRegex.ReplaceAllString(content, "")
	hash := sha256.Sum256([]byte(content))
	now := time.Now()

	state, exists := m.panes[paneID]
	if !exists {
		// First time seeing this pane
		m.panes[paneID] = &paneState{
			hash:           hash,
			lastChangeTime: now,
		}
		state = m.panes[paneID]
	} else {
		// Update hash and timestamp if content changed
		if hash != state.hash {
			state.hash = hash
			state.lastChangeTime = now
		}
	}

	// Get patterns for this app (app-specific replace globals)
	promptEndings, waitingStrings, busyStrings := m.config.GetPatternsForApp(appName)

	// Check for busy strings first - they take precedence over waiting detection
	if m.hasBusyString(content, busyStrings) {
		return Busy
	}

	// Check for waiting pattern matches
	if m.hasWaitingString(content, waitingStrings) {
		return Waiting
	}
	if m.hasPromptEnding(content, promptEndings) {
		return Waiting
	}

	// Check idle timeout (if enabled)
	if m.idleTimeout > 0 {
		idleDuration := now.Sub(state.lastChangeTime)
		if idleDuration >= m.idleTimeout {
			return Waiting
		}
	}

	return Busy
}

// hasBusyString checks if content contains any of the busy strings.
// Busy strings take precedence over waiting strings.
func (m *Monitor) hasBusyString(content string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(content, pattern) {
			return true
		}
	}
	return false
}

// hasWaitingString checks if content contains any of the waiting strings.
func (m *Monitor) hasWaitingString(content string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(content, pattern) {
			return true
		}
	}
	return false
}

// hasPromptEnding checks if the last non-empty line ends with any prompt ending.
func (m *Monitor) hasPromptEnding(content string, endings []string) bool {
	if len(endings) == 0 {
		return false
	}

	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	if len(lines) == 0 {
		return false
	}

	lastLine := strings.TrimSpace(lines[len(lines)-1])
	if len(lastLine) == 0 {
		return false
	}

	lastChar := string(lastLine[len(lastLine)-1])
	for _, ending := range endings {
		if lastChar == ending {
			return true
		}
	}
	return false
}
