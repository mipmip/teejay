## ADDED Requirements

### Requirement: Delete all panes from watchlist

The user SHALL be able to delete all panes from the watchlist with a single keybinding and confirmation.

#### Scenario: Delete all with confirmation
- **WHEN** the user presses `D` (shift+d)
- **AND** the watchlist is not empty
- **THEN** a confirmation prompt SHALL be shown: "Delete all N panes? (y/n)"

#### Scenario: Confirm delete all
- **WHEN** the confirmation prompt is shown
- **AND** the user presses `y`
- **THEN** all panes SHALL be removed from the watchlist
- **AND** the watchlist SHALL be saved
- **AND** the empty state SHALL be shown

#### Scenario: Cancel delete all
- **WHEN** the confirmation prompt is shown
- **AND** the user presses `n` or `Esc`
- **THEN** the delete SHALL be cancelled and the watchlist SHALL remain unchanged

#### Scenario: Delete all on empty watchlist
- **WHEN** the user presses `D`
- **AND** the watchlist is empty
- **THEN** nothing SHALL happen
