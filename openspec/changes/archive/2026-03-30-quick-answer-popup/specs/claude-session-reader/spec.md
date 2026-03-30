## ADDED Requirements

### Requirement: Map tmux pane to Claude session

The system SHALL map a tmux pane running Claude Code to its session transcript file via the PID chain.

#### Scenario: Successful pane-to-session mapping
- **WHEN** a tmux pane has a child process with command name "claude"
- **AND** a session file exists at `~/.claude/sessions/<claude_pid>.json`
- **THEN** the system SHALL extract the `sessionId` and `cwd` from that file

#### Scenario: Claude process not found
- **WHEN** a tmux pane has no child process named "claude"
- **THEN** the session lookup SHALL return no result

#### Scenario: Session file missing
- **WHEN** a tmux pane has a claude child process
- **AND** no session file exists for that PID
- **THEN** the session lookup SHALL return no result

### Requirement: Locate transcript file from session ID

The system SHALL derive the transcript file path from the session ID and working directory.

#### Scenario: Transcript file found
- **WHEN** a session ID and cwd are known
- **AND** the project directory `~/.claude/projects/<project-hash>/` exists
- **AND** a `.jsonl` file matching the session ID exists in that directory
- **THEN** the system SHALL return the transcript file path

#### Scenario: Project hash derivation
- **WHEN** deriving the project hash from a cwd like `/home/pim/cVibeCoding/teejay`
- **THEN** the project hash SHALL be the path with `/` replaced by `-` and prefixed with `-` (e.g., `-home-pim-cVibeCoding-teejay`)

### Requirement: Read last assistant message from transcript

The system SHALL read the last assistant message from a Claude session transcript to determine the current prompt state.

#### Scenario: Last message is tool_use
- **WHEN** the last assistant entry in the transcript has `stop_reason: "tool_use"`
- **THEN** the system SHALL extract all `tool_use` content blocks with their `name` and `input` fields

#### Scenario: Last message is end_turn
- **WHEN** the last assistant entry in the transcript has `stop_reason: "end_turn"`
- **THEN** the system SHALL report that the agent is at the main prompt

#### Scenario: Efficient reading of large transcripts
- **WHEN** reading a transcript file
- **THEN** the system SHALL seek to the tail of the file (last ~64KB) rather than parsing the entire file
- **AND** SHALL scan backwards from the tail to find the last assistant entry
