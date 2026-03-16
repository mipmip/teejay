## ADDED Requirements

### Requirement: Config path flag

The CLI SHALL accept `--config` or `-c` flag to specify an alternative config.yaml path.

#### Scenario: Config flag with valid path
- **WHEN** user runs `tj --config /path/to/config.yaml`
- **THEN** configuration is loaded from the specified path

#### Scenario: Config flag with short form
- **WHEN** user runs `tj -c /path/to/config.yaml`
- **THEN** configuration is loaded from the specified path

#### Scenario: Config flag with non-existent file
- **WHEN** user runs `tj --config /nonexistent/config.yaml`
- **THEN** default configuration values are used
- **AND** a warning is logged

### Requirement: Watchlist path flag

The CLI SHALL accept `--watchlist` or `-w` flag to specify an alternative watchlist.json path.

#### Scenario: Watchlist flag with valid path
- **WHEN** user runs `tj --watchlist /path/to/watchlist.json`
- **THEN** watchlist is loaded from the specified path
- **AND** saves write back to the same path

#### Scenario: Watchlist flag with short form
- **WHEN** user runs `tj -w /path/to/watchlist.json`
- **THEN** watchlist is loaded from the specified path

#### Scenario: Watchlist flag with non-existent file
- **WHEN** user runs `tj --watchlist /new/watchlist.json`
- **THEN** an empty watchlist is created
- **AND** saves write to the specified path

### Requirement: Flags work with subcommands

The CLI SHALL accept path flags before any subcommand.

#### Scenario: Config flag with add subcommand
- **WHEN** user runs `tj --config /path/config.yaml add`
- **THEN** the add command uses the specified config

#### Scenario: Watchlist flag with add subcommand
- **WHEN** user runs `tj --watchlist /path/watchlist.json add`
- **THEN** the add command adds to the specified watchlist

#### Scenario: Both flags combined
- **WHEN** user runs `tj -c /config.yaml -w /watchlist.json`
- **THEN** both custom paths are used
