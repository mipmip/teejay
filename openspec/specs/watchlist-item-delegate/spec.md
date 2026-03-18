# watchlist-item-delegate Specification

## Purpose
TBD - created by archiving change watchlist-item-styling. Update Purpose after archive.
## Requirements
### Requirement: Watchlist items display with background styling

The main watchlist panel SHALL render each pane item with a styled background using lipgloss. Unselected items SHALL have a dark grey background (#333333). The selected item SHALL have a lighter grey background (#555555).

#### Scenario: Unselected item background
- **WHEN** an item in the watchlist is not selected
- **THEN** the item SHALL be rendered with a dark grey (#333333) background

#### Scenario: Selected item background
- **WHEN** an item in the watchlist is selected
- **THEN** the item SHALL be rendered with a lighter grey (#555555) background

### Requirement: Watchlist items have visual separation

Items in the watchlist SHALL have visual separation between them. Each item SHALL take 3 lines: title, description (breadcrumb), and a margin line for spacing.

#### Scenario: Item spacing
- **WHEN** multiple items are displayed in the watchlist
- **THEN** each item SHALL be separated by a 1-line margin

#### Scenario: Description shows breadcrumb instead of plain process
- **WHEN** an item is displayed in the watchlist
- **THEN** the description line SHALL show the breadcrumb trail (`session > window : process`) instead of only the process name

### Requirement: Mouse click detection matches item height

Mouse clicks in the watchlist panel SHALL correctly identify the clicked item based on the 3-line item height.

#### Scenario: Click selects correct item
- **WHEN** user clicks on an item in the watchlist
- **THEN** the item at that position SHALL become selected based on itemHeight=3

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

