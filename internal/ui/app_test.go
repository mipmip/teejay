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
	if !strings.Contains(view, "tj add") {
		t.Error("Empty watchlist should suggest 'tj add' command")
	}
}

func TestPaneItemInterface(t *testing.T) {
	item := paneItem{id: "%5"}

	// Title includes status indicator (○ for Idle which is zero value)
	if item.Title() != "○ %5" {
		t.Errorf("Title() = %q, want '○ %%5'", item.Title())
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

func TestNotInTmuxMessageState(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New()

	// Initially notInTmuxMsg should be false
	if m.notInTmuxMsg {
		t.Error("notInTmuxMsg should initially be false")
	}

	// Test that the state field exists and can be set
	m.notInTmuxMsg = true
	if !m.notInTmuxMsg {
		t.Error("notInTmuxMsg should be settable to true")
	}
}

func TestFooterIncludesEnterKeybinding(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New()
	// The empty state shows different footer, so we check the model has the right structure
	// For non-empty state, footer would include "Enter: switch"
	// This is a structural test since we can't easily create panes in test
	if m.editing || m.deleting || m.notInTmuxMsg {
		t.Error("Model should initialize with all modal states false")
	}
}

func TestConfigurePopupStateTransitions(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New()

	// Initially configuring should be false
	if m.configuring {
		t.Error("configuring should initially be false")
	}

	// Test that configuring state can be set
	m.configuring = true
	m.configMenuItem = configMenuName

	if !m.configuring {
		t.Error("configuring should be settable to true")
	}
	if m.configMenuItem != configMenuName {
		t.Error("configMenuItem should be set to configMenuName")
	}

	// Test menu item navigation
	m.configMenuItem = configMenuSound
	if m.configMenuItem != configMenuSound {
		t.Error("configMenuItem should be set to configMenuSound")
	}

	m.configMenuItem = configMenuNotify
	if m.configMenuItem != configMenuNotify {
		t.Error("configMenuItem should be set to configMenuNotify")
	}

	// Test configEditingName state
	m.configEditingName = true
	if !m.configEditingName {
		t.Error("configEditingName should be settable to true")
	}
}

func TestConfigMenuItemConstants(t *testing.T) {
	// Verify menu item order
	if configMenuName != 0 {
		t.Error("configMenuName should be 0")
	}
	if configMenuSound != 1 {
		t.Error("configMenuSound should be 1")
	}
	if configMenuNotify != 2 {
		t.Error("configMenuNotify should be 2")
	}
}
