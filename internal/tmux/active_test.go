package tmux

import (
	"os"
	"testing"
)

func TestGetActivePaneID_NotInTmux(t *testing.T) {
	original := os.Getenv("TMUX")
	defer os.Setenv("TMUX", original)

	os.Unsetenv("TMUX")

	result := GetActivePaneID()
	if result != "" {
		t.Errorf("GetActivePaneID() should return empty string when not in tmux, got %q", result)
	}
}

func TestGetActivePaneID_InTmux(t *testing.T) {
	if os.Getenv("TMUX") == "" {
		t.Skip("Skipping: not running inside tmux")
	}

	result := GetActivePaneID()
	if result == "" {
		t.Error("GetActivePaneID() should return a pane ID when running inside tmux")
	}
	if result[0] != '%' {
		t.Errorf("GetActivePaneID() should return a pane ID starting with '%%', got %q", result)
	}
}
