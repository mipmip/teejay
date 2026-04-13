# config-init Specification

## Purpose
TBD - created by archiving change config-init-command. Update Purpose after archive.
## Requirements
### Requirement: Init command creates config file

The `tj init` command SHALL create a configuration file at the default config path (`~/.config/teejay/config.yaml`), creating parent directories as needed.

#### Scenario: Config directory does not exist
- **WHEN** user runs `tj init`
- **AND** `~/.config/teejay/` does not exist
- **THEN** the directory is created
- **AND** the config file is written

#### Scenario: Config directory exists but no config file
- **WHEN** user runs `tj init`
- **AND** `~/.config/teejay/` exists but `config.yaml` does not
- **THEN** the config file is written to `~/.config/teejay/config.yaml`

### Requirement: Interactive wizard prompts for settings

The init command SHALL prompt the user for key settings via stdin, with sensible defaults shown in brackets.

#### Scenario: Sound alerts prompt
- **WHEN** the wizard runs
- **THEN** it SHALL ask whether to enable sound alerts when a pane becomes ready
- **AND** the default answer SHALL be "no"

#### Scenario: Desktop notifications prompt
- **WHEN** the wizard runs
- **THEN** it SHALL ask whether to enable desktop notifications when a pane becomes ready
- **AND** the default answer SHALL be "no"

#### Scenario: Layout prompt
- **WHEN** the wizard runs
- **THEN** it SHALL ask whether the default layout is single-column or multi-column
- **AND** the default answer SHALL be "single-column"

#### Scenario: Sort order prompt
- **WHEN** the wizard runs
- **THEN** it SHALL ask whether to sort by activity or watchlist order
- **AND** the default answer SHALL be "watchlist order"

#### Scenario: User accepts defaults by pressing Enter
- **WHEN** the user presses Enter without typing at a prompt
- **THEN** the default value for that setting is used

### Requirement: Handle existing config file

When a config file already exists at the target path, the init command SHALL prompt the user before proceeding.

#### Scenario: User chooses to overwrite
- **WHEN** user runs `tj init`
- **AND** a config file already exists
- **AND** user chooses "overwrite"
- **THEN** the existing file is replaced with the new config

#### Scenario: User chooses to backup
- **WHEN** user runs `tj init`
- **AND** a config file already exists
- **AND** user chooses "backup"
- **THEN** the existing file is renamed to `config.yaml.bak`
- **AND** the new config file is written

#### Scenario: User chooses to cancel
- **WHEN** user runs `tj init`
- **AND** a config file already exists
- **AND** user chooses "cancel"
- **THEN** no files are modified
- **AND** a message confirms no changes were made

### Requirement: Generated config file format

The generated config file SHALL be valid YAML with comments explaining each setting.

#### Scenario: File contains user-selected values
- **WHEN** the config file is written
- **THEN** it SHALL contain the settings chosen during the wizard under the correct YAML keys (`alerts.sound_on_ready`, `alerts.notify_on_ready`, `display.layout_mode`, `display.sort_by_activity`)

#### Scenario: File includes comments
- **WHEN** the config file is written
- **THEN** each setting SHALL have a comment explaining what it does

#### Scenario: File references advanced configuration
- **WHEN** the config file is written
- **THEN** it SHALL include a comment pointing users to `tj --help` or the example config for advanced options like detection patterns

### Requirement: Init command registered in CLI

The `init` subcommand SHALL be recognized by the CLI parser alongside existing commands (`add`, `del`, `scan`).

#### Scenario: Running tj init
- **WHEN** user runs `tj init`
- **THEN** the init wizard starts

#### Scenario: Init shown in help
- **WHEN** user runs `tj --help`
- **THEN** `init` appears in the list of available commands

