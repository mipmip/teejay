package watchlist

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Pane struct {
	ID      string    `json:"id"`
	AddedAt time.Time `json:"added_at"`
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
