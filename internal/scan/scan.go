package scan

import (
	"tj/internal/config"
	"tj/internal/naming"
	"tj/internal/tmux"
	"tj/internal/watchlist"
)

// ScanResult holds the outcome of a scan operation.
type ScanResult struct {
	Found   int
	Added   int
	Skipped int
}

// ScanAndAdd scans the given panes for those running configured agent apps
// and adds new ones to the watchlist. Returns the scan result.
func ScanAndAdd(wl *watchlist.Watchlist, cfg *config.Config, allPanes []tmux.PaneInfo) ScanResult {
	var result ScanResult

	for _, pane := range allPanes {
		if _, ok := cfg.Detection.Apps[pane.Command]; !ok {
			continue
		}
		result.Found++

		if wl.Contains(pane.ID) {
			result.Skipped++
			continue
		}

		name, _ := naming.GuessName(pane)
		wl.AddWithName(pane.ID, name)
		result.Added++
	}

	return result
}
