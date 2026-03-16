# native-sounds Specification

## Purpose
TBD - created by archiving change native-sound-selection. Update Purpose after archive.
## Requirements
### Requirement: Five selectable notification sounds

The system SHALL provide 5 distinct notification sounds to choose from.

#### Scenario: Available sound types
- **WHEN** the system lists available sounds
- **THEN** the following sound types are available: "chime", "bell", "ping", "pop", "ding"

### Requirement: Play selected sound type

The system SHALL play the specified sound type when triggered.

#### Scenario: Play chime sound
- **WHEN** a sound alert is triggered with sound_type "chime"
- **THEN** the chime sound is played

#### Scenario: Play bell sound
- **WHEN** a sound alert is triggered with sound_type "bell"
- **THEN** the bell sound is played

#### Scenario: Play ping sound
- **WHEN** a sound alert is triggered with sound_type "ping"
- **THEN** the ping sound is played

#### Scenario: Play pop sound
- **WHEN** a sound alert is triggered with sound_type "pop"
- **THEN** the pop sound is played

#### Scenario: Play ding sound
- **WHEN** a sound alert is triggered with sound_type "ding"
- **THEN** the ding sound is played

### Requirement: Default to chime sound

The system SHALL use "chime" as the default sound type when none is specified.

#### Scenario: No sound type specified
- **WHEN** a sound alert is triggered
- **AND** no sound_type is configured
- **THEN** the chime sound is played

### Requirement: Fallback on audio failure

The system SHALL fall back to terminal bell if native audio playback fails.

#### Scenario: Audio initialization fails
- **WHEN** native audio playback cannot be initialized
- **THEN** the system logs a warning
- **AND** falls back to terminal bell for alerts

### Requirement: Embedded sound files

The system SHALL embed all sound files in the binary.

#### Scenario: Sound available without external files
- **WHEN** the application is run
- **THEN** all 5 sounds are playable without any external sound files

