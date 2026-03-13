## ADDED Requirements

### Requirement: README documents keyboard shortcuts

The README.md SHALL include a "Keyboard Shortcuts" section documenting all available keybindings.

#### Scenario: Shortcuts section exists
- **WHEN** a user reads README.md
- **THEN** they SHALL find a "Keyboard Shortcuts" section

### Requirement: Shortcuts organized by context

The shortcuts documentation SHALL group keybindings by their applicable context (main view, popups, etc.).

#### Scenario: Main view shortcuts documented
- **WHEN** a user reads the shortcuts section
- **THEN** they SHALL find shortcuts for: quit (q), add pane (a), configure (c), delete (d), switch pane (Enter), navigation (↑/↓)

#### Scenario: Browser popup shortcuts documented
- **WHEN** a user reads the shortcuts section
- **THEN** they SHALL find shortcuts for: select (Enter), back/close (Esc), quit (q)

#### Scenario: Configure popup shortcuts documented
- **WHEN** a user reads the shortcuts section
- **THEN** they SHALL find shortcuts for: navigate (↑/↓), toggle/edit (Enter/Space), close (Esc)
