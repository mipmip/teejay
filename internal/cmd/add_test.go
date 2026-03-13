package cmd

import (
	"errors"
	"os"
	"strings"
	"testing"

	"tj/internal/watchlist"
)

func TestGetTmuxPaneID_Set(t *testing.T) {
	t.Setenv("TMUX_PANE", "%42")

	paneID, err := GetTmuxPaneID()
	if err != nil {
		t.Fatalf("GetTmuxPaneID() error = %v, want nil", err)
	}
	if paneID != "%42" {
		t.Errorf("GetTmuxPaneID() = %q, want %%42", paneID)
	}
}

func TestGetTmuxPaneID_NotSet(t *testing.T) {
	t.Setenv("TMUX_PANE", "")

	_, err := GetTmuxPaneID()
	if err == nil {
		t.Fatal("GetTmuxPaneID() error = nil, want ErrNotInTmux")
	}
	if !errors.Is(err, ErrNotInTmux) {
		t.Errorf("GetTmuxPaneID() error = %v, want ErrNotInTmux", err)
	}
}

func TestAddPane_WithMockedStdin(t *testing.T) {
	// Skip if not in tmux - need real pane for name guessing
	if os.Getenv("TMUX") == "" {
		t.Skip("Not running inside tmux")
	}

	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)
	// Use the actual pane ID from the test environment
	paneID := os.Getenv("TMUX_PANE")
	t.Setenv("TMUX_PANE", paneID)

	// Mock stdin with custom name (in case the guessed name is generic)
	Stdin = strings.NewReader("my-custom-name\n")
	defer func() { Stdin = os.Stdin }()

	err := AddPane()
	if err != nil {
		t.Fatalf("AddPane() error = %v", err)
	}

	// Verify pane was added with a name
	wl, _ := watchlist.Load()
	if len(wl.Panes) != 1 {
		t.Fatalf("After add: %d panes, want 1", len(wl.Panes))
	}
	if wl.Panes[0].ID != paneID {
		t.Errorf("Pane ID = %q, want %q", wl.Panes[0].ID, paneID)
	}
	// Pane should have a name (either guessed or from stdin)
	if wl.Panes[0].Name == "" {
		t.Error("Pane should have a name set")
	}
}

func TestAddPane_Duplicate(t *testing.T) {
	// Skip if not in tmux - need real pane for name guessing
	if os.Getenv("TMUX") == "" {
		t.Skip("Not running inside tmux")
	}

	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)
	paneID := os.Getenv("TMUX_PANE")
	t.Setenv("TMUX_PANE", paneID)

	// Mock stdin
	Stdin = strings.NewReader("test-name\n")
	defer func() { Stdin = os.Stdin }()

	// First add should succeed
	err := AddPane()
	if err != nil {
		t.Fatalf("First AddPane() error = %v, want nil", err)
	}

	// Verify pane was added
	wl, _ := watchlist.Load()
	if len(wl.Panes) != 1 {
		t.Fatalf("After first add: %d panes, want 1", len(wl.Panes))
	}

	// Reset stdin for second call
	Stdin = strings.NewReader("another-name\n")

	// Second add of same pane should not add duplicate
	err = AddPane()
	if err != nil {
		t.Fatalf("Second AddPane() error = %v, want nil", err)
	}

	// Verify no duplicate was added
	wl, _ = watchlist.Load()
	if len(wl.Panes) != 1 {
		t.Errorf("After second add: %d panes, want 1 (no duplicate)", len(wl.Panes))
	}
}

func TestAddPane_NotInTmux(t *testing.T) {
	t.Setenv("TMUX_PANE", "")

	err := AddPane()
	if err == nil {
		t.Fatal("AddPane() should fail when not in tmux")
	}
	if !strings.Contains(err.Error(), "not running inside tmux") {
		t.Errorf("error should mention 'not running inside tmux', got: %v", err)
	}
}
