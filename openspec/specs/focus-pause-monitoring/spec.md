# focus-pause-monitoring Specification

## Purpose
TBD - created by archiving change suppress-focused-pane-status. Update Purpose after archive.
## Requirements
### Requirement: Pause status monitoring for focused pane
The system SHALL NOT update the status indicator for a pane that is currently focused by the user in tmux. The pane SHALL retain its last known status while focused.

#### Scenario: User typing in monitored pane
- **WHEN** the user is focused on a monitored pane and types input
- **THEN** the pane's status indicator SHALL remain unchanged (not flip to Busy)

#### Scenario: Pane status preserved while focused
- **WHEN** a pane was Waiting before the user focused it
- **THEN** the pane SHALL continue showing the Waiting indicator while focused

#### Scenario: Preview still updates while focused
- **WHEN** the user is focused on a monitored pane
- **THEN** the preview panel SHALL still show the live pane content

### Requirement: Grace period after defocus
The system SHALL wait 2 seconds after the user switches focus away from a pane before resuming status monitoring for that pane.

#### Scenario: User switches away from pane
- **WHEN** the user switches focus from pane A to another pane
- **THEN** pane A's status SHALL remain frozen for 2 seconds before monitoring resumes

#### Scenario: Grace period expires
- **WHEN** 2 seconds have elapsed since the user left a pane
- **THEN** the system SHALL resume normal status monitoring for that pane

#### Scenario: User returns during grace period
- **WHEN** the user switches back to a pane before the 2-second grace period expires
- **THEN** the pane SHALL remain paused (grace period resets)

