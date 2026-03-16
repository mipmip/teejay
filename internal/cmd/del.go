package cmd

import (
	"fmt"

	"tj/internal/naming"
	"tj/internal/tmux"
	"tj/internal/watchlist"
)

func DelPane() error {
	paneID, err := GetTmuxPaneID()
	if err != nil {
		return fmt.Errorf("cannot delete pane: %w", err)
	}

	wl, err := watchlist.Load()
	if err != nil {
		return fmt.Errorf("failed to load watchlist: %w", err)
	}

	if !wl.Contains(paneID) {
		fmt.Printf("Pane %s is not being watched\n", paneID)
		return nil
	}

	// Get pane name before removal for feedback
	var paneName string
	if pane := wl.GetPane(paneID); pane != nil && pane.Name != "" {
		paneName = pane.Name
	} else {
		// Fall back to guessing from tmux metadata
		paneInfo, err := tmux.GetPaneByID(paneID)
		if err == nil && paneInfo != nil {
			paneName, _ = naming.GuessName(*paneInfo)
		}
		if paneName == "" {
			paneName = paneID
		}
	}

	wl.Remove(paneID)

	if err := wl.Save(); err != nil {
		return fmt.Errorf("failed to save watchlist: %w", err)
	}

	fmt.Printf("Removed '%s' from watchlist\n", paneName)
	return nil
}
