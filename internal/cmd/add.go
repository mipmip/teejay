package cmd

import (
	"errors"
	"fmt"
	"os"

	"tmon/internal/watchlist"
)

var ErrNotInTmux = errors.New("not running inside tmux")

func GetTmuxPaneID() (string, error) {
	paneID := os.Getenv("TMUX_PANE")
	if paneID == "" {
		return "", ErrNotInTmux
	}
	return paneID, nil
}

func AddPane() error {
	paneID, err := GetTmuxPaneID()
	if err != nil {
		return fmt.Errorf("cannot add pane: %w", err)
	}

	wl, err := watchlist.Load()
	if err != nil {
		return fmt.Errorf("failed to load watchlist: %w", err)
	}

	if wl.Contains(paneID) {
		fmt.Printf("Pane %s is already being watched\n", paneID)
		return nil
	}

	wl.Add(paneID)

	if err := wl.Save(); err != nil {
		return fmt.Errorf("failed to save watchlist: %w", err)
	}

	fmt.Printf("Added pane %s to watchlist\n", paneID)
	return nil
}
