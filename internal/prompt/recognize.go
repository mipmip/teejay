package prompt

import "tj/internal/tmux"

// Recognize determines the prompt state of a waiting pane.
// For Claude panes, uses structured transcript data.
// For other agents, falls back to screen scraping.
func Recognize(paneID string, appName string) PromptInfo {
	if appName == "claude" {
		return RecognizeClaude(paneID)
	}

	// Fallback: capture pane content and scrape
	content, err := tmux.CapturePane(paneID)
	if err != nil {
		return PromptInfo{Type: Unknown}
	}
	return ScrapePrompt(content)
}
