## ADDED Requirements

### Requirement: Guess pane name from tmux metadata

The system SHALL provide a function to guess a meaningful name for a pane based on tmux metadata.

#### Scenario: Name from running command
- **WHEN** guessing a name for a pane running `nvim`
- **THEN** the guessed name is `nvim`

#### Scenario: Name from window name
- **WHEN** the pane command is generic (e.g., `zsh`) but window name is `api-server`
- **THEN** the guessed name is `api-server`

#### Scenario: Name from session name
- **WHEN** both pane command and window name are generic but session name is `project-x`
- **THEN** the guessed name is `project-x`

#### Scenario: All names are generic
- **WHEN** pane command is `bash`, window name is `0`, and session name is `main`
- **THEN** the function indicates the name is generic

### Requirement: Identify generic names

The system SHALL identify names that are too generic to be useful for identification.

#### Scenario: Shell names are generic
- **WHEN** checking if `bash`, `zsh`, `fish`, or `sh` are generic
- **THEN** they are identified as generic

#### Scenario: Tool names are generic
- **WHEN** checking if `tmux`, `screen`, `claude`, `opencode`, or `aider` are generic
- **THEN** they are identified as generic

#### Scenario: Numeric names are generic
- **WHEN** checking if `0`, `1`, `2` or similar single digits are generic
- **THEN** they are identified as generic

#### Scenario: Common tmux defaults are generic
- **WHEN** checking if `main`, `default`, `new`, or `window` are generic
- **THEN** they are identified as generic

#### Scenario: Distinctive names are not generic
- **WHEN** checking if `api-server`, `frontend`, `nvim`, or `cargo` are generic
- **THEN** they are NOT identified as generic

### Requirement: Fetch window name from tmux

The system SHALL fetch the window name as additional metadata for name guessing.

#### Scenario: Window name available
- **WHEN** listing pane information
- **THEN** the window name is included in the metadata
