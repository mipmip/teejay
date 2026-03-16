## ADDED Requirements

### Requirement: Help flag displays usage information

The CLI SHALL display usage information when invoked with `--help` or `-h` flag.

#### Scenario: User runs tj --help
- **WHEN** user runs `tj --help`
- **THEN** the CLI displays usage information including available commands and flags
- **THEN** the CLI exits with code 0

#### Scenario: User runs tj -h
- **WHEN** user runs `tj -h`
- **THEN** the CLI displays the same usage information as `--help`
- **THEN** the CLI exits with code 0

### Requirement: Help text includes all commands

The help output SHALL list all available commands with descriptions.

#### Scenario: Commands shown in help
- **WHEN** user views help output
- **THEN** the output includes the `add` command with description
- **THEN** the output includes the `del` command with description

### Requirement: Help text includes all flags

The help output SHALL list all global flags with descriptions.

#### Scenario: Flags shown in help
- **WHEN** user views help output
- **THEN** the output includes `--help, -h` flag
- **THEN** the output includes `--version, -v` flag
- **THEN** the output includes `--config, -c` flag with path description
- **THEN** the output includes `--watchlist, -w` flag with path description

### Requirement: Help flag takes precedence

The `--help` flag SHALL be processed before other arguments when it appears anywhere in the argument list.

#### Scenario: Help with other arguments
- **WHEN** user runs `tj add --help`
- **THEN** the CLI displays help information (not add command)
- **THEN** the CLI exits with code 0
