## MODIFIED Requirements

### Requirement: Switch to previewed pane

The TUI SHALL allow users to interact with the currently previewed pane by pressing Enter. When running inside tmux 3.2+, this opens a popup proxy. Otherwise, it switches directly to the pane. TJ SHALL remain open to allow continued monitoring.

#### Scenario: Open proxy popup when running in tmux 3.2+
- **WHEN** user presses Enter with a pane selected
- **AND** TJ is running inside a tmux session
- **AND** tmux version is 3.2 or higher
- **THEN** a tmux popup opens running `tj proxy <pane-id>`
- **AND** TJ remains open underneath the popup

#### Scenario: Fall back to switch when tmux < 3.2
- **WHEN** user presses Enter with a pane selected
- **AND** TJ is running inside a tmux session
- **AND** tmux version is below 3.2
- **THEN** tmux switches to the selected pane's session/window/pane
- **AND** TJ remains open and continues monitoring

#### Scenario: Switch when not in tmux
- **WHEN** user presses Enter with a pane selected
- **AND** TJ is NOT running inside a tmux session
- **THEN** a message is displayed indicating that switching requires running inside tmux
- **AND** TJ does NOT exit

#### Scenario: No pane selected
- **WHEN** user presses Enter
- **AND** no pane is selected (empty watchlist)
- **THEN** nothing happens
