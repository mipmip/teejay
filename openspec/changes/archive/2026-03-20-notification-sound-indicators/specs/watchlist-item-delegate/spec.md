## ADDED Requirements

### Requirement: Per-pane alert override indicators
Watchlist pane items SHALL display alert override indicators on the description line when the pane has explicit per-pane alert overrides. The indicators SHALL appear after the breadcrumb text, separated by two spaces.

#### Scenario: Pane with sound override enabled
- **WHEN** a pane has an explicit `sound_on_ready` override set to true
- **THEN** the description line SHALL show `♪` in green after the breadcrumb

#### Scenario: Pane with sound override disabled
- **WHEN** a pane has an explicit `sound_on_ready` override set to false
- **THEN** the description line SHALL show `♪` in dim gray after the breadcrumb

#### Scenario: Pane with notification override enabled
- **WHEN** a pane has an explicit `notify_on_ready` override set to true
- **THEN** the description line SHALL show `✉` in yellow after the breadcrumb

#### Scenario: Pane with both overrides
- **WHEN** a pane has both `sound_on_ready` and `notify_on_ready` overrides set
- **THEN** the description line SHALL show both `♪` and `✉` symbols after the breadcrumb, each colored by their respective state

### Requirement: No indicators for default-inherited panes
Watchlist pane items SHALL NOT display alert indicators when the pane has no explicit overrides and is inheriting global defaults.

#### Scenario: Pane using global defaults
- **WHEN** a pane has no explicit `sound_on_ready` or `notify_on_ready` overrides (both nil)
- **THEN** the description line SHALL show only the breadcrumb without any alert indicators
