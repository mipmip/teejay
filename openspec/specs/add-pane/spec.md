# add-pane Specification

## Purpose
TBD - created by archiving change add-pane-command. Update Purpose after archive.
## Requirements
### Requirement: Add pane subcommand

The system SHALL provide a `tj add` subcommand that adds the current tmux pane to the watchlist with an auto-guessed or user-provided name.

#### Scenario: Successfully add pane with auto-guessed name
- **WHEN** running `tj add` from within a tmux pane running a distinctive command (e.g., `nvim`)
- **THEN** the current pane ID is added to the watchlist
- **AND** the pane is automatically named based on the running command
- **AND** a success message including the assigned name is printed to stdout

#### Scenario: Add pane with generic command prompts for name
- **WHEN** running `tj add` from within a tmux pane where all names are generic
- **THEN** the user is prompted to enter a name for the pane
- **AND** the pane is added with the user-provided name
- **AND** a success message is printed to stdout

#### Scenario: User can skip naming prompt
- **WHEN** prompted for a name and the user enters empty input
- **THEN** the pane is added with no name (will display as pane ID)

#### Scenario: Add pane outside tmux
- **WHEN** running `tj add` outside of a tmux session
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

