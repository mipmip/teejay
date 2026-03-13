package tmux

import (
	"os"
	"testing"
)

func TestIsInsideTmux(t *testing.T) {
	// Save original value
	original := os.Getenv("TMUX")
	defer os.Setenv("TMUX", original)

	// Test when TMUX is set
	os.Setenv("TMUX", "/tmp/tmux-1000/default,12345,0")
	if !IsInsideTmux() {
		t.Error("IsInsideTmux() should return true when TMUX is set")
	}

	// Test when TMUX is empty
	os.Setenv("TMUX", "")
	if IsInsideTmux() {
		t.Error("IsInsideTmux() should return false when TMUX is empty")
	}

	// Test when TMUX is unset
	os.Unsetenv("TMUX")
	if IsInsideTmux() {
		t.Error("IsInsideTmux() should return false when TMUX is unset")
	}
}

func TestSwitchToPane_NotInTmux(t *testing.T) {
	// Save original value
	original := os.Getenv("TMUX")
	defer os.Setenv("TMUX", original)

	// Ensure we're not in tmux
	os.Unsetenv("TMUX")

	err := SwitchToPane("%1")
	if err == nil {
		t.Error("SwitchToPane() should return error when not in tmux")
	}
	if err.Error() != "not running inside tmux" {
		t.Errorf("unexpected error message: %v", err)
	}
}
