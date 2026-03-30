package prompt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"tj/internal/tmux"
)

// SessionInfo holds the Claude Code session mapping.
type SessionInfo struct {
	PID       int    `json:"pid"`
	SessionID string `json:"sessionId"`
	Cwd       string `json:"cwd"`
}

// TranscriptEntry represents a parsed entry from a Claude transcript .jsonl file.
type TranscriptEntry struct {
	Type    string          `json:"type"`
	Message json.RawMessage `json:"message"`
}

// assistantMessage is the parsed assistant message from the transcript.
type assistantMessage struct {
	StopReason string        `json:"stop_reason"`
	Content    []contentItem `json:"content"`
}

// contentItem represents a content block in an assistant message.
type contentItem struct {
	Type  string          `json:"type"`
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Input json.RawMessage `json:"input"`
}

// askUserQuestionInput represents the input to AskUserQuestion tool.
type askUserQuestionInput struct {
	Questions []askUserQuestion `json:"questions"`
}

type askUserQuestion struct {
	Question string                 `json:"question"`
	Options  []askUserQuestionOption `json:"options"`
}

type askUserQuestionOption struct {
	Label       string `json:"label"`
	Description string `json:"description"`
}

// ReadClaudeSession reads the session file for a Claude process and returns session info.
func ReadClaudeSession(claudePID int) (*SessionInfo, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot get home dir: %w", err)
	}

	sessionPath := filepath.Join(homeDir, ".claude", "sessions", fmt.Sprintf("%d.json", claudePID))
	data, err := os.ReadFile(sessionPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read session file: %w", err)
	}

	var info SessionInfo
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, fmt.Errorf("cannot parse session file: %w", err)
	}
	return &info, nil
}

// FindTranscript locates the transcript .jsonl file for a given session.
func FindTranscript(sessionID, cwd string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot get home dir: %w", err)
	}

	// Derive project hash: replace "/" with "-", prefix with "-"
	projectHash := "-" + strings.ReplaceAll(strings.TrimPrefix(cwd, "/"), "/", "-")
	transcriptPath := filepath.Join(homeDir, ".claude", "projects", projectHash, sessionID+".jsonl")

	if _, err := os.Stat(transcriptPath); err != nil {
		return "", fmt.Errorf("transcript not found: %w", err)
	}
	return transcriptPath, nil
}

// ReadLastAssistant reads the last assistant entry from a transcript file.
// Only reads the tail of the file (~64KB) for efficiency, since the last
// assistant entry is always near the end.
func ReadLastAssistant(transcriptPath string) (*assistantMessage, error) {
	f, err := os.Open(transcriptPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Seek to the tail of the file — 64KB is plenty for the last few entries
	const tailSize = 64 * 1024
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	offset := info.Size() - tailSize
	if offset < 0 {
		offset = 0
	}
	if offset > 0 {
		f.Seek(offset, 0)
	}

	// Read lines from the tail
	var lines []string
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// If we seeked into the middle of a line, the first line is partial — skip it
	if offset > 0 && len(lines) > 0 {
		lines = lines[1:]
	}

	// Search backwards for the last assistant entry
	for i := len(lines) - 1; i >= 0; i-- {
		var entry TranscriptEntry
		if err := json.Unmarshal([]byte(lines[i]), &entry); err != nil {
			continue
		}
		if entry.Type != "assistant" || entry.Message == nil {
			continue
		}

		var msg assistantMessage
		if err := json.Unmarshal(entry.Message, &msg); err != nil {
			continue
		}
		return &msg, nil
	}
	return nil, fmt.Errorf("no assistant entry found in transcript")
}

// ParsePrompt converts a Claude transcript assistant message to PromptInfo.
func ParsePrompt(msg *assistantMessage) PromptInfo {
	if msg == nil {
		return PromptInfo{Type: Unknown}
	}

	// end_turn means agent finished and is at the main prompt
	if msg.StopReason == "end_turn" {
		return PromptInfo{Type: FreeInput}
	}

	// tool_use means agent proposed a tool call
	if msg.StopReason == "tool_use" {
		for _, c := range msg.Content {
			if c.Type != "tool_use" {
				continue
			}

			// AskUserQuestion is a special case — it's a question, not a permission
			if c.Name == "AskUserQuestion" {
				var input askUserQuestionInput
				if err := json.Unmarshal(c.Input, &input); err == nil && len(input.Questions) > 0 {
					q := input.Questions[0] // Handle first question
					if len(q.Options) > 0 {
						// Options are displayed in order by Claude Code,
						// plus an automatic "Other" option at the end
						opts := make([]Option, len(q.Options))
						for i, o := range q.Options {
							label := o.Label
							if o.Description != "" {
								label += " — " + o.Description
							}
							opts[i] = Option{
								Key:   fmt.Sprintf("%d", i+1),
								Label: label,
							}
						}
						return PromptInfo{
							Type:         Choice,
							QuestionText: q.Question,
							Options:      opts,
							ToolUseID:    c.ID,
						}
					}
					return PromptInfo{
						Type:         Question,
						QuestionText: q.Question,
						ToolUseID:    c.ID,
					}
				}
			}

			// Any other tool is a permission prompt.
			// Options will be scraped from the screen in RecognizeClaude.
			// Fallback options here in case scraping fails.
			summary := summarizeToolInput(c.Name, c.Input)
			return PromptInfo{
				Type:        Permission,
				ToolName:    c.Name,
				ToolSummary: summary,
				ToolUseID:   c.ID,
				Options: []Option{
					{Key: "1", Label: "Yes"},
					{Key: "2", Label: "Yes, always"},
					{Key: "3", Label: "No"},
				},
			}
		}
	}

	return PromptInfo{Type: Unknown}
}

// summarizeToolInput extracts key parameters from tool input for display.
func summarizeToolInput(toolName string, rawInput json.RawMessage) string {
	var input map[string]interface{}
	if err := json.Unmarshal(rawInput, &input); err != nil {
		return ""
	}

	switch toolName {
	case "Read":
		if fp, ok := input["file_path"].(string); ok {
			return fp
		}
	case "Edit", "Write":
		if fp, ok := input["file_path"].(string); ok {
			return fp
		}
	case "Bash":
		if cmd, ok := input["command"].(string); ok {
			if len(cmd) > 80 {
				cmd = cmd[:77] + "..."
			}
			return cmd
		}
	case "Glob":
		if p, ok := input["pattern"].(string); ok {
			return p
		}
	case "Grep":
		if p, ok := input["pattern"].(string); ok {
			return p
		}
	}
	return ""
}

// RecognizeClaude performs the full PID chain lookup and prompt parsing for a Claude pane.
// Uses the transcript for prompt type detection and tool info, but scrapes the actual
// menu options from the screen — this is version-resistant since we read what Claude
// actually renders rather than guessing the menu structure.
func RecognizeClaude(paneID string) PromptInfo {
	// Step 1: Get shell PID for this pane
	shellPID, err := tmux.GetPanePID(paneID)
	if err != nil {
		return PromptInfo{Type: FreeInput}
	}

	// Step 2: Find claude child process
	claudePID, err := tmux.GetChildPID(shellPID, "claude")
	if err != nil {
		return PromptInfo{Type: FreeInput}
	}

	// Step 3: Read session file
	session, err := ReadClaudeSession(claudePID)
	if err != nil {
		return PromptInfo{Type: FreeInput}
	}

	// Step 4: Find transcript
	transcriptPath, err := FindTranscript(session.SessionID, session.Cwd)
	if err != nil {
		return PromptInfo{Type: FreeInput}
	}

	// Step 5: Read last assistant message
	msg, err := ReadLastAssistant(transcriptPath)
	if err != nil {
		return PromptInfo{Type: FreeInput}
	}

	// Step 6: Parse transcript into PromptInfo (gives us type + tool info)
	info := ParsePrompt(msg)

	// Step 7: For Permission and Choice prompts, scrape actual menu options
	// from the screen instead of using hardcoded/transcript options.
	// This makes us resilient to Claude Code UI changes.
	if info.Type == Permission || info.Type == Choice {
		content, err := tmux.CapturePane(paneID)
		if err == nil {
			question, scraped := ScrapeMenuOptions(content)
			if len(scraped) > 0 {
				info.Options = scraped
				if question != "" && info.QuestionText == "" {
					info.QuestionText = question
				}
			}
		}
	}

	return info
}
