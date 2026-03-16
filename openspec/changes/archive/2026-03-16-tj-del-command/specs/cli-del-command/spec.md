## ADDED Requirements

### Requirement: Delete current pane from watchlist
The system SHALL provide a `tj del` CLI command that removes the current tmux pane from the watchlist.

#### Scenario: Successfully delete watched pane
- **WHEN** user runs `tj del` inside a tmux pane that is in the watchlist
- **THEN** the pane is removed from the watchlist
- **AND** a success message is displayed showing the pane's name: "Removed '<name>' from watchlist"

#### Scenario: Pane not in watchlist
- **WHEN** user runs `tj del` inside a tmux pane that is NOT in the watchlist
- **THEN** an informational message is displayed: "Pane %s is not being watched"
- **AND** no error exit code is returned

#### Scenario: Not running in tmux
- **WHEN** user runs `tj del` outside of tmux (TMUX_PANE not set)
- **THEN** an error is displayed: "cannot delete pane: not running inside tmux"
- **AND** the program exits with error code 1

### Requirement: Named feedback for delete operations
The system SHALL display the pane's human-readable name in deletion feedback messages, using the stored name or guessing from tmux metadata.

#### Scenario: Pane has stored name
- **WHEN** deleting a pane that has a custom name stored in the watchlist
- **THEN** the feedback message uses the stored name

#### Scenario: Pane has no stored name
- **WHEN** deleting a pane that has no custom name in the watchlist
- **THEN** the system uses `naming.GuessName()` to determine a display name from tmux metadata
- **AND** the feedback message uses the guessed name
