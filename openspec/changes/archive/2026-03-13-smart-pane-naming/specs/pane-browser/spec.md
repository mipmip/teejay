## MODIFIED Requirements

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
