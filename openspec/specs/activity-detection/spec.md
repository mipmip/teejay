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

The system SHALL detect when a pane is waiting for user input by matching known prompt patterns.

#### Scenario: Claude Code prompt detected
- **WHEN** pane content contains "No, and tell Claude what to do differently"
- **THEN** the pane status is set to Waiting

#### Scenario: Aider prompt detected
- **WHEN** pane content contains "(Y)es/(N)o"
- **THEN** the pane status is set to Waiting

#### Scenario: No prompt detected
- **WHEN** pane content does not match any known prompt patterns
- **THEN** the pane status is set to Busy

### Requirement: Display status indicator in pane list

The system SHALL display a visual status indicator next to each pane in the list.

#### Scenario: Busy pane indicator
- **WHEN** a pane has status Busy
- **THEN** an animated spinner is displayed

#### Scenario: Waiting pane indicator
- **WHEN** a pane has status Waiting
- **THEN** a green "●" indicator is displayed

