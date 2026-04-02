## MODIFIED Requirements

### Requirement: Dynamic column calculation

In multi-column mode, the UI SHALL calculate the number of columns dynamically based on terminal width, a minimum column width of 45 characters, and the number of pane items.

#### Scenario: Wide terminal shows multiple columns
- **WHEN** the terminal width is 120 characters in multi-column mode
- **THEN** the number of columns SHALL be `floor(availableWidth / 45)`
- **AND** each column SHALL have equal width

#### Scenario: Columns capped to item count
- **WHEN** there are fewer items than the calculated number of columns
- **THEN** the number of columns SHALL equal the number of items
- **AND** each column SHALL be wider to fill the available space

#### Scenario: Narrow terminal shows single column
- **WHEN** the terminal width allows only one column (less than 90 characters)
- **THEN** the layout SHALL show a single full-width column of pane items

#### Scenario: Terminal resize recalculates columns
- **WHEN** the terminal is resized while in multi-column mode
- **THEN** the column count SHALL be recalculated on the next render
