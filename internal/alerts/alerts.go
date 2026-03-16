package alerts

import (
	"fmt"
	"os"
	"os/exec"

	"tj/internal/alerts/sounds"
)

// PlayBell plays the default notification sound (chime).
// This maintains backwards compatibility with existing code.
func PlayBell() {
	sounds.PlaySound(sounds.DefaultSound)
}

// PlaySound plays the specified notification sound.
// Falls back to terminal bell if audio playback fails.
func PlaySound(soundType string) {
	sounds.PlaySound(soundType)
}

// PlayTerminalBell plays the terminal bell sound (fallback).
func PlayTerminalBell() {
	fmt.Fprint(os.Stdout, "\a")
}

// SendNotification sends a desktop notification using notify-send.
// Fails silently if notify-send is not available.
func SendNotification(title, message string) {
	cmd := exec.Command("notify-send", title, message)
	_ = cmd.Run() // Ignore errors - fail silently if not available
}
