## ADDED Requirements

### Requirement: Show preview below columns when space allows

In multi-column mode, the UI SHALL display the preview panel below the pane item grid when sufficient vertical space is available.

#### Scenario: Enough vertical space for preview
- **WHEN** the terminal height minus the grid height minus the footer height is 8 lines or more
- **THEN** the preview panel SHALL be rendered below the column grid
- **AND** the preview SHALL span the full terminal width

#### Scenario: Not enough vertical space
- **WHEN** the remaining vertical space below the grid is less than 8 lines
- **THEN** the preview panel SHALL NOT be displayed
- **AND** the multi-column layout SHALL render as before (columns only)

#### Scenario: Preview shows selected pane content
- **WHEN** the below-preview is visible
- **THEN** it SHALL display the content of the currently selected pane
- **AND** it SHALL update when the selection changes

#### Scenario: Terminal resize adjusts preview visibility
- **WHEN** the terminal is resized
- **THEN** the preview visibility SHALL be recalculated on the next render
- **AND** the preview SHALL appear or disappear based on the new available space
