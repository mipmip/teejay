package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"tj/internal/naming"
	"tj/internal/tmux"
	"tj/internal/watchlist"
)

var ErrNotInTmux = errors.New("not running inside tmux")

// Stdin can be overridden for testing
var Stdin io.Reader = os.Stdin

func GetTmuxPaneID() (string, error) {
	paneID := os.Getenv("TMUX_PANE")
	if paneID == "" {
		return "", ErrNotInTmux
	}
	return paneID, nil
}

// AddPane adds the current tmux pane to the watchlist.
// If watchlistPath is provided, uses that path instead of the default.
func AddPane(watchlistPath ...string) error {
	paneID, err := GetTmuxPaneID()
	if err != nil {
		return fmt.Errorf("cannot add pane: %w", err)
	}

	var wl *watchlist.Watchlist
	if len(watchlistPath) > 0 && watchlistPath[0] != "" {
		wl, err = watchlist.Load(watchlistPath[0])
	} else {
		wl, err = watchlist.Load()
	}
	if err != nil {
		return fmt.Errorf("failed to load watchlist: %w", err)
	}

	if wl.Contains(paneID) {
		fmt.Printf("Pane %s is already being watched\n", paneID)
		return nil
	}

	// Get pane info for name guessing
	paneInfo, err := tmux.GetPaneByID(paneID)
	if err != nil {
		return fmt.Errorf("failed to get pane info: %w", err)
	}

	var paneName string
	if paneInfo != nil {
		guessedName, isGeneric := naming.GuessName(*paneInfo)
		if isGeneric {
			// Prompt user for a name
			fmt.Printf("Enter a name for this pane (suggested: %s): ", guessedName)
			reader := bufio.NewReader(Stdin)
			input, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				return fmt.Errorf("failed to read input: %w", err)
			}
			input = strings.TrimSpace(input)
			if input != "" {
				paneName = input
			} else {
				paneName = guessedName
			}
		} else {
			paneName = guessedName
		}
	}

	wl.AddWithName(paneID, paneName)

	if err := wl.Save(); err != nil {
		return fmt.Errorf("failed to save watchlist: %w", err)
	}

	fmt.Printf("Added pane %s as '%s' to watchlist\n", paneID, paneName)
	return nil
}
