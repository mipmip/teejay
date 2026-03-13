## ADDED Requirements

### Requirement: Watchlist items display with background styling

The main watchlist panel SHALL render each pane item with a styled background using lipgloss. Unselected items SHALL have a dark grey background (#333333). The selected item SHALL have a lighter grey background (#555555).

#### Scenario: Unselected item background
- **WHEN** an item in the watchlist is not selected
- **THEN** the item SHALL be rendered with a dark grey (#333333) background

#### Scenario: Selected item background
- **WHEN** an item in the watchlist is selected
- **THEN** the item SHALL be rendered with a lighter grey (#555555) background

### Requirement: Watchlist items have visual separation

Items in the watchlist SHALL have visual separation between them. Each item SHALL take 3 lines: title, description, and a margin line for spacing.

#### Scenario: Item spacing
- **WHEN** multiple items are displayed in the watchlist
- **THEN** each item SHALL be separated by a 1-line margin

### Requirement: Mouse click detection matches item height

Mouse clicks in the watchlist panel SHALL correctly identify the clicked item based on the 3-line item height.

#### Scenario: Click selects correct item
- **WHEN** user clicks on an item in the watchlist
- **THEN** the item at that position SHALL become selected based on itemHeight=3
