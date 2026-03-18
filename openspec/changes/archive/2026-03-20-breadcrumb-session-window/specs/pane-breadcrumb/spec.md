## ADDED Requirements

### Requirement: Pane breadcrumb displays session, window, and process
The watchlist pane item description SHALL display a breadcrumb trail in the format `session > window : process`, where `session` is the tmux session name, `window` is the tmux window name, and `process` is the current foreground process.

#### Scenario: Full breadcrumb with all components
- **WHEN** a pane belongs to session "technative-docs", window "proposals", running process "claude"
- **THEN** the description SHALL display `technative-docs > proposals : claude`

#### Scenario: Breadcrumb without active process
- **WHEN** a pane belongs to session "main", window "dev", and has no foreground process (or process is empty)
- **THEN** the description SHALL display `main > dev`

#### Scenario: Breadcrumb updates on process change
- **WHEN** the foreground process in a watched pane changes from "claude" to "bash"
- **THEN** the breadcrumb SHALL update to reflect the new process on the next refresh tick
