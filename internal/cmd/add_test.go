package cmd

import (
	"errors"
	"testing"
)

func TestGetTmuxPaneID_Set(t *testing.T) {
	t.Setenv("TMUX_PANE", "%42")

	paneID, err := GetTmuxPaneID()
	if err != nil {
		t.Fatalf("GetTmuxPaneID() error = %v, want nil", err)
	}
	if paneID != "%42" {
		t.Errorf("GetTmuxPaneID() = %q, want %%42", paneID)
	}
}

func TestGetTmuxPaneID_NotSet(t *testing.T) {
	t.Setenv("TMUX_PANE", "")

	_, err := GetTmuxPaneID()
	if err == nil {
		t.Fatal("GetTmuxPaneID() error = nil, want ErrNotInTmux")
	}
	if !errors.Is(err, ErrNotInTmux) {
		t.Errorf("GetTmuxPaneID() error = %v, want ErrNotInTmux", err)
	}
}
