# pane-browser Specification

## Purpose
TBD - created by archiving change in-app-pane-browser. Update Purpose after archive.
## Requirements
### Requirement: Open pane browser with keyboard shortcut

The system SHALL open a pane browser popup when the user presses `a` in the main view.

#### Scenario: Open browser
- **WHEN** user presses `a` key in the main view
- **THEN** a popup overlay appears showing available tmux panes

#### Scenario: Browser not opened during edit mode
- **WHEN** user presses `a` while in edit mode
- **THEN** the pane browser does NOT open

### Requirement: Navigate pane list

The system SHALL allow keyboard navigation of the pane list.

#### Scenario: Navigate with arrow keys
- **WHEN** the pane browser is open
- **THEN** up/down arrow keys move selection through the pane list

### Requirement: Select and add pane

The system SHALL add the selected pane to the watchlist with an auto-guessed name when confirmed.

#### Scenario: Add pane with Enter
- **WHEN** user presses Enter on a selected pane in the pane list
- **THEN** the pane is added to the watchlist with an auto-guessed name
- **AND** the popup closes
- **AND** the main list refreshes to show the new pane with its guessed name

#### Scenario: Add pane with generic name
- **WHEN** user selects a pane where the guessed name is generic
- **THEN** the pane is added to the watchlist with the guessed name anyway
- **AND** the user can rename it later via the configure popup

### Requirement: Cancel pane browser

The system SHALL close the popup without changes when cancelled.

#### Scenario: Cancel with Escape from session list
- **WHEN** user presses Escape while viewing session list
- **THEN** the popup closes
- **AND** no pane is added to the watchlist

#### Scenario: Cancel with q key
- **WHEN** user presses `q` at any level in the pane browser
- **THEN** the popup closes
- **AND** no pane is added to the watchlist

### Requirement: List tmux panes

The system SHALL provide a function to list all tmux panes with metadata.

#### Scenario: Get all panes
- **WHEN** listing tmux panes
- **THEN** returns pane ID, session name, window index, pane index, and current command for each pane

### Requirement: Display session list first

The system SHALL display a list of sessions as the first step when opening the pane browser.

#### Scenario: Show sessions on browser open
- **WHEN** user presses `a` to open pane browser
- **THEN** a list of tmux sessions is displayed
- **AND** each session shows its name and count of available (unwatched) panes

#### Scenario: Session with no available panes is hidden
- **WHEN** all panes in a session are already watched
- **THEN** that session is NOT shown in the session list

### Requirement: Select session to view panes

The system SHALL show panes for a selected session when the user presses Enter.

#### Scenario: Enter on session shows pane list
- **WHEN** user presses Enter on a selected session
- **THEN** the view switches to show panes within that session
- **AND** the popup title indicates the selected session name

#### Scenario: Pane list shows window and command
- **WHEN** viewing panes for a session
- **THEN** each pane displays its window.pane index and running command

### Requirement: Navigate back from pane list

The system SHALL allow navigating back from pane list to session list with Escape.

#### Scenario: Escape from pane list returns to sessions
- **WHEN** user presses Escape while viewing pane list
- **THEN** the view returns to the session list
- **AND** no pane is added to the watchlist

### Requirement: Show preview panel in pane browser

The system SHALL display a preview panel showing the currently selected pane's content when viewing panes in the browser.

#### Scenario: Preview shown when viewing panes
- **WHEN** user is viewing the pane list (not session list) in the browser
- **THEN** a preview panel is displayed next to the pane list
- **AND** the preview shows the content of the currently selected pane

#### Scenario: Preview not shown during session selection
- **WHEN** user is viewing the session list in the browser
- **THEN** no preview panel is displayed
- **AND** only the session list is shown

### Requirement: Update preview on navigation

The system SHALL update the preview panel content when the user navigates to a different pane.

#### Scenario: Navigate to different pane
- **WHEN** user presses up/down arrow to select a different pane
- **THEN** the preview panel updates to show the newly selected pane's content

#### Scenario: Enter pane list from session
- **WHEN** user presses Enter on a session to view its panes
- **THEN** the first pane in the list is selected
- **AND** the preview panel shows that pane's content

### Requirement: Handle preview capture errors

The system SHALL gracefully handle errors when capturing pane content for preview.

#### Scenario: Pane capture fails
- **WHEN** capturing pane content for preview fails
- **THEN** the preview panel displays an error message
- **AND** the browser remains functional for selection

