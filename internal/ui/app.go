package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"tj/internal/alerts"
	"tj/internal/monitor"
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
)

// previewTickMsg is sent periodically to trigger preview refresh
type previewTickMsg struct{}

// tickCmd returns a command that sends a previewTickMsg after 100ms
func tickCmd() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return previewTickMsg{}
	})
}

// paneItem implements list.Item for watchlist panes
type paneItem struct {
	id      string
	name    string
	addedAt time.Time
	status  monitor.PaneStatus
	frame   int
}

func (p paneItem) Title() string {
	indicator := p.status.IndicatorAnimated(p.frame)
	// Apply green styling for Ready status
	if p.status == monitor.Ready {
		indicator = readyIndicatorStyle.Render(indicator)
	}
	displayName := p.id
	if p.name != "" {
		displayName = p.name
	}
	return indicator + " " + displayName
}
func (p paneItem) Description() string { return p.id + " • added " + p.addedAt.Format("2006-01-02 15:04") }
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
	configMenuNotify
)

type Model struct {
	list           list.Model
	viewport       viewport.Model
	textInput      textinput.Model
	watchlist      *watchlist.Watchlist
	monitor        *monitor.Monitor
	paneStatuses   map[string]monitor.PaneStatus
	empty          bool
	loadErr        error
	selectedPaneID string
	previewContent string
	previewErr     error
	width          int
	height         int
	editing        bool
	deleting       bool
	notInTmuxMsg   bool
	browsing        bool
	browserList     list.Model
	browserEmpty    bool
	tickFrame       int
	browsingSession bool              // true when showing sessions, false when showing panes
	selectedSession string            // session name selected for pane browsing
	allBrowserPanes []tmux.PaneInfo   // cached panes for filtering by session
	configuring     bool              // true when configure popup is open
	configMenuItem  configMenuItem    // selected menu item in configure popup
	configEditingName bool            // true when editing name in configure popup
	watchlistMtime  time.Time         // last known modification time of watchlist file
}

func New() Model {
	wl, err := watchlist.Load()
	if err != nil {
		return Model{loadErr: err}
	}

	// Get initial mtime for watchlist file
	var wlMtime time.Time
	if path, err := watchlist.ConfigPath(); err == nil {
		if info, err := os.Stat(path); err == nil {
			wlMtime = info.ModTime()
		}
	}

	items := make([]list.Item, len(wl.Panes))
	for i, p := range wl.Panes {
		items[i] = paneItem{id: p.ID, name: p.Name, addedAt: p.AddedAt, status: monitor.Idle}
	}

	l := list.New(items, list.NewDefaultDelegate(), 30, 20)
	l.Title = "Watched Panes"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = titleStyle

	vp := viewport.New(50, 20)

	ti := textinput.New()
	ti.Placeholder = "Enter name..."
	ti.CharLimit = 50

	m := Model{
		list:           l,
		viewport:       vp,
		textInput:      ti,
		watchlist:      wl,
		monitor:        monitor.New(),
		paneStatuses:   make(map[string]monitor.PaneStatus),
		empty:          len(items) == 0,
		watchlistMtime: wlMtime,
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

func (m *Model) captureSelectedPane() {
	if m.selectedPaneID == "" {
		m.previewContent = ""
		m.previewErr = nil
		return
	}

	content, err := tmux.CapturePane(m.selectedPaneID)
	if err != nil {
		m.previewErr = err
		m.previewContent = ""
	} else {
		m.previewErr = nil
		m.previewContent = content
	}
	m.viewport.SetContent(m.previewContent)
}

func (m Model) Init() tea.Cmd {
	if !m.empty {
		return tickCmd()
	}
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle preview tick - refresh pane content periodically
	if _, ok := msg.(previewTickMsg); ok {
		m.tickFrame++
		// Skip refresh when in modal modes
		if !m.editing && !m.deleting && !m.browsing && !m.configuring {
			// Check if watchlist file has been modified externally
			m.checkWatchlistFileChange()

			// Only refresh preview if we have a pane selected
			if !m.empty && m.selectedPaneID != "" {
				m.captureSelectedPane()
				// Update status for selected pane
				prevStatus := m.paneStatuses[m.selectedPaneID]
				status := m.monitor.Update(m.selectedPaneID, m.previewContent)
				if prevStatus != status {
					m.paneStatuses[m.selectedPaneID] = status
					// Check for Running -> Ready transition and trigger alerts
					if prevStatus == monitor.Running && status == monitor.Ready {
						m.triggerAlerts(m.selectedPaneID)
					}
				}
				// Always refresh list to update spinner animation
				m.refreshListWithFrame(m.tickFrame)
			}
		}
		return m, tickCmd()
	}

	// Handle edit mode
	if m.editing {
		return m.updateEditing(msg)
	}

	// Handle delete confirmation mode
	if m.deleting {
		return m.updateDeleting(msg)
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
			if listWidth < 20 {
				listWidth = 20
			}
			// Check if click is within the list panel area (accounting for border)
			if msg.X >= 1 && msg.X <= listWidth+1 {
				// Calculate which item was clicked
				// List starts rendering items at y=3 (border + title + spacing)
				// Each item takes 2 lines in the default delegate (title + description)
				itemHeight := 2
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
					// Switch to the pane but keep app running
					tmux.SwitchToPane(m.selectedPaneID)
					return m, nil
				}
				// Show "not in tmux" message
				m.notInTmuxMsg = true
				return m, nil
			}
		case "esc":
			// Clear any temporary messages
			if m.notInTmuxMsg {
				m.notInTmuxMsg = false
				return m, nil
			}
		case "a":
			// Open pane browser
			m.loadBrowserPanes()
			m.browsing = true
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Calculate panel sizes (30% list, 70% preview, minus borders)
		listWidth := m.width*30/100 - 2
		previewWidth := m.width*70/100 - 2
		panelHeight := m.height - 4

		m.list.SetWidth(listWidth)
		m.list.SetHeight(panelHeight)
		m.viewport.Width = previewWidth
		m.viewport.Height = panelHeight
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
				itemHeight := 2 // title + description
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
					// Pane selected - add to watchlist
					if item, ok := m.browserList.SelectedItem().(browserItem); ok {
						m.watchlist.Add(item.paneInfo.ID)
						m.watchlist.Save()
						m.refreshList()
						m.empty = false
						// Select the newly added pane
						m.selectedPaneID = item.paneInfo.ID
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

	// Update browser list for navigation
	var cmd tea.Cmd
	m.browserList, cmd = m.browserList.Update(msg)
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
				// Items: title (line 0), blank (line 1), Name (line 2), Sound (line 3), Notify (line 4)
				// Accounting for border (1) + padding (1) = 2 offset
				relativeY := msg.Y - popupY - 2

				switch relativeY {
				case 2: // Name row
					m.configMenuItem = configMenuName
				case 3: // Sound row
					m.configMenuItem = configMenuSound
				case 4: // Notify row
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
				// Toggle sound
				m.watchlist.SetSound(m.selectedPaneID, !pane.SoundOnReady)
				m.watchlist.Save()
				return m, nil
			case configMenuNotify:
				// Toggle notification
				m.watchlist.SetNotify(m.selectedPaneID, !pane.NotifyOnReady)
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

	if pane.SoundOnReady {
		alerts.PlayBell()
	}

	if pane.NotifyOnReady {
		displayName := pane.DisplayName()
		alerts.SendNotification("Teejay", displayName+" is ready")
	}
}

func (m *Model) refreshList() {
	m.refreshListWithFrame(m.tickFrame)
}

func (m *Model) refreshListWithFrame(frame int) {
	items := make([]list.Item, len(m.watchlist.Panes))
	for i, p := range m.watchlist.Panes {
		status := m.paneStatuses[p.ID]
		items[i] = paneItem{id: p.ID, name: p.Name, addedAt: p.AddedAt, status: status, frame: frame}
	}
	m.list.SetItems(items)
	m.empty = len(items) == 0
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

	// Create session list
	delegate := list.NewDefaultDelegate()
	m.browserList = list.New(items, delegate, 50, 15)
	m.browserList.Title = "Select Session"
	m.browserList.SetShowStatusBar(false)
	m.browserList.SetFilteringEnabled(false)
	m.browserList.SetShowHelp(false)
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

	// Create pane list
	delegate := list.NewDefaultDelegate()
	m.browserList = list.New(items, delegate, 50, 15)
	m.browserList.Title = "Select Pane (" + sessionName + ")"
	m.browserList.SetShowStatusBar(false)
	m.browserList.SetFilteringEnabled(false)
	m.browserList.SetShowHelp(false)
	m.browserList.Styles.Title = browserTitleStyle

	m.browsingSession = false
	m.selectedSession = sessionName
	m.browserEmpty = len(items) == 0
}

func (m Model) View() string {
	if m.loadErr != nil {
		return fmt.Sprintf("Error loading watchlist: %v\n\nPress q to quit.\n", m.loadErr)
	}

	if m.empty && !m.browsing {
		return titleStyle.Render("Teejay") + "\n\n" +
			emptyStyle.Render("No panes are being watched.") + "\n\n" +
			helpStyle.Render("Press 'a' to browse and add panes, or run 'tj add' in a tmux pane.") + "\n\n" +
			helpStyle.Render("Press q to quit.")
	}

	// Handle browsing popup
	if m.browsing {
		return m.renderBrowserPopup()
	}

	// Handle configure popup
	if m.configuring {
		return m.renderConfigurePopup()
	}

	// Calculate panel widths
	listWidth := m.width*30/100 - 2
	previewWidth := m.width*70/100 - 2
	if listWidth < 20 {
		listWidth = 20
	}
	if previewWidth < 20 {
		previewWidth = 20
	}

	// Build list panel
	listPanel := listPanelStyle.
		Width(listWidth).
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

	previewTitle := previewTitleStyle.Render("Preview: " + m.selectedPaneID)
	previewPanel := previewPanelStyle.
		Width(previewWidth).
		Render(previewTitle + "\n" + previewContent)

	// Join panels horizontally
	layout := lipgloss.JoinHorizontal(lipgloss.Top, listPanel, previewPanel)

	// Show mode-specific help/input
	var footer string
	if m.editing {
		footer = "Rename: " + m.textInput.View() + "\n" + helpStyle.Render("Enter: save • Esc: cancel")
	} else if m.deleting {
		paneName := m.selectedPaneID
		if item, ok := m.list.SelectedItem().(paneItem); ok {
			paneName = item.Title()
		}
		footer = errorStyle.Render(fmt.Sprintf("Delete %s? (y/n)", paneName))
	} else if m.notInTmuxMsg {
		footer = errorStyle.Render("Cannot switch: not running inside tmux") + "\n" + helpStyle.Render("Press Esc to dismiss")
	} else {
		footer = helpStyle.Render("↑/↓: navigate • Enter: switch • a: add • c: configure • d: delete • q: quit")
	}

	return layout + "\n" + footer
}

func (m Model) renderBrowserPopup() string {
	var content string
	if m.browserEmpty {
		if m.browsingSession {
			content = emptyStyle.Render("No additional panes available.\nAll tmux panes are already being watched.")
		} else {
			content = emptyStyle.Render("No panes available in this session.")
		}
	} else {
		content = m.browserList.View()
	}

	popup := browserPopupStyle.Render(content)

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

	// Sound toggle row
	soundStatus := "[ ]"
	if pane.SoundOnReady {
		soundStatus = "[x]"
	}
	if m.configMenuItem == configMenuSound {
		lines = append(lines, "> Sound on Ready: "+soundStatus)
	} else {
		lines = append(lines, "  Sound on Ready: "+soundStatus)
	}

	// Notification toggle row
	notifyStatus := "[ ]"
	if pane.NotifyOnReady {
		notifyStatus = "[x]"
	}
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
