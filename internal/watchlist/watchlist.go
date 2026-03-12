package watchlist

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Pane struct {
	ID      string    `json:"id"`
	Name    string    `json:"name,omitempty"`
	AddedAt time.Time `json:"added_at"`
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

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "tmon", "watchlist.json"), nil
}

func Load() (*Watchlist, error) {
	path, err := configPath()
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
	path, err := configPath()
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
