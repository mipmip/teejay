## ADDED Requirements

### Requirement: Preview title shows pane name

The preview panel title SHALL display the pane's display name instead of the raw pane ID.

#### Scenario: Pane with custom name
- **WHEN** a pane with custom name "My Project" is selected
- **THEN** the preview title shows "Preview: My Project"

#### Scenario: Pane without custom name
- **WHEN** a pane without a custom name (ID "%5") is selected
- **THEN** the preview title shows "Preview: %5"

#### Scenario: Title updates on selection change
- **WHEN** the user navigates from pane "Frontend" to pane "Backend"
- **THEN** the preview title updates to "Preview: Backend"
