package alerts

import (
	"bytes"
	"os"
	"testing"
)

func TestPlayBell(t *testing.T) {
	// Capture stdout to verify bell character is written
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PlayBell()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)

	if buf.String() != "\a" {
		t.Errorf("PlayBell() output = %q, want %q", buf.String(), "\a")
	}
}

func TestSendNotification(t *testing.T) {
	// Just verify it doesn't panic when notify-send is not available
	// This is a smoke test - actual notification is system-dependent
	SendNotification("Test", "Test message")
	// If we get here without panicking, the test passes
}
