package prompt

import (
	"regexp"
	"strings"
)

// ansiRegex matches ANSI escape sequences for stripping before parsing.
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]|\x1b\][^\x1b]*\x1b\\|\x1b\[[\?]?[0-9;]*[a-zA-Z]`)

// menuItemRegex matches lines like "❯ 1. Yes" or "  2. No, deny" or "  3. Something"
var menuItemRegex = regexp.MustCompile(`^\s*[❯>]?\s*(\d+)\.\s+(.+)$`)

// ScrapePrompt extracts basic prompt context from captured pane content.
// This is the fallback for non-Claude agents.
func ScrapePrompt(capturedContent string) PromptInfo {
	lines := strings.Split(strings.TrimRight(capturedContent, "\n"), "\n")

	// Extract last few non-empty lines as context
	var context []string
	for i := len(lines) - 1; i >= 0 && len(context) < 5; i-- {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			context = append([]string{line}, context...)
		}
	}

	return PromptInfo{
		Type:         FreeInput,
		QuestionText: strings.Join(context, "\n"),
	}
}

// ScrapeMenuOptions parses an interactive numbered menu from captured pane content.
// Returns the options in display order, and any question/context text above the menu.
// Returns nil options if no menu pattern is found.
func ScrapeMenuOptions(capturedContent string) (question string, options []Option) {
	// Strip ANSI codes for reliable matching
	clean := ansiRegex.ReplaceAllString(capturedContent, "")
	lines := strings.Split(clean, "\n")

	// Find all numbered menu items by scanning from the end
	type menuItem struct {
		index int    // line index
		num   string // the number (e.g., "1", "2")
		label string // the text after the number
	}
	var items []menuItem

	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		m := menuItemRegex.FindStringSubmatch(line)
		if m != nil {
			items = append([]menuItem{{index: i, num: m[1], label: strings.TrimSpace(m[2])}}, items...)
		} else if len(items) > 0 {
			// We've moved past the menu block — stop
			break
		}
	}

	if len(items) == 0 {
		return "", nil
	}

	// Extract question text: non-empty lines above the first menu item
	var questionLines []string
	for i := items[0].index - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			if len(questionLines) > 0 {
				break // blank line gap above question
			}
			continue
		}
		questionLines = append([]string{line}, questionLines...)
		if len(questionLines) >= 3 {
			break
		}
	}

	options = make([]Option, len(items))
	for i, item := range items {
		options[i] = Option{
			Key:   item.num,
			Label: item.label,
		}
	}

	return strings.Join(questionLines, "\n"), options
}
