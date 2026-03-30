## MODIFIED Requirements

### Requirement: Split-panel layout

The TUI SHALL display a split-panel layout with the pane list on the left and preview on the right when in default layout mode.

#### Scenario: Layout displays correctly
- **WHEN** the TUI starts with panes in the watchlist
- **THEN** the pane list is displayed on the left side (~30% width)
- **AND** the preview panel is displayed on the right side (~70% width)

#### Scenario: Empty watchlist layout
- **WHEN** the TUI starts with an empty watchlist
- **THEN** the empty state message is displayed (no split panel needed)

### Requirement: Hide preview panel on narrow terminals
The main view SHALL hide the preview panel when the calculated sidebar width (30% of terminal width minus borders) would be less than 25 characters, giving the full terminal width to the watchlist sidebar.

#### Scenario: Narrow terminal hides preview
- **WHEN** the terminal width results in a sidebar width less than 25 characters at the 30% split
- **THEN** the preview panel SHALL be hidden and the sidebar SHALL use the full terminal width minus borders

#### Scenario: Wide terminal shows preview
- **WHEN** the terminal width results in a sidebar width of 25 characters or more at the 30% split
- **AND** the layout mode is default
- **THEN** both sidebar (30%) and preview (70%) panels SHALL be displayed

#### Scenario: Resizing toggles preview
- **WHEN** the terminal is resized across the breakpoint threshold
- **THEN** the preview panel SHALL appear or disappear accordingly on the next render

#### Scenario: Multi-column mode hides preview
- **WHEN** the layout mode is multi-column
- **THEN** the preview panel SHALL NOT be displayed regardless of terminal width
