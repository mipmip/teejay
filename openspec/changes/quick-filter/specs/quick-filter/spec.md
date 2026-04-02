## ADDED Requirements

### Requirement: Enter filter mode with /

The user SHALL be able to press `/` to enter filter mode with a text input.

#### Scenario: Activate filter
- **WHEN** the user presses `/`
- **THEN** a text input SHALL appear in the footer
- **AND** the text input SHALL be focused for typing

#### Scenario: Filter not activated in modal
- **WHEN** the user is in a modal (editing, deleting, browsing, configuring, quick-answer)
- **THEN** pressing `/` SHALL NOT activate the filter

### Requirement: Real-time filtering of pane list

As the user types in filter mode, the pane list SHALL be filtered in real-time.

#### Scenario: Filter matches name
- **WHEN** the user types "claude" in the filter
- **THEN** only panes whose name, session, window name, or command contain "claude" (case-insensitive) SHALL be shown

#### Scenario: No matches
- **WHEN** the filter query matches no panes
- **THEN** an empty list SHALL be shown

#### Scenario: Empty filter shows all
- **WHEN** the filter query is empty
- **THEN** all panes SHALL be shown

### Requirement: Confirm or cancel filter

The user SHALL be able to confirm the filter (keep it active) or cancel it (clear and show all).

#### Scenario: Confirm filter with Enter
- **WHEN** the user presses Enter in filter mode
- **THEN** the filter SHALL remain active
- **AND** the text input SHALL be dismissed
- **AND** normal navigation SHALL resume

#### Scenario: Cancel filter with Esc
- **WHEN** the user presses Esc in filter mode
- **THEN** the filter SHALL be cleared
- **AND** all panes SHALL be shown
- **AND** the text input SHALL be dismissed

#### Scenario: Re-enter filter mode
- **WHEN** a filter is active (confirmed)
- **AND** the user presses `/`
- **THEN** the filter input SHALL reappear with the current query for editing

### Requirement: Filter indicator in footer

When a filter is active (confirmed), the footer SHALL show the filter query as a reminder.

#### Scenario: Active filter shown
- **WHEN** a filter is active and the user is not in filter-edit mode
- **THEN** the footer SHALL show "Filter: <query>" with a hint to edit or clear
