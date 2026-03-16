package watchlist

import (
	"os"
	"path/filepath"
	"testing"

	"tj/internal/config"
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

	tr := true
	wl.SetSound("%1", &tr)

	if wl.Panes[1].SoundOnReady == nil || !*wl.Panes[1].SoundOnReady {
		t.Error("SetSound(%1, true) did not set SoundOnReady")
	}
	if wl.Panes[0].SoundOnReady != nil {
		t.Error("SetSound() changed wrong pane")
	}

	fl := false
	wl.SetSound("%1", &fl)
	if wl.Panes[1].SoundOnReady == nil || *wl.Panes[1].SoundOnReady {
		t.Error("SetSound(%1, false) did not unset SoundOnReady")
	}

	// Test clearing to nil (use default)
	wl.SetSound("%1", nil)
	if wl.Panes[1].SoundOnReady != nil {
		t.Error("SetSound(%1, nil) did not clear SoundOnReady")
	}
}

func TestSetNotify(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0"},
			{ID: "%1"},
		},
	}

	tr := true
	wl.SetNotify("%0", &tr)

	if wl.Panes[0].NotifyOnReady == nil || !*wl.Panes[0].NotifyOnReady {
		t.Error("SetNotify(%0, true) did not set NotifyOnReady")
	}
	if wl.Panes[1].NotifyOnReady != nil {
		t.Error("SetNotify() changed wrong pane")
	}

	fl := false
	wl.SetNotify("%0", &fl)
	if wl.Panes[0].NotifyOnReady == nil || *wl.Panes[0].NotifyOnReady {
		t.Error("SetNotify(%0, false) did not unset NotifyOnReady")
	}

	// Test clearing to nil (use default)
	wl.SetNotify("%0", nil)
	if wl.Panes[0].NotifyOnReady != nil {
		t.Error("SetNotify(%0, nil) did not clear NotifyOnReady")
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
	tr := true
	wl.SetSound("%0", &tr)
	wl.SetNotify("%0", &tr)

	if err := wl.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if loaded.Panes[0].SoundOnReady == nil || !*loaded.Panes[0].SoundOnReady {
		t.Error("SoundOnReady not persisted")
	}
	if loaded.Panes[0].NotifyOnReady == nil || !*loaded.Panes[0].NotifyOnReady {
		t.Error("NotifyOnReady not persisted")
	}
}

func TestAddWithName(t *testing.T) {
	wl := &Watchlist{}

	wl.AddWithName("%5", "my-editor")

	if len(wl.Panes) != 1 {
		t.Fatalf("AddWithName() resulted in %d panes, want 1", len(wl.Panes))
	}
	if wl.Panes[0].ID != "%5" {
		t.Errorf("Panes[0].ID = %q, want %%5", wl.Panes[0].ID)
	}
	if wl.Panes[0].Name != "my-editor" {
		t.Errorf("Panes[0].Name = %q, want 'my-editor'", wl.Panes[0].Name)
	}
	if wl.Panes[0].AddedAt.IsZero() {
		t.Error("Panes[0].AddedAt is zero, want non-zero")
	}
}

func TestAddWithNamePersistence(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	wl := &Watchlist{}
	wl.AddWithName("%0", "nvim-editor")

	if err := wl.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if loaded.Panes[0].Name != "nvim-editor" {
		t.Errorf("Name not persisted: got %q, want 'nvim-editor'", loaded.Panes[0].Name)
	}
}

func TestGetEffectiveSound(t *testing.T) {
	cfg := config.Default()

	// Pane with no override - use global default (false)
	pane := &Pane{ID: "%0"}
	if pane.GetEffectiveSound(cfg) {
		t.Error("expected false when no override and default is false")
	}

	// Change global default to true
	cfg.Alerts.SoundOnReady = true
	if !pane.GetEffectiveSound(cfg) {
		t.Error("expected true when no override and default is true")
	}

	// Pane with explicit false override
	fl := false
	pane.SoundOnReady = &fl
	if pane.GetEffectiveSound(cfg) {
		t.Error("expected false when override is false (regardless of default)")
	}

	// Pane with explicit true override
	tr := true
	pane.SoundOnReady = &tr
	cfg.Alerts.SoundOnReady = false // default is now false
	if !pane.GetEffectiveSound(cfg) {
		t.Error("expected true when override is true (regardless of default)")
	}
}

func TestGetEffectiveNotify(t *testing.T) {
	cfg := config.Default()

	// Pane with no override - use global default (false)
	pane := &Pane{ID: "%0"}
	if pane.GetEffectiveNotify(cfg) {
		t.Error("expected false when no override and default is false")
	}

	// Change global default to true
	cfg.Alerts.NotifyOnReady = true
	if !pane.GetEffectiveNotify(cfg) {
		t.Error("expected true when no override and default is true")
	}

	// Pane with explicit false override
	fl := false
	pane.NotifyOnReady = &fl
	if pane.GetEffectiveNotify(cfg) {
		t.Error("expected false when override is false (regardless of default)")
	}

	// Pane with explicit true override
	tr := true
	pane.NotifyOnReady = &tr
	cfg.Alerts.NotifyOnReady = false // default is now false
	if !pane.GetEffectiveNotify(cfg) {
		t.Error("expected true when override is true (regardless of default)")
	}
}

func TestSetSoundType(t *testing.T) {
	wl := &Watchlist{
		Panes: []Pane{
			{ID: "%0"},
			{ID: "%1"},
		},
	}

	chime := "chime"
	wl.SetSoundType("%1", &chime)

	if wl.Panes[1].SoundType == nil || *wl.Panes[1].SoundType != "chime" {
		t.Error("SetSoundType(%1, chime) did not set SoundType")
	}
	if wl.Panes[0].SoundType != nil {
		t.Error("SetSoundType() changed wrong pane")
	}

	bell := "bell"
	wl.SetSoundType("%1", &bell)
	if wl.Panes[1].SoundType == nil || *wl.Panes[1].SoundType != "bell" {
		t.Error("SetSoundType(%1, bell) did not update SoundType")
	}

	// Test clearing to nil (use default)
	wl.SetSoundType("%1", nil)
	if wl.Panes[1].SoundType != nil {
		t.Error("SetSoundType(%1, nil) did not clear SoundType")
	}
}

func TestGetEffectiveSoundType(t *testing.T) {
	cfg := config.Default()
	cfg.Alerts.SoundType = "chime"

	// Pane with no override - use global default
	pane := &Pane{ID: "%0"}
	if pane.GetEffectiveSoundType(cfg) != "chime" {
		t.Errorf("expected chime when no override, got %s", pane.GetEffectiveSoundType(cfg))
	}

	// Change global default
	cfg.Alerts.SoundType = "bell"
	if pane.GetEffectiveSoundType(cfg) != "bell" {
		t.Errorf("expected bell when no override and default is bell, got %s", pane.GetEffectiveSoundType(cfg))
	}

	// Pane with explicit override
	ping := "ping"
	pane.SoundType = &ping
	if pane.GetEffectiveSoundType(cfg) != "ping" {
		t.Errorf("expected ping when override is ping, got %s", pane.GetEffectiveSoundType(cfg))
	}

	// Pane with empty string override - should use default
	empty := ""
	pane.SoundType = &empty
	if pane.GetEffectiveSoundType(cfg) != "bell" {
		t.Errorf("expected bell when override is empty string, got %s", pane.GetEffectiveSoundType(cfg))
	}
}

func TestSoundTypePersistence(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	wl := &Watchlist{}
	wl.Add("%0")
	ding := "ding"
	wl.SetSoundType("%0", &ding)

	if err := wl.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if loaded.Panes[0].SoundType == nil || *loaded.Panes[0].SoundType != "ding" {
		t.Error("SoundType not persisted")
	}
}

func TestBackwardCompatibilityWithOldWatchlistJSON(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	// Create config directory
	configDir := filepath.Join(tmpDir, ".config", "teejay")
	os.MkdirAll(configDir, 0755)

	// Simulate old watchlist.json with bool values (not pointers)
	// When JSON has "sound_on_ready": false, Go's json.Unmarshal will set *bool to nil
	// When JSON has "sound_on_ready": true, Go's json.Unmarshal will set *bool to &true
	oldJSON := `{
  "panes": [
    {"id": "%0", "added_at": "2024-01-01T00:00:00Z"},
    {"id": "%1", "added_at": "2024-01-01T00:00:00Z", "sound_on_ready": true},
    {"id": "%2", "added_at": "2024-01-01T00:00:00Z", "sound_on_ready": false}
  ]
}`
	os.WriteFile(filepath.Join(configDir, "watchlist.json"), []byte(oldJSON), 0644)

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	// Pane with no field should have nil (use default)
	if loaded.Panes[0].SoundOnReady != nil {
		t.Error("expected nil for pane without sound_on_ready field")
	}

	// Pane with true should have *true
	if loaded.Panes[1].SoundOnReady == nil || !*loaded.Panes[1].SoundOnReady {
		t.Error("expected *true for pane with sound_on_ready: true")
	}

	// Pane with false should have *false (explicit override)
	if loaded.Panes[2].SoundOnReady == nil || *loaded.Panes[2].SoundOnReady {
		t.Error("expected *false for pane with sound_on_ready: false")
	}
}
