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

func TestContains(t *testing.T) {
	wl := &Watchlist{}
	wl.Add("%0")
	wl.Add("%1")

	if !wl.Contains("%0") {
		t.Error("Contains(%0) = false, want true")
	}
	if !wl.Contains("%1") {
		t.Error("Contains(%1) = false, want true")
	}
	if wl.Contains("%2") {
		t.Error("Contains(%2) = true, want false")
	}
}

func TestContainsEmpty(t *testing.T) {
	wl := &Watchlist{}

	if wl.Contains("%0") {
		t.Error("Contains(%0) on empty watchlist = true, want false")
	}
}

func TestDeduplicate(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0"},
			{ID: "%1"},
			{ID: "%0"}, // duplicate
			{ID: "%2"},
			{ID: "%1"}, // duplicate
		},
	}

	wl.Deduplicate()

	if len(wl.Panes) != 3 {
		t.Fatalf("Deduplicate() resulted in %d panes, want 3", len(wl.Panes))
	}
	expected := []string{"%0", "%1", "%2"}
	for i, want := range expected {
		if wl.Panes[i].ID != want {
			t.Errorf("Panes[%d].ID = %q, want %q", i, wl.Panes[i].ID, want)
		}
	}
}

func TestDeduplicateNoDuplicates(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0"},
			{ID: "%1"},
			{ID: "%2"},
		},
	}

	wl.Deduplicate()

	if len(wl.Panes) != 3 {
		t.Fatalf("Deduplicate() resulted in %d panes, want 3", len(wl.Panes))
	}
}

func TestDeduplicateEmpty(t *testing.T) {
	wl := &Watchlist{}

	wl.Deduplicate()

	if len(wl.Panes) != 0 {
		t.Fatalf("Deduplicate() on empty resulted in %d panes, want 0", len(wl.Panes))
	}
}

func TestRemove(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0"},
			{ID: "%1"},
			{ID: "%2"},
		},
	}

	wl.Remove("%1")

	if len(wl.Panes) != 2 {
		t.Fatalf("Remove() resulted in %d panes, want 2", len(wl.Panes))
	}
	if wl.Panes[0].ID != "%0" || wl.Panes[1].ID != "%2" {
		t.Errorf("Remove() left wrong panes: %v", wl.Panes)
	}
}

func TestRemoveNonExistent(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0"},
			{ID: "%1"},
		},
	}

	wl.Remove("%99")

	if len(wl.Panes) != 2 {
		t.Fatalf("Remove() non-existent resulted in %d panes, want 2", len(wl.Panes))
	}
}

func TestRename(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0"},
			{ID: "%1"},
		},
	}

	wl.Rename("%1", "my-pane")

	if wl.Panes[1].Name != "my-pane" {
		t.Errorf("Rename() resulted in Name = %q, want 'my-pane'", wl.Panes[1].Name)
	}
	if wl.Panes[0].Name != "" {
		t.Errorf("Rename() changed wrong pane, Panes[0].Name = %q", wl.Panes[0].Name)
	}
}

func TestRenameNonExistent(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0"},
		},
	}

	wl.Rename("%99", "test")

	// Should not panic or modify existing panes
	if wl.Panes[0].Name != "" {
		t.Errorf("Rename() non-existent modified wrong pane")
	}
}

func TestDisplayName(t *testing.T) {
	paneWithName := Pane{ID: "%0", Name: "my-pane"}
	paneWithoutName := Pane{ID: "%1"}

	if paneWithName.DisplayName() != "my-pane" {
		t.Errorf("DisplayName() = %q, want 'my-pane'", paneWithName.DisplayName())
	}
	if paneWithoutName.DisplayName() != "%1" {
		t.Errorf("DisplayName() = %q, want '%%1'", paneWithoutName.DisplayName())
	}
}
