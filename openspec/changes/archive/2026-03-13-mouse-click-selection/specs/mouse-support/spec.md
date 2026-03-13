## ADDED Requirements

### Requirement: Mouse mode enabled
The application SHALL enable mouse cell motion mode when starting.

#### Scenario: Program starts with mouse support
- **WHEN** tmon starts
- **THEN** mouse events are captured by the application

### Requirement: Click to select in main list
The main pane list SHALL allow mouse clicks to select items.

#### Scenario: Click pane in list
- **WHEN** user clicks on a pane item in the watched panes list
- **THEN** that pane becomes selected
- **AND** the preview updates to show the clicked pane

### Requirement: Click to select in browser popup
The browser popup lists (sessions and panes) SHALL allow mouse clicks to select items.

#### Scenario: Click session in browser
- **WHEN** user clicks on a session in the session list
- **THEN** that session becomes selected

#### Scenario: Click pane in browser
- **WHEN** user clicks on a pane in the pane list
- **THEN** that pane becomes selected

### Requirement: Click to select in configure popup
The configure popup menu SHALL allow mouse clicks to select menu items.

#### Scenario: Click menu item in configure
- **WHEN** user clicks on a menu item in the configure popup
- **THEN** that menu item becomes selected
