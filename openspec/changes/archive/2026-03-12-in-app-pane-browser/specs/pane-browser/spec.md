## ADDED Requirements

### Requirement: Open pane browser with keyboard shortcut

The system SHALL open a pane browser popup when the user presses `a` in the main view.

#### Scenario: Open browser
- **WHEN** user presses `a` key in the main view
- **THEN** a popup overlay appears showing available tmux panes

#### Scenario: Browser not opened during edit mode
- **WHEN** user presses `a` while in edit mode
- **THEN** the pane browser does NOT open

### Requirement: Display all available tmux panes

The system SHALL display a list of all tmux panes with their session and window context.

#### Scenario: List panes with context
- **WHEN** the pane browser is open
- **THEN** each pane is displayed with format "session:window.pane command"

#### Scenario: Filter out already watched panes
- **WHEN** the pane browser is open
- **THEN** panes already in the watchlist are NOT shown in the list

#### Scenario: Show message when no panes available
- **WHEN** the pane browser is open and all panes are already watched
- **THEN** a message indicates no additional panes are available

### Requirement: Navigate pane list

The system SHALL allow keyboard navigation of the pane list.

#### Scenario: Navigate with arrow keys
- **WHEN** the pane browser is open
- **THEN** up/down arrow keys move selection through the pane list

### Requirement: Select and add pane

The system SHALL add the selected pane to the watchlist when confirmed.

#### Scenario: Add pane with Enter
- **WHEN** user presses Enter on a selected pane
- **THEN** the pane is added to the watchlist
- **AND** the popup closes
- **AND** the main list refreshes to show the new pane

### Requirement: Cancel pane browser

The system SHALL close the popup without changes when cancelled.

#### Scenario: Cancel with Escape
- **WHEN** user presses Escape in the pane browser
- **THEN** the popup closes
- **AND** no pane is added to the watchlist

### Requirement: List tmux panes

The system SHALL provide a function to list all tmux panes with metadata.

#### Scenario: Get all panes
- **WHEN** listing tmux panes
- **THEN** returns pane ID, session name, window index, pane index, and current command for each pane
