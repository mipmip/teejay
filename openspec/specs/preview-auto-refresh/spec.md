# preview-auto-refresh Specification

## Purpose
TBD - created by archiving change auto-update-preview. Update Purpose after archive.
## Requirements
### Requirement: Preview auto-refresh ticker

The TUI SHALL automatically refresh the preview panel content at regular intervals.

#### Scenario: Auto-refresh starts on launch
- **WHEN** the TUI starts with panes in the watchlist
- **THEN** the preview panel begins auto-refreshing at 100ms intervals

#### Scenario: Preview updates without user action
- **WHEN** the selected pane's content changes in tmux
- **THEN** the preview panel reflects the new content within 100ms
- **AND** no user input is required

#### Scenario: No refresh on empty watchlist
- **WHEN** the TUI starts with an empty watchlist
- **THEN** no refresh ticker is started

### Requirement: Refresh respects modal states

The preview refresh SHALL pause during modal operations to avoid visual disruption.

#### Scenario: No refresh during edit mode
- **WHEN** the user is in edit mode (renaming a pane)
- **THEN** the preview content does not refresh

#### Scenario: No refresh during delete confirmation
- **WHEN** the user is confirming a delete operation
- **THEN** the preview content does not refresh

#### Scenario: Refresh resumes after modal closes
- **WHEN** the user exits edit or delete mode
- **THEN** preview auto-refresh resumes on the next tick

