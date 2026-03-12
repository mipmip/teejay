# watchlist-management Specification

## Purpose
TBD - created by archiving change watchlist-edit-delete. Update Purpose after archive.
## Requirements
### Requirement: Edit pane name

The system SHALL allow users to edit a pane's display name by pressing `e` on the selected pane.

#### Scenario: Enter edit mode
- **WHEN** user presses `e` on a selected pane
- **THEN** a text input appears with the current name (or pane ID if no name set)
- **AND** the cursor is at the end of the text

#### Scenario: Save edited name
- **WHEN** user is in edit mode and presses `Enter`
- **THEN** the new name is saved to the watchlist
- **AND** the list updates to show the new name
- **AND** edit mode exits

#### Scenario: Cancel edit
- **WHEN** user is in edit mode and presses `Escape`
- **THEN** edit mode exits without saving
- **AND** the original name is preserved

### Requirement: Delete pane from watchlist

The system SHALL allow users to delete a pane by pressing `d` on the selected pane.

#### Scenario: Request delete
- **WHEN** user presses `d` on a selected pane
- **THEN** a confirmation prompt appears asking "Delete [pane name]? (y/n)"

#### Scenario: Confirm delete
- **WHEN** user is in delete confirmation mode and presses `y`
- **THEN** the pane is removed from the watchlist
- **AND** the watchlist is saved
- **AND** the list updates to reflect the removal

#### Scenario: Cancel delete
- **WHEN** user is in delete confirmation mode and presses `n` or `Escape`
- **THEN** delete is cancelled
- **AND** the pane remains in the watchlist

### Requirement: Display pane name

The system SHALL display the pane's custom name if set, otherwise the pane ID.

#### Scenario: Pane has custom name
- **WHEN** displaying a pane that has a Name field set
- **THEN** the custom name is displayed in the list

#### Scenario: Pane has no custom name
- **WHEN** displaying a pane that has no Name field or empty Name
- **THEN** the pane ID is displayed in the list

### Requirement: Remove pane from watchlist data

The watchlist SHALL provide a method to remove a pane by ID.

#### Scenario: Remove existing pane
- **WHEN** removing a pane ID that exists in the watchlist
- **THEN** the pane is removed from the list

#### Scenario: Remove non-existent pane
- **WHEN** removing a pane ID that does not exist in the watchlist
- **THEN** the watchlist remains unchanged

### Requirement: Rename pane in watchlist data

The watchlist SHALL provide a method to rename a pane by ID.

#### Scenario: Rename existing pane
- **WHEN** renaming a pane ID that exists in the watchlist
- **THEN** the pane's Name field is updated

