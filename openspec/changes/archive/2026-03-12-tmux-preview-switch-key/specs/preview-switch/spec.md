## ADDED Requirements

### Requirement: Switch to previewed pane

The TUI SHALL allow users to switch tmux to the currently previewed pane by pressing Enter.

#### Scenario: Switch when running in tmux
- **WHEN** user presses Enter with a pane selected
- **AND** tmon is running inside a tmux session
- **THEN** tmux switches to the selected pane's session/window/pane
- **AND** tmon exits

#### Scenario: Switch when not in tmux
- **WHEN** user presses Enter with a pane selected
- **AND** tmon is NOT running inside a tmux session
- **THEN** a message is displayed indicating that switching requires running inside tmux
- **AND** tmon does NOT exit

#### Scenario: No pane selected
- **WHEN** user presses Enter
- **AND** no pane is selected (empty watchlist)
- **THEN** nothing happens
