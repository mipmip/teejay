# pane-preview Specification

## Purpose
TBD - created by archiving change tui-pane-preview. Update Purpose after archive.
## Requirements
### Requirement: Split-panel layout

The TUI SHALL display a split-panel layout with the pane list on the left and preview on the right when in default layout mode.

#### Scenario: Layout displays correctly
- **WHEN** the TUI starts with panes in the watchlist
- **THEN** the pane list is displayed on the left side (~30% width)
- **AND** the preview panel is displayed on the right side (~70% width)

#### Scenario: Empty watchlist layout
- **WHEN** the TUI starts with an empty watchlist
- **THEN** the empty state message is displayed (no split panel needed)

### Requirement: Pane content capture

The system SHALL capture tmux pane content using `tmux capture-pane`.

#### Scenario: Capture pane content
- **WHEN** a pane is selected in the list
- **THEN** the system runs `tmux capture-pane -p -t <pane-id>`
- **AND** the output is displayed in the preview panel

#### Scenario: Pane no longer exists
- **WHEN** capturing content for a pane that no longer exists
- **THEN** an error message is displayed in the preview panel
- **AND** the TUI remains functional

### Requirement: Preview updates on selection

The preview panel SHALL update when the user selects a different pane AND automatically at regular intervals.

#### Scenario: Change selection
- **WHEN** the user navigates to a different pane in the list
- **THEN** the preview panel immediately updates to show the newly selected pane's content

#### Scenario: Automatic refresh
- **WHEN** the selected pane remains the same
- **THEN** the preview panel re-captures and displays content every 100ms

#### Scenario: Switch to pane
- **WHEN** the user presses Enter on a selected pane
- **THEN** tmux switches to that pane (if running in tmux)
- **AND** tmon exits

### Requirement: Scrollable preview

The preview panel SHALL be scrollable when content exceeds the viewport.

#### Scenario: Scroll preview content
- **WHEN** the pane content is longer than the preview viewport
- **THEN** the user can scroll using Page Up/Down or arrow keys when focused on preview

### Requirement: Preview title shows pane name

The preview panel title SHALL display the pane's display name instead of the raw pane ID.

#### Scenario: Pane with custom name
- **WHEN** a pane with custom name "My Project" is selected
- **THEN** the preview title shows "Preview: My Project"

#### Scenario: Pane without custom name
- **WHEN** a pane without a custom name (ID "%5") is selected
- **THEN** the preview title shows "Preview: %5"

#### Scenario: Title updates on selection change
- **WHEN** the user navigates from pane "Frontend" to pane "Backend"
- **THEN** the preview title updates to "Preview: Backend"

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

### Requirement: Preview in horizontal orientation

The preview panel SHALL support rendering below the content area (horizontal orientation) in addition to the existing side-by-side orientation.

#### Scenario: Below-preview uses full width
- **WHEN** the preview is rendered in horizontal orientation (multi-column mode)
- **THEN** the preview panel SHALL use the full terminal width minus borders

#### Scenario: Below-preview uses remaining height
- **WHEN** the preview is rendered in horizontal orientation
- **THEN** the preview panel height SHALL fill the remaining vertical space after the column grid and footer

#### Scenario: Visual consistency with side preview
- **WHEN** the below-preview is rendered
- **THEN** it SHALL use the same border style and title format as the side preview ("Preview: \<pane name\>")

### Requirement: Global preview toggle

The preview panel SHALL be globally toggleable via config and CLI flag.

#### Scenario: Preview disabled in default layout
- **WHEN** `display.show_preview` is `false`
- **THEN** the default layout SHALL show a full-width pane list without the right-side preview panel

#### Scenario: Preview disabled in multi-column layout
- **WHEN** `display.show_preview` is `false`
- **THEN** the multi-column layout SHALL NOT show the bottom preview panel regardless of available space

#### Scenario: Preview enabled (default)
- **WHEN** `display.show_preview` is `true` (or unspecified)
- **THEN** preview panels SHALL render as before (when space allows)

#### Scenario: CLI flag overrides config
- **WHEN** config has `display.show_preview: true`
- **AND** the user runs `tj --no-preview`
- **THEN** the preview SHALL be hidden

#### Scenario: CLI flag enables preview
- **WHEN** config has `display.show_preview: false`
- **AND** the user runs `tj --preview`
- **THEN** the preview SHALL be shown (when space allows)

