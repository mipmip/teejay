package scan

import (
	"testing"

	"tj/internal/config"
	"tj/internal/tmux"
	"tj/internal/watchlist"
)

func newTestWatchlist(t *testing.T) *watchlist.Watchlist {
	t.Helper()
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)
	wl, err := watchlist.Load()
	if err != nil {
		t.Fatalf("failed to load watchlist: %v", err)
	}
	return wl
}

func TestScanAndAddFindsAgents(t *testing.T) {
	cfg := config.Default()
	wl := newTestWatchlist(t)

	panes := []tmux.PaneInfo{
		{ID: "%0", Session: "work", WindowName: "code", Command: "claude"},
		{ID: "%1", Session: "work", WindowName: "edit", Command: "vim"},
		{ID: "%2", Session: "dev", WindowName: "ai", Command: "aider"},
	}

	result := ScanAndAdd(wl, cfg, panes)

	if result.Found != 2 {
		t.Errorf("Found = %d, want 2", result.Found)
	}
	if result.Added != 2 {
		t.Errorf("Added = %d, want 2", result.Added)
	}
	if result.Skipped != 0 {
		t.Errorf("Skipped = %d, want 0", result.Skipped)
	}
	if !wl.Contains("%0") || !wl.Contains("%2") {
		t.Error("agent panes should be in watchlist")
	}
	if wl.Contains("%1") {
		t.Error("non-agent pane should not be in watchlist")
	}
}

func TestScanAndAddSkipsAlreadyWatched(t *testing.T) {
	cfg := config.Default()
	wl := newTestWatchlist(t)
	wl.AddWithName("%0", "existing")

	panes := []tmux.PaneInfo{
		{ID: "%0", Session: "work", WindowName: "code", Command: "claude"},
		{ID: "%1", Session: "dev", WindowName: "ai", Command: "aider"},
	}

	result := ScanAndAdd(wl, cfg, panes)

	if result.Found != 2 {
		t.Errorf("Found = %d, want 2", result.Found)
	}
	if result.Added != 1 {
		t.Errorf("Added = %d, want 1", result.Added)
	}
	if result.Skipped != 1 {
		t.Errorf("Skipped = %d, want 1", result.Skipped)
	}
}

func TestScanAndAddNoAgents(t *testing.T) {
	cfg := config.Default()
	wl := newTestWatchlist(t)

	panes := []tmux.PaneInfo{
		{ID: "%0", Session: "work", WindowName: "code", Command: "vim"},
		{ID: "%1", Session: "work", WindowName: "shell", Command: "fish"},
	}

	result := ScanAndAdd(wl, cfg, panes)

	if result.Found != 0 {
		t.Errorf("Found = %d, want 0", result.Found)
	}
	if result.Added != 0 {
		t.Errorf("Added = %d, want 0", result.Added)
	}
}
