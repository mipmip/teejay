package prompt

// PromptType classifies what a waiting agent is asking for.
type PromptType int

const (
	Unknown    PromptType = iota
	FreeInput             // Agent idle at main prompt
	Permission            // Tool permission prompt (y/n/a)
	Question              // Free-text question (no preset options)
	Choice                // Multiple choice with options
)

// String returns a human-readable name for the prompt type.
func (t PromptType) String() string {
	switch t {
	case FreeInput:
		return "FreeInput"
	case Permission:
		return "Permission"
	case Question:
		return "Question"
	case Choice:
		return "Choice"
	default:
		return "Unknown"
	}
}

// IsActionable returns true if this prompt type represents a specific question
// that the user should respond to (not just idle).
func (t PromptType) IsActionable() bool {
	return t == Permission || t == Question || t == Choice
}

// Option represents a selectable choice in a prompt.
type Option struct {
	Key   string // The keystroke to send (e.g., "y", "1", "a")
	Label string // Display text (e.g., "Allow once", "Fix the bug")
}

// PromptInfo holds the parsed prompt state of a waiting pane.
type PromptInfo struct {
	Type         PromptType
	QuestionText string   // The question or context text
	Options      []Option // Selectable options (empty for FreeInput/Question)
	ToolName     string   // Tool name for Permission prompts
	ToolSummary  string   // Brief summary of tool input parameters
	ToolUseID    string   // Claude's tool_use ID for freshness checking
}
