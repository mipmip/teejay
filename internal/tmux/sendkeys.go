package tmux

import (
	"fmt"
	"os/exec"
	"strings"
)

// SendKeys sends keystrokes to a tmux pane followed by Enter.
// Use for free-text input where the agent expects a line of text.
func SendKeys(paneID string, keys string) error {
	cmd := exec.Command("tmux", "send-keys", "-t", paneID, keys, "Enter")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("tmux send-keys failed: %s", strings.TrimSpace(string(output)))
	}
	return nil
}

// SendKeysThenEnter sends keystrokes and Enter as two separate tmux calls.
// Some TUI apps (e.g., opencode) swallow Enter when it arrives in the same
// send-keys invocation as the text, because the app is still processing the
// literal characters. Splitting into two calls adds enough delay.
func SendKeysThenEnter(paneID string, keys string) error {
	if keys != "" {
		cmd := exec.Command("tmux", "send-keys", "-t", paneID, keys)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("tmux send-keys failed: %s", strings.TrimSpace(string(output)))
		}
	}
	cmd := exec.Command("tmux", "send-keys", "-t", paneID, "Enter")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("tmux send-keys Enter failed: %s", strings.TrimSpace(string(output)))
	}
	return nil
}

// SendRawKey sends a single key to a tmux pane without appending Enter.
// Use for TUI prompts that read single keypresses (e.g., y/n permission prompts).
func SendRawKey(paneID string, key string) error {
	cmd := exec.Command("tmux", "send-keys", "-t", paneID, key)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("tmux send-keys failed: %s", strings.TrimSpace(string(output)))
	}
	return nil
}

// SendArrowAndEnter sends N down-arrow presses followed by Enter.
// Use for interactive selection menus where you navigate with arrows then confirm.
func SendArrowAndEnter(paneID string, downPresses int) error {
	for i := 0; i < downPresses; i++ {
		cmd := exec.Command("tmux", "send-keys", "-t", paneID, "Down")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("tmux send-keys Down failed: %s", strings.TrimSpace(string(output)))
		}
	}
	cmd := exec.Command("tmux", "send-keys", "-t", paneID, "Enter")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("tmux send-keys Enter failed: %s", strings.TrimSpace(string(output)))
	}
	return nil
}
