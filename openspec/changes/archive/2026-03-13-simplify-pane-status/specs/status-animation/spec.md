# status-animation Delta Spec

## Changes

### MODIFY: Rename status states

Replace status names throughout:
- `Running` → `Busy`
- `Ready` → `Waiting`

### MODIFY: Requirement "Animated spinner for running panes"

Rename to "Animated spinner for busy panes":

#### Scenario: Spinner animates while busy
- **WHEN** a pane has Busy status
- **THEN** the status indicator displays a braille spinner character
- **AND** the spinner cycles through frames on each UI tick (100ms)

#### Scenario: Spinner stops when prompt detected
- **WHEN** a Busy pane transitions to Waiting
- **THEN** the spinner stops and the green indicator is shown

### MODIFY: Requirement "Green indicator for ready panes"

Rename to "Green indicator for waiting panes":

#### Scenario: Waiting pane shows green circle
- **WHEN** a pane has Waiting status (waiting for input)
- **THEN** the status indicator displays a green "●" character

### REMOVE: Requirement "Static indicator for idle panes"

Remove the entire requirement. Idle state no longer exists.

**Reason:** With the simplified two-state model, there is no Idle state. Panes are either Busy (animated) or Waiting (green).
