# activity-detection Delta Spec

## Changes

### MODIFY: Rename status states

Replace status names throughout:
- `Running` → `Busy`
- `Ready` → `Waiting`
- `Idle` → (removed)

### REMOVE: Requirement "Track idle state after inactivity"

Remove the entire requirement for tracking idle state. The idle counter and 2-second threshold are no longer needed.

**Reason:** With only two states (Busy/Waiting), we no longer need to track how long content has been stable. If no prompt is detected, the pane is simply Busy.

### MODIFY: Requirement "Detect content changes via hash comparison"

Update scenarios to use new state names:

#### Scenario: Content has changed
- **WHEN** the current pane content hash differs from the previous hash
- **THEN** the pane status is set to Busy

#### Scenario: Content is unchanged
- **WHEN** the current pane content hash matches the previous hash
- **THEN** the pane status remains Busy (unless prompt is detected)

### MODIFY: Requirement "Detect prompt patterns indicating waiting for input"

Update scenarios to use new state names:

#### Scenario: Claude Code prompt detected
- **WHEN** pane content contains "No, and tell Claude what to do differently"
- **THEN** the pane status is set to Waiting

#### Scenario: Aider prompt detected
- **WHEN** pane content contains "(Y)es/(N)o"
- **THEN** the pane status is set to Waiting

#### Scenario: No prompt detected
- **WHEN** pane content does not match any known prompt patterns
- **THEN** the pane status is set to Busy

### MODIFY: Requirement "Display status indicator in pane list"

Update to two states only:

#### Scenario: Busy pane indicator
- **WHEN** a pane has status Busy
- **THEN** an animated spinner is displayed

#### Scenario: Waiting pane indicator
- **WHEN** a pane has status Waiting
- **THEN** a green "●" indicator is displayed

#### REMOVE: Scenario "Idle pane indicator"
This scenario is removed as Idle state no longer exists.
