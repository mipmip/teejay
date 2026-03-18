# vhs-showcase Specification

## Purpose
TBD - created by archiving change vhs-showcase-video. Update Purpose after archive.
## Requirements
### Requirement: VHS tape file exists

The project SHALL include a `demo.tape` file in the repository root that generates a showcase GIF using VHS.

#### Scenario: Tape file is valid
- **WHEN** running `vhs demo.tape`
- **THEN** a `demo.gif` file SHALL be generated without errors

### Requirement: Demo showcases core features

The generated demo SHALL visually demonstrate the following Teejay features in sequence:

#### Scenario: Shows main TUI with panes
- **WHEN** the demo starts
- **THEN** it SHALL show the Teejay TUI with at least one watched pane visible

#### Scenario: Shows pane browser
- **WHEN** the user presses 'a' in the demo
- **THEN** the session/pane browser popup SHALL be displayed

#### Scenario: Shows adding a pane
- **WHEN** the user selects a pane from the browser
- **THEN** the pane SHALL be added to the watchlist

#### Scenario: Shows status indicator
- **WHEN** a watched pane is displayed
- **THEN** the status indicator (busy/ready) SHALL be visible

### Requirement: Demo is appropriately sized

The generated GIF SHALL be optimized for GitHub README display.

#### Scenario: Reasonable dimensions
- **WHEN** the demo is generated
- **THEN** the terminal dimensions SHALL be approximately 1200x700 pixels

#### Scenario: Reasonable duration
- **WHEN** the demo plays
- **THEN** it SHALL complete in under 30 seconds

