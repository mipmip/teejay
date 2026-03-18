## MODIFIED Requirements

### Requirement: Watchlist items have visual separation

Items in the watchlist SHALL have visual separation between them. Each item SHALL take 3 lines: title, description (breadcrumb), and a margin line for spacing.

#### Scenario: Item spacing
- **WHEN** multiple items are displayed in the watchlist
- **THEN** each item SHALL be separated by a 1-line margin

#### Scenario: Description shows breadcrumb instead of plain process
- **WHEN** an item is displayed in the watchlist
- **THEN** the description line SHALL show the breadcrumb trail (`session > window : process`) instead of only the process name
