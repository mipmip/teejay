## MODIFIED Requirements

### Requirement: Preview updates on selection

The preview panel SHALL update when the user selects a different pane AND automatically at regular intervals.

#### Scenario: Change selection
- **WHEN** the user navigates to a different pane in the list
- **THEN** the preview panel immediately updates to show the newly selected pane's content

#### Scenario: Automatic refresh
- **WHEN** the selected pane remains the same
- **THEN** the preview panel re-captures and displays content every 100ms
