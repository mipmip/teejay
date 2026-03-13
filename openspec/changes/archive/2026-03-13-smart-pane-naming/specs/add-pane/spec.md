## MODIFIED Requirements

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
