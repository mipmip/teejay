package cmd

import (
	"fmt"

	"tj/internal/config"
	"tj/internal/scan"
	"tj/internal/tmux"
	"tj/internal/watchlist"
)

// ScanPanes scans all tmux panes for configured agent apps and adds them to the watchlist.
func ScanPanes(cfg *config.Config, watchlistPath ...string) error {
	var wl *watchlist.Watchlist
	var err error
	if len(watchlistPath) > 0 && watchlistPath[0] != "" {
		wl, err = watchlist.Load(watchlistPath[0])
	} else {
		wl, err = watchlist.Load()
	}
	if err != nil {
		return fmt.Errorf("failed to load watchlist: %w", err)
	}

	allPanes, err := tmux.ListAllPanes()
	if err != nil {
		return fmt.Errorf("failed to list tmux panes: %w", err)
	}

	result := scan.ScanAndAdd(wl, cfg, allPanes)

	if result.Added > 0 {
		if err := wl.Save(); err != nil {
			return fmt.Errorf("failed to save watchlist: %w", err)
		}
	}

	if result.Found == 0 {
		fmt.Println("No agent panes found")
	} else {
		fmt.Printf("Found %d agent panes: added %d, skipped %d (already watched)\n", result.Found, result.Added, result.Skipped)
	}

	return nil
}
