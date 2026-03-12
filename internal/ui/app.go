package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
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

// paneItem implements list.Item for watchlist panes
type paneItem struct {
	id      string
	addedAt time.Time
}

func (p paneItem) Title() string       { return p.id }
func (p paneItem) Description() string { return "added " + p.addedAt.Format("2006-01-02 15:04") }
func (p paneItem) FilterValue() string { return p.id }

type Model struct {
	list           list.Model
	viewport       viewport.Model
	empty          bool
	loadErr        error
	selectedPaneID string
	previewContent string
	previewErr     error
	width          int
	height         int
}

func New() Model {
	wl, err := watchlist.Load()
	if err != nil {
		return Model{loadErr: err}
	}

	items := make([]list.Item, len(wl.Panes))
	for i, p := range wl.Panes {
		items[i] = paneItem{id: p.ID, addedAt: p.AddedAt}
	}

	l := list.New(items, list.NewDefaultDelegate(), 30, 20)
	l.Title = "Watched Panes"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = titleStyle

	vp := viewport.New(50, 20)

	m := Model{
		list:     l,
		viewport: vp,
		empty:    len(items) == 0,
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
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
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

	return layout + "\n" + helpStyle.Render("↑/↓: navigate • q: quit")
}
