## ADDED Requirements

### Requirement: Temporary message display

The system SHALL support displaying temporary messages that auto-dismiss after a timeout.

#### Scenario: Show temporary message
- **WHEN** a temporary message is set (e.g., error or status notification)
- **THEN** the message is displayed in the footer area
- **AND** an auto-dismiss timer is started

#### Scenario: New message replaces existing
- **WHEN** a new temporary message is set while another is displayed
- **THEN** the new message replaces the old one
- **AND** the auto-dismiss timer resets for the new message

### Requirement: Auto-dismiss after timeout

The system SHALL automatically dismiss temporary messages after a timeout period.

#### Scenario: Message auto-dismissed after timeout
- **WHEN** a temporary message is displayed
- **AND** 3 seconds elapse without the message being manually dismissed
- **THEN** the message is automatically dismissed
- **AND** the normal footer is restored

### Requirement: Manual dismiss with Esc

The system SHALL allow users to dismiss temporary messages early by pressing Esc.

#### Scenario: User dismisses message with Esc
- **WHEN** a temporary message is displayed
- **AND** user presses Esc
- **THEN** the message is immediately dismissed
- **AND** the pending auto-dismiss timer is cancelled
- **AND** the normal footer is restored

### Requirement: Navigation does not dismiss

The system SHALL NOT dismiss temporary messages when users perform navigation actions.

#### Scenario: User navigates while message displayed
- **WHEN** a temporary message is displayed
- **AND** user performs navigation (up/down arrow keys)
- **THEN** the message remains visible
- **AND** the auto-dismiss timer continues unchanged

### Requirement: Not-in-tmux error uses temporary message

The system SHALL display the "not running inside tmux" error as a temporary message.

#### Scenario: Switch attempted outside tmux
- **WHEN** user attempts to switch panes while not running inside tmux
- **THEN** the error "Cannot switch: not running inside tmux" is shown as a temporary message
- **AND** the message auto-dismisses after the standard timeout
