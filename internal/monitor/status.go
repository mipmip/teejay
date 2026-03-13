package monitor

// spinnerFrames contains the braille dot spinner animation frames.
var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// PaneStatus represents the activity state of a pane.
type PaneStatus int

const (
	// Busy means the pane is active or has no detected prompt.
	Busy PaneStatus = iota
	// Waiting means the pane is waiting for user input (prompt detected).
	Waiting
)

// String returns a human-readable status name.
func (s PaneStatus) String() string {
	switch s {
	case Busy:
		return "Busy"
	case Waiting:
		return "Waiting"
	default:
		return "Unknown"
	}
}

// Indicator returns a single-character status indicator for display.
func (s PaneStatus) Indicator() string {
	switch s {
	case Busy:
		return spinnerFrames[0] // First spinner frame for static display
	case Waiting:
		return "●"
	default:
		return " "
	}
}

// IndicatorAnimated returns a status indicator, with animation for Busy state.
// The frame parameter controls which animation frame to display (cycles through spinner).
func (s PaneStatus) IndicatorAnimated(frame int) string {
	switch s {
	case Busy:
		return spinnerFrames[frame%len(spinnerFrames)]
	case Waiting:
		return "●"
	default:
		return " "
	}
}
