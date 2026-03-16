package monitor

import (
	"strings"
	"testing"
	"time"

	"tj/internal/config"
)

func TestNewMonitor(t *testing.T) {
	m := New(nil)
	if m == nil {
		t.Fatal("New() returned nil")
	}
	if m.panes == nil {
		t.Fatal("panes map is nil")
	}
}

func TestNewMonitorWithConfig(t *testing.T) {
	cfg := config.Default()
	cfg.Detection.IdleTimeout = 5 * time.Second
	m := New(cfg)

	if m.idleTimeout != 5*time.Second {
		t.Errorf("expected idle timeout 5s, got %v", m.idleTimeout)
	}
}

func TestUpdateFirstSeen(t *testing.T) {
	cfg := config.Default()
	cfg.Detection.IdleTimeout = 0 // Disable idle timeout for this test
	m := New(cfg)

	// First time seeing a pane should return Busy (no pattern match)
	status := m.Update("%0", "some output", "fish")
	if status != Busy {
		t.Errorf("First update without pattern: got %v, want Busy", status)
	}
}

func TestUpdateWithClaudePattern(t *testing.T) {
	m := New(nil) // Uses defaults which include claude patterns

	// Claude's "? for shortcuts" should trigger Waiting
	status := m.Update("%0", "some content\n? for shortcuts", "claude")
	if status != Waiting {
		t.Errorf("Claude pattern: got %v, want Waiting", status)
	}
}

func TestUpdateWithAiderPattern(t *testing.T) {
	m := New(nil)

	// Aider's "(Y)es/(N)o" should trigger Waiting
	status := m.Update("%0", "Apply changes? (Y)es/(N)o", "aider")
	if status != Waiting {
		t.Errorf("Aider pattern: got %v, want Waiting", status)
	}
}

func TestAppSpecificPatternsReplaceGlobals(t *testing.T) {
	cfg := config.Default()
	cfg.Detection.WaitingStrings = []string{"global pattern"}
	m := New(cfg)

	// Claude pane should NOT match global pattern (app patterns replace globals)
	status := m.Update("%0", "global pattern", "claude")
	// Since claude has its own patterns, it won't match "global pattern"
	// and idle timeout hasn't passed, so should be Busy
	if status != Busy {
		t.Errorf("Claude should ignore global patterns: got %v, want Busy", status)
	}

	// Unknown app should match global pattern
	status = m.Update("%1", "global pattern", "unknown-app")
	if status != Waiting {
		t.Errorf("Unknown app should use global patterns: got %v, want Waiting", status)
	}
}

func TestPromptEndingDetection(t *testing.T) {
	cfg := config.Default()
	cfg.Detection.IdleTimeout = 0
	cfg.Detection.PromptEndings = []string{"$", ">", "#"}
	m := New(cfg)

	tests := []struct {
		content string
		appName string
		want    PaneStatus
	}{
		{"user@host:~$", "bash", Waiting},
		{">>> ", "python", Waiting},
		{"root#", "bash", Waiting},
		{"just some output", "bash", Busy},
	}

	for i, tt := range tests {
		// Use different pane IDs to avoid state interference
		paneID := string(rune('0' + i))
		got := m.Update(paneID, tt.content, tt.appName)
		if got != tt.want {
			t.Errorf("content=%q app=%q: got %v, want %v", tt.content, tt.appName, got, tt.want)
		}
	}
}

func TestIdleTimeoutDetection(t *testing.T) {
	cfg := config.Default()
	cfg.Detection.IdleTimeout = 100 * time.Millisecond
	m := New(cfg)

	// First update - should be Busy
	status := m.Update("%0", "some output", "fish")
	if status != Busy {
		t.Errorf("First update: got %v, want Busy", status)
	}

	// Same content immediately - should still be Busy (timeout not reached)
	status = m.Update("%0", "some output", "fish")
	if status != Busy {
		t.Errorf("Immediate update: got %v, want Busy", status)
	}

	// Wait for idle timeout
	time.Sleep(150 * time.Millisecond)

	// Same content after timeout - should be Waiting
	status = m.Update("%0", "some output", "fish")
	if status != Waiting {
		t.Errorf("After idle timeout: got %v, want Waiting", status)
	}

	// Content changes - should reset to Busy
	status = m.Update("%0", "new output", "fish")
	if status != Busy {
		t.Errorf("After content change: got %v, want Busy", status)
	}
}

func TestIdleTimeoutDisabled(t *testing.T) {
	cfg := config.Default()
	cfg.Detection.IdleTimeout = 0 // Disabled
	m := New(cfg)

	// First update
	m.Update("%0", "some output", "fish")

	// Even after waiting, should remain Busy because timeout is disabled
	time.Sleep(50 * time.Millisecond)
	status := m.Update("%0", "some output", "fish")
	if status != Busy {
		t.Errorf("With disabled timeout: got %v, want Busy", status)
	}
}

func TestHasWaitingString(t *testing.T) {
	m := New(nil)

	tests := []struct {
		content  string
		patterns []string
		want     bool
	}{
		{"contains ? for shortcuts here", []string{"? for shortcuts"}, true},
		{"no match", []string{"? for shortcuts"}, false},
		{"", []string{"? for shortcuts"}, false},
		{"anything", []string{}, false},
	}

	for _, tt := range tests {
		got := m.hasWaitingString(tt.content, tt.patterns)
		if got != tt.want {
			t.Errorf("hasWaitingString(%q, %v) = %v, want %v", tt.content, tt.patterns, got, tt.want)
		}
	}
}

func TestHasPromptEnding(t *testing.T) {
	m := New(nil)

	tests := []struct {
		content string
		endings []string
		want    bool
	}{
		{"user@host:~$", []string{"$"}, true},
		{">>> ", []string{">"}, true},
		{"no prompt here", []string{"$", ">"}, false},
		{"", []string{"$"}, false},
		{"anything", []string{}, false},
		{"ends with space $", []string{"$"}, true}, // TrimSpace should handle this
	}

	for _, tt := range tests {
		got := m.hasPromptEnding(tt.content, tt.endings)
		if got != tt.want {
			t.Errorf("hasPromptEnding(%q, %v) = %v, want %v", tt.content, tt.endings, got, tt.want)
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

	// Test Waiting returns filled circle (not animated)
	if got := Waiting.IndicatorAnimated(0); got != "●" {
		t.Errorf("Waiting.IndicatorAnimated(0) = %q, want %q", got, "●")
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
	m := New(nil)

	// Track multiple panes independently
	m.Update("%0", "pane 0 content", "fish")
	m.Update("%1", "? for shortcuts", "claude")

	status0 := m.Update("%0", "pane 0 content", "fish")
	status1 := m.Update("%1", "? for shortcuts", "claude")

	// Pane 0 should be Busy (no pattern, idle timeout not reached)
	if status0 != Busy {
		t.Errorf("Pane 0 without pattern: got %v, want Busy", status0)
	}

	// Pane 1 should be Waiting (has claude pattern)
	if status1 != Waiting {
		t.Errorf("Pane 1 with claude pattern: got %v, want Waiting", status1)
	}
}

func TestLongContent(t *testing.T) {
	cfg := config.Default()
	cfg.Detection.IdleTimeout = 0 // Disable idle timeout
	m := New(cfg)

	// Test with large content (no pattern match)
	longContent := strings.Repeat("line of output\n", 1000)
	status := m.Update("%0", longContent, "fish")
	if status != Busy {
		t.Errorf("Long content: got %v, want Busy", status)
	}
}
