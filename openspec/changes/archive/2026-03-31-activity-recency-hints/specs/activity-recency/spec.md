## ADDED Requirements

### Requirement: Recency color gradient on waiting indicator

The waiting pane indicator SHALL display a color gradient based on how recently the pane was last active, fading from bright green to dim green over time.

#### Scenario: Just finished (0-10 seconds)
- **WHEN** a pane has been Waiting for 10 seconds or less
- **THEN** the `●` indicator SHALL be rendered in bright neon green (`#00FF00`)

#### Scenario: Recently finished (10-30 seconds)
- **WHEN** a pane has been Waiting for between 10 and 30 seconds
- **THEN** the `●` indicator SHALL be rendered in bright green (`#00DD00`)

#### Scenario: Moderately idle (30 seconds to 2 minutes)
- **WHEN** a pane has been Waiting for between 30 seconds and 2 minutes
- **THEN** the `●` indicator SHALL be rendered in medium green (`#00BB00`)

#### Scenario: Idle (2-5 minutes)
- **WHEN** a pane has been Waiting for between 2 and 5 minutes
- **THEN** the `●` indicator SHALL be rendered in dim green (`#009900`)

#### Scenario: Long idle (5+ minutes)
- **WHEN** a pane has been Waiting for more than 5 minutes
- **THEN** the `●` indicator SHALL be rendered in very dim green (`#006600`)

#### Scenario: Actionable prompt indicator unaffected
- **WHEN** a pane has a `?` indicator (actionable prompt)
- **THEN** the indicator SHALL remain yellow regardless of recency

#### Scenario: Recency color disabled
- **WHEN** the config has `display.recency_color: false`
- **THEN** all waiting indicators SHALL use the default green (`#00FF00`)

### Requirement: Toggle activity sort order

The user SHALL be able to toggle between watchlist order and activity-sorted order using the `o` keybind.

#### Scenario: Toggle to activity sort
- **WHEN** the user presses `o` while in watchlist order
- **THEN** panes SHALL be sorted with busy panes first (most recently active first), then waiting panes (most recently finished first)

#### Scenario: Toggle back to watchlist order
- **WHEN** the user presses `o` while in activity sort order
- **THEN** panes SHALL return to the original watchlist order

#### Scenario: Help text shows sort toggle
- **WHEN** the help footer is displayed
- **THEN** `o` SHALL be listed as the sort toggle keybinding

#### Scenario: Default sort from config
- **WHEN** the application starts
- **AND** config has `display.sort_by_activity: true`
- **THEN** the initial sort order SHALL be activity sort

### Requirement: Config options for display hints

The config file SHALL support a `display` section with options for recency color and default sort order.

#### Scenario: Default config values
- **WHEN** no display config is specified
- **THEN** `recency_color` SHALL default to `true`
- **AND** `sort_by_activity` SHALL default to `false`
