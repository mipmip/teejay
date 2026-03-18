## ADDED Requirements

### Requirement: Scan detects panes running configured agents
The system SHALL scan all tmux panes and identify those whose foreground command matches a key in `config.detection.apps`. The match SHALL be an exact string comparison between the pane's current command and the app name.

#### Scenario: Pane running a known agent
- **WHEN** a tmux pane has foreground command "claude" and "claude" is a key in `config.detection.apps`
- **THEN** the scan SHALL identify this pane as an agent pane

#### Scenario: Pane running an unknown process
- **WHEN** a tmux pane has foreground command "vim" and "vim" is not in `config.detection.apps`
- **THEN** the scan SHALL NOT identify this pane as an agent pane

#### Scenario: Pane running a shell
- **WHEN** a tmux pane has foreground command "fish" and "fish" is not in `config.detection.apps`
- **THEN** the scan SHALL NOT identify this pane as an agent pane

### Requirement: Scan adds detected agent panes to watchlist
The system SHALL add each detected agent pane to the watchlist using `naming.GuessName()` for the display name. Panes already in the watchlist SHALL be skipped.

#### Scenario: New agent pane found
- **WHEN** a scan detects an agent pane that is not in the watchlist
- **THEN** the pane SHALL be added to the watchlist with an auto-generated name

#### Scenario: Agent pane already watched
- **WHEN** a scan detects an agent pane that is already in the watchlist
- **THEN** the pane SHALL be skipped and not duplicated

#### Scenario: Multiple agents across sessions
- **WHEN** three panes run "claude" across two sessions and one pane runs "aider"
- **THEN** all four panes SHALL be detected and added (if not already watched)

### Requirement: Scan reports results
The system SHALL report the scan outcome: total agent panes found, number added, and number skipped (already watched).

#### Scenario: Some panes added, some skipped
- **WHEN** a scan finds 5 agent panes, 3 are new and 2 are already watched
- **THEN** the result SHALL indicate 5 found, 3 added, 2 skipped

#### Scenario: No agent panes found
- **WHEN** a scan finds no panes running configured agents
- **THEN** the result SHALL indicate 0 found, 0 added, 0 skipped

### Requirement: TUI scan keybinding
The system SHALL provide an `s` keybinding in the main watchlist view that triggers the scan. The result SHALL be displayed as a status message.

#### Scenario: User presses s in main view
- **WHEN** the user presses `s` in the main watchlist view
- **THEN** the scan SHALL execute and display a status message with the results

#### Scenario: Scan result message
- **WHEN** a scan adds 3 panes and skips 1
- **THEN** the status message SHALL read "Scan: added 3 panes (1 already watched)"

#### Scenario: Scan finds nothing
- **WHEN** a scan finds no agent panes
- **THEN** the status message SHALL read "Scan: no agent panes found"

### Requirement: CLI scan command
The system SHALL provide a `tj scan` CLI command that runs the scan non-interactively and prints results to stdout.

#### Scenario: CLI scan with results
- **WHEN** the user runs `tj scan`
- **THEN** the command SHALL print the number of panes found and added, then exit

#### Scenario: CLI scan with custom config
- **WHEN** the user runs `tj scan -c /path/to/config.yaml`
- **THEN** the scan SHALL use the specified config file for agent detection
