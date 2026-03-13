## ADDED Requirements

### Requirement: Detect external watchlist changes

The system SHALL detect when the watchlist file is modified by an external process.

#### Scenario: File modified by tmon add
- **WHEN** `tmon add` is run in another terminal while TUI is running
- **THEN** the TUI detects the file modification within 200ms

#### Scenario: No detection during modal modes
- **WHEN** user is in edit, delete, or browse mode
- **THEN** file change detection is skipped until modal is closed

### Requirement: Auto-reload watchlist on change

The system SHALL reload the watchlist when external changes are detected.

#### Scenario: Reload after external add
- **WHEN** a new pane is added via `tmon add`
- **THEN** the TUI reloads the watchlist from disk
- **AND** the pane list updates to show the new pane

#### Scenario: Reload after external remove
- **WHEN** the watchlist file is modified externally (pane removed)
- **THEN** the TUI reloads and reflects the removal

### Requirement: Preserve selection after refresh

The system SHALL preserve the user's current selection when possible.

#### Scenario: Selected pane still exists
- **WHEN** watchlist is reloaded
- **AND** the currently selected pane still exists
- **THEN** that pane remains selected

#### Scenario: Selected pane was removed
- **WHEN** watchlist is reloaded
- **AND** the currently selected pane no longer exists
- **THEN** the first pane in the list is selected

### Requirement: Expose watchlist config path

The system SHALL provide access to the watchlist file path for monitoring.

#### Scenario: Get config path
- **WHEN** code needs to monitor the watchlist file
- **THEN** the config path is accessible via exported function
