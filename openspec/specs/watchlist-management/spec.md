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

The watchlist SHALL provide a method to remove a pane by ID, optionally returning whether a pane was actually removed.

#### Scenario: Remove existing pane
- **WHEN** removing a pane ID that exists in the watchlist
- **THEN** the pane is removed from the list

#### Scenario: Remove non-existent pane
- **WHEN** removing a pane ID that does not exist in the watchlist
- **THEN** the watchlist remains unchanged

#### Scenario: Check if pane was removed
- **WHEN** removing a pane programmatically
- **THEN** the caller can determine if a pane was actually removed (for logging/notification purposes)

### Requirement: Rename pane in watchlist data

The watchlist SHALL provide a method to rename a pane by ID.

#### Scenario: Rename existing pane
- **WHEN** renaming a pane ID that exists in the watchlist
- **THEN** the pane's Name field is updated

### Requirement: Load watchlist from custom path

The watchlist module SHALL accept an optional custom path for loading, defaulting to `~/.config/teejay/watchlist.json`.

#### Scenario: Load from default path
- **WHEN** Load() is called without a path argument
- **THEN** watchlist is loaded from `~/.config/teejay/watchlist.json`

#### Scenario: Load from custom path
- **WHEN** Load() is called with a custom path argument
- **THEN** watchlist is loaded from the specified path
- **AND** the watchlist remembers the path for subsequent saves

### Requirement: Save watchlist to custom path

The watchlist SHALL save to the same path it was loaded from.

#### Scenario: Save to default path
- **WHEN** watchlist was loaded from default path
- **AND** Save() is called
- **THEN** watchlist is saved to `~/.config/teejay/watchlist.json`

#### Scenario: Save to custom path
- **WHEN** watchlist was loaded from a custom path
- **AND** Save() is called
- **THEN** watchlist is saved to the custom path (not the default)

#### Scenario: Save creates parent directories
- **WHEN** Save() is called
- **AND** the parent directory does not exist
- **THEN** parent directories are created before saving

