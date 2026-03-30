## ADDED Requirements

### Requirement: Toggle between layouts

The UI SHALL support toggling between "default" (list + preview) and "multi-column" (columns only) layout modes via the `v` keybinding.

#### Scenario: Toggle to multi-column layout
- **WHEN** the user presses `v` in default layout mode
- **THEN** the preview panel SHALL be hidden
- **AND** the pane items SHALL be displayed in multiple columns filling the available width

#### Scenario: Toggle back to default layout
- **WHEN** the user presses `v` in multi-column layout mode
- **THEN** the layout SHALL return to the default list + preview split

#### Scenario: Help text shows toggle key
- **WHEN** the help footer is displayed
- **THEN** `v` SHALL be listed as the layout toggle keybinding

### Requirement: Dynamic column calculation

In multi-column mode, the UI SHALL calculate the number of columns dynamically based on terminal width and a minimum column width of 30 characters.

#### Scenario: Wide terminal shows multiple columns
- **WHEN** the terminal width is 120 characters in multi-column mode
- **THEN** the number of columns SHALL be `floor(availableWidth / 30)`
- **AND** each column SHALL have equal width

#### Scenario: Narrow terminal shows single column
- **WHEN** the terminal width allows only one column (less than 60 characters)
- **THEN** the layout SHALL show a single full-width column of pane items

#### Scenario: Terminal resize recalculates columns
- **WHEN** the terminal is resized while in multi-column mode
- **THEN** the column count SHALL be recalculated on the next render

### Requirement: Column fill order

In multi-column mode, pane items SHALL fill columns top-to-bottom, then left-to-right.

#### Scenario: Items distributed across columns
- **WHEN** there are 8 pane items and 3 columns
- **THEN** column 1 SHALL contain items 1-3, column 2 SHALL contain items 4-6, column 3 SHALL contain items 7-8

### Requirement: Navigation in multi-column mode

In multi-column mode, up/down keys SHALL navigate within a column and left/right keys SHALL navigate between columns.

#### Scenario: Navigate down within column
- **WHEN** the user presses down arrow in multi-column mode
- **THEN** the selection moves to the next item below in the same column

#### Scenario: Navigate up within column
- **WHEN** the user presses up arrow in multi-column mode
- **THEN** the selection moves to the previous item above in the same column

#### Scenario: Navigate to next column
- **WHEN** the user presses right arrow in multi-column mode
- **THEN** the selection moves to the same row in the next column to the right
- **AND** if the target position has no item, the selection moves to the last item in that column

#### Scenario: Navigate to previous column
- **WHEN** the user presses left arrow in multi-column mode
- **THEN** the selection moves to the same row in the previous column to the left

#### Scenario: Wrap at column boundaries
- **WHEN** the user presses right arrow on the rightmost column
- **THEN** the selection SHALL NOT wrap (stays in place)

#### Scenario: Selected item is highlighted
- **WHEN** an item is selected in multi-column mode
- **THEN** the item SHALL be visually highlighted regardless of which column it appears in
