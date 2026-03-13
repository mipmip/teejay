# activity-detection Specification

## Purpose
TBD - created by archiving change pane-activity-detection. Update Purpose after archive.
## Requirements
### Requirement: Detect content changes via hash comparison

The system SHALL detect when pane content has changed by comparing SHA256 hashes.

#### Scenario: Content has changed
- **WHEN** the current pane content hash differs from the previous hash
- **THEN** the pane status is set to Running

#### Scenario: Content is unchanged
- **WHEN** the current pane content hash matches the previous hash
- **THEN** the pane status is NOT set to Running based on this check alone

### Requirement: Detect prompt patterns indicating waiting for input

The system SHALL detect when a pane is waiting for user input by matching known prompt patterns.

#### Scenario: Claude Code prompt detected
- **WHEN** pane content contains "No, and tell Claude what to do differently"
- **THEN** the pane status is set to Ready

#### Scenario: Aider prompt detected
- **WHEN** pane content contains "(Y)es/(N)o"
- **THEN** the pane status is set to Ready

#### Scenario: No prompt detected
- **WHEN** pane content does not match any known prompt patterns
- **THEN** the pane status is NOT set to Ready based on this check alone

### Requirement: Track idle state after inactivity

The system SHALL set pane status to Idle when content has been stable and no prompt is detected.

#### Scenario: Pane becomes idle
- **WHEN** pane content hash has not changed for 2 seconds
- **AND** no prompt pattern is detected
- **THEN** the pane status is set to Idle

#### Scenario: Pane resumes activity
- **WHEN** a pane in Idle state has content changes
- **THEN** the pane status is set to Running

### Requirement: Display status indicator in pane list

The system SHALL display a visual status indicator next to each pane in the list.

#### Scenario: Running pane indicator
- **WHEN** a pane has status Running
- **THEN** a running indicator is displayed (e.g., spinner or "●")

#### Scenario: Ready pane indicator
- **WHEN** a pane has status Ready
- **THEN** a ready indicator is displayed (e.g., "?" or different color)

#### Scenario: Idle pane indicator
- **WHEN** a pane has status Idle
- **THEN** an idle indicator is displayed (e.g., "○" or muted color)

