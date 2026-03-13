## ADDED Requirements

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

## MODIFIED Requirements

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

### Requirement: Select and add pane

The system SHALL add the selected pane to the watchlist when confirmed.

#### Scenario: Add pane with Enter
- **WHEN** user presses Enter on a selected pane in the pane list
- **THEN** the pane is added to the watchlist
- **AND** the popup closes
- **AND** the main list refreshes to show the new pane

## REMOVED Requirements

### Requirement: Display all available tmux panes

**Reason**: Replaced by hierarchical session→pane display
**Migration**: Panes are now shown per-session after selecting a session from the session list
