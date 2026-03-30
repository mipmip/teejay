## MODIFIED Requirements

### Requirement: Display status indicator in pane list

The system SHALL display a visual status indicator next to each pane in the list, with a distinct indicator for panes waiting with an actionable question.

#### Scenario: Busy pane indicator
- **WHEN** a pane has status Busy
- **THEN** an animated spinner is displayed

#### Scenario: Waiting pane indicator (idle)
- **WHEN** a pane has status Waiting
- **AND** the prompt type is `FreeInput` or `Unknown`
- **THEN** a green "●" indicator is displayed

#### Scenario: Waiting pane indicator (question)
- **WHEN** a pane has status Waiting
- **AND** the prompt type is `Permission`, `Question`, or `Choice`
- **THEN** a yellow "?" indicator is displayed
