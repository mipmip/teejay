package watchlist

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadNonExistent(t *testing.T) {
	// Temporarily override HOME to a temp dir
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	wl, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v, want nil", err)
	}
	if len(wl.Panes) != 0 {
		t.Errorf("Load() returned %d panes, want 0", len(wl.Panes))
	}
}

func TestSaveAndLoad(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	wl := &Watchlist{}
	wl.Add("%0")
	wl.Add("%1")

	if err := wl.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	// Verify file exists
	path := filepath.Join(tmpDir, ".config", "tmon", "watchlist.json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("watchlist.json not created")
	}

	// Load and verify
	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if len(loaded.Panes) != 2 {
		t.Errorf("Load() returned %d panes, want 2", len(loaded.Panes))
	}
	if loaded.Panes[0].ID != "%0" {
		t.Errorf("Panes[0].ID = %q, want %%0", loaded.Panes[0].ID)
	}
	if loaded.Panes[1].ID != "%1" {
		t.Errorf("Panes[1].ID = %q, want %%1", loaded.Panes[1].ID)
	}
}

func TestAdd(t *testing.T) {
	wl := &Watchlist{}

	wl.Add("%5")

	if len(wl.Panes) != 1 {
		t.Fatalf("Add() resulted in %d panes, want 1", len(wl.Panes))
	}
	if wl.Panes[0].ID != "%5" {
		t.Errorf("Panes[0].ID = %q, want %%5", wl.Panes[0].ID)
	}
	if wl.Panes[0].AddedAt.IsZero() {
		t.Error("Panes[0].AddedAt is zero, want non-zero")
	}
}
