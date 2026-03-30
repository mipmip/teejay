# clear-screen-on-quit Specification

## Purpose
TBD - created by archiving change clear-screen-on-quit. Update Purpose after archive.
## Requirements
### Requirement: Clean terminal on application exit

The application SHALL use the terminal's alternate screen buffer so that quitting restores the previous terminal state.

#### Scenario: Quit restores terminal
- **WHEN** the user quits the application
- **THEN** the terminal SHALL return to the state it was in before the application started
- **AND** no TUI output SHALL remain visible

