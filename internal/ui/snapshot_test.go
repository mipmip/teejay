package ui

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"tj/internal/monitor"
)

var update = flag.Bool("update", false, "update golden files")

func init() {
	// Force color output in tests so ANSI codes are captured in golden files
	lipgloss.SetColorProfile(termenv.ANSI256)
}

// goldenPath returns the path to a golden file in testdata/.
func goldenPath(name string) string {
	return filepath.Join("testdata", name+".golden")
}

// updateGolden writes the actual output to the golden file.
func updateGolden(t *testing.T, name string, actual []byte) {
	t.Helper()
	path := goldenPath(name)
	if err := os.WriteFile(path, actual, 0644); err != nil {
		t.Fatalf("failed to update golden file %s: %v", path, err)
	}
}

// assertGolden compares actual output against the golden file.
// If -update is set, overwrites the golden file instead.
func assertGolden(t *testing.T, name string, actual []byte) {
	t.Helper()
	if *update {
		updateGolden(t, name, actual)
		return
	}
	path := goldenPath(name)
	expected, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("golden file %s not found (run with -update to create): %v", path, err)
	}
	if !bytes.Equal(expected, actual) {
		t.Errorf("output does not match golden file %s\n--- expected ---\n%q\n--- actual ---\n%q", path, expected, actual)
	}
}

// renderDelegate renders a pane item using the browserItemDelegate and returns the output.
func renderDelegate(t *testing.T, item paneItem, selectedIndex int, itemIndex int) []byte {
	t.Helper()
	items := []list.Item{item}
	l := list.New(items, browserItemDelegate{}, 40, 10)
	if selectedIndex == 0 {
		l.Select(0)
	}

	var buf bytes.Buffer
	delegate := browserItemDelegate{}
	delegate.Render(&buf, l, itemIndex, item)
	return buf.Bytes()
}

func TestDelegateRenderUnselected(t *testing.T) {
	// Busy pane, not selected, with breadcrumb, no alert overrides
	item := paneItem{
		id:         "%1",
		name:       "my-project",
		status:     monitor.Busy,
		frame:      0,
		command:    "claude",
		session:    "work",
		windowName: "code",
	}
	actual := renderDelegate(t, item, -1, 0)
	assertGolden(t, "delegate_unselected", actual)
}

func TestDelegateRenderSelected(t *testing.T) {
	item := paneItem{
		id:         "%1",
		name:       "my-project",
		status:     monitor.Busy,
		frame:      0,
		command:    "claude",
		session:    "work",
		windowName: "code",
	}
	actual := renderDelegate(t, item, 0, 0)
	assertGolden(t, "delegate_selected", actual)
}

func TestDelegateRenderWaiting(t *testing.T) {
	item := paneItem{
		id:         "%2",
		name:       "assistant",
		status:     monitor.Waiting,
		frame:      0,
		command:    "claude",
		session:    "dev",
		windowName: "ai",
	}
	actual := renderDelegate(t, item, -1, 0)
	assertGolden(t, "delegate_waiting", actual)
}

func TestDelegateRenderAlerts(t *testing.T) {
	soundOn := true
	notifyOff := false
	item := paneItem{
		id:             "%3",
		name:           "alerts-pane",
		status:         monitor.Busy,
		frame:          0,
		command:        "aider",
		session:        "main",
		windowName:     "dev",
		soundOverride:  &soundOn,
		notifyOverride: &notifyOff,
	}
	actual := renderDelegate(t, item, -1, 0)
	assertGolden(t, "delegate_alerts", actual)
}

func TestDelegateRenderBoldTitle(t *testing.T) {
	item := paneItem{
		id:         "%1",
		name:       "my-project",
		status:     monitor.Busy,
		frame:      0,
		command:    "claude",
		session:    "work",
		windowName: "code",
	}
	actual := renderDelegate(t, item, 0, 0)
	// Verify the output contains bold ANSI escape code (may be combined with other codes, e.g. \x1b[1;48;5;59m)
	if !strings.Contains(string(actual), "\x1b[1;") {
		t.Error("rendered title should contain bold ANSI escape code (\\x1b[1;...)")
	}
}
