package alerts

import (
	"testing"
)

func TestPlayBell(t *testing.T) {
	// PlayBell now uses native sound playback (via sounds.PlaySound)
	// This is a smoke test - actual audio is system-dependent
	// If we get here without panicking, the test passes
	PlayBell()
}

func TestPlaySound(t *testing.T) {
	// PlaySound uses native sound playback
	// This is a smoke test - actual audio is system-dependent
	// If we get here without panicking, the test passes
	PlaySound("chime")
	PlaySound("bell")
	PlaySound("invalid") // should fall back gracefully
}

func TestSendNotification(t *testing.T) {
	// Just verify it doesn't panic when notify-send is not available
	// This is a smoke test - actual notification is system-dependent
	SendNotification("Test", "Test message")
	// If we get here without panicking, the test passes
}
