# status-animation Specification

## Purpose
TBD - created by archiving change status-indicator-animation. Update Purpose after archive.
## Requirements
### Requirement: Animated spinner for busy panes

The pane list SHALL display an animated spinner indicator for panes in Busy status.

#### Scenario: Spinner animates while busy
- **WHEN** a pane has Busy status
- **THEN** the status indicator displays a braille spinner character
- **AND** the spinner cycles through frames on each UI tick (100ms)

#### Scenario: Spinner stops when prompt detected
- **WHEN** a Busy pane transitions to Waiting
- **THEN** the spinner stops and the green indicator is shown

### Requirement: Green indicator for waiting panes

The pane list SHALL display a green circle indicator for panes in Waiting status.

#### Scenario: Waiting pane shows green circle
- **WHEN** a pane has Waiting status (waiting for input)
- **THEN** the status indicator displays a green "●" character

