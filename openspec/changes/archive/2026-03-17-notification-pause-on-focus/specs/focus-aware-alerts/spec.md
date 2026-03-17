## ADDED Requirements

### Requirement: Detect active tmux pane

The system SHALL be able to query which tmux pane is currently focused by the user.

#### Scenario: Get active pane ID
- **WHEN** the system queries the active tmux pane
- **THEN** the currently focused pane ID is returned (e.g., "%0")

#### Scenario: tmux command fails
- **WHEN** the active pane query fails (e.g., tmux not running)
- **THEN** an empty string is returned
- **AND** alert suppression is skipped (alerts fire normally)

### Requirement: Suppress alerts for focused pane

The system SHALL NOT fire alerts (sound or notification) for a pane that is currently focused by the user in tmux.

#### Scenario: Pane transitions while focused
- **WHEN** a pane transitions from Busy to Waiting
- **AND** that pane is the user's currently active tmux pane
- **THEN** no sound is played
- **AND** no desktop notification is sent

#### Scenario: Pane transitions while not focused
- **WHEN** a pane transitions from Busy to Waiting
- **AND** that pane is NOT the user's currently active tmux pane
- **THEN** alerts fire normally based on pane/global settings

#### Scenario: Status indicator still updates
- **WHEN** a pane transitions from Busy to Waiting
- **AND** that pane is the user's currently active tmux pane
- **THEN** the pane status indicator in the UI still updates to Waiting (green dot)
- **AND** only sound/notification alerts are suppressed
