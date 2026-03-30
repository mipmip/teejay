## ADDED Requirements

### Requirement: Only q and ctrl+c quit the application

The application SHALL only exit when the user presses `q` or `ctrl+c` in the main view. The `esc` key SHALL NOT cause the application to exit under any circumstances.

#### Scenario: Escape pressed in main view with no popup open
- **WHEN** user presses Escape in the main view
- **AND** no popup or overlay is open
- **THEN** the application SHALL NOT exit
- **AND** any temporary message SHALL be cleared

#### Scenario: Escape pressed after dismissing a popup
- **WHEN** user presses Escape to close a popup
- **AND** presses Escape again in the main view
- **THEN** the application SHALL NOT exit

#### Scenario: Q pressed in main view
- **WHEN** user presses `q` in the main view
- **AND** no popup or text input is active
- **THEN** the application SHALL exit
