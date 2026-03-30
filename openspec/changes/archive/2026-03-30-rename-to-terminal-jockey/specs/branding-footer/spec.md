## MODIFIED Requirements

### Requirement: Display branded app name in footer

The UI SHALL display "Terminal Jockey" text in the bottom-right corner of the screen with neon-style coloring.

#### Scenario: Branding visible on normal terminal
- **WHEN** the terminal width is 80 columns or more
- **THEN** "Terminal Jockey" is displayed in the bottom-right corner with bright neon styling

#### Scenario: Branding hidden on small terminal
- **WHEN** the terminal width is less than 80 columns
- **THEN** the branding footer is not displayed to preserve space

### Requirement: Display version number in footer

The UI SHALL display the current version number alongside the app branding.

#### Scenario: Version displayed with branding
- **WHEN** the branding footer is visible
- **THEN** the version number is displayed below or next to "Terminal Jockey"

#### Scenario: Dev version display
- **WHEN** the app is running in development mode (version = "dev")
- **THEN** the version shows as "dev"

### Requirement: Global alert status in branding footer
The branding footer SHALL display the global alert configuration status as compact symbols after the version number. The display format SHALL be `Terminal Jockey v0.x.x ♪ ✉` where each symbol is colored according to the global config state.

#### Scenario: Both alerts enabled globally
- **WHEN** global config has `sound_on_ready: true` and `notify_on_ready: true`
- **THEN** the footer SHALL show `♪` in green and `✉` in yellow after the version

#### Scenario: Both alerts disabled globally
- **WHEN** global config has `sound_on_ready: false` and `notify_on_ready: false`
- **THEN** the footer SHALL show `♪` in dim gray and `✉` in dim gray after the version

#### Scenario: Mixed alert state globally
- **WHEN** global config has `sound_on_ready: true` and `notify_on_ready: false`
- **THEN** the footer SHALL show `♪` in green and `✉` in dim gray after the version

#### Scenario: Systray hidden on narrow terminals
- **WHEN** the terminal width is less than 80 columns
- **THEN** the alert status symbols SHALL NOT be displayed (same as branding footer)
