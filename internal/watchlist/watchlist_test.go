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
	path := filepath.Join(tmpDir, ".config", "teejay", "watchlist.json")
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

func TestSetSound(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0"},
			{ID: "%1"},
		},
	}

	wl.SetSound("%1", true)

	if !wl.Panes[1].SoundOnReady {
		t.Error("SetSound(%1, true) did not set SoundOnReady")
	}
	if wl.Panes[0].SoundOnReady {
		t.Error("SetSound() changed wrong pane")
	}

	wl.SetSound("%1", false)
	if wl.Panes[1].SoundOnReady {
		t.Error("SetSound(%1, false) did not unset SoundOnReady")
	}
}

func TestSetNotify(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0"},
			{ID: "%1"},
		},
	}

	wl.SetNotify("%0", true)

	if !wl.Panes[0].NotifyOnReady {
		t.Error("SetNotify(%0, true) did not set NotifyOnReady")
	}
	if wl.Panes[1].NotifyOnReady {
		t.Error("SetNotify() changed wrong pane")
	}

	wl.SetNotify("%0", false)
	if wl.Panes[0].NotifyOnReady {
		t.Error("SetNotify(%0, false) did not unset NotifyOnReady")
	}
}

func TestGetPane(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0", Name: "first"},
			{ID: "%1", Name: "second"},
		},
	}

	p := wl.GetPane("%1")
	if p == nil {
		t.Fatal("GetPane(%1) returned nil")
	}
	if p.Name != "second" {
		t.Errorf("GetPane(%%1).Name = %q, want 'second'", p.Name)
	}

	// Modify through pointer
	p.Name = "modified"
	if wl.Panes[1].Name != "modified" {
		t.Error("GetPane() did not return pointer to actual pane")
	}

	// Non-existent pane
	if wl.GetPane("%99") != nil {
		t.Error("GetPane(%99) should return nil")
	}
}

func TestAlertFieldsPersistence(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	wl := &Watchlist{}
	wl.Add("%0")
	wl.SetSound("%0", true)
	wl.SetNotify("%0", true)

	if err := wl.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if !loaded.Panes[0].SoundOnReady {
		t.Error("SoundOnReady not persisted")
	}
	if !loaded.Panes[0].NotifyOnReady {
		t.Error("NotifyOnReady not persisted")
	}
}
