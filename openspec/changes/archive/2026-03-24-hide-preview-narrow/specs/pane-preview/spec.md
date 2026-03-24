## ADDED Requirements

### Requirement: Hide preview panel on narrow terminals
The main view SHALL hide the preview panel when the calculated sidebar width (30% of terminal width minus borders) would be less than 25 characters, giving the full terminal width to the watchlist sidebar.

#### Scenario: Narrow terminal hides preview
- **WHEN** the terminal width results in a sidebar width less than 25 characters at the 30% split
- **THEN** the preview panel SHALL be hidden and the sidebar SHALL use the full terminal width minus borders

#### Scenario: Wide terminal shows preview
- **WHEN** the terminal width results in a sidebar width of 25 characters or more at the 30% split
- **THEN** both sidebar (30%) and preview (70%) panels SHALL be displayed

#### Scenario: Resizing toggles preview
- **WHEN** the terminal is resized across the breakpoint threshold
- **THEN** the preview panel SHALL appear or disappear accordingly on the next render
