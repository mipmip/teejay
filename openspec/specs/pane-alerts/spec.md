# pane-alerts Specification

## Purpose
TBD - created by archiving change pane-configure-popup. Update Purpose after archive.
## Requirements
### Requirement: Play sound on ready transition

The system SHALL play a sound when a pane with sound enabled transitions to Ready status.

#### Scenario: Sound alert triggered
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has sound_on_ready enabled
- **THEN** a terminal bell sound is played

#### Scenario: No sound when disabled
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has sound_on_ready disabled
- **THEN** no sound is played

#### Scenario: No sound on repeated ready
- **WHEN** a pane is already in Ready status
- **AND** the status is checked again (still Ready)
- **THEN** no sound is played (only on transition)

### Requirement: Send notification on ready transition

The system SHALL send a desktop notification when a pane with notify enabled transitions to Ready status.

#### Scenario: Notification triggered
- **WHEN** a pane transitions from Running to Ready
- **AND** the pane has notify_on_ready enabled
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

### Requirement: Store alert settings

The watchlist SHALL store sound and notification settings per pane.

#### Scenario: Save sound setting
- **WHEN** a pane's sound_on_ready setting is changed
- **THEN** the setting is persisted in watchlist.json

#### Scenario: Save notify setting
- **WHEN** a pane's notify_on_ready setting is changed
- **THEN** the setting is persisted in watchlist.json

#### Scenario: Load alert settings
- **WHEN** the watchlist is loaded
- **THEN** each pane's sound_on_ready and notify_on_ready settings are restored

