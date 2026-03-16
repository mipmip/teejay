## ADDED Requirements

### Requirement: Show preview panel in pane browser

The system SHALL display a preview panel showing the currently selected pane's content when viewing panes in the browser.

#### Scenario: Preview shown when viewing panes
- **WHEN** user is viewing the pane list (not session list) in the browser
- **THEN** a preview panel is displayed next to the pane list
- **AND** the preview shows the content of the currently selected pane

#### Scenario: Preview not shown during session selection
- **WHEN** user is viewing the session list in the browser
- **THEN** no preview panel is displayed
- **AND** only the session list is shown

### Requirement: Update preview on navigation

The system SHALL update the preview panel content when the user navigates to a different pane.

#### Scenario: Navigate to different pane
- **WHEN** user presses up/down arrow to select a different pane
- **THEN** the preview panel updates to show the newly selected pane's content

#### Scenario: Enter pane list from session
- **WHEN** user presses Enter on a session to view its panes
- **THEN** the first pane in the list is selected
- **AND** the preview panel shows that pane's content

### Requirement: Handle preview capture errors

The system SHALL gracefully handle errors when capturing pane content for preview.

#### Scenario: Pane capture fails
- **WHEN** capturing pane content for preview fails
- **THEN** the preview panel displays an error message
- **AND** the browser remains functional for selection
