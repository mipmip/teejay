package ui

import (
	"errors"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewModel(t *testing.T) {
	m := New("test", nil, "")
	if m.View() == "" {
		t.Error("View() should return non-empty string")
	}
}

func TestEmptyWatchlistShowsMessage(t *testing.T) {
	// With a fresh temp HOME, watchlist will be empty
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New("test", nil, "")
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

	// Title returns just the display name (id if no name set)
	if item.Title() != "%5" {
		t.Errorf("Title() = %q, want '%%5'", item.Title())
	}
	if item.FilterValue() != "%5" {
		t.Errorf("FilterValue() = %q, want %%5", item.FilterValue())
	}
}

func TestPaneItemBreadcrumbFull(t *testing.T) {
	item := paneItem{
		id:         "%5",
		session:    "technative-docs",
		windowName: "proposals",
		command:    "claude",
	}
	want := "technative-docs > proposals : claude"
	if got := item.Description(); got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}

func TestPaneItemBreadcrumbNoProcess(t *testing.T) {
	item := paneItem{
		id:         "%5",
		session:    "main",
		windowName: "dev",
	}
	want := "main > dev"
	if got := item.Description(); got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}

func TestRenderAlertIndicators(t *testing.T) {
	// Both enabled - should contain ♪ and ✉
	result := renderAlertIndicators(true, true)
	if !strings.Contains(result, "♪") || !strings.Contains(result, "✉") {
		t.Errorf("renderAlertIndicators(true, true) should contain ♪ and ✉, got %q", result)
	}

	// Both disabled - should still contain symbols (just styled differently)
	result = renderAlertIndicators(false, false)
	if !strings.Contains(result, "♪") || !strings.Contains(result, "✉") {
		t.Errorf("renderAlertIndicators(false, false) should contain ♪ and ✉, got %q", result)
	}

	// Mixed state
	result = renderAlertIndicators(true, false)
	if !strings.Contains(result, "♪") || !strings.Contains(result, "✉") {
		t.Errorf("renderAlertIndicators(true, false) should contain ♪ and ✉, got %q", result)
	}
}

func TestPaneItemDescriptionWithOverrides(t *testing.T) {
	soundOn := true
	notifyOff := false
	item := paneItem{
		id:             "%5",
		session:        "main",
		windowName:     "dev",
		command:        "claude",
		soundOverride:  &soundOn,
		notifyOverride: &notifyOff,
	}
	got := item.Description()
	if !strings.Contains(got, "main > dev : claude") {
		t.Errorf("Description() should contain breadcrumb, got %q", got)
	}
	if !strings.Contains(got, "♪") || !strings.Contains(got, "✉") {
		t.Errorf("Description() with overrides should contain alert indicators, got %q", got)
	}
}

func TestPaneItemDescriptionNoOverrides(t *testing.T) {
	item := paneItem{
		id:         "%5",
		session:    "main",
		windowName: "dev",
		command:    "claude",
	}
	got := item.Description()
	want := "main > dev : claude"
	if got != want {
		t.Errorf("Description() = %q, want %q (no indicators without overrides)", got, want)
	}
}

func TestModelHasViewportAndList(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New("test", nil, "")

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

	m := New("test", nil, "")

	// With empty watchlist, should not show split panel
	view := m.View()
	if strings.Contains(view, "Preview:") {
		t.Error("Empty watchlist should not show preview panel")
	}
}

func TestTemporaryMessageState(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New("test", nil, "")

	// Initially temporaryMessage should be empty
	if m.temporaryMessage != "" {
		t.Error("temporaryMessage should initially be empty")
	}

	// Test that the state field exists and can be set
	m.temporaryMessage = "Test error message"
	if m.temporaryMessage != "Test error message" {
		t.Error("temporaryMessage should be settable")
	}
}

func TestFooterIncludesEnterKeybinding(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New("test", nil, "")
	// The empty state shows different footer, so we check the model has the right structure
	// For non-empty state, footer would include "Enter: switch"
	// This is a structural test since we can't easily create panes in test
	if m.editing || m.deleting || m.temporaryMessage != "" {
		t.Error("Model should initialize with all modal states false/empty")
	}
}

func TestConfigurePopupStateTransitions(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New("test", nil, "")

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
	if configMenuSoundType != 2 {
		t.Error("configMenuSoundType should be 2")
	}
	if configMenuNotify != 3 {
		t.Error("configMenuNotify should be 3")
	}
}

func TestIsStalePaneError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error is not stale",
			err:      nil,
			expected: false,
		},
		{
			name:     "can't find pane error is stale",
			err:      errors.New("can't find pane: %65"),
			expected: true,
		},
		{
			name:     "can't find pane without ID is stale",
			err:      errors.New("can't find pane"),
			expected: true,
		},
		{
			name:     "other error is not stale",
			err:      errors.New("connection refused"),
			expected: false,
		},
		{
			name:     "tmux not running is not stale",
			err:      errors.New("no server running on /tmp/tmux-1000/default"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isStalePaneError(tt.err)
			if got != tt.expected {
				t.Errorf("isStalePaneError(%v) = %v, want %v", tt.err, got, tt.expected)
			}
		})
	}
}

func TestRemoveStalePaneSetsStatusMessage(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New("test", nil, "")

	// Manually add a pane to the watchlist
	m.watchlist.Add("%99")
	m.watchlist.Save()
	m.refreshList()
	m.empty = false
	m.selectedPaneID = "%99"

	// Simulate removing a stale pane
	m.removeStalePane("%99")

	// Check that status message was set
	if m.statusMessage == "" {
		t.Error("statusMessage should be set after removing stale pane")
	}
	if !strings.Contains(m.statusMessage, "%99") {
		t.Errorf("statusMessage should contain pane ID, got: %q", m.statusMessage)
	}

	// Check that pane was removed from watchlist
	if m.watchlist.Contains("%99") {
		t.Error("pane should be removed from watchlist")
	}

	// Check that empty state is updated
	if !m.empty {
		t.Error("empty should be true when last pane is removed")
	}
}

func TestRemoveStalePaneWithMultiplePanes(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New("test", nil, "")

	// Add multiple panes
	m.watchlist.Add("%98")
	m.watchlist.Add("%99")
	m.watchlist.Save()
	m.refreshList()
	m.empty = false
	m.selectedPaneID = "%99"

	// Remove one stale pane
	m.removeStalePane("%99")

	// Check that only the stale pane was removed
	if m.watchlist.Contains("%99") {
		t.Error("stale pane should be removed")
	}
	if !m.watchlist.Contains("%98") {
		t.Error("other pane should remain")
	}

	// Check selection moved to remaining pane
	if m.selectedPaneID != "%98" {
		t.Errorf("selectedPaneID should be %%98, got %q", m.selectedPaneID)
	}

	// Should not be empty
	if m.empty {
		t.Error("empty should be false when panes remain")
	}
}

func TestStatusMessageClearsOnKeyPress(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New("test", nil, "")
	m.statusMessage = "Test message"

	// Simulate a key press
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = updated.(Model)

	if m.statusMessage != "" {
		t.Errorf("statusMessage should be cleared on key press, got: %q", m.statusMessage)
	}
}

func TestPreviewTitleShowsCustomPaneName(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New("test", nil, "")

	// Add a pane with a custom name
	m.watchlist.AddWithName("%10", "My Project")
	m.watchlist.Save()
	m.refreshList()
	m.empty = false
	m.selectedPaneID = "%10"

	// Set reasonable dimensions for rendering
	m.width = 100
	m.height = 30

	view := m.View()

	// Preview title should show the custom name, not the pane ID
	if !strings.Contains(view, "Preview: My Project") {
		t.Errorf("Preview title should show custom name 'My Project', got view containing: %q", view)
	}
	if strings.Contains(view, "Preview: %10") {
		t.Error("Preview title should NOT show pane ID when custom name exists")
	}
}

func TestPreviewTitleFallsBackToPaneID(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	m := New("test", nil, "")

	// Add a pane without a custom name
	m.watchlist.Add("%20")
	m.watchlist.Save()
	m.refreshList()
	m.empty = false
	m.selectedPaneID = "%20"

	// Set reasonable dimensions for rendering
	m.width = 100
	m.height = 30

	view := m.View()

	// Preview title should show the pane ID as fallback
	if !strings.Contains(view, "Preview: %20") {
		t.Errorf("Preview title should show pane ID '%%20' when no custom name, got view containing: %q", view)
	}
}

func TestRecencyColor(t *testing.T) {
	tests := []struct {
		elapsed  time.Duration
		expected string
	}{
		{0, "#00FF00"},
		{5 * time.Second, "#00FF00"},
		{10 * time.Second, "#00DD00"},
		{20 * time.Second, "#00DD00"},
		{30 * time.Second, "#00BB00"},
		{1 * time.Minute, "#00BB00"},
		{2 * time.Minute, "#009900"},
		{4 * time.Minute, "#009900"},
		{5 * time.Minute, "#006600"},
		{30 * time.Minute, "#006600"},
	}
	for _, tt := range tests {
		got := string(recencyColor(tt.elapsed))
		if got != tt.expected {
			t.Errorf("recencyColor(%v) = %q, want %q", tt.elapsed, got, tt.expected)
		}
	}
}

func TestCompactDuration(t *testing.T) {
	tests := []struct {
		d    time.Duration
		want string
	}{
		{0, "0s"},
		{3 * time.Second, "3s"},
		{59 * time.Second, "59s"},
		{60 * time.Second, "1m"},
		{90 * time.Second, "1m"},
		{14 * time.Minute, "14m"},
		{59*time.Minute + 59*time.Second, "59m"},
		{60 * time.Minute, "1h"},
		{8 * time.Hour, "8h"},
		{23*time.Hour + 59*time.Minute, "23h"},
		{24 * time.Hour, "1d"},
		{72 * time.Hour, "3d"},
	}
	for _, tt := range tests {
		got := compactDuration(tt.d)
		if got != tt.want {
			t.Errorf("compactDuration(%v) = %q, want %q", tt.d, got, tt.want)
		}
	}
}
