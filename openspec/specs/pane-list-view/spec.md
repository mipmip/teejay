# pane-list-view Specification

## Purpose
TBD - created by archiving change tui-pane-list. Update Purpose after archive.
## Requirements
### Requirement: Display pane list

The TUI SHALL display all watched panes in a vertical list.

#### Scenario: Panes exist in watchlist
- **WHEN** the TUI starts with panes in the watchlist
- **THEN** each pane is displayed as a list item
- **AND** the list shows the pane ID (e.g., "%0")
- **AND** the list shows when the pane was added

#### Scenario: No panes in watchlist
- **WHEN** the TUI starts with an empty watchlist
- **THEN** a message is displayed indicating no panes are being watched
- **AND** the message suggests how to add panes (e.g., "Run 'tmon add' in a tmux pane")

### Requirement: Keyboard navigation

The TUI SHALL support keyboard navigation through the pane list.

#### Scenario: Navigate down
- **WHEN** the user presses down arrow or 'j'
- **THEN** the selection moves to the next pane in the list

#### Scenario: Navigate up
- **WHEN** the user presses up arrow or 'k'
- **THEN** the selection moves to the previous pane in the list

#### Scenario: Quit application
- **WHEN** the user presses 'q' or Ctrl+C
- **THEN** the application exits

### Requirement: Load watchlist on startup

The TUI SHALL load the watchlist from disk when starting.

#### Scenario: Watchlist loads successfully
- **WHEN** the TUI starts
- **THEN** panes from ~/.config/tmon/watchlist.json are loaded and displayed

#### Scenario: Watchlist file missing
- **WHEN** the TUI starts and no watchlist file exists
- **THEN** an empty list is displayed (same as empty watchlist)

