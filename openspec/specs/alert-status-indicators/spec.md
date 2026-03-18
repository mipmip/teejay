# alert-status-indicators Specification

## Purpose
TBD - created by archiving change notification-sound-indicators. Update Purpose after archive.
## Requirements
### Requirement: Sound indicator symbol
The system SHALL use `♪` as the visual indicator for sound alert configuration. The symbol SHALL be rendered in green (`#00FF00`) when sound is enabled and dim gray (`#555555`) when disabled.

#### Scenario: Sound enabled indicator
- **WHEN** sound alerts are enabled (globally or per-pane)
- **THEN** the `♪` symbol SHALL be rendered in green (`#00FF00`)

#### Scenario: Sound disabled indicator
- **WHEN** sound alerts are disabled (globally or per-pane)
- **THEN** the `♪` symbol SHALL be rendered in dim gray (`#555555`)

### Requirement: Notification indicator symbol
The system SHALL use `✉` as the visual indicator for notification configuration. The symbol SHALL be rendered in yellow (`#FFD700`) when notifications are enabled and dim gray (`#555555`) when disabled.

#### Scenario: Notification enabled indicator
- **WHEN** desktop notifications are enabled (globally or per-pane)
- **THEN** the `✉` symbol SHALL be rendered in yellow (`#FFD700`)

#### Scenario: Notification disabled indicator
- **WHEN** desktop notifications are disabled (globally or per-pane)
- **THEN** the `✉` symbol SHALL be rendered in dim gray (`#555555`)

