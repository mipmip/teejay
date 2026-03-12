## ADDED Requirements

### Requirement: Split-panel layout

The TUI SHALL display a split-panel layout with the pane list on the left and preview on the right.

#### Scenario: Layout displays correctly
- **WHEN** the TUI starts with panes in the watchlist
- **THEN** the pane list is displayed on the left side (~30% width)
- **AND** the preview panel is displayed on the right side (~70% width)

#### Scenario: Empty watchlist layout
- **WHEN** the TUI starts with an empty watchlist
- **THEN** the empty state message is displayed (no split panel needed)

### Requirement: Pane content capture

The system SHALL capture tmux pane content using `tmux capture-pane`.

#### Scenario: Capture pane content
- **WHEN** a pane is selected in the list
- **THEN** the system runs `tmux capture-pane -p -t <pane-id>`
- **AND** the output is displayed in the preview panel

#### Scenario: Pane no longer exists
- **WHEN** capturing content for a pane that no longer exists
- **THEN** an error message is displayed in the preview panel
- **AND** the TUI remains functional

### Requirement: Preview updates on selection

The preview panel SHALL update when the user selects a different pane.

#### Scenario: Change selection
- **WHEN** the user navigates to a different pane in the list
- **THEN** the preview panel updates to show the newly selected pane's content

### Requirement: Scrollable preview

The preview panel SHALL be scrollable when content exceeds the viewport.

#### Scenario: Scroll preview content
- **WHEN** the pane content is longer than the preview viewport
- **THEN** the user can scroll using Page Up/Down or arrow keys when focused on preview
