## ADDED Requirements

### Requirement: Elapsed time label for waiting panes

The browser list delegate SHALL display a compact elapsed-time label on the title row of each waiting pane, positioned right-aligned before the status indicator.

#### Scenario: Waiting pane shows elapsed time
- **WHEN** a pane has status Waiting
- **AND** `lastActivity` is not zero
- **THEN** a compact duration label (e.g., "3s", "14m", "2h", "1d") is displayed on the title row before the status indicator

#### Scenario: Busy pane shows no elapsed time
- **WHEN** a pane has status Busy
- **THEN** no elapsed-time label is displayed

#### Scenario: Pane with zero lastActivity
- **WHEN** a pane has status Waiting
- **AND** `lastActivity` is the zero time
- **THEN** no elapsed-time label is displayed

### Requirement: Compact duration format

The elapsed-time label SHALL use the single largest time unit in compact form.

#### Scenario: Seconds range
- **WHEN** elapsed time is less than 60 seconds
- **THEN** label shows seconds with "s" suffix (e.g., "3s", "45s")

#### Scenario: Minutes range
- **WHEN** elapsed time is 60 seconds or more but less than 60 minutes
- **THEN** label shows minutes with "m" suffix (e.g., "1m", "14m")

#### Scenario: Hours range
- **WHEN** elapsed time is 60 minutes or more but less than 24 hours
- **THEN** label shows hours with "h" suffix (e.g., "1h", "8h")

#### Scenario: Days range
- **WHEN** elapsed time is 24 hours or more
- **THEN** label shows days with "d" suffix (e.g., "1d", "3d")

### Requirement: Dim styling

The elapsed-time label SHALL be styled dim so it does not compete visually with the pane name or status indicator.

#### Scenario: Label is visually subdued
- **WHEN** the elapsed-time label is rendered
- **THEN** it uses a dimmed/muted text color
