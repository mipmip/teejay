## ADDED Requirements

### Requirement: Display branded app name in footer

The UI SHALL display "Terminal Junkie" text in the bottom-right corner of the screen with neon-style coloring.

#### Scenario: Branding visible on normal terminal
- **WHEN** the terminal width is 80 columns or more
- **THEN** "Terminal Junkie" is displayed in the bottom-right corner with bright neon styling

#### Scenario: Branding hidden on small terminal
- **WHEN** the terminal width is less than 80 columns
- **THEN** the branding footer is not displayed to preserve space

### Requirement: Display version number in footer

The UI SHALL display the current version number alongside the app branding.

#### Scenario: Version displayed with branding
- **WHEN** the branding footer is visible
- **THEN** the version number is displayed below or next to "Terminal Junkie"

#### Scenario: Dev version display
- **WHEN** the app is running in development mode (version = "dev")
- **THEN** the version shows as "dev"
