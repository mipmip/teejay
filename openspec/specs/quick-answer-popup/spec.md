# quick-answer-popup Specification

## Purpose
TBD - created by archiving change quick-answer-popup. Update Purpose after archive.
## Requirements
### Requirement: Open quick-answer popup with space key

The system SHALL open a quick-answer popup when the user presses `space` on a Waiting pane.

#### Scenario: Space on waiting pane
- **WHEN** the user presses `space`
- **AND** the selected pane is in Waiting state
- **THEN** a quick-answer popup SHALL be displayed with the detected prompt information

#### Scenario: Space on busy pane
- **WHEN** the user presses `space`
- **AND** the selected pane is in Busy state
- **THEN** no popup SHALL be displayed

#### Scenario: Space on empty list
- **WHEN** the user presses `space`
- **AND** no pane is selected
- **THEN** no popup SHALL be displayed

### Requirement: Permission prompt popup

The quick-answer popup SHALL display tool permission prompts with selectable options.

#### Scenario: Display permission prompt
- **WHEN** the popup opens for a `Permission` prompt
- **THEN** the popup SHALL show the tool name and key input parameters
- **AND** SHALL show selectable options: y (Allow once), n (Deny), a (Always allow)

#### Scenario: Select and send permission response
- **WHEN** the user selects an option and presses Enter
- **THEN** the selected key (y, n, or a) SHALL be sent to the pane via `tmux send-keys`

### Requirement: Choice prompt popup

The quick-answer popup SHALL display multiple-choice prompts with selectable options.

#### Scenario: Display choice prompt
- **WHEN** the popup opens for a `Choice` prompt
- **THEN** the popup SHALL show the question text
- **AND** SHALL show the numbered/lettered options as selectable items

#### Scenario: Select and send choice response
- **WHEN** the user selects a choice and presses Enter
- **THEN** the corresponding key (number or letter) SHALL be sent to the pane via `tmux send-keys`

### Requirement: Question and free-text popup

The quick-answer popup SHALL display a text input field for questions and free input.

#### Scenario: Display question prompt
- **WHEN** the popup opens for a `Question` prompt
- **THEN** the popup SHALL show the question text
- **AND** SHALL show a text input field

#### Scenario: Display free input prompt
- **WHEN** the popup opens for a `FreeInput` prompt
- **THEN** the popup SHALL show a text input field
- **AND** MAY show a minimal context line

#### Scenario: Send free-text response
- **WHEN** the user types text and presses Enter
- **THEN** the typed text SHALL be sent to the pane via `tmux send-keys` followed by Enter

### Requirement: Cancel popup without sending

The user SHALL be able to close the popup without sending any response.

#### Scenario: Cancel with Esc
- **WHEN** the user presses `Esc` while the popup is open
- **THEN** the popup SHALL close
- **AND** no keystrokes SHALL be sent to the target pane

### Requirement: Freshness check before sending

The system SHALL verify that the prompt is still active before sending a response.

#### Scenario: Prompt still active
- **WHEN** the user confirms a response
- **AND** re-capture confirms the pane is still Waiting
- **AND** (for Claude) the transcript shows the same tool_use ID
- **THEN** the response SHALL be sent

#### Scenario: Prompt expired
- **WHEN** the user confirms a response
- **AND** re-capture shows the pane is no longer Waiting (or the tool_use ID changed)
- **THEN** the response SHALL NOT be sent
- **AND** a "Prompt expired" message SHALL be shown
- **AND** the popup SHALL close

### Requirement: Send response via tmux send-keys

The system SHALL send user responses to the target pane using `tmux send-keys`, adapting the send method to the prompt type.

#### Scenario: Send permission or choice response
- **WHEN** sending a response for a Permission or Choice prompt
- **THEN** the system SHALL send arrow-down keypresses to navigate to the selected item followed by Enter to confirm
- **AND** the number of arrow-down presses SHALL equal the selected option's zero-based index

#### Scenario: Send free-text response
- **WHEN** sending a response for a Question or FreeInput prompt
- **THEN** the system SHALL execute `tmux send-keys -t <pane_id> <text> Enter`

