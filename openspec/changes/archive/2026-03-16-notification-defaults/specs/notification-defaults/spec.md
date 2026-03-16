## ADDED Requirements

### Requirement: Global alert defaults in configuration

The system SHALL support global default settings for `sound_on_ready` and `notify_on_ready` in the configuration file.

#### Scenario: Default configuration values
- **WHEN** no config file exists
- **THEN** global defaults are `sound_on_ready: false` and `notify_on_ready: false`

#### Scenario: User configures global defaults
- **WHEN** user sets `alerts.sound_on_ready: true` in config.yaml
- **THEN** panes without explicit settings use sound alerts

### Requirement: Per-pane override of global defaults

Per-pane settings in watchlist.json SHALL override the global defaults.

#### Scenario: Pane has explicit setting
- **WHEN** a pane has `sound_on_ready` explicitly set to false
- **AND** global default is true
- **THEN** the pane uses its explicit setting (false)

#### Scenario: Pane has no explicit setting
- **WHEN** a pane has no `sound_on_ready` setting (null/missing)
- **AND** global default is true
- **THEN** the pane uses the global default (true)

### Requirement: New panes inherit global defaults

New panes added to the watchlist SHALL inherit global default behavior without explicit settings.

#### Scenario: Adding new pane with global defaults enabled
- **WHEN** user adds a new pane to watchlist
- **AND** global `notify_on_ready` is true
- **THEN** the new pane has no explicit setting
- **AND** the pane uses the global default (true)

### Requirement: UI indicates override state

The UI SHALL clearly indicate whether a pane uses default or explicit settings.

#### Scenario: Pane using default
- **WHEN** a pane has no explicit setting
- **THEN** the UI shows an indicator that default is in use

#### Scenario: Pane with explicit override
- **WHEN** a pane has an explicit setting
- **THEN** the UI shows the explicit enabled or disabled state

### Requirement: Toggle cycles through states

Toggling a pane's alert setting SHALL cycle through: default → enabled → disabled → default.

#### Scenario: Toggle from default to enabled
- **WHEN** user toggles a pane that uses default
- **THEN** the pane setting becomes explicitly enabled

#### Scenario: Toggle from enabled to disabled
- **WHEN** user toggles a pane that is explicitly enabled
- **THEN** the pane setting becomes explicitly disabled

#### Scenario: Toggle from disabled to default
- **WHEN** user toggles a pane that is explicitly disabled
- **THEN** the pane setting becomes unset (uses default)
