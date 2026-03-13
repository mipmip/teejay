# browser-item-delegate Specification

## Purpose
TBD - created by archiving change browser-item-styling. Update Purpose after archive.
## Requirements
### Requirement: Item background styling

Browser popup items SHALL be rendered with background colors to distinguish them visually.

#### Scenario: Unselected item appearance
- **WHEN** an item in the browser list is not selected
- **THEN** it is displayed with a dark grey background

#### Scenario: Selected item appearance
- **WHEN** an item in the browser list is selected
- **THEN** it is displayed with a lighter grey background

### Requirement: Item spacing

Browser popup items SHALL have visual spacing between them.

#### Scenario: Items have margin
- **WHEN** multiple items are displayed in the browser list
- **THEN** there is a 1-line gap between each item

### Requirement: Full-width clickable area

The entire item background area SHALL be clickable for mouse selection.

#### Scenario: Click anywhere on item
- **WHEN** user clicks anywhere within an item's background area
- **THEN** that item becomes selected

### Requirement: Consistent styling across lists

Both session list and pane list in the browser popup SHALL use the same item styling.

#### Scenario: Session list styling
- **WHEN** the browser shows the session selection list
- **THEN** items are displayed with background colors and spacing

#### Scenario: Pane list styling
- **WHEN** the browser shows the pane selection list for a session
- **THEN** items are displayed with background colors and spacing

