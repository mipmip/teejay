## ADDED Requirements

### Requirement: Boolean flags for alert settings

The CLI SHALL support flags to override alert settings from the config file.

#### Scenario: Enable sound via flag
- **WHEN** the user runs `tj --sound`
- **THEN** sound alerts SHALL be enabled regardless of config file value

#### Scenario: Disable sound via flag
- **WHEN** the user runs `tj --no-sound`
- **THEN** sound alerts SHALL be disabled regardless of config file value

#### Scenario: Enable notifications via flag
- **WHEN** the user runs `tj --notify`
- **THEN** desktop notifications SHALL be enabled regardless of config file value

#### Scenario: Disable notifications via flag
- **WHEN** the user runs `tj --no-notify`
- **THEN** desktop notifications SHALL be disabled regardless of config file value

### Requirement: Boolean flags for display settings

The CLI SHALL support flags to override display settings from the config file.

#### Scenario: Start with activity sort
- **WHEN** the user runs `tj --sort-activity`
- **THEN** the TUI SHALL start with activity sort enabled

#### Scenario: Start with watchlist sort
- **WHEN** the user runs `tj --sort-watchlist`
- **THEN** the TUI SHALL start with watchlist order (default)

#### Scenario: Start in column layout
- **WHEN** the user runs `tj --columns`
- **THEN** the TUI SHALL start in multi-column layout mode

#### Scenario: Enable recency color
- **WHEN** the user runs `tj --recency-color`
- **THEN** the recency color gradient SHALL be enabled

#### Scenario: Disable recency color
- **WHEN** the user runs `tj --no-recency-color`
- **THEN** the recency color gradient SHALL be disabled

### Requirement: Flags override config file

CLI flags SHALL take precedence over config file values, but only when explicitly specified.

#### Scenario: Flag overrides config
- **WHEN** the config file has `alerts.sound_on_ready: true`
- **AND** the user runs `tj --no-sound`
- **THEN** sound alerts SHALL be disabled

#### Scenario: Unspecified flag preserves config
- **WHEN** the config file has `alerts.sound_on_ready: true`
- **AND** the user runs `tj` without any sound flags
- **THEN** sound alerts SHALL remain enabled (from config)

### Requirement: Picker mode flag

The CLI SHALL support a `--picker` flag that changes Enter behavior to switch to the selected pane and quit the application.

#### Scenario: Picker mode enabled
- **WHEN** the user runs `tj --picker`
- **AND** presses Enter on a selected pane
- **THEN** tmux SHALL switch to the selected pane
- **AND** the application SHALL quit

#### Scenario: Picker mode disabled (default)
- **WHEN** the user runs `tj` without `--picker`
- **AND** presses Enter on a selected pane
- **THEN** tmux SHALL switch to the selected pane
- **AND** the application SHALL continue running

### Requirement: Help text includes all flags

The `--help` output SHALL document all available flags.

#### Scenario: Help shows new flags
- **WHEN** the user runs `tj --help`
- **THEN** the output SHALL list all boolean flags with descriptions

### Requirement: Multiple flags can be combined

The CLI SHALL support combining multiple flags in a single invocation.

#### Scenario: Multiple flags
- **WHEN** the user runs `tj --columns --sort-activity --sound`
- **THEN** all three overrides SHALL be applied
