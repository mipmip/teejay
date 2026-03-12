# pane-deduplication Specification

## Purpose
TBD - created by archiving change deduplicate-panes. Update Purpose after archive.
## Requirements
### Requirement: Prevent adding duplicate panes

The system SHALL prevent adding a pane that is already in the watchlist.

#### Scenario: Add pane that already exists
- **WHEN** running `tmon add` for a pane already in the watchlist
- **THEN** a message "Pane %s is already being watched" is displayed
- **AND** the watchlist is not modified
- **AND** the command exits with code 0 (not an error)

#### Scenario: Add new pane
- **WHEN** running `tmon add` for a pane not in the watchlist
- **THEN** the pane is added to the watchlist
- **AND** a success message is displayed

### Requirement: Check if pane exists in watchlist

The watchlist SHALL provide a method to check if a pane ID exists.

#### Scenario: Pane exists
- **WHEN** checking for a pane ID that is in the watchlist
- **THEN** the check returns true

#### Scenario: Pane does not exist
- **WHEN** checking for a pane ID that is not in the watchlist
- **THEN** the check returns false

### Requirement: Deduplicate on load

The system SHALL remove duplicate pane entries when loading the watchlist.

#### Scenario: Watchlist has duplicates
- **WHEN** loading a watchlist with duplicate pane IDs
- **THEN** only the first occurrence of each pane ID is kept
- **AND** the deduplicated list is returned

#### Scenario: Watchlist has no duplicates
- **WHEN** loading a watchlist with unique pane IDs
- **THEN** all entries are preserved unchanged

