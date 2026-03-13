# pane-foreground-process Specification

## Purpose
TBD - created by archiving change show-foreground-process. Update Purpose after archive.
## Requirements
### Requirement: Description shows foreground process

The pane item description line SHALL display the current foreground process running in the tmux pane.

#### Scenario: Process is displayed
- **WHEN** a pane is displayed in the watchlist
- **THEN** the description line SHALL show the current foreground command name

#### Scenario: Pane ID is visible
- **WHEN** a pane is displayed in the watchlist
- **THEN** the description line SHALL include the pane ID for reference

### Requirement: Process updates dynamically

The displayed process SHALL update as the foreground process changes in the tmux pane.

#### Scenario: Process change is reflected
- **WHEN** the foreground process in a watched pane changes
- **THEN** the description SHALL update to show the new process within the refresh interval

#### Scenario: Graceful handling of errors
- **WHEN** fetching the current process fails
- **THEN** the last known process SHALL be displayed (no UI breakage)

