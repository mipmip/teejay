# status-animation Specification

## Purpose
TBD - created by archiving change status-indicator-animation. Update Purpose after archive.
## Requirements
### Requirement: Animated spinner for running panes

The pane list SHALL display an animated spinner indicator for panes in Running status.

#### Scenario: Spinner animates while running
- **WHEN** a pane has Running status
- **THEN** the status indicator displays a braille spinner character
- **AND** the spinner cycles through frames on each UI tick (100ms)

#### Scenario: Spinner stops when status changes
- **WHEN** a Running pane transitions to Ready or Idle
- **THEN** the spinner stops and the appropriate static indicator is shown

### Requirement: Green indicator for ready panes

The pane list SHALL display a green circle indicator for panes in Ready status.

#### Scenario: Ready pane shows green circle
- **WHEN** a pane has Ready status (waiting for input)
- **THEN** the status indicator displays a green "●" character

### Requirement: Static indicator for idle panes

The pane list SHALL display a static circle indicator for panes in Idle status.

#### Scenario: Idle pane shows empty circle
- **WHEN** a pane has Idle status
- **THEN** the status indicator displays "○"

