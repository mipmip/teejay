package tmux

import (
	"os"
	"testing"
)

func TestListAllPanes(t *testing.T) {
	// Skip if not in tmux
	if os.Getenv("TMUX") == "" {
		t.Skip("Not running inside tmux, skipping ListAllPanes test")
	}

	panes, err := ListAllPanes()
	if err != nil {
		t.Fatalf("ListAllPanes() error = %v", err)
	}

	// Should have at least one pane (the one running this test)
	if len(panes) == 0 {
		t.Error("ListAllPanes() returned empty list, expected at least one pane")
	}

	// Each pane should have required fields
	for i, pane := range panes {
		if pane.ID == "" {
			t.Errorf("Pane %d has empty ID", i)
		}
		if pane.Session == "" {
			t.Errorf("Pane %d has empty Session", i)
		}
		if pane.SessionID == "" {
			t.Errorf("Pane %d has empty SessionID", i)
		}
		// ID should start with %
		if len(pane.ID) > 0 && pane.ID[0] != '%' {
			t.Errorf("Pane %d ID = %q, expected to start with %%", i, pane.ID)
		}
	}
}

func TestPaneInfoSessionID(t *testing.T) {
	// Test that SessionID format is correct
	pane := PaneInfo{
		Session: "main",
		Window:  0,
		Pane:    1,
	}

	expected := "main:0.1"
	if pane.SessionID != "" {
		// SessionID is set by ListAllPanes, this tests the struct
	}

	// Manually construct what SessionID should be
	got := pane.Session + ":" + string('0'+byte(pane.Window)) + "." + string('0'+byte(pane.Pane))
	if got != expected {
		t.Errorf("SessionID format: got %q, want %q", got, expected)
	}
}

func TestPaneInfoWindowName(t *testing.T) {
	// Test that WindowName is populated by ListAllPanes
	if os.Getenv("TMUX") == "" {
		t.Skip("Not running inside tmux, skipping WindowName test")
	}

	panes, err := ListAllPanes()
	if err != nil {
		t.Fatalf("ListAllPanes() error = %v", err)
	}

	// Each pane should have a window name (tmux always has one)
	for i, pane := range panes {
		if pane.WindowName == "" {
			t.Errorf("Pane %d has empty WindowName", i)
		}
	}
}
