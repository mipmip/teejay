package monitor

import (
	"strings"
	"testing"
)

func TestNewMonitor(t *testing.T) {
	m := New()
	if m == nil {
		t.Fatal("New() returned nil")
	}
	if m.panes == nil {
		t.Fatal("panes map is nil")
	}
}

func TestUpdateFirstSeen(t *testing.T) {
	m := New()

	// First time seeing a pane should return Running (no prompt)
	status := m.Update("%0", "some output")
	if status != Running {
		t.Errorf("First update without prompt: got %v, want Running", status)
	}
}

func TestUpdateFirstSeenWithPrompt(t *testing.T) {
	m := New()

	// First time with prompt should return Ready
	status := m.Update("%0", "Do you want to proceed?")
	if status != Ready {
		t.Errorf("First update with prompt: got %v, want Ready", status)
	}
}

func TestUpdateContentChanged(t *testing.T) {
	m := New()

	m.Update("%0", "output 1")
	status := m.Update("%0", "output 2")
	if status != Running {
		t.Errorf("Content changed: got %v, want Running", status)
	}
}

func TestUpdatePromptDetected(t *testing.T) {
	m := New()

	m.Update("%0", "working...")
	status := m.Update("%0", "No, and tell Claude what to do differently")
	if status != Ready {
		t.Errorf("Claude prompt: got %v, want Ready", status)
	}
}

func TestUpdateAiderPrompt(t *testing.T) {
	m := New()

	status := m.Update("%0", "Apply changes? (Y)es/(N)o")
	if status != Ready {
		t.Errorf("Aider prompt: got %v, want Ready", status)
	}
}

func TestUpdateShellPrompt(t *testing.T) {
	m := New()

	status := m.Update("%0", "user@host:~$")
	if status != Ready {
		t.Errorf("Shell prompt $: got %v, want Ready", status)
	}

	m2 := New()
	status = m2.Update("%1", ">>> ")
	if status != Ready {
		t.Errorf("Shell prompt >: got %v, want Ready", status)
	}
}

func TestUpdateIdleTransition(t *testing.T) {
	m := New()

	// Initial state
	m.Update("%0", "stable content")

	// Content stays the same for idleThreshold ticks
	var status PaneStatus
	for i := 0; i < idleThreshold; i++ {
		status = m.Update("%0", "stable content")
	}

	if status != Idle {
		t.Errorf("After %d stable ticks: got %v, want Idle", idleThreshold, status)
	}
}

func TestUpdateIdleToRunning(t *testing.T) {
	m := New()

	// Get to idle state
	m.Update("%0", "stable content")
	for i := 0; i < idleThreshold; i++ {
		m.Update("%0", "stable content")
	}

	// Content changes - should go back to Running
	status := m.Update("%0", "new content")
	if status != Running {
		t.Errorf("Content changed from idle: got %v, want Running", status)
	}
}

func TestHasPromptPatterns(t *testing.T) {
	tests := []struct {
		content string
		want    bool
	}{
		{"No, and tell Claude what to do differently", true},
		{"Do you want to proceed?", true},
		{"(Y)es/(N)o", true},
		{"(y)es/(n)o", true},
		{"user@host:~$", true},
		{">>>", true},
		{"root#", true},
		{"just some output", false},
		{"building...", false},
	}

	for _, tt := range tests {
		got := hasPrompt(tt.content)
		if got != tt.want {
			t.Errorf("hasPrompt(%q) = %v, want %v", tt.content, got, tt.want)
		}
	}
}

func TestStatusIndicator(t *testing.T) {
	tests := []struct {
		status PaneStatus
		want   string
	}{
		{Running, "●"},
		{Ready, "?"},
		{Idle, "○"},
	}

	for _, tt := range tests {
		got := tt.status.Indicator()
		if got != tt.want {
			t.Errorf("%v.Indicator() = %q, want %q", tt.status, got, tt.want)
		}
	}
}

func TestStatusIndicatorAnimated(t *testing.T) {
	// Test Running cycles through spinner frames
	expectedFrames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	for i, expected := range expectedFrames {
		got := Running.IndicatorAnimated(i)
		if got != expected {
			t.Errorf("Running.IndicatorAnimated(%d) = %q, want %q", i, got, expected)
		}
	}

	// Test frame wraps around
	if got := Running.IndicatorAnimated(10); got != "⠋" {
		t.Errorf("Running.IndicatorAnimated(10) = %q, want %q (should wrap)", got, "⠋")
	}
	if got := Running.IndicatorAnimated(15); got != "⠴" {
		t.Errorf("Running.IndicatorAnimated(15) = %q, want %q (should wrap)", got, "⠴")
	}

	// Test Ready returns filled circle (not animated)
	if got := Ready.IndicatorAnimated(0); got != "●" {
		t.Errorf("Ready.IndicatorAnimated(0) = %q, want %q", got, "●")
	}
	if got := Ready.IndicatorAnimated(5); got != "●" {
		t.Errorf("Ready.IndicatorAnimated(5) = %q, want %q", got, "●")
	}

	// Test Idle returns empty circle (not animated)
	if got := Idle.IndicatorAnimated(0); got != "○" {
		t.Errorf("Idle.IndicatorAnimated(0) = %q, want %q", got, "○")
	}
}

func TestStatusString(t *testing.T) {
	if Running.String() != "Running" {
		t.Error("Running.String() != Running")
	}
	if Ready.String() != "Ready" {
		t.Error("Ready.String() != Ready")
	}
	if Idle.String() != "Idle" {
		t.Error("Idle.String() != Idle")
	}
}

func TestMultiplePanes(t *testing.T) {
	m := New()

	// Track multiple panes independently
	m.Update("%0", "pane 0 content")
	m.Update("%1", "Do you want to proceed?")

	status0 := m.Update("%0", "pane 0 content")
	status1 := m.Update("%1", "Do you want to proceed?")

	// Pane 0 should be Running (first update was different content)
	// Actually second update is same content, so it's still running but moving toward idle
	if status0 == Idle {
		t.Error("Pane 0 should not be Idle after just 2 updates")
	}

	// Pane 1 should be Ready (has prompt)
	if status1 != Ready {
		t.Errorf("Pane 1 with prompt: got %v, want Ready", status1)
	}
}

func TestLongContent(t *testing.T) {
	m := New()

	// Test with large content
	longContent := strings.Repeat("line of output\n", 1000)
	status := m.Update("%0", longContent)
	if status != Running {
		t.Errorf("Long content: got %v, want Running", status)
	}

	// Same content should move toward idle
	for i := 0; i < idleThreshold; i++ {
		status = m.Update("%0", longContent)
	}
	if status != Idle {
		t.Errorf("Long stable content: got %v, want Idle", status)
	}
}
