## MODIFIED Requirements

### Requirement: Watchlist items display with background styling

The main watchlist panel SHALL render each pane item with a styled background using lipgloss. Unselected items SHALL have a dark grey background (#333333). The selected item SHALL have a lighter grey background (#555555). The pane title text SHALL be rendered in bold.

#### Scenario: Unselected item background
- **WHEN** an item in the watchlist is not selected
- **THEN** the item SHALL be rendered with a dark grey (#333333) background

#### Scenario: Selected item background
- **WHEN** an item in the watchlist is selected
- **THEN** the item SHALL be rendered with a lighter grey (#555555) background

#### Scenario: Pane title is bold
- **WHEN** a pane item is rendered in the watchlist
- **THEN** the title text SHALL be rendered with bold styling
