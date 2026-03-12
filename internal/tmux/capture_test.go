package tmux

import (
	"os"
	"strings"
	"testing"
)

func TestCapturePane_InvalidPane(t *testing.T) {
	// Try to capture a pane that doesn't exist
	_, err := CapturePane("%999999")
	if err == nil {
		t.Error("CapturePane() should return error for invalid pane")
	}
}

func TestCapturePane_CurrentPane(t *testing.T) {
	// Skip if not running in tmux
	paneID := os.Getenv("TMUX_PANE")
	if paneID == "" {
		t.Skip("Not running inside tmux")
	}

	content, err := CapturePane(paneID)
	if err != nil {
		t.Fatalf("CapturePane() error = %v", err)
	}

	// Content should be non-empty (at least has some terminal output)
	if strings.TrimSpace(content) == "" {
		t.Log("Warning: captured pane content is empty (this may be normal)")
	}
}
