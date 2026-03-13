package monitor

// spinnerFrames contains the braille dot spinner animation frames.
var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// PaneStatus represents the activity state of a pane.
type PaneStatus int

const (
	// Idle means the pane content has been stable with no prompt detected.
	Idle PaneStatus = iota
	// Running means the pane content is actively changing.
	Running
	// Ready means the pane is waiting for user input (prompt detected).
	Ready
)

// String returns a human-readable status name.
func (s PaneStatus) String() string {
	switch s {
	case Running:
		return "Running"
	case Ready:
		return "Ready"
	case Idle:
		return "Idle"
	default:
		return "Unknown"
	}
}

// Indicator returns a single-character status indicator for display.
func (s PaneStatus) Indicator() string {
	switch s {
	case Running:
		return "●"
	case Ready:
		return "?"
	case Idle:
		return "○"
	default:
		return " "
	}
}

// IndicatorAnimated returns a status indicator, with animation for Running state.
// The frame parameter controls which animation frame to display (cycles through spinner).
func (s PaneStatus) IndicatorAnimated(frame int) string {
	switch s {
	case Running:
		return spinnerFrames[frame%len(spinnerFrames)]
	case Ready:
		return "●"
	case Idle:
		return "○"
	default:
		return " "
	}
}
