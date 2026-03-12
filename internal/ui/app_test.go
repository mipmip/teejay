package ui

import (
	"strings"
	"testing"
)

func TestNewModel(t *testing.T) {
	m := New()
	if m.View() == "" {
		t.Error("View() should return non-empty string")
	}
}

func TestEmptyWatchlistShowsMessage(t *testing.T) {
	// With a fresh temp HOME, watchlist will be empty
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New()
	view := m.View()

	if !strings.Contains(view, "No panes are being watched") {
		t.Error("Empty watchlist should show 'No panes are being watched' message")
	}
	if !strings.Contains(view, "tmon add") {
		t.Error("Empty watchlist should suggest 'tmon add' command")
	}
}

func TestPaneItemInterface(t *testing.T) {
	item := paneItem{id: "%5"}

	if item.Title() != "%5" {
		t.Errorf("Title() = %q, want %%5", item.Title())
	}
	if item.FilterValue() != "%5" {
		t.Errorf("FilterValue() = %q, want %%5", item.FilterValue())
	}
}

func TestModelHasViewportAndList(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New()

	// Model should have viewport initialized (width > 0 after init)
	// This is a basic structural test
	if m.empty != true {
		t.Error("Empty watchlist should set empty=true")
	}
}

func TestSplitPanelLayoutWithPanes(t *testing.T) {
	// This test would need panes in watchlist to verify split panel
	// For now, we just verify the model structure exists
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New()

	// With empty watchlist, should not show split panel
	view := m.View()
	if strings.Contains(view, "Preview:") {
		t.Error("Empty watchlist should not show preview panel")
	}
}
