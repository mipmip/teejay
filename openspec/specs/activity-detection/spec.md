# activity-detection Specification

## Purpose
TBD - created by archiving change pane-activity-detection. Update Purpose after archive.
## Requirements
### Requirement: Detect content changes via hash comparison

The system SHALL detect when pane content has changed by comparing SHA256 hashes.

#### Scenario: Content has changed
- **WHEN** the current pane content hash differs from the previous hash
- **THEN** the pane status is set to Busy

#### Scenario: Content is unchanged
- **WHEN** the current pane content hash matches the previous hash
- **THEN** the pane status remains Busy (unless prompt is detected)

### Requirement: Detect prompt patterns indicating waiting for input

The system SHALL detect when a pane is waiting for user input by matching configurable prompt patterns, with app-specific overrides.

#### Scenario: App-specific pattern detected
- **WHEN** pane is running application "claude"
- **AND** config has patterns for "claude"
- **AND** pane content matches an app-specific pattern
- **THEN** the pane status is set to Waiting

#### Scenario: App-specific patterns replace globals
- **WHEN** pane is running application "claude"
- **AND** config has patterns for "claude"
- **THEN** only app-specific patterns are checked (global patterns ignored)

#### Scenario: Global pattern detected
- **WHEN** pane is running an application without app-specific config
- **AND** pane content matches a global pattern
- **THEN** the pane status is set to Waiting

#### Scenario: Prompt ending detected
- **WHEN** the last non-empty line ends with a configured prompt ending character
- **THEN** the pane status is set to Waiting

#### Scenario: Waiting string detected
- **WHEN** pane content contains a configured waiting string
- **THEN** the pane status is set to Waiting

#### Scenario: No pattern match
- **WHEN** pane content does not match any configured patterns
- **THEN** pattern matching does not trigger Waiting (idle timeout may still apply)

### Requirement: Display status indicator in pane list

The system SHALL display a visual status indicator next to each pane in the list.

#### Scenario: Busy pane indicator
- **WHEN** a pane has status Busy
- **THEN** an animated spinner is displayed

#### Scenario: Waiting pane indicator
- **WHEN** a pane has status Waiting
- **THEN** a green "●" indicator is displayed

### Requirement: Detect idle state via content stability

The system SHALL detect when a pane is idle by tracking how long content has remained unchanged.

#### Scenario: Content stable for idle timeout
- **WHEN** pane content hash has not changed
- **AND** time since last change exceeds configured idle_timeout
- **THEN** the pane status is set to Waiting

#### Scenario: Content changed recently
- **WHEN** pane content hash changed
- **THEN** the last change timestamp is updated
- **AND** the pane status is set to Busy (unless pattern matches)

#### Scenario: Idle timeout disabled
- **WHEN** idle_timeout is set to 0
- **THEN** idle detection is disabled
- **AND** only pattern matching determines Waiting state

### Requirement: Track per-pane state for idle detection

The system SHALL maintain per-pane state including content hash and last change timestamp.

#### Scenario: First content capture
- **WHEN** a pane is first monitored
- **THEN** its content hash is recorded
- **AND** last change time is set to current time

#### Scenario: Content update
- **WHEN** pane content is captured
- **AND** hash differs from previous
- **THEN** hash is updated
- **AND** last change time is updated to current time

### Requirement: Pass application name to monitor

The system SHALL provide the current foreground application name when updating pane status.

#### Scenario: Update with app name
- **WHEN** monitor.Update is called
- **THEN** the pane's current command (app name) is provided
- **AND** app-specific patterns can be matched

