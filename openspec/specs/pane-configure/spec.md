# pane-configure Specification

## Purpose
TBD - created by archiving change pane-configure-popup. Update Purpose after archive.
## Requirements
### Requirement: Open configure popup

The TUI SHALL allow users to open a configure popup by pressing `c` on a selected pane.

#### Scenario: Open configure popup
- **WHEN** user presses `c` on a selected pane
- **THEN** a modal popup appears with configuration options
- **AND** the popup shows the pane name/ID in the title

#### Scenario: Configure on empty list
- **WHEN** user presses `c` with no panes in watchlist
- **THEN** nothing happens

### Requirement: Configure popup menu

The configure popup SHALL display a menu with Name, Sound on Ready, and Notify on Ready options.

#### Scenario: Display menu items
- **WHEN** the configure popup is open
- **THEN** it shows "Name: [current name]"
- **AND** it shows "Sound on Ready: [On/Off]"
- **AND** it shows "Notify on Ready: [On/Off]"

#### Scenario: Navigate menu
- **WHEN** user presses up/down arrows in configure popup
- **THEN** the selection moves between menu items

### Requirement: Edit name from configure popup

The user SHALL be able to edit the pane name from the configure popup.

#### Scenario: Enter name edit mode
- **WHEN** user selects "Name" and presses Enter
- **THEN** a text input appears with the current name
- **AND** the user can type a new name

#### Scenario: Save name
- **WHEN** user is editing name and presses Enter
- **THEN** the new name is saved to the watchlist
- **AND** the popup updates to show the new name

#### Scenario: Cancel name edit
- **WHEN** user is editing name and presses Escape
- **THEN** the edit is cancelled
- **AND** the original name is preserved

### Requirement: Toggle sound setting

The user SHALL be able to toggle sound alerts for a pane.

#### Scenario: Toggle sound on
- **WHEN** user selects "Sound on Ready" and presses Enter
- **AND** sound is currently off
- **THEN** sound is turned on
- **AND** the display updates to show "On"

#### Scenario: Toggle sound off
- **WHEN** user selects "Sound on Ready" and presses Enter
- **AND** sound is currently on
- **THEN** sound is turned off
- **AND** the display updates to show "Off"

### Requirement: Toggle notification setting

The user SHALL be able to toggle notification alerts for a pane.

#### Scenario: Toggle notify on
- **WHEN** user selects "Notify on Ready" and presses Enter
- **AND** notify is currently off
- **THEN** notify is turned on
- **AND** the display updates to show "On"

#### Scenario: Toggle notify off
- **WHEN** user selects "Notify on Ready" and presses Enter
- **AND** notify is currently on
- **THEN** notify is turned off
- **AND** the display updates to show "Off"

### Requirement: Close configure popup

The user SHALL be able to close the configure popup.

#### Scenario: Close with Escape
- **WHEN** user presses Escape in configure popup (not editing)
- **THEN** the popup closes
- **AND** all settings are saved

#### Scenario: Close with 'c' again
- **WHEN** user presses 'c' while configure popup is open
- **THEN** the popup closes

