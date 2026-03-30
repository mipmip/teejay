## ADDED Requirements

### Requirement: Determine prompt type from structured or scraped data

The system SHALL classify the prompt state of a Waiting pane into one of: `Permission`, `Question`, `Choice`, `FreeInput`, or `Unknown`.

#### Scenario: Claude pane with tool permission pending
- **WHEN** a Claude pane is Waiting
- **AND** the last assistant transcript entry has `stop_reason: "tool_use"` with a tool other than `AskUserQuestion`
- **THEN** the prompt type SHALL be `Permission`
- **AND** the tool name and key input parameters SHALL be extracted from the transcript
- **AND** the actual menu options SHALL be scraped from the rendered pane content

#### Scenario: Claude pane with AskUserQuestion
- **WHEN** a Claude pane is Waiting
- **AND** the last assistant transcript entry has `stop_reason: "tool_use"` with tool `AskUserQuestion`
- **AND** the input contains options
- **THEN** the prompt type SHALL be `Choice`
- **AND** the question text SHALL be extracted from the transcript
- **AND** the actual menu options SHALL be scraped from the rendered pane content

#### Scenario: Claude pane with free-text AskUserQuestion
- **WHEN** a Claude pane is Waiting
- **AND** the last assistant transcript entry has `stop_reason: "tool_use"` with tool `AskUserQuestion`
- **AND** the input contains no preset options
- **THEN** the prompt type SHALL be `Question`
- **AND** the question text SHALL be extracted

#### Scenario: Claude pane idle at main prompt
- **WHEN** a Claude pane is Waiting
- **AND** the last assistant transcript entry has `stop_reason: "end_turn"`
- **THEN** the prompt type SHALL be `FreeInput`

#### Scenario: Non-Claude pane waiting
- **WHEN** a non-Claude pane is Waiting
- **THEN** the prompt type SHALL be `FreeInput`
- **AND** a basic context extraction from captured pane content MAY be attempted

#### Scenario: Transcript unreadable or missing
- **WHEN** the Claude session transcript cannot be read or parsed
- **THEN** the prompt type SHALL fall back to `FreeInput`

### Requirement: Scrape menu options from rendered pane content

For Permission and Choice prompts, the system SHALL scrape the actual menu options from the terminal output rather than hardcoding them. This ensures the popup displays exactly what the agent shows.

#### Scenario: Numbered menu detected
- **WHEN** the captured pane content contains lines matching numbered menu items (e.g., `❯ 1. Yes`, `  2. No`)
- **THEN** the options SHALL be extracted in display order
- **AND** any question/context text above the menu SHALL be extracted

#### Scenario: No menu detected
- **WHEN** the captured pane content does not contain a recognizable numbered menu
- **THEN** fallback options from the transcript parsing SHALL be used

#### Scenario: Menu options override transcript options
- **WHEN** both transcript-based and screen-scraped options are available
- **THEN** the screen-scraped options SHALL take precedence

### Requirement: Periodic prompt check for waiting panes

The system SHALL periodically check the prompt state of Waiting panes to determine if they have an actionable question.

#### Scenario: Prompt check interval
- **WHEN** a pane transitions to Waiting state
- **THEN** a prompt check SHALL be performed
- **AND** subsequent checks SHALL occur at a regular interval (approximately every 2 seconds)

#### Scenario: Prompt check only for waiting panes
- **WHEN** a pane is in Busy state
- **THEN** no prompt check SHALL be performed for that pane

#### Scenario: Prompt check runs asynchronously
- **WHEN** a periodic prompt check is triggered
- **THEN** the recognition SHALL run in a background goroutine
- **AND** the UI thread SHALL NOT be blocked during recognition
