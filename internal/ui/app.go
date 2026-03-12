package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"tmon/internal/tmux"
	"tmon/internal/watchlist"
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
}

func (p paneItem) Title() string {
	if p.name != "" {
		return p.name
	}
	return p.id
}
func (p paneItem) Description() string { return p.id + " • added " + p.addedAt.Format("2006-01-02 15:04") }
func (p paneItem) FilterValue() string { return p.id }

type Model struct {
	list           list.Model
	viewport       viewport.Model
	textInput      textinput.Model
	watchlist      *watchlist.Watchlist
	empty          bool
	loadErr        error
	selectedPaneID string
	previewContent string
	previewErr     error
	width          int
	height         int
	editing        bool
	deleting       bool
}

func New() Model {
	wl, err := watchlist.Load()
	if err != nil {
		return Model{loadErr: err}
	}

	items := make([]list.Item, len(wl.Panes))
	for i, p := range wl.Panes {
		items[i] = paneItem{id: p.ID, name: p.Name, addedAt: p.AddedAt}
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
		list:      l,
		viewport:  vp,
		textInput: ti,
		watchlist: wl,
		empty:     len(items) == 0,
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
		// Skip refresh when in modal modes or no pane selected
		if !m.editing && !m.deleting && !m.empty && m.selectedPaneID != "" {
			m.captureSelectedPane()
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

	switch msg := msg.(type) {
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

func (m *Model) refreshList() {
	items := make([]list.Item, len(m.watchlist.Panes))
	for i, p := range m.watchlist.Panes {
		items[i] = paneItem{id: p.ID, name: p.Name, addedAt: p.AddedAt}
	}
	m.list.SetItems(items)
	m.empty = len(items) == 0
}

func (m Model) View() string {
	if m.loadErr != nil {
		return fmt.Sprintf("Error loading watchlist: %v\n\nPress q to quit.\n", m.loadErr)
	}

	if m.empty {
		return titleStyle.Render("tmon") + "\n\n" +
			emptyStyle.Render("No panes are being watched.") + "\n\n" +
			helpStyle.Render("Run 'tmon add' in a tmux pane to start watching it.") + "\n\n" +
			helpStyle.Render("Press q to quit.")
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
	} else {
		footer = helpStyle.Render("↑/↓: navigate • e: edit • d: delete • q: quit")
	}

	return layout + "\n" + footer
}
