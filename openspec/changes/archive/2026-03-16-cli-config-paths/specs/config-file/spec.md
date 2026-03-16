## MODIFIED Requirements

### Requirement: Load configuration from YAML file

The application SHALL load configuration from a specified path, defaulting to `~/.config/teejay/config.yaml` when no path is provided.

#### Scenario: Config file exists at default path
- **WHEN** the application starts without --config flag
- **AND** `~/.config/teejay/config.yaml` exists
- **THEN** configuration is loaded from the default path

#### Scenario: Config file exists at custom path
- **WHEN** the application starts with `--config /custom/path.yaml`
- **AND** the file exists
- **THEN** configuration is loaded from the custom path

#### Scenario: Config file missing at default path
- **WHEN** the application starts without --config flag
- **AND** `~/.config/teejay/config.yaml` does not exist
- **THEN** default configuration values are used

#### Scenario: Config file missing at custom path
- **WHEN** the application starts with `--config /custom/path.yaml`
- **AND** the file does not exist
- **THEN** default configuration values are used
- **AND** a warning is logged

#### Scenario: Config file malformed
- **WHEN** the application starts
- **AND** the config file (default or custom) contains invalid YAML
- **THEN** a warning is logged
- **AND** default configuration values are used
