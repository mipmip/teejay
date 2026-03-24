## ADDED Requirements

### Requirement: Golden file snapshot tests for delegate rendering
The project SHALL have golden-file snapshot tests that capture the full ANSI-styled output of `browserItemDelegate.Render()` for various pane item states.

#### Scenario: Unselected pane item rendering
- **WHEN** a pane item is rendered in unselected state
- **THEN** the output SHALL match the golden file `testdata/delegate_unselected.golden`

#### Scenario: Selected pane item rendering
- **WHEN** a pane item is rendered in selected state
- **THEN** the output SHALL match the golden file `testdata/delegate_selected.golden`

#### Scenario: Pane item with waiting status
- **WHEN** a pane item with Waiting status is rendered
- **THEN** the output SHALL match the golden file `testdata/delegate_waiting.golden`

#### Scenario: Pane item with alert overrides
- **WHEN** a pane item with alert override indicators is rendered
- **THEN** the output SHALL match the golden file `testdata/delegate_alerts.golden`

### Requirement: Golden file update mechanism
The snapshot tests SHALL support a `-update` flag to regenerate golden files when intentional visual changes are made.

#### Scenario: Update golden files
- **WHEN** tests are run with `go test -update`
- **THEN** the golden files SHALL be overwritten with the current rendered output

#### Scenario: Normal test run
- **WHEN** tests are run without `-update`
- **THEN** the rendered output SHALL be compared against existing golden files and fail on mismatch
