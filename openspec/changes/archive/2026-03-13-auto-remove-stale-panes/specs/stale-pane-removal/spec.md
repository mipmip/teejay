## ADDED Requirements

### Requirement: Detect stale panes

The system SHALL detect when a watched pane no longer exists in tmux by checking for "can't find pane" in capture errors.

#### Scenario: Pane capture fails with missing pane error
- **WHEN** capturing a pane returns an error containing "can't find pane"
- **THEN** the system identifies the pane as stale

#### Scenario: Pane capture fails with other error
- **WHEN** capturing a pane returns an error NOT containing "can't find pane"
- **THEN** the system does NOT identify the pane as stale
- **AND** the error is displayed to the user as normal

### Requirement: Auto-remove stale panes

The system SHALL automatically remove stale panes from the watchlist when detected.

#### Scenario: Stale pane detected during capture
- **WHEN** a pane is identified as stale during capture
- **THEN** the pane is removed from the watchlist
- **AND** the watchlist is saved to disk
- **AND** the pane list UI is refreshed

#### Scenario: Selected pane becomes stale
- **WHEN** the currently selected pane is removed as stale
- **THEN** selection moves to an adjacent pane if available
- **AND** preview updates to show the new selection

### Requirement: Notify user of stale pane removal

The system SHALL briefly notify the user when a stale pane is automatically removed.

#### Scenario: Stale pane removed
- **WHEN** a stale pane is auto-removed
- **THEN** a status message is displayed (e.g., "Removed stale pane %65")
- **AND** the message clears on the next user action or after a brief delay
