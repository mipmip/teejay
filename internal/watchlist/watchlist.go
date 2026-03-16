package watchlist

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"tj/internal/config"
)

type Pane struct {
	ID            string    `json:"id"`
	Name          string    `json:"name,omitempty"`
	AddedAt       time.Time `json:"added_at"`
	SoundOnReady  *bool     `json:"sound_on_ready,omitempty"`
	NotifyOnReady *bool     `json:"notify_on_ready,omitempty"`
	SoundType     *string   `json:"sound_type,omitempty"`
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
	path  string // path to the watchlist file (not serialized)
}

// ConfigPath returns the path to the watchlist configuration file.
func ConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "teejay", "watchlist.json"), nil
}

// Load reads the watchlist from the specified path, or ~/.config/teejay/watchlist.json if not provided.
// The watchlist remembers its path for subsequent Save() calls.
func Load(customPath ...string) (*Watchlist, error) {
	var path string
	var err error

	if len(customPath) > 0 && customPath[0] != "" {
		path = customPath[0]
	} else {
		path, err = ConfigPath()
		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &Watchlist{Panes: []Pane{}, path: path}, nil
	}
	if err != nil {
		return nil, err
	}

	var wl Watchlist
	if err := json.Unmarshal(data, &wl); err != nil {
		return nil, err
	}
	wl.path = path
	wl.Deduplicate()
	return &wl, nil
}

// Save writes the watchlist to the path it was loaded from.
// If the watchlist was created without loading, it saves to the default path.
func (wl *Watchlist) Save() error {
	path := wl.path
	if path == "" {
		var err error
		path, err = ConfigPath()
		if err != nil {
			return err
		}
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

// AddWithName adds a new pane to the watchlist with a specified name.
func (wl *Watchlist) AddWithName(paneID, name string) {
	wl.Panes = append(wl.Panes, Pane{
		ID:      paneID,
		Name:    name,
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
// Pass nil to clear the override and use the global default.
func (wl *Watchlist) SetSound(paneID string, enabled *bool) {
	for i := range wl.Panes {
		if wl.Panes[i].ID == paneID {
			wl.Panes[i].SoundOnReady = enabled
			return
		}
	}
}

// SetNotify sets the NotifyOnReady flag for a pane by ID.
// Pass nil to clear the override and use the global default.
func (wl *Watchlist) SetNotify(paneID string, enabled *bool) {
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

// GetEffectiveSound returns whether sound alerts should be used for a pane,
// considering the pane's override and the global default.
func (p *Pane) GetEffectiveSound(cfg *config.Config) bool {
	if p.SoundOnReady != nil {
		return *p.SoundOnReady
	}
	return cfg.Alerts.SoundOnReady
}

// GetEffectiveNotify returns whether desktop notifications should be used for a pane,
// considering the pane's override and the global default.
func (p *Pane) GetEffectiveNotify(cfg *config.Config) bool {
	if p.NotifyOnReady != nil {
		return *p.NotifyOnReady
	}
	return cfg.Alerts.NotifyOnReady
}

// GetEffectiveSoundType returns the sound type to use for a pane,
// considering the pane's override and the global default.
func (p *Pane) GetEffectiveSoundType(cfg *config.Config) string {
	if p.SoundType != nil && *p.SoundType != "" {
		return *p.SoundType
	}
	return cfg.Alerts.SoundType
}

// SetSoundType sets the sound type override for a pane by ID.
// Pass nil to clear the override and use the global default.
func (wl *Watchlist) SetSoundType(paneID string, soundType *string) {
	for i := range wl.Panes {
		if wl.Panes[i].ID == paneID {
			wl.Panes[i].SoundType = soundType
			return
		}
	}
}
