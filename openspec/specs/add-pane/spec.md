# add-pane Specification

## Purpose
TBD - created by archiving change add-pane-command. Update Purpose after archive.
## Requirements
### Requirement: Add pane subcommand

The system SHALL provide a `tmon add` subcommand that adds the current tmux pane to the watchlist.

#### Scenario: Successfully add pane inside tmux
- **WHEN** running `tmon add` from within a tmux pane
- **THEN** the current pane ID is added to the watchlist
- **AND** a success message is printed to stdout

#### Scenario: Add pane outside tmux
- **WHEN** running `tmon add` outside of a tmux session
- **THEN** an error message is printed to stderr
- **AND** the command exits with a non-zero status

### Requirement: Tmux pane detection

The system SHALL detect the current tmux pane ID by reading the `$TMUX_PANE` environment variable.

#### Scenario: Pane ID available
- **WHEN** `$TMUX_PANE` is set (e.g., `%0`)
- **THEN** the pane ID is captured correctly

#### Scenario: Pane ID not available
- **WHEN** `$TMUX_PANE` is not set or empty
- **THEN** the system reports that it is not running inside tmux

### Requirement: Watchlist persistence

The system SHALL persist the watchlist to `~/.config/tmon/watchlist.json`.

#### Scenario: First pane added
- **WHEN** adding a pane and no watchlist file exists
- **THEN** the config directory is created if needed
- **AND** a new watchlist.json file is created with the pane entry

#### Scenario: Subsequent pane added
- **WHEN** adding a pane and watchlist file already exists
- **THEN** the new pane is appended to the existing list
- **AND** existing entries are preserved

### Requirement: Subcommand routing

The system SHALL route commands based on arguments: no arguments launches the TUI, `add` runs the add-pane command.

#### Scenario: No arguments
- **WHEN** running `tmon` with no arguments
- **THEN** the TUI is launched

#### Scenario: Add subcommand
- **WHEN** running `tmon add`
- **THEN** the add-pane command is executed instead of the TUI

