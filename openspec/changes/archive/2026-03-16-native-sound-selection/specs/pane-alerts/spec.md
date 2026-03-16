## ADDED Requirements

### Requirement: Configure sound type globally

The system SHALL allow configuring the default sound type in the global config.

#### Scenario: Set global sound type in config
- **WHEN** config.yaml contains `alerts.sound_type: "ping"`
- **THEN** panes without a per-pane sound_type override use "ping" sound

#### Scenario: Global default when not configured
- **WHEN** config.yaml does not specify `alerts.sound_type`
- **THEN** the default sound type "chime" is used

### Requirement: Configure sound type per pane

The system SHALL allow configuring the sound type per pane, overriding the global default.

#### Scenario: Per-pane sound type override
- **WHEN** a pane has sound_type set to "bell"
- **AND** global sound_type is "chime"
- **THEN** that pane uses "bell" sound for alerts

#### Scenario: Per-pane uses global default
- **WHEN** a pane has no sound_type configured (nil)
- **THEN** that pane uses the global default sound type

### Requirement: Store per-pane sound type

The watchlist SHALL store the sound_type setting per pane.

#### Scenario: Save sound type setting
- **WHEN** a pane's sound_type setting is changed
- **THEN** the setting is persisted in watchlist.json

#### Scenario: Load sound type setting
- **WHEN** the watchlist is loaded
- **THEN** each pane's sound_type setting is restored

## MODIFIED Requirements

### Requirement: Play sound on ready transition

The system SHALL play the configured sound type when a pane with sound enabled transitions to Ready status.

#### Scenario: Sound alert triggered
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has sound_on_ready enabled
- **THEN** the pane's effective sound type is played

#### Scenario: No sound when disabled
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has sound_on_ready disabled
- **THEN** no sound is played

#### Scenario: No sound on repeated ready
- **WHEN** a pane is already in Ready status
- **AND** the status is checked again (still Ready)
- **THEN** no sound is played (only on transition)
