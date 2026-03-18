## ADDED Requirements

### Requirement: Global alert status in branding footer
The branding footer SHALL display the global alert configuration status as compact symbols after the version number. The display format SHALL be `Terminal Junkie v0.x.x ♪ ✉` where each symbol is colored according to the global config state.

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
