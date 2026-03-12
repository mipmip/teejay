package cmd

import (
	"errors"
	"testing"

	"tmon/internal/watchlist"
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

func TestAddPane_Duplicate(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)
	t.Setenv("TMUX_PANE", "%99")

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
