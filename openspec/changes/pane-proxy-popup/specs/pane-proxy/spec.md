## ADDED Requirements

### Requirement: Proxy displays target pane content

The proxy SHALL continuously capture and display the content of the target tmux pane, including ANSI escape sequences for colors and styling.

#### Scenario: Display pane content with colors
- **WHEN** the proxy is started with a valid pane ID
- **THEN** the proxy displays the pane's visible content
- **AND** ANSI colors and styling are preserved

#### Scenario: Content updates in real-time
- **WHEN** the target pane's content changes
- **THEN** the proxy display updates within 100ms

### Requirement: Proxy forwards keyboard input

The proxy SHALL forward all keyboard input to the target pane via tmux send-keys.

#### Scenario: Forward regular characters
- **WHEN** user types alphanumeric characters
- **THEN** those characters are sent to the target pane

#### Scenario: Forward special keys
- **WHEN** user presses Ctrl-C, arrow keys, or other special keys
- **THEN** the corresponding key sequence is sent to the target pane

#### Scenario: Forward escape sequences
- **WHEN** user presses Escape (single)
- **THEN** Escape is sent to the target pane

### Requirement: Proxy exit on double-Escape

The proxy SHALL exit when the user presses Escape twice within 500ms.

#### Scenario: Double-Escape exits
- **WHEN** user presses Escape twice within 500ms
- **THEN** the proxy exits cleanly

#### Scenario: Slow double-Escape does not exit
- **WHEN** user presses Escape, waits more than 500ms, then presses Escape again
- **THEN** both Escapes are forwarded to the target pane
- **AND** the proxy continues running

### Requirement: Proxy shows cursor position

The proxy SHALL position the cursor to match the target pane's cursor location.

#### Scenario: Cursor position matches target
- **WHEN** the target pane has cursor at position (x, y)
- **THEN** the proxy positions its cursor at the same (x, y)

### Requirement: Proxy CLI subcommand

The application SHALL provide a `proxy` subcommand that runs the proxy for a specified pane.

#### Scenario: Start proxy via CLI
- **WHEN** user runs `tj proxy %42`
- **THEN** the proxy starts for pane %42

#### Scenario: Invalid pane ID
- **WHEN** user runs `tj proxy` with a non-existent pane ID
- **THEN** an error message is displayed
- **AND** the proxy exits with non-zero status
