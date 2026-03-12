package ui

import "testing"

func TestNewModel(t *testing.T) {
	m := New()
	if m.View() == "" {
		t.Error("View() should return non-empty string")
	}
}
