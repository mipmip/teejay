package ui

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"tj/internal/alerts"
	"tj/internal/alerts/sounds"
	"tj/internal/config"
	"tj/internal/monitor"
	"tj/internal/naming"
	"tj/internal/prompt"
	"tj/internal/scan"
	"tj/internal/tmux"
	"tj/internal/watchlist"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262"))

	emptyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Italic(true)

	listPanelStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4"))

	previewPanelStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#626262"))

	previewTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#7D56F4")).
				MarginBottom(1)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Italic(true)

	browserPopupStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#7D56F4")).
				Padding(1, 2)

	browserTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#7D56F4")).
				MarginBottom(1)

	readyIndicatorStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#00FF00"))

	browserItemStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#333333")).
				Padding(1, 1, 1, 2). // top, right, bottom, left
				Margin(0, 1)         // vertical, horizontal

	browserItemSelectedStyle = lipgloss.NewStyle().
					Background(lipgloss.Color("#555555")).
					Padding(1, 1, 1, 2). // top, right, bottom, left
					Margin(0, 1)         // vertical, horizontal

	itemTitleStyle = lipgloss.NewStyle().
			Bold(true)

	// Branding footer styles
	brandingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#39FF14")). // Neon green
			Bold(true)

	versionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")) // Muted gray

	// Alert indicator styles
	soundEnabledStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#00FF00")) // Green
	notifyEnabledStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFD700")) // Yellow
	alertDisabledStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#555555")) // Dim gray
)

// recencyColor returns a green color that fades with elapsed time.
func recencyColor(elapsed time.Duration) lipgloss.Color {
	switch {
	case elapsed < 10*time.Second:
		return lipgloss.Color("#00FF00") // Bright neon green
	case elapsed < 30*time.Second:
		return lipgloss.Color("#00DD00") // Bright green
	case elapsed < 2*time.Minute:
		return lipgloss.Color("#00BB00") // Medium green
	case elapsed < 5*time.Minute:
		return lipgloss.Color("#009900") // Dim green
	default:
		return lipgloss.Color("#006600") // Very dim green
	}
}

// browserItemDelegate implements list.ItemDelegate for styled browser items
type browserItemDelegate struct{}

func (d browserItemDelegate) Height() int  { return 5 }
func (d browserItemDelegate) Spacing() int { return 0 }
func (d browserItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d browserItemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(list.DefaultItem)
	if !ok {
		return
	}

	title := i.Title()
	desc := i.Description()

	width := m.Width()
	contentWidth := width - 2 // subtract left and right margin

	bgColor := lipgloss.Color("#333333")
	if index == m.Index() {
		bgColor = lipgloss.Color("#555555")
	}

	// Two-column layout: left (auto-width) + right (fixed width).
	// Each cell is independently styled — no nested ANSI, no "bar" artifacts.
	const rightColWidth = 7 // indicator / "♪ ✉" + padding
	leftColWidth := contentWidth - rightColWidth

	blankLine := lipgloss.NewStyle().Background(bgColor).Width(contentWidth).Render("")

	leftBase := lipgloss.NewStyle().Background(bgColor).PaddingLeft(2).Width(leftColWidth)
	rightBase := lipgloss.NewStyle().Background(bgColor).Width(rightColWidth).PaddingLeft(1).PaddingRight(1).Align(lipgloss.Right)

	maxTextWidth := leftColWidth - 2 // subtract PaddingLeft
	titleLeft := leftBase.Bold(true).Render(truncateWithEllipsis(title, maxTextWidth))
	descLeft := leftBase.Bold(false).Render(desc)
	indicatorRight := rightBase.Render(" ")
	symbolsRight := rightBase.Render(" ")

	if p, ok := item.(paneItem); ok {
		// Right column: indicator with color
		indicatorText := p.status.IndicatorAnimated(p.frame)
		indicatorStyle := rightBase
		if p.status == monitor.Waiting {
			if p.promptInfo.Type.IsActionable() {
				indicatorText = "?"
				indicatorStyle = indicatorStyle.Foreground(lipgloss.Color("#FFD700")) // Yellow
			} else if p.recencyColorOn && !p.lastActivity.IsZero() {
				indicatorStyle = indicatorStyle.Foreground(recencyColor(time.Since(p.lastActivity)))
			} else {
				indicatorStyle = indicatorStyle.Foreground(lipgloss.Color("#00FF00"))
			}
		}
		indicatorRight = indicatorStyle.Render(indicatorText)

		// Right column: alert symbols
		if p.soundOverride != nil || p.notifyOverride != nil {
			symbolsRight = rightBase.Render("♪ ✉ ")
		}

		// Left column: plain text description, truncated with ellipsis if too long
		breadcrumb := p.session + " > " + p.windowName
		if p.command != "" {
			breadcrumb += " : " + p.command
		}
		maxWidth := leftColWidth - 2 // subtract PaddingLeft
		breadcrumb = truncateWithEllipsis(breadcrumb, maxWidth)
		descLeft = leftBase.Bold(false).Render(breadcrumb)
	}

	// Join columns per row, then stack rows
	titleRow := lipgloss.JoinHorizontal(lipgloss.Top, titleLeft, indicatorRight)
	descRow := lipgloss.JoinHorizontal(lipgloss.Top, descLeft, symbolsRight)

	output := " " + blankLine + "\n" +
		" " + titleRow + "\n" +
		" " + descRow + "\n" +
		" " + blankLine
	fmt.Fprintln(w, output)
}

// truncateWithEllipsis truncates text to maxWidth, replacing the end with "…" if needed.
func truncateWithEllipsis(text string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}
	w := lipgloss.Width(text)
	if w <= maxWidth {
		return text
	}
	runes := []rune(text)
	for len(runes) > 0 {
		runes = runes[:len(runes)-1]
		if lipgloss.Width(string(runes)+"…") <= maxWidth {
			return string(runes) + "…"
		}
	}
	return "…"
}

// Layout modes
const (
	layoutDefault     = 0
	layoutMultiColumn = 1
)

// previewTickMsg is sent periodically to trigger preview refresh
type previewTickMsg struct{}

// promptCheckResultMsg carries the results of an async prompt recognition check.
type promptCheckResultMsg struct {
	results map[string]prompt.PromptInfo
}

// scanResultMsg carries the result of an async startup scan.
type scanResultMsg struct {
	result scan.ScanResult
	err    error
}

// dismissTemporaryMsg is sent to auto-dismiss temporary messages after timeout
type dismissTemporaryMsg struct{}

// tickCmd returns a command that sends a previewTickMsg after 100ms
func tickCmd() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return previewTickMsg{}
	})
}

// dismissTemporaryCmd returns a command that sends a dismissTemporaryMsg after 3 seconds
func dismissTemporaryCmd() tea.Cmd {
	return tea.Tick(3*time.Second, func(t time.Time) tea.Msg {
		return dismissTemporaryMsg{}
	})
}

// showTemporaryMessage sets a temporary message and returns the dismiss command
func (m *Model) showTemporaryMessage(msg string) tea.Cmd {
	m.temporaryMessage = msg
	return dismissTemporaryCmd()
}

// paneItem implements list.Item for watchlist panes
type paneItem struct {
	id             string
	name           string
	addedAt        time.Time
	status         monitor.PaneStatus
	frame          int
	command        string // current foreground process
	session        string // tmux session name
	windowName     string // tmux window name
	soundOverride  *bool  // per-pane sound override (nil = inheriting global)
	notifyOverride *bool  // per-pane notification override (nil = inheriting global)
	promptInfo       prompt.PromptInfo // parsed prompt state when waiting
	lastActivity     time.Time         // last content change time from monitor
	recencyColorOn   bool              // whether to use recency-based color gradient
}

func (p paneItem) Title() string {
	displayName := p.id
	if p.name != "" {
		displayName = p.name
	}
	return displayName
}

func (p paneItem) Indicator() string {
	indicator := p.status.IndicatorAnimated(p.frame)
	if p.status == monitor.Waiting {
		indicator = readyIndicatorStyle.Render(indicator)
	}
	return indicator
}
func (p paneItem) Description() string {
	breadcrumb := p.session + " > " + p.windowName
	if p.command != "" {
		breadcrumb += " : " + p.command
	}
	if p.soundOverride != nil || p.notifyOverride != nil {
		soundOn := p.soundOverride != nil && *p.soundOverride
		notifyOn := p.notifyOverride != nil && *p.notifyOverride
		breadcrumb += "  " + renderAlertIndicators(soundOn, notifyOn)
	}
	return breadcrumb
}
func (p paneItem) FilterValue() string { return p.id }

// sessionItem implements list.Item for session browser selection
type sessionItem struct {
	name      string
	paneCount int
}

func (s sessionItem) Title() string       { return s.name }
func (s sessionItem) Description() string { return fmt.Sprintf("%d pane(s)", s.paneCount) }
func (s sessionItem) FilterValue() string { return s.name }

// browserItem implements list.Item for pane browser selection
type browserItem struct {
	paneInfo tmux.PaneInfo
}

func (b browserItem) Title() string {
	return fmt.Sprintf("%d.%d %s", b.paneInfo.Window, b.paneInfo.Pane, b.paneInfo.Command)
}
func (b browserItem) Description() string {
	return b.paneInfo.ID
}
func (b browserItem) FilterValue() string { return b.paneInfo.ID }

// configMenuItem represents a menu item in the configure popup
type configMenuItem int

const (
	configMenuName configMenuItem = iota
	configMenuSound
	configMenuSoundType
	configMenuNotify
)

// cycleTriState cycles through: nil (default) → true → false → nil
func cycleTriState(current *bool) *bool {
	if current == nil {
		t := true
		return &t
	}
	if *current {
		f := false
		return &f
	}
	return nil
}

// triStateIndicator returns a display string for a tri-state setting.
// [D:x] = using default (shown as x or blank), [x] = explicitly enabled, [ ] = explicitly disabled
func triStateIndicator(override *bool, defaultVal bool) string {
	if override == nil {
		if defaultVal {
			return "[D:x]"
		}
		return "[D: ]"
	}
	if *override {
		return "[x]"
	}
	return "[ ]"
}

type Model struct {
	list           list.Model
	viewport       viewport.Model
	textInput      textinput.Model
	watchlist      *watchlist.Watchlist
	monitor        *monitor.Monitor
	config         *config.Config
	paneStatuses   map[string]monitor.PaneStatus
	paneCommands   map[string]string     // current foreground command per pane
	paneSessions   map[string]string     // cached tmux session name per pane
	paneWindows    map[string]string     // cached tmux window name per pane
	empty          bool
	loadErr        error
	selectedPaneID string
	previewContent string
	previewErr     error
	width          int
	height         int
	panelHeight    int
	editing          bool
	deleting         bool
	temporaryMessage string // auto-dismissing error/status message
	browsing         bool
	browserList     list.Model
	browserEmpty    bool
	tickFrame       int
	browsingSession       bool              // true when showing sessions, false when showing panes
	selectedSession       string            // session name selected for pane browsing
	allBrowserPanes       []tmux.PaneInfo   // cached panes for filtering by session
	browserPreviewContent string            // preview content for selected browser pane
	browserPreviewErr     error             // error from capturing browser preview
	configuring     bool              // true when configure popup is open
	configMenuItem  configMenuItem    // selected menu item in configure popup
	configEditingName bool            // true when editing name in configure popup
	watchlistMtime  time.Time         // last known modification time of watchlist file
	statusMessage    string            // temporary status message to display to user
	version          string            // app version for footer display
	brandingShimmer  int               // shimmer animation frame (0 = not animating)
	lastActivePanes  map[string]bool          // previously active pane IDs
	paneFocusLostAt  map[string]time.Time     // when each pane lost focus (for grace period)
	layoutMode       int                      // 0 = default (list+preview), 1 = multi-column
	panePrompts      map[string]prompt.PromptInfo // cached prompt info per pane
	promptCheckTick  int                      // counter for periodic prompt checking
	sortByActivity   bool                     // sort panes by last activity (most recent first)
	pickerMode       bool                     // Enter switches pane and quits
	deletingAll      bool                     // true when "delete all" confirmation is shown
	filtering        bool                     // true when filter input is active
	filterQuery      string                   // current filter query (persists after confirm)
	filterInput      textinput.Model          // text input for filter
	quickAnswering     bool              // true when quick-answer popup is open
	quickAnswerPane    string            // pane ID being answered
	quickAnswerPrompt  prompt.PromptInfo // prompt being answered
	quickAnswerSelected int             // selected option index (for Permission/Choice)
	quickAnswerInput   textinput.Model  // text input for Question/FreeInput
}

// New creates a new Model with the given version, config, and optional watchlist path.
// If cfg is nil, config is loaded from the default path.
// If watchlistPath is empty, the default watchlist path is used.
func New(version string, cfg *config.Config, watchlistPath string) Model {
	// Load config if not provided
	if cfg == nil {
		cfg = config.Load()
	}

	// Load watchlist from custom path if provided
	var wl *watchlist.Watchlist
	var err error
	if watchlistPath != "" {
		wl, err = watchlist.Load(watchlistPath)
	} else {
		wl, err = watchlist.Load()
	}
	if err != nil {
		// Still need to initialize list/viewport to avoid nil panics on WindowSizeMsg
		l := list.New([]list.Item{}, browserItemDelegate{}, 30, 20)
		l.Title = "Watched Panes"
		l.SetShowStatusBar(false)
		l.SetFilteringEnabled(false)
		l.SetShowHelp(false)
		l.KeyMap.Quit.SetEnabled(false)
		vp := viewport.New(50, 20)
		return Model{loadErr: err, version: version, config: cfg, list: l, viewport: vp}
	}

	// Get initial mtime for watchlist file
	var wlMtime time.Time
	if watchlistPath != "" {
		if info, err := os.Stat(watchlistPath); err == nil {
			wlMtime = info.ModTime()
		}
	} else if path, err := watchlist.ConfigPath(); err == nil {
		if info, err := os.Stat(path); err == nil {
			wlMtime = info.ModTime()
		}
	}

	items := make([]list.Item, len(wl.Panes))
	for i, p := range wl.Panes {
		items[i] = paneItem{id: p.ID, name: p.Name, addedAt: p.AddedAt, status: monitor.Busy}
	}

	l := list.New(items, browserItemDelegate{}, 30, 20)
	l.Title = "Watched Panes"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.KeyMap.Quit.SetEnabled(false)
	l.SetShowPagination(true)
	l.Styles.Title = titleStyle

	vp := viewport.New(50, 20)

	ti := textinput.New()
	ti.Placeholder = "Enter name..."
	ti.CharLimit = 50

	fi := textinput.New()
	fi.Placeholder = "Filter..."
	fi.CharLimit = 100

	m := Model{
		list:           l,
		viewport:       vp,
		textInput:      ti,
		filterInput:    fi,
		watchlist:      wl,
		monitor:        monitor.New(cfg),
		config:         cfg,
		paneStatuses:   make(map[string]monitor.PaneStatus),
		paneCommands:    make(map[string]string),
		paneSessions:   make(map[string]string),
		paneWindows:    make(map[string]string),
		lastActivePanes:  make(map[string]bool),
		paneFocusLostAt: make(map[string]time.Time),
		panePrompts:     make(map[string]prompt.PromptInfo),
		sortByActivity:  cfg.Display.SortByActivity,
		pickerMode:      cfg.Display.PickerMode,
		empty:           len(items) == 0,
		watchlistMtime:  wlMtime,
		version:         version,
	}

	// Initialize layout mode from config
	if cfg.Display.LayoutMode == "columns" {
		m.layoutMode = layoutMultiColumn
	}

	// Capture initial pane content if there are panes
	if len(items) > 0 {
		if item, ok := items[0].(paneItem); ok {
			m.selectedPaneID = item.id
			m.captureSelectedPane()
		}
	}

	return m
}

// isStalePaneError checks if the error indicates a pane no longer exists in tmux.
func isStalePaneError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "can't find pane")
}

func (m *Model) captureSelectedPane() {
	if m.selectedPaneID == "" {
		m.previewContent = ""
		m.previewErr = nil
		return
	}

	content, err := tmux.CapturePane(m.selectedPaneID)
	if err != nil {
		// Check if pane no longer exists in tmux
		if isStalePaneError(err) {
			m.removeStalePane(m.selectedPaneID)
			return
		}
		m.previewErr = err
		m.previewContent = ""
	} else {
		m.previewErr = nil
		m.previewContent = content
	}
	m.viewport.SetContent(m.previewContent)
	m.viewport.GotoBottom()
}

// captureBrowserPreview captures pane content for the currently selected browser item.
func (m *Model) captureBrowserPreview() {
	if m.browsingSession || m.browserEmpty {
		m.browserPreviewContent = ""
		m.browserPreviewErr = nil
		return
	}

	item, ok := m.browserList.SelectedItem().(browserItem)
	if !ok {
		m.browserPreviewContent = ""
		m.browserPreviewErr = nil
		return
	}

	content, err := tmux.CapturePane(item.paneInfo.ID)
	if err != nil {
		m.browserPreviewErr = err
		m.browserPreviewContent = ""
	} else {
		m.browserPreviewErr = nil
		m.browserPreviewContent = content
	}
}

// removeStalePane removes a pane that no longer exists in tmux from the watchlist
func (m *Model) removeStalePane(paneID string) {
	m.statusMessage = fmt.Sprintf("Removed stale pane %s", paneID)
	m.watchlist.Remove(paneID)
	m.watchlist.Save()
	m.refreshList()

	// Update selection if the removed pane was selected
	if len(m.watchlist.Panes) == 0 {
		m.empty = true
		m.selectedPaneID = ""
		m.previewContent = ""
		m.previewErr = nil
	} else {
		// Select first available pane
		m.selectedPaneID = m.watchlist.Panes[0].ID
		m.list.Select(0)
	}
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	if !m.empty {
		cmds = append(cmds, tickCmd())
	}
	if m.config.Display.ScanOnStart {
		cfg := m.config
		wl := m.watchlist
		cmds = append(cmds, func() tea.Msg {
			allPanes, err := tmux.ListAllPanes()
			if err != nil {
				return scanResultMsg{err: err}
			}
			return scanResultMsg{result: scan.ScanAndAdd(wl, cfg, allPanes)}
		})
	}
	if len(cmds) == 0 {
		return nil
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Clear status message on any key press
	if _, ok := msg.(tea.KeyMsg); ok {
		m.statusMessage = ""
	}

	// Handle auto-dismiss of temporary messages
	if _, ok := msg.(dismissTemporaryMsg); ok {
		m.temporaryMessage = ""
		return m, nil
	}

	// Handle preview tick - refresh pane content periodically
	if _, ok := msg.(previewTickMsg); ok {
		m.tickFrame++

		// Update branding shimmer animation
		if m.brandingShimmer > 0 {
			m.brandingShimmer++
			if m.brandingShimmer > 25 { // Animation complete after ~2.5 seconds
				m.brandingShimmer = 0
			}
		} else if rand.Intn(200) == 0 { // ~0.5% chance per tick to start shimmer
			m.brandingShimmer = 1
		}

		// Skip refresh when in modal modes
		var promptCmd tea.Cmd
		if !m.editing && !m.deleting && !m.deletingAll && !m.browsing && !m.configuring {
			// Check if watchlist file has been modified externally
			m.checkWatchlistFileChange()

			if !m.empty {
				// Get all currently focused panes (may be multiple with multiple attached sessions)
				activePanes := tmux.GetActivePaneIDs()

				// Panes that were active last tick but are no longer → start grace period
				for id := range m.lastActivePanes {
					if !activePanes[id] {
						m.paneFocusLostAt[id] = time.Now()
					}
				}
				// Panes that gained focus → cancel any grace period
				for id := range activePanes {
					delete(m.paneFocusLostAt, id)
				}
				// Remember current active set for next tick
				m.lastActivePanes = activePanes

				// Clean up expired grace periods (> 2s) and reset monitor baseline
				// so user-made content changes during the pause don't trigger Busy
				for id, lostAt := range m.paneFocusLostAt {
					if time.Since(lostAt) > 2*time.Second {
						if content, err := tmux.CapturePane(id); err == nil {
							m.monitor.ResetBaseline(id, content)
						}
						delete(m.paneFocusLostAt, id)
					}
				}

				// Update status for ALL panes in the watchlist
				for _, p := range m.watchlist.Panes {
					content, err := tmux.CapturePane(p.ID)
					if err != nil {
						if isStalePaneError(err) {
							m.removeStalePane(p.ID)
							continue
						}
						continue
					}

					// Skip status monitoring for focused panes and panes in grace period
					// (still capture content for preview)
					isPaused := activePanes[p.ID]
					if !isPaused {
						if _, inGrace := m.paneFocusLostAt[p.ID]; inGrace {
							isPaused = true
						}
					}

					if !isPaused {
						// Get app name for this pane
						appName := m.paneCommands[p.ID]

						// Update status
						prevStatus := m.paneStatuses[p.ID]
						status := m.monitor.Update(p.ID, content, appName)
						if prevStatus != status {
							m.paneStatuses[p.ID] = status
							// Check for Busy -> Waiting transition and trigger alerts
							if prevStatus == monitor.Busy && status == monitor.Waiting {
								if !activePanes[p.ID] {
									m.triggerAlerts(p.ID)
								}
							}
						}
					}

					// Update preview content if this is the selected pane
					if p.ID == m.selectedPaneID {
						m.previewContent = content
						m.previewErr = nil
						m.viewport.SetContent(m.previewContent)
						m.viewport.GotoBottom()
					}
				}

				// Periodic prompt check (~2s = every 20 ticks) — async
				m.promptCheckTick++
				if m.promptCheckTick >= 20 {
					m.promptCheckTick = 0
					// Snapshot waiting panes and their app names for the goroutine
					type paneCheck struct {
						id      string
						appName string
					}
					var checks []paneCheck
					for _, p := range m.watchlist.Panes {
						if m.paneStatuses[p.ID] == monitor.Waiting {
							checks = append(checks, paneCheck{id: p.ID, appName: m.paneCommands[p.ID]})
						} else {
							delete(m.panePrompts, p.ID)
						}
					}
					if len(checks) > 0 {
						promptCmd = func() tea.Msg {
							results := make(map[string]prompt.PromptInfo, len(checks))
							for _, c := range checks {
								results[c.id] = prompt.Recognize(c.id, c.appName)
							}
							return promptCheckResultMsg{results: results}
						}
					}
				}

				// Always refresh list to update spinner animation
				m.refreshListWithFrame(m.tickFrame)
			}
		}
		if promptCmd != nil {
			return m, tea.Batch(tickCmd(), promptCmd)
		}
		return m, tickCmd()
	}

	// Handle scanResultMsg — apply async startup scan results
	if msg, ok := msg.(scanResultMsg); ok {
		if msg.err != nil {
			m.statusMessage = fmt.Sprintf("Scan failed: %v", msg.err)
		} else if msg.result.Found == 0 {
			m.statusMessage = "Scan: no agent panes found"
		} else {
			if msg.result.Skipped > 0 {
				m.statusMessage = fmt.Sprintf("Scan: added %d panes (%d already watched)", msg.result.Added, msg.result.Skipped)
			} else {
				m.statusMessage = fmt.Sprintf("Scan: added %d panes", msg.result.Added)
			}
			if msg.result.Added > 0 {
				m.watchlist.Save()
				m.refreshList()
				m.empty = len(m.watchlist.Panes) == 0
				if m.empty {
					return m, nil
				}
				return m, tickCmd()
			}
		}
		return m, nil
	}

	// Handle promptCheckResultMsg — apply async prompt recognition results
	if msg, ok := msg.(promptCheckResultMsg); ok {
		for id, info := range msg.results {
			m.panePrompts[id] = info
		}
		m.refreshListWithFrame(m.tickFrame)
		return m, nil
	}

	// Handle edit mode
	if m.editing {
		return m.updateEditing(msg)
	}

	// Handle filter mode
	if m.filtering {
		return m.updateFiltering(msg)
	}

	// Handle delete-all confirmation mode
	if m.deletingAll {
		return m.updateDeletingAll(msg)
	}

	// Handle delete confirmation mode
	if m.deleting {
		return m.updateDeleting(msg)
	}

	// Handle quick-answer mode
	if m.quickAnswering {
		return m.updateQuickAnswer(msg)
	}

	// Handle browsing mode
	if m.browsing {
		return m.updateBrowsing(msg)
	}

	// Handle configuring mode
	if m.configuring {
		return m.updateConfiguring(msg)
	}

	switch msg := msg.(type) {
	case tea.MouseMsg:
		// Handle mouse clicks in the main pane list
		if msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonLeft {
			// The list is rendered in the left panel with a 1-cell border
			// List panel starts at x=0, items start after title (around y=3 accounting for border and title)
			listWidth := m.width*30/100 - 2
			if listWidth < 25 {
				listWidth = m.width - 4 // full width when preview hidden
			}
			if listWidth < 20 {
				listWidth = 20
			}
			// Check if click is within the list panel area (accounting for border)
			if msg.X >= 1 && msg.X <= listWidth+1 {
				// Calculate which item was clicked
				// List starts rendering items at y=3 (border + title + spacing)
				// Each item takes 5 lines with custom delegate (top padding + title + description + bottom padding + spacing)
				itemHeight := 5
				headerOffset := 3 // border (1) + title (1) + spacing (1)

				if msg.Y >= headerOffset {
					clickedIndex := (msg.Y - headerOffset) / itemHeight
					items := m.list.Items()
					if clickedIndex >= 0 && clickedIndex < len(items) {
						m.list.Select(clickedIndex)
						// Update selection
						if item, ok := m.list.SelectedItem().(paneItem); ok {
							if item.id != m.selectedPaneID {
								m.selectedPaneID = item.id
								m.captureSelectedPane()
							}
						}
					}
				}
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "e":
			if !m.empty && m.selectedPaneID != "" {
				m.editing = true
				// Set text input to current display name
				if item, ok := m.list.SelectedItem().(paneItem); ok {
					m.textInput.SetValue(item.Title())
				}
				m.textInput.Focus()
				return m, textinput.Blink
			}
		case "d":
			if !m.empty && m.selectedPaneID != "" {
				m.deleting = true
				return m, nil
			}
		case "D":
			if !m.empty {
				m.deletingAll = true
				return m, nil
			}
		case "c":
			if !m.empty && m.selectedPaneID != "" {
				m.configuring = true
				m.configMenuItem = configMenuName
				m.configEditingName = false
				return m, nil
			}
		case "enter":
			if !m.empty && m.selectedPaneID != "" {
				if tmux.IsInsideTmux() {
					tmux.SwitchToPane(m.selectedPaneID)
					if m.pickerMode {
						return m, tea.Quit
					}
					return m, nil
				}
				return m, m.showTemporaryMessage("Cannot switch: not running inside tmux")
			}
		case "esc":
			// Clear active filter, or temporary messages
			if m.filterQuery != "" {
				m.filterQuery = ""
				m.refreshListWithFrame(m.tickFrame)
				return m, nil
			}
			if m.temporaryMessage != "" {
				m.temporaryMessage = ""
			}
			return m, nil
		case "/":
			// Enter filter mode
			m.filtering = true
			m.filterInput.SetValue(m.filterQuery)
			m.filterInput.Focus()
			return m, textinput.Blink
		case "p":
			// Toggle preview panel
			m.config.Display.ShowPreview = !m.config.Display.ShowPreview
			return m, nil
		case "o":
			// Toggle sort order
			m.sortByActivity = !m.sortByActivity
			m.refreshListWithFrame(m.tickFrame)
			return m, nil
		case "a":
			// Open pane browser
			m.loadBrowserPanes()
			m.browsing = true
			return m, nil
		case "v":
			// Toggle layout mode
			if m.layoutMode == layoutDefault {
				m.layoutMode = layoutMultiColumn
			} else {
				m.layoutMode = layoutDefault
			}
			return m, nil
		case " ":
			// Open quick-answer popup for waiting panes
			if !m.empty && m.selectedPaneID != "" && m.paneStatuses[m.selectedPaneID] == monitor.Waiting {
				appName := m.paneCommands[m.selectedPaneID]
				info := prompt.Recognize(m.selectedPaneID, appName)
				m.quickAnswering = true
				m.quickAnswerPane = m.selectedPaneID
				m.quickAnswerPrompt = info
				m.quickAnswerSelected = 0
				m.quickAnswerInput = textinput.New()
				m.quickAnswerInput.Focus()
				m.quickAnswerInput.Placeholder = "Type your response..."
				return m, textinput.Blink
			}
		case "s":
			// Scan for agent panes and auto-add
			allPanes, err := tmux.ListAllPanes()
			if err != nil {
				m.statusMessage = fmt.Sprintf("Scan failed: %v", err)
				return m, nil
			}
			result := scan.ScanAndAdd(m.watchlist, m.config, allPanes)
			if result.Found == 0 {
				m.statusMessage = "Scan: no agent panes found"
			} else {
				if result.Skipped > 0 {
					m.statusMessage = fmt.Sprintf("Scan: added %d panes (%d already watched)", result.Added, result.Skipped)
				} else {
					m.statusMessage = fmt.Sprintf("Scan: added %d panes", result.Added)
				}
				if result.Added > 0 {
					m.watchlist.Save()
					m.refreshList()
					m.empty = len(m.watchlist.Panes) == 0
				}
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.panelHeight = m.height - 4 // leave room for footer (help text may wrap on narrow terminals)

		if m.layoutMode == layoutMultiColumn {
			// Multi-column: list gets full width (rendering is manual)
			listWidth := m.width - 4
			m.list.SetWidth(listWidth)
			m.list.SetHeight(m.panelHeight - 5)
			// Set viewport width for below-preview (height set dynamically in render)
			m.viewport.Width = m.width - 6 // full width minus border + panel border
		} else {
			// Default: 30% list, 70% preview
			listWidth := m.width*30/100 - 2
			if listWidth < 25 {
				listWidth = m.width - 4 // full width when preview hidden
			}
			previewWidth := m.width*70/100 - 2
			m.list.SetWidth(listWidth)
			m.list.SetHeight(m.panelHeight - 5)
			m.viewport.Width = previewWidth
			m.viewport.Height = m.panelHeight - 4 // border (2) + title (2)
		}
	}

	// In multi-column mode, intercept arrow keys for spatial navigation
	if m.layoutMode == layoutMultiColumn {
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			items := m.list.Items()
			if len(items) > 0 {
				mc := calcMultiColumn(m.width-4, len(items))
				col, row := mc.colRow(m.list.Index())
				handled := false

				switch keyMsg.String() {
				case "up", "k":
					if row > 0 {
						m.list.Select(mc.flatIndex(col, row-1))
					}
					handled = true
				case "down", "j":
					newIdx := mc.flatIndex(col, row+1)
					if row+1 < mc.itemsPerCol && newIdx < len(items) {
						m.list.Select(newIdx)
					}
					handled = true
				case "left", "h":
					if col > 0 {
						newIdx := mc.flatIndex(col-1, row)
						if newIdx < len(items) {
							m.list.Select(newIdx)
						}
					}
					handled = true
				case "right", "l":
					if col+1 < mc.numColumns {
						newIdx := mc.flatIndex(col+1, row)
						if newIdx >= len(items) {
							// Jump to last item in that column
							lastInCol := mc.flatIndex(col+1, mc.itemsPerCol-1)
							for lastInCol >= len(items) {
								lastInCol--
							}
							if lastInCol > mc.flatIndex(col+1, 0)-1 {
								m.list.Select(lastInCol)
							}
						} else {
							m.list.Select(newIdx)
						}
					}
					handled = true
				}

				if handled {
					if item, ok := m.list.SelectedItem().(paneItem); ok {
						if item.id != m.selectedPaneID {
							m.selectedPaneID = item.id
							m.captureSelectedPane()
						}
					}
					return m, nil
				}
			}
		}
	}

	// Track previous selection
	prevSelected := m.selectedPaneID

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	// Check if selection changed
	if item, ok := m.list.SelectedItem().(paneItem); ok {
		if item.id != prevSelected {
			m.selectedPaneID = item.id
			m.captureSelectedPane()
		}
	}

	return m, cmd
}

func (m Model) updateEditing(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Save the new name
			newName := m.textInput.Value()
			m.watchlist.Rename(m.selectedPaneID, newName)
			m.watchlist.Save()
			m.refreshList()
			m.editing = false
			m.textInput.Blur()
			return m, nil
		case "esc":
			// Cancel edit
			m.editing = false
			m.textInput.Blur()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) updateQuickAnswer(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		hasOptions := len(m.quickAnswerPrompt.Options) > 0

		switch msg.String() {
		case "esc":
			m.quickAnswering = false
			return m, nil
		case "enter":
			// Determine what we'd send
			var response string
			if hasOptions {
				if m.quickAnswerSelected >= 0 && m.quickAnswerSelected < len(m.quickAnswerPrompt.Options) {
					response = m.quickAnswerPrompt.Options[m.quickAnswerSelected].Key
				}
			} else {
				response = m.quickAnswerInput.Value()
			}

			// Freshness check: re-verify pane is still waiting
			status := m.paneStatuses[m.quickAnswerPane]
			if status != monitor.Waiting {
				m.quickAnswering = false
				return m, m.showTemporaryMessage(fmt.Sprintf("Prompt expired — status=%v (not Waiting), would send %q", status, response))
			}

			// For Claude: re-check transcript for same tool_use ID
			if m.quickAnswerPrompt.ToolUseID != "" {
				appName := m.paneCommands[m.quickAnswerPane]
				freshInfo := prompt.Recognize(m.quickAnswerPane, appName)
				if freshInfo.ToolUseID != m.quickAnswerPrompt.ToolUseID {
					m.quickAnswering = false
					return m, m.showTemporaryMessage(fmt.Sprintf("Prompt expired — tool_use_id changed: %q → %q", m.quickAnswerPrompt.ToolUseID[:8], freshInfo.ToolUseID))
				}
			}

			// Send response using the appropriate method based on prompt type
			var err error
			switch m.quickAnswerPrompt.Type {
			case prompt.Permission, prompt.Choice:
				// Both permission and choice prompts are interactive lists in Claude Code —
				// navigate with arrow keys then press Enter to select
				err = tmux.SendArrowAndEnter(m.quickAnswerPane, m.quickAnswerSelected)
			default:
				// Free text: send the full line + Enter
				if response != "" {
					err = tmux.SendKeys(m.quickAnswerPane, response)
				}
			}
			if err != nil {
				m.quickAnswering = false
				return m, m.showTemporaryMessage(fmt.Sprintf("Send failed: %v", err))
			}
			m.quickAnswering = false
			m.statusMessage = fmt.Sprintf("Sent %q to %s", response, m.quickAnswerPane)
			return m, nil
		case "up", "k":
			if hasOptions {
				if m.quickAnswerSelected > 0 {
					m.quickAnswerSelected--
				}
				return m, nil
			}
		case "down", "j":
			if hasOptions {
				if m.quickAnswerSelected < len(m.quickAnswerPrompt.Options)-1 {
					m.quickAnswerSelected++
				}
				return m, nil
			}
		}

		// Forward to text input when no options (free text mode)
		if !hasOptions {
			var cmd tea.Cmd
			m.quickAnswerInput, cmd = m.quickAnswerInput.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m Model) updateFiltering(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Confirm filter, keep query active
			m.filtering = false
			m.filterQuery = m.filterInput.Value()
			m.filterInput.Blur()
			m.applyFilterToList()
			return m, nil
		case "esc":
			// Cancel and clear filter
			m.filtering = false
			m.filterQuery = ""
			m.filterInput.Blur()
			m.applyFilterToList()
			return m, nil
		}
	}

	// Forward all other keys to the text input
	var cmd tea.Cmd
	m.filterInput, cmd = m.filterInput.Update(msg)
	// Update filter query live as user types
	m.filterQuery = m.filterInput.Value()
	m.applyFilterToList()
	return m, cmd
}

// applyFilterToList re-filters the cached item list without re-fetching tmux data.
// Uses cached session/window/command data — no subprocess calls, instant response.
func (m *Model) applyFilterToList() {
	items := make([]list.Item, 0, len(m.watchlist.Panes))
	query := strings.ToLower(m.filterQuery)

	for _, p := range m.watchlist.Panes {
		item := paneItem{
			id: p.ID, name: p.Name, addedAt: p.AddedAt,
			status: m.paneStatuses[p.ID], frame: m.tickFrame,
			command: m.paneCommands[p.ID],
			session: m.paneSessions[p.ID], windowName: m.paneWindows[p.ID],
			soundOverride: p.SoundOnReady, notifyOverride: p.NotifyOnReady,
			promptInfo: m.panePrompts[p.ID],
			lastActivity: m.monitor.LastChangeTime(p.ID),
			recencyColorOn: m.config.Display.RecencyColor,
		}
		if query == "" || strings.Contains(strings.ToLower(item.Title()+" "+item.session+" "+item.windowName+" "+item.command), query) {
			items = append(items, item)
		}
	}

	m.list.SetItems(items)
}

func (m Model) updateDeletingAll(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "Y":
			m.watchlist.Panes = nil
			m.watchlist.Save()
			m.refreshList()
			m.deletingAll = false
			m.empty = true
			m.selectedPaneID = ""
			return m, nil
		case "n", "N", "esc":
			m.deletingAll = false
			return m, nil
		}
	}
	return m, nil
}

func (m Model) updateDeleting(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "Y":
			// Confirm delete
			m.watchlist.Remove(m.selectedPaneID)
			m.watchlist.Save()
			m.refreshList()
			m.deleting = false
			// Update selection after delete
			if len(m.watchlist.Panes) == 0 {
				m.empty = true
				m.selectedPaneID = ""
			} else if item, ok := m.list.SelectedItem().(paneItem); ok {
				m.selectedPaneID = item.id
				m.captureSelectedPane()
			}
			return m, nil
		case "n", "N", "esc":
			// Cancel delete
			m.deleting = false
			return m, nil
		}
	}
	return m, nil
}

func (m Model) updateBrowsing(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		// Handle mouse clicks in the browser popup list
		if msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonLeft {
			// The popup is centered, we need to calculate its position
			// Popup is approximately 54 wide (50 + padding) and 19 tall (15 + padding + border)
			popupWidth := 54
			popupHeight := 19
			popupX := (m.width - popupWidth) / 2
			popupY := (m.height - popupHeight) / 2

			// Check if click is within the popup
			if msg.X >= popupX && msg.X < popupX+popupWidth &&
				msg.Y >= popupY && msg.Y < popupY+popupHeight {
				// Calculate which item was clicked
				// Items start after border (1) + padding (1) + title (1) + spacing (1) = 4
				itemHeight := 5 // top padding + title + description + bottom padding + spacing
				headerOffset := 4

				relativeY := msg.Y - popupY
				if relativeY >= headerOffset {
					clickedIndex := (relativeY - headerOffset) / itemHeight
					items := m.browserList.Items()
					if clickedIndex >= 0 && clickedIndex < len(items) {
						m.browserList.Select(clickedIndex)
					}
				}
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if !m.browserEmpty {
				if m.browsingSession {
					// Session selected - show panes for this session
					if item, ok := m.browserList.SelectedItem().(sessionItem); ok {
						m.loadPaneListForSession(item.name)
					}
				} else {
					// Pane selected - add to watchlist with guessed name
					if item, ok := m.browserList.SelectedItem().(browserItem); ok {
						// Use naming package to guess a name (TUI doesn't prompt, user can rename later)
						guessedName, _ := naming.GuessName(item.paneInfo)
						m.watchlist.AddWithName(item.paneInfo.ID, guessedName)
						m.watchlist.Save()
						m.refreshList()
						m.empty = false
						// Select the newly added pane (it's at the end of the list)
						m.selectedPaneID = item.paneInfo.ID
						m.list.Select(len(m.watchlist.Panes) - 1)
						m.captureSelectedPane()
					}
					m.browsing = false
				}
			}
			return m, nil
		case "esc":
			if m.browsingSession {
				// At session level - close browser
				m.browsing = false
			} else {
				// At pane level - go back to session list
				m.loadSessionList()
			}
			return m, nil
		case "q":
			// Always close browser
			m.browsing = false
			return m, nil
		}
	}

	// Track selection before update to detect navigation changes
	prevIndex := m.browserList.Index()

	// Update browser list for navigation
	var cmd tea.Cmd
	m.browserList, cmd = m.browserList.Update(msg)

	// If selection changed and viewing panes, update preview
	if !m.browsingSession && m.browserList.Index() != prevIndex {
		m.captureBrowserPreview()
	}

	return m, cmd
}

func (m Model) updateConfiguring(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle name editing mode within configure popup
	if m.configEditingName {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				// Save the new name
				newName := m.textInput.Value()
				m.watchlist.Rename(m.selectedPaneID, newName)
				m.watchlist.Save()
				m.refreshList()
				m.configEditingName = false
				m.textInput.Blur()
				return m, nil
			case "esc":
				// Cancel name edit, stay in configure popup
				m.configEditingName = false
				m.textInput.Blur()
				return m, nil
			}
		}
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.MouseMsg:
		// Handle mouse clicks in the configure popup menu
		if msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonLeft {
			// The popup is centered, calculate its approximate position
			// Configure popup is approximately 40 wide and 10 tall
			popupWidth := 40
			popupHeight := 10
			popupX := (m.width - popupWidth) / 2
			popupY := (m.height - popupHeight) / 2

			// Check if click is within the popup
			if msg.X >= popupX && msg.X < popupX+popupWidth &&
				msg.Y >= popupY && msg.Y < popupY+popupHeight {
				// Calculate which menu item was clicked
				// Items: title (line 0), blank (line 1), Name (line 2), Sound (line 3), SoundType (line 4), Notify (line 5)
				// Accounting for border (1) + padding (1) = 2 offset
				relativeY := msg.Y - popupY - 2

				switch relativeY {
				case 2: // Name row
					m.configMenuItem = configMenuName
				case 3: // Sound row
					m.configMenuItem = configMenuSound
				case 4: // Sound Type row
					m.configMenuItem = configMenuSoundType
				case 5: // Notify row
					m.configMenuItem = configMenuNotify
				}
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.configMenuItem > configMenuName {
				m.configMenuItem--
			}
			return m, nil
		case "down", "j":
			if m.configMenuItem < configMenuNotify {
				m.configMenuItem++
			}
			return m, nil
		case "enter", " ":
			pane := m.watchlist.GetPane(m.selectedPaneID)
			if pane == nil {
				return m, nil
			}
			switch m.configMenuItem {
			case configMenuName:
				// Enter name editing mode
				m.configEditingName = true
				m.textInput.SetValue(pane.Name)
				m.textInput.Focus()
				return m, textinput.Blink
			case configMenuSound:
				// Cycle sound: default → enabled → disabled → default
				m.watchlist.SetSound(m.selectedPaneID, cycleTriState(pane.SoundOnReady))
				m.watchlist.Save()
				return m, nil
			case configMenuSoundType:
				// Cycle sound type: chime → bell → ping → pop → ding → chime
				currentType := pane.GetEffectiveSoundType(m.config)
				nextType := sounds.NextSound(currentType)
				m.watchlist.SetSoundType(m.selectedPaneID, &nextType)
				m.watchlist.Save()
				alerts.PlaySound(nextType)
				return m, nil
			case configMenuNotify:
				// Cycle notification: default → enabled → disabled → default
				m.watchlist.SetNotify(m.selectedPaneID, cycleTriState(pane.NotifyOnReady))
				m.watchlist.Save()
				return m, nil
			}
		case "esc", "q":
			m.configuring = false
			return m, nil
		}
	}
	return m, nil
}

// checkWatchlistFileChange checks if the watchlist file has been modified
// and reloads it if necessary, preserving the current selection when possible.
func (m *Model) checkWatchlistFileChange() {
	path, err := watchlist.ConfigPath()
	if err != nil {
		return
	}

	info, err := os.Stat(path)
	if err != nil {
		return
	}

	currentMtime := info.ModTime()
	if currentMtime.Equal(m.watchlistMtime) {
		return
	}

	// File has changed - reload watchlist
	newWl, err := watchlist.Load()
	if err != nil {
		return
	}

	// Store current selection
	prevSelectedID := m.selectedPaneID

	// Update model with new watchlist
	m.watchlist = newWl
	m.watchlistMtime = currentMtime
	m.refreshList()

	// Check if previous selection still exists
	if m.watchlist.Contains(prevSelectedID) {
		m.selectedPaneID = prevSelectedID
	} else if len(m.watchlist.Panes) > 0 {
		// Select first pane if previous selection was removed
		m.selectedPaneID = m.watchlist.Panes[0].ID
	} else {
		m.selectedPaneID = ""
		m.empty = true
	}

	// Update empty state
	m.empty = len(m.watchlist.Panes) == 0

	// Re-capture preview for selected pane
	if m.selectedPaneID != "" {
		m.captureSelectedPane()
	}
}

func (m *Model) triggerAlerts(paneID string) {
	pane := m.watchlist.GetPane(paneID)
	if pane == nil {
		return
	}

	if !m.config.Alerts.MuteSound && pane.GetEffectiveSound(m.config) {
		soundType := pane.GetEffectiveSoundType(m.config)
		alerts.PlaySound(soundType)
	}

	if !m.config.Alerts.MuteNotify && pane.GetEffectiveNotify(m.config) {
		displayName := pane.DisplayName()
		alerts.SendNotification("Teejay", displayName+" is ready")
	}
}

func (m *Model) refreshList() {
	m.refreshListWithFrame(m.tickFrame)
}

func (m *Model) refreshListWithFrame(frame int) {
	// Fetch current commands and pane info for all panes
	paneInfoMap := make(map[string]*tmux.PaneInfo)
	for _, p := range m.watchlist.Panes {
		if paneInfo, err := tmux.GetPaneByID(p.ID); err == nil && paneInfo != nil {
			m.paneCommands[p.ID] = paneInfo.Command
			m.paneSessions[p.ID] = paneInfo.Session
			m.paneWindows[p.ID] = paneInfo.WindowName
			paneInfoMap[p.ID] = paneInfo
		}
		// On error, keep last known command (graceful degradation)
	}

	items := make([]list.Item, len(m.watchlist.Panes))
	for i, p := range m.watchlist.Panes {
		status := m.paneStatuses[p.ID]
		command := m.paneCommands[p.ID]
		session := m.paneSessions[p.ID]
		windowName := m.paneWindows[p.ID]
		items[i] = paneItem{id: p.ID, name: p.Name, addedAt: p.AddedAt, status: status, frame: frame, command: command, session: session, windowName: windowName, soundOverride: p.SoundOnReady, notifyOverride: p.NotifyOnReady, promptInfo: m.panePrompts[p.ID], lastActivity: m.monitor.LastChangeTime(p.ID), recencyColorOn: m.config.Display.RecencyColor}
	}

	// Sort by activity if enabled: busy first (most recent), then waiting (most recent)
	if m.sortByActivity {
		sort.SliceStable(items, func(i, j int) bool {
			a := items[i].(paneItem)
			b := items[j].(paneItem)
			// Busy panes before waiting panes
			if a.status != b.status {
				return a.status == monitor.Busy
			}
			// Within same status, most recently active first
			return a.lastActivity.After(b.lastActivity)
		})
	}

	// Apply filter if active
	if m.filterQuery != "" {
		query := strings.ToLower(m.filterQuery)
		filtered := make([]list.Item, 0, len(items))
		for _, item := range items {
			p := item.(paneItem)
			searchText := strings.ToLower(p.Title() + " " + p.session + " " + p.windowName + " " + p.command)
			if strings.Contains(searchText, query) {
				filtered = append(filtered, item)
			}
		}
		items = filtered
	}

	m.list.SetItems(items)
	// empty reflects the actual watchlist, not the filtered view
	m.empty = len(m.watchlist.Panes) == 0
}

// browserListWidth calculates the appropriate list width for the browser popup.
// On narrow screens (< 80 cols) shows list only, on wider screens shows split with preview.
func (m *Model) browserListWidth() int {
	popupWidth := m.width * 90 / 100
	if popupWidth < 40 {
		popupWidth = m.width - 4
	}

	// On narrow screens, use full popup width (minus padding)
	if m.width < 80 {
		return popupWidth - 6 // account for border and padding
	}

	// On wider screens, use 35% for list panel
	listWidth := popupWidth * 35 / 100
	if listWidth < 30 {
		listWidth = 30
	}
	return listWidth - 6 // account for border and padding
}

func (m *Model) loadBrowserPanes() {
	allPanes, err := tmux.ListAllPanes()
	if err != nil {
		m.browserEmpty = true
		return
	}

	// Build set of already watched pane IDs
	watched := make(map[string]bool)
	for _, p := range m.watchlist.Panes {
		watched[p.ID] = true
	}

	// Filter out already watched panes and cache them
	m.allBrowserPanes = make([]tmux.PaneInfo, 0)
	for _, p := range allPanes {
		if !watched[p.ID] {
			m.allBrowserPanes = append(m.allBrowserPanes, p)
		}
	}

	// Start with session list
	m.browsingSession = true
	m.selectedSession = ""
	m.loadSessionList()
}

func (m *Model) loadSessionList() {
	// Count panes per session
	sessionPanes := make(map[string]int)
	for _, p := range m.allBrowserPanes {
		sessionPanes[p.Session]++
	}

	// Build session items (preserve order from first pane appearance)
	seen := make(map[string]bool)
	items := make([]list.Item, 0)
	for _, p := range m.allBrowserPanes {
		if !seen[p.Session] {
			seen[p.Session] = true
			items = append(items, sessionItem{
				name:      p.Session,
				paneCount: sessionPanes[p.Session],
			})
		}
	}

	// Create session list with custom styled delegate
	// Calculate list width based on screen size
	listWidth := m.browserListWidth()
	delegate := browserItemDelegate{}
	m.browserList = list.New(items, delegate, listWidth, 15)
	m.browserList.Title = "Select Session"
	m.browserList.SetShowStatusBar(false)
	m.browserList.SetFilteringEnabled(false)
	m.browserList.SetShowHelp(false)
	m.browserList.KeyMap.Quit.SetEnabled(false)
	m.browserList.Styles.Title = browserTitleStyle

	m.browserEmpty = len(items) == 0
}

func (m *Model) loadPaneListForSession(sessionName string) {
	// Filter panes for this session
	items := make([]list.Item, 0)
	for _, p := range m.allBrowserPanes {
		if p.Session == sessionName {
			items = append(items, browserItem{paneInfo: p})
		}
	}

	// Create pane list with custom styled delegate
	// Calculate list width based on screen size
	listWidth := m.browserListWidth()
	delegate := browserItemDelegate{}
	m.browserList = list.New(items, delegate, listWidth, 15)
	m.browserList.Title = "Select Pane (" + sessionName + ")"
	m.browserList.SetShowStatusBar(false)
	m.browserList.SetFilteringEnabled(false)
	m.browserList.SetShowHelp(false)
	m.browserList.KeyMap.Quit.SetEnabled(false)
	m.browserList.Styles.Title = browserTitleStyle

	m.browsingSession = false
	m.selectedSession = sessionName
	m.browserEmpty = len(items) == 0

	// Capture preview for first pane
	m.captureBrowserPreview()
}

func (m Model) View() string {
	if m.loadErr != nil {
		return fmt.Sprintf("Error loading watchlist: %v\n\nPress q to quit.\n", m.loadErr)
	}

	if m.empty && !m.browsing && m.filterQuery == "" && !m.filtering {
		return titleStyle.Render("Teejay") + "\n\n" +
			emptyStyle.Render("No panes are being watched.") + "\n\n" +
			helpStyle.Render("Press 'a' to browse and add panes, or run 'tj add' in a tmux pane.") + "\n\n" +
			helpStyle.Render("Press q to quit.")
	}

	// Handle quick-answer popup
	if m.quickAnswering {
		return m.renderQuickAnswerPopup()
	}

	// Handle browsing popup
	if m.browsing {
		return m.renderBrowserPopup()
	}

	// Handle configure popup
	if m.configuring {
		return m.renderConfigurePopup()
	}

	// Estimate footer height: the help text is ~140 chars and wraps on narrow terminals
	helpTextLen := 140 // approximate length of the longest help line
	footerLineCount := 2 + (helpTextLen / max(m.width, 1)) // base 2 lines + wrapping
	if footerLineCount < 2 {
		footerLineCount = 2
	}

	var layout string
	if m.layoutMode == layoutMultiColumn {
		// Multi-column mode: no preview, fill width with columns
		layout = m.renderMultiColumnLayout(footerLineCount)
	} else {
		// Default mode: list + preview
		listWidth := m.width*30/100 - 2
		hasItems := len(m.list.Items()) > 0
		showPreview := m.config.Display.ShowPreview && listWidth >= 25 && hasItems

		if showPreview {
			previewWidth := m.width*70/100 - 2
			if previewWidth < 20 {
				previewWidth = 20
			}

			// Build list panel
			listPanel := listPanelStyle.
				Width(listWidth).
				Height(m.panelHeight).
				Render(m.list.View())

			// Build preview panel
			var previewContent string
			if m.previewErr != nil {
				previewContent = errorStyle.Render(fmt.Sprintf("Error: %v", m.previewErr))
			} else if m.previewContent == "" {
				previewContent = emptyStyle.Render("No content")
			} else {
				previewContent = m.viewport.View()
			}

			previewName := m.selectedPaneID
			if item, ok := m.list.SelectedItem().(paneItem); ok {
				previewName = item.Title()
			}
			previewTitle := previewTitleStyle.Render("Preview: " + previewName)
			previewPanel := previewPanelStyle.
				Width(previewWidth).
				Height(m.panelHeight).
				Render(previewTitle + "\n" + previewContent)

			layout = lipgloss.JoinHorizontal(lipgloss.Top, listPanel, previewPanel)
		} else {
			// Narrow terminal: sidebar only, full width
			listWidth = m.width - 4
			if listWidth < 20 {
				listWidth = 20
			}
			listPanel := listPanelStyle.
				Width(listWidth).
				Render(m.list.View())
			layout = listPanel
		}
	}

	// Build footer first so we can measure its height for layout capping
	var footer string
	if m.filtering {
		footer = "/ " + m.filterInput.View() + "\n" + helpStyle.Render("Enter: confirm • Esc: clear")
	} else if m.editing {
		footer = "Rename: " + m.textInput.View() + "\n" + helpStyle.Render("Enter: save • Esc: cancel")
	} else if m.deletingAll {
		footer = errorStyle.Render(fmt.Sprintf("Delete all %d panes? (y/n)", len(m.watchlist.Panes)))
	} else if m.deleting {
		paneName := m.selectedPaneID
		if item, ok := m.list.SelectedItem().(paneItem); ok {
			paneName = item.Title()
		}
		footer = errorStyle.Render(fmt.Sprintf("Delete %s? (y/n)", paneName))
	} else if m.filterQuery != "" {
		footer = helpStyle.Render("Filter: "+m.filterQuery) + "\n" + helpStyle.Render("/ to edit • Esc to clear")
	} else if m.temporaryMessage != "" {
		footer = errorStyle.Render(m.temporaryMessage) + "\n" + helpStyle.Render("Press Esc to dismiss")
	} else if m.statusMessage != "" {
		footer = helpStyle.Render(m.statusMessage) + "\n" + helpStyle.Render("↑/↓: navigate • Enter: switch • /: filter • space: answer • v: view • p: preview • o: order • a: add • s: scan • c: configure • d/D: delete • q: quit")
	} else {
		footer = helpStyle.Render("↑/↓: navigate • Enter: switch • /: filter • space: answer • v: view • p: preview • o: order • a: add • s: scan • c: configure • d/D: delete • q: quit")
	}

	// Add branding to footer line if terminal is wide enough
	if m.width >= 80 {
		branding := m.renderBrandingFooter()
		footerWidth := lipgloss.Width(footer)
		brandingWidth := lipgloss.Width(branding)
		padding := m.width - footerWidth - brandingWidth
		if padding > 0 {
			footer = footer + strings.Repeat(" ", padding) + branding
		}
	}

	return layout + "\n" + footer
}

// multiColumnInfo holds layout calculations for the multi-column view.
type multiColumnInfo struct {
	numColumns    int
	itemsPerCol   int
	colWidth      int
	totalItems    int
}

// calcMultiColumn computes column layout from available width and item count.
func calcMultiColumn(availableWidth, totalItems int) multiColumnInfo {
	const minColWidth = 40
	numColumns := availableWidth / minColWidth
	if numColumns < 1 {
		numColumns = 1
	}
	// Don't create more columns than items
	if totalItems > 0 && numColumns > totalItems {
		numColumns = totalItems
	}
	colWidth := availableWidth / numColumns
	itemsPerCol := totalItems / numColumns
	if totalItems%numColumns != 0 {
		itemsPerCol++
	}
	return multiColumnInfo{
		numColumns:  numColumns,
		itemsPerCol: itemsPerCol,
		colWidth:    colWidth,
		totalItems:  totalItems,
	}
}

// multiColumnIndex converts (column, row) to flat list index.
func (mc multiColumnInfo) flatIndex(col, row int) int {
	return col*mc.itemsPerCol + row
}

// multiColumnPos converts a flat list index to (column, row).
func (mc multiColumnInfo) colRow(index int) (int, int) {
	if mc.itemsPerCol == 0 {
		return 0, 0
	}
	return index / mc.itemsPerCol, index % mc.itemsPerCol
}

const minPreviewBelowHeight = 8 // minimum lines to show the below-preview

// renderMultiColumnLayout renders all pane items in a multi-column grid,
// with an optional preview panel below when vertical space allows.
// footerLines is the estimated number of lines the footer will occupy.
func (m Model) renderMultiColumnLayout(footerLines int) string {
	items := m.list.Items()
	if len(items) == 0 {
		if m.filterQuery != "" {
			return emptyStyle.Render("No panes match filter: " + m.filterQuery)
		}
		return emptyStyle.Render("No panes are being watched.")
	}

	availableWidth := m.width - 4 // border space
	mc := calcMultiColumn(availableWidth, len(items))
	selectedIndex := m.list.Index()

	columns := make([]string, mc.numColumns)
	for col := 0; col < mc.numColumns; col++ {
		var rows []string
		for row := 0; row < mc.itemsPerCol; row++ {
			idx := mc.flatIndex(col, row)
			if idx >= len(items) {
				// Render empty space to keep columns aligned
				blankLine := lipgloss.NewStyle().Width(mc.colWidth - 2).Render("")
				rows = append(rows, blankLine+"\n"+blankLine+"\n"+blankLine+"\n"+blankLine)
				continue
			}

			item := items[idx]
			isSelected := idx == selectedIndex
			rows = append(rows, m.renderMultiColumnItem(item, mc.colWidth-2, isSelected))
		}
		columns[col] = lipgloss.JoinVertical(lipgloss.Left, rows...)
	}

	grid := lipgloss.JoinHorizontal(lipgloss.Top, columns...)
	gridHeight := lipgloss.Height(grid)

	// Reserve space for footer + the \n separator between layout and footer
	remainingHeight := m.height - gridHeight - footerLines - 1

	if m.config.Display.ShowPreview && remainingHeight >= minPreviewBelowHeight {
		// Render preview panel below the grid
		previewWidth := m.width - 4
		previewPanelHeight := remainingHeight - 2 // subtract panel border
		viewportHeight := previewPanelHeight - 2  // subtract title + margin

		if viewportHeight < 1 {
			viewportHeight = 1
		}

		var previewContent string
		if m.previewErr != nil {
			previewContent = errorStyle.Render(fmt.Sprintf("Error: %v", m.previewErr))
		} else if m.previewContent == "" {
			previewContent = emptyStyle.Render("No content")
		} else {
			m.viewport.Width = previewWidth - 2 // subtract border
			m.viewport.Height = viewportHeight
			previewContent = m.viewport.View()
		}

		previewName := m.selectedPaneID
		if item, ok := m.list.SelectedItem().(paneItem); ok {
			previewName = item.Title()
		}
		previewTitle := previewTitleStyle.Render("Preview: " + previewName)
		previewPanel := previewPanelStyle.
			Width(previewWidth).
			Height(previewPanelHeight).
			Render(previewTitle + "\n" + previewContent)

		return lipgloss.JoinVertical(lipgloss.Left, grid, previewPanel)
	}

	return grid
}

// renderMultiColumnItem renders a single pane item for the multi-column layout.
func (m Model) renderMultiColumnItem(item list.Item, contentWidth int, isSelected bool) string {
	i, ok := item.(paneItem)
	if !ok {
		return ""
	}

	bgColor := lipgloss.Color("#333333")
	if isSelected {
		bgColor = lipgloss.Color("#555555")
	}

	const rightColWidth = 7
	leftColWidth := contentWidth - rightColWidth

	blankLine := lipgloss.NewStyle().Background(bgColor).Width(contentWidth).Render("")
	leftBase := lipgloss.NewStyle().Background(bgColor).PaddingLeft(2).Width(leftColWidth)
	rightBase := lipgloss.NewStyle().Background(bgColor).Width(rightColWidth).PaddingLeft(1).PaddingRight(1).Align(lipgloss.Right)

	maxTextWidth := leftColWidth - 2
	titleLeft := leftBase.Bold(true).Render(truncateWithEllipsis(i.Title(), maxTextWidth))

	// Indicator
	indicatorText := i.status.IndicatorAnimated(i.frame)
	indicatorStyle := rightBase
	if i.status == monitor.Waiting {
		if i.promptInfo.Type.IsActionable() {
			indicatorText = "?"
			indicatorStyle = indicatorStyle.Foreground(lipgloss.Color("#FFD700"))
		} else if i.recencyColorOn && !i.lastActivity.IsZero() {
			indicatorStyle = indicatorStyle.Foreground(recencyColor(time.Since(i.lastActivity)))
		} else {
			indicatorStyle = indicatorStyle.Foreground(lipgloss.Color("#00FF00"))
		}
	}
	indicatorRight := indicatorStyle.Render(indicatorText)

	// Alert symbols
	symbolsRight := rightBase.Render(" ")
	if i.soundOverride != nil || i.notifyOverride != nil {
		symbolsRight = rightBase.Render("♪ ✉ ")
	}

	// Breadcrumb
	breadcrumb := i.session + " > " + i.windowName
	if i.command != "" {
		breadcrumb += " : " + i.command
	}
	breadcrumb = truncateWithEllipsis(breadcrumb, maxTextWidth)
	descLeft := leftBase.Bold(false).Render(breadcrumb)

	titleRow := lipgloss.JoinHorizontal(lipgloss.Top, titleLeft, indicatorRight)
	descRow := lipgloss.JoinHorizontal(lipgloss.Top, descLeft, symbolsRight)

	return blankLine + "\n" + titleRow + "\n" + descRow + "\n" + blankLine
}

// renderAlertIndicators returns styled ♪ ✉ symbols based on enabled state.
func renderAlertIndicators(soundEnabled, notifyEnabled bool) string {
	var sound, notify string
	if soundEnabled {
		sound = soundEnabledStyle.Render("♪")
	} else {
		sound = alertDisabledStyle.Render("♪")
	}
	if notifyEnabled {
		notify = notifyEnabledStyle.Render("✉")
	} else {
		notify = alertDisabledStyle.Render("✉")
	}
	return sound + " " + notify
}

// renderBrandingFooter returns the "Terminal Jockey" branding with version
func (m Model) renderBrandingFooter() string {
	text := "Terminal Jockey"
	var brand string

	if m.brandingShimmer > 0 {
		// Shimmer animation - gradient sweep across text
		baseColor := lipgloss.Color("#39FF14")    // Neon green
		shimmerColor := lipgloss.Color("#AFFFAF") // Bright mint green
		midColor := lipgloss.Color("#6FFF6F")     // Light green

		for i, ch := range text {
			// Calculate distance from shimmer position
			shimmerPos := float64(m.brandingShimmer-1) * 0.8
			dist := shimmerPos - float64(i)
			if dist < 0 {
				dist = -dist
			}

			var charStyle lipgloss.Style
			if dist < 2 {
				charStyle = lipgloss.NewStyle().Foreground(shimmerColor).Bold(true)
			} else if dist < 4 {
				charStyle = lipgloss.NewStyle().Foreground(midColor).Bold(true)
			} else {
				charStyle = lipgloss.NewStyle().Foreground(baseColor).Bold(true)
			}
			brand += charStyle.Render(string(ch))
		}
	} else {
		brand = brandingStyle.Render(text)
	}

	ver := versionStyle.Render(" " + m.version)
	alerts := " " + renderAlertIndicators(m.config.Alerts.SoundOnReady, m.config.Alerts.NotifyOnReady)
	return brand + ver + alerts
}

// renderQuickAnswerPopup renders the quick-answer popup for responding to agent prompts.
func (m Model) renderQuickAnswerPopup() string {
	p := m.quickAnswerPrompt
	popupWidth := m.width * 60 / 100
	if popupWidth < 40 {
		popupWidth = m.width - 4
	}

	// Title
	paneName := m.quickAnswerPane
	for _, wp := range m.watchlist.Panes {
		if wp.ID == m.quickAnswerPane && wp.Name != "" {
			paneName = wp.Name
			break
		}
	}
	title := browserTitleStyle.Render("Quick Answer: " + paneName)

	// Context/question
	var content string
	switch p.Type {
	case prompt.Permission:
		toolDesc := p.ToolName
		if p.ToolSummary != "" {
			summary := p.ToolSummary
			maxW := popupWidth - 10
			if len(summary) > maxW {
				summary = summary[:maxW-3] + "..."
			}
			toolDesc += " on " + summary
		}
		content = lipgloss.NewStyle().Bold(true).Render(toolDesc) + "\n\n"
	case prompt.Question, prompt.Choice:
		if p.QuestionText != "" {
			content = p.QuestionText + "\n\n"
		}
	case prompt.FreeInput:
		content = helpStyle.Render("Agent is waiting for input") + "\n\n"
	default:
		content = helpStyle.Render("Agent is waiting") + "\n\n"
	}

	// Options or text input
	hasOptions := len(p.Options) > 0
	if hasOptions {
		for i, opt := range p.Options {
			cursor := "  "
			style := lipgloss.NewStyle()
			if i == m.quickAnswerSelected {
				cursor = ">"
				style = style.Bold(true).Foreground(lipgloss.Color("#7D56F4"))
			}
			content += fmt.Sprintf(" %s %s  %s\n", cursor, opt.Key, style.Render(opt.Label))
		}
		content += "\n" + helpStyle.Render("↑/↓: select • Enter: send • Esc: cancel")
	} else {
		content += m.quickAnswerInput.View() + "\n\n"
		content += helpStyle.Render("Enter: send • Esc: cancel")
	}

	popup := browserPopupStyle.
		Width(popupWidth).
		Render(title + "\n" + content)

	// Center the popup
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, popup)
}

func (m Model) renderBrowserPopup() string {
	var popup string

	// Session list: single panel layout
	if m.browsingSession {
		// Calculate popup width - use most of screen but not more than available
		popupWidth := m.width * 80 / 100
		if popupWidth < 30 {
			popupWidth = m.width - 4
		}

		var content string
		if m.browserEmpty {
			content = emptyStyle.Render("No additional panes available.\nAll tmux panes are already being watched.")
		} else {
			content = m.browserList.View()
		}
		popup = browserPopupStyle.Width(popupWidth).Render(content)
	} else {
		// Pane list: split layout with preview (on wide screens) or single panel (narrow)
		// Use 90% of terminal width
		popupWidth := m.width * 90 / 100
		if popupWidth < 40 {
			popupWidth = m.width - 4
		}

		// On narrow screens (< 80 cols), show list only without preview
		showPreview := m.width >= 80

		var listWidth, previewWidth int
		if showPreview {
			// Split: 35% for list, 65% for preview
			listWidth = popupWidth * 35 / 100
			if listWidth < 30 {
				listWidth = 30
			}
			previewWidth = popupWidth - listWidth - 8 // account for borders and gaps
		} else {
			// Single panel - use full popup width
			listWidth = popupWidth - 4
			previewWidth = 0
		}

		var listContent string
		if m.browserEmpty {
			listContent = emptyStyle.Render("No panes available in this session.")
		} else {
			listContent = m.browserList.View()
		}

		// Render list panel
		listPanel := browserPopupStyle.
			Width(listWidth).
			Render(listContent)

		if showPreview {
			// Render preview panel
			var previewContent string
			if m.browserPreviewErr != nil {
				previewContent = errorStyle.Render(fmt.Sprintf("Error: %v", m.browserPreviewErr))
			} else if m.browserPreviewContent == "" {
				previewContent = emptyStyle.Render("No content")
			} else {
				// Truncate lines that are too wide and limit height
				lines := strings.Split(m.browserPreviewContent, "\n")
				maxLines := 18
				if len(lines) > maxLines {
					lines = lines[len(lines)-maxLines:]
				}
				// Truncate each line to fit preview width
				maxLineWidth := previewWidth - 4 // account for padding
				for i, line := range lines {
					if len(line) > maxLineWidth {
						lines[i] = line[:maxLineWidth]
					}
				}
				previewContent = strings.Join(lines, "\n")
			}

			previewPanel := previewPanelStyle.
				Width(previewWidth).
				Render(previewTitleStyle.Render("Preview") + "\n" + previewContent)

			// Join panels horizontally
			popup = lipgloss.JoinHorizontal(lipgloss.Top, listPanel, previewPanel)
		} else {
			// Narrow screen - list only
			popup = listPanel
		}
	}

	// Center the popup
	centered := lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		popup,
	)

	// Add footer help - different for session vs pane selection
	var footer string
	if m.browsingSession {
		footer = helpStyle.Render("↑/↓: navigate • Enter: select session • Esc: cancel")
	} else {
		footer = helpStyle.Render("↑/↓: navigate • Enter: add pane • Esc: back to sessions")
	}

	return centered + "\n" + footer
}

func (m Model) renderConfigurePopup() string {
	pane := m.watchlist.GetPane(m.selectedPaneID)
	if pane == nil {
		return "Error: pane not found"
	}

	// Build menu items
	var lines []string

	// Title
	displayName := pane.DisplayName()
	lines = append(lines, browserTitleStyle.Render("Configure: "+displayName))
	lines = append(lines, "")

	// Name editing row
	if m.configEditingName {
		lines = append(lines, "Name: "+m.textInput.View())
	} else {
		nameValue := pane.Name
		if nameValue == "" {
			nameValue = emptyStyle.Render("(none)")
		}
		if m.configMenuItem == configMenuName {
			lines = append(lines, "> Name: "+nameValue)
		} else {
			lines = append(lines, "  Name: "+nameValue)
		}
	}

	// Sound toggle row - show tri-state: [D] default, [x] enabled, [ ] disabled
	soundStatus := triStateIndicator(pane.SoundOnReady, m.config.Alerts.SoundOnReady)
	if m.configMenuItem == configMenuSound {
		lines = append(lines, "> Sound on Ready: "+soundStatus)
	} else {
		lines = append(lines, "  Sound on Ready: "+soundStatus)
	}

	// Sound type row - show current sound type with indicator for default
	soundType := pane.GetEffectiveSoundType(m.config)
	soundTypeDisplay := soundType
	if pane.SoundType == nil || *pane.SoundType == "" {
		soundTypeDisplay = "[D:" + soundType + "]"
	}
	if m.configMenuItem == configMenuSoundType {
		lines = append(lines, "> Sound Type: "+soundTypeDisplay)
	} else {
		lines = append(lines, "  Sound Type: "+soundTypeDisplay)
	}

	// Notification toggle row - show tri-state: [D] default, [x] enabled, [ ] disabled
	notifyStatus := triStateIndicator(pane.NotifyOnReady, m.config.Alerts.NotifyOnReady)
	if m.configMenuItem == configMenuNotify {
		lines = append(lines, "> Notify on Ready: "+notifyStatus)
	} else {
		lines = append(lines, "  Notify on Ready: "+notifyStatus)
	}

	content := ""
	for _, line := range lines {
		content += line + "\n"
	}

	popup := browserPopupStyle.Render(content)

	// Center the popup
	centered := lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		popup,
	)

	// Footer
	var footer string
	if m.configEditingName {
		footer = helpStyle.Render("Enter: save • Esc: cancel")
	} else {
		footer = helpStyle.Render("↑/↓: navigate • Enter/Space: toggle/edit • Esc: close")
	}

	return centered + "\n" + footer
}
