package watchlist

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Pane struct {
	ID            string    `json:"id"`
	Name          string    `json:"name,omitempty"`
	AddedAt       time.Time `json:"added_at"`
	SoundOnReady  bool      `json:"sound_on_ready,omitempty"`
	NotifyOnReady bool      `json:"notify_on_ready,omitempty"`
}

// DisplayName returns the custom name if set, otherwise the pane ID.
func (p Pane) DisplayName() string {
	if p.Name != "" {
		return p.Name
	}
	return p.ID
}

type Watchlist struct {
	Panes []Pane `json:"panes"`
}

// ConfigPath returns the path to the watchlist configuration file.
func ConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "teejay", "watchlist.json"), nil
}

func Load() (*Watchlist, error) {
	path, err := ConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &Watchlist{Panes: []Pane{}}, nil
	}
	if err != nil {
		return nil, err
	}

	var wl Watchlist
	if err := json.Unmarshal(data, &wl); err != nil {
		return nil, err
	}
	wl.Deduplicate()
	return &wl, nil
}

func (wl *Watchlist) Save() error {
	path, err := ConfigPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(wl, "", "  ")
	if err != nil {
		return err
	}

	// Atomic write: write to temp file, then rename
	tmpPath := path + ".tmp"
	if err := os.WriteFile(tmpPath, data, 0644); err != nil {
		return err
	}
	return os.Rename(tmpPath, path)
}

func (wl *Watchlist) Add(paneID string) {
	wl.Panes = append(wl.Panes, Pane{
		ID:      paneID,
		AddedAt: time.Now(),
	})
}

// Contains returns true if the pane ID is already in the watchlist.
func (wl *Watchlist) Contains(paneID string) bool {
	for _, p := range wl.Panes {
		if p.ID == paneID {
			return true
		}
	}
	return false
}

// Deduplicate removes duplicate pane entries, keeping the first occurrence.
func (wl *Watchlist) Deduplicate() {
	seen := make(map[string]bool)
	unique := make([]Pane, 0, len(wl.Panes))
	for _, p := range wl.Panes {
		if !seen[p.ID] {
			seen[p.ID] = true
			unique = append(unique, p)
		}
	}
	wl.Panes = unique
}

// Remove removes a pane by ID from the watchlist.
func (wl *Watchlist) Remove(paneID string) {
	filtered := make([]Pane, 0, len(wl.Panes))
	for _, p := range wl.Panes {
		if p.ID != paneID {
			filtered = append(filtered, p)
		}
	}
	wl.Panes = filtered
}

// Rename sets the display name for a pane by ID.
func (wl *Watchlist) Rename(paneID, name string) {
	for i := range wl.Panes {
		if wl.Panes[i].ID == paneID {
			wl.Panes[i].Name = name
			return
		}
	}
}

// SetSound sets the SoundOnReady flag for a pane by ID.
func (wl *Watchlist) SetSound(paneID string, enabled bool) {
	for i := range wl.Panes {
		if wl.Panes[i].ID == paneID {
			wl.Panes[i].SoundOnReady = enabled
			return
		}
	}
}

// SetNotify sets the NotifyOnReady flag for a pane by ID.
func (wl *Watchlist) SetNotify(paneID string, enabled bool) {
	for i := range wl.Panes {
		if wl.Panes[i].ID == paneID {
			wl.Panes[i].NotifyOnReady = enabled
			return
		}
	}
}

// GetPane returns a pointer to a Pane by ID, or nil if not found.
func (wl *Watchlist) GetPane(paneID string) *Pane {
	for i := range wl.Panes {
		if wl.Panes[i].ID == paneID {
			return &wl.Panes[i]
		}
	}
	return nil
}
