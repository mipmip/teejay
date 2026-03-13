package alerts

import (
	"fmt"
	"os"
	"os/exec"
)

// PlayBell plays the terminal bell sound.
func PlayBell() {
	fmt.Fprint(os.Stdout, "\a")
}

// SendNotification sends a desktop notification using notify-send.
// Fails silently if notify-send is not available.
func SendNotification(title, message string) {
	cmd := exec.Command("notify-send", title, message)
	_ = cmd.Run() // Ignore errors - fail silently if not available
}
