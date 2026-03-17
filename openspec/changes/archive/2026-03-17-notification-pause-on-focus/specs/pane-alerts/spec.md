## MODIFIED Requirements

### Requirement: Play sound on ready transition

The system SHALL play the configured sound type when a pane with sound enabled transitions to Ready status, unless the pane is currently focused by the user in tmux.

#### Scenario: Sound alert triggered
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has sound_on_ready enabled
- **AND** the pane is NOT the user's active tmux pane
- **THEN** the pane's effective sound type is played

#### Scenario: No sound when disabled
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has sound_on_ready disabled
- **THEN** no sound is played

#### Scenario: No sound on repeated ready
- **WHEN** a pane is already in Ready status
- **AND** the status is checked again (still Ready)
- **THEN** no sound is played (only on transition)

#### Scenario: No sound when pane is focused
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has sound_on_ready enabled
- **AND** the pane is the user's active tmux pane
- **THEN** no sound is played

### Requirement: Send notification on ready transition

The system SHALL send a desktop notification when a pane with notify enabled transitions to Ready status, unless the pane is currently focused by the user in tmux.

#### Scenario: Notification triggered
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has notify_on_ready enabled
- **AND** the pane is NOT the user's active tmux pane
- **THEN** a desktop notification is sent
- **AND** the notification shows the pane name

#### Scenario: No notification when disabled
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has notify_on_ready disabled
- **THEN** no notification is sent

#### Scenario: No notification on repeated ready
- **WHEN** a pane is already in Ready status
- **AND** the status is checked again (still Ready)
- **THEN** no notification is sent (only on transition)

#### Scenario: No notification when pane is focused
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has notify_on_ready enabled
- **AND** the pane is the user's active tmux pane
- **THEN** no notification is sent
