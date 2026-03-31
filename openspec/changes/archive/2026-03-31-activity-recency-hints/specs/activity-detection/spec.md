## ADDED Requirements

### Requirement: Expose last activity time

The monitor SHALL expose the last content change timestamp for each pane.

#### Scenario: Query last change time
- **WHEN** `LastChangeTime(paneID)` is called for a monitored pane
- **THEN** the timestamp of the last content change for that pane SHALL be returned

#### Scenario: Unknown pane
- **WHEN** `LastChangeTime(paneID)` is called for a pane not yet monitored
- **THEN** the zero time SHALL be returned
