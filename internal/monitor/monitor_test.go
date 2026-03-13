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

	// First time seeing a pane should return Busy (no prompt)
	status := m.Update("%0", "some output")
	if status != Busy {
		t.Errorf("First update without prompt: got %v, want Busy", status)
	}
}

func TestUpdateFirstSeenWithPrompt(t *testing.T) {
	m := New()

	// First time with prompt should return Waiting
	status := m.Update("%0", "Do you want to proceed?")
	if status != Waiting {
		t.Errorf("First update with prompt: got %v, want Waiting", status)
	}
}

func TestUpdateContentChanged(t *testing.T) {
	m := New()

	m.Update("%0", "output 1")
	status := m.Update("%0", "output 2")
	if status != Busy {
		t.Errorf("Content changed: got %v, want Busy", status)
	}
}

func TestUpdatePromptDetected(t *testing.T) {
	m := New()

	m.Update("%0", "working...")
	status := m.Update("%0", "No, and tell Claude what to do differently")
	if status != Waiting {
		t.Errorf("Claude prompt: got %v, want Waiting", status)
	}
}

func TestUpdateAiderPrompt(t *testing.T) {
	m := New()

	status := m.Update("%0", "Apply changes? (Y)es/(N)o")
	if status != Waiting {
		t.Errorf("Aider prompt: got %v, want Waiting", status)
	}
}

func TestUpdateShellPrompt(t *testing.T) {
	m := New()

	status := m.Update("%0", "user@host:~$")
	if status != Waiting {
		t.Errorf("Shell prompt $: got %v, want Waiting", status)
	}

	m2 := New()
	status = m2.Update("%1", ">>> ")
	if status != Waiting {
		t.Errorf("Shell prompt >: got %v, want Waiting", status)
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
		{Busy, "⠋"},    // First spinner frame for Busy
		{Waiting, "●"}, // Green circle for Waiting
	}

	for _, tt := range tests {
		got := tt.status.Indicator()
		if got != tt.want {
			t.Errorf("%v.Indicator() = %q, want %q", tt.status, got, tt.want)
		}
	}
}

func TestStatusIndicatorAnimated(t *testing.T) {
	// Test Busy cycles through spinner frames
	expectedFrames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	for i, expected := range expectedFrames {
		got := Busy.IndicatorAnimated(i)
		if got != expected {
			t.Errorf("Busy.IndicatorAnimated(%d) = %q, want %q", i, got, expected)
		}
	}

	// Test frame wraps around
	if got := Busy.IndicatorAnimated(10); got != "⠋" {
		t.Errorf("Busy.IndicatorAnimated(10) = %q, want %q (should wrap)", got, "⠋")
	}
	if got := Busy.IndicatorAnimated(15); got != "⠴" {
		t.Errorf("Busy.IndicatorAnimated(15) = %q, want %q (should wrap)", got, "⠴")
	}

	// Test Waiting returns filled circle (not animated)
	if got := Waiting.IndicatorAnimated(0); got != "●" {
		t.Errorf("Waiting.IndicatorAnimated(0) = %q, want %q", got, "●")
	}
	if got := Waiting.IndicatorAnimated(5); got != "●" {
		t.Errorf("Waiting.IndicatorAnimated(5) = %q, want %q", got, "●")
	}
}

func TestStatusString(t *testing.T) {
	if Busy.String() != "Busy" {
		t.Error("Busy.String() != Busy")
	}
	if Waiting.String() != "Waiting" {
		t.Error("Waiting.String() != Waiting")
	}
}

func TestMultiplePanes(t *testing.T) {
	m := New()

	// Track multiple panes independently
	m.Update("%0", "pane 0 content")
	m.Update("%1", "Do you want to proceed?")

	status0 := m.Update("%0", "pane 0 content")
	status1 := m.Update("%1", "Do you want to proceed?")

	// Pane 0 should be Busy (no prompt detected)
	if status0 != Busy {
		t.Errorf("Pane 0 without prompt: got %v, want Busy", status0)
	}

	// Pane 1 should be Waiting (has prompt)
	if status1 != Waiting {
		t.Errorf("Pane 1 with prompt: got %v, want Waiting", status1)
	}
}

func TestLongContent(t *testing.T) {
	m := New()

	// Test with large content (no prompt)
	longContent := strings.Repeat("line of output\n", 1000)
	status := m.Update("%0", longContent)
	if status != Busy {
		t.Errorf("Long content: got %v, want Busy", status)
	}

	// Same content still returns Busy (no idle state anymore)
	status = m.Update("%0", longContent)
	if status != Busy {
		t.Errorf("Long stable content: got %v, want Busy", status)
	}
}
