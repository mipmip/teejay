# config-file Specification

## Purpose
TBD - created by archiving change activity-detection-v2. Update Purpose after archive.
## Requirements
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

### Requirement: Detection configuration structure

The configuration SHALL support a `detection` section with idle timeout and pattern settings.

#### Scenario: Configure idle timeout
- **WHEN** config contains `detection.idle_timeout: 3s`
- **THEN** idle timeout is set to 3 seconds

#### Scenario: Configure global prompt endings
- **WHEN** config contains `detection.prompt_endings: ["$", ">"]`
- **THEN** those characters are used for prompt detection

#### Scenario: Configure global waiting strings
- **WHEN** config contains `detection.waiting_strings: ["Continue?"]`
- **THEN** those strings are used for waiting detection

#### Scenario: Configure app-specific patterns
- **WHEN** config contains `detection.apps.claude.waiting_strings: ["? for shortcuts"]`
- **THEN** those patterns are used for panes running "claude"

### Requirement: Default configuration values

The application SHALL provide sensible defaults when configuration keys are missing.

#### Scenario: Default idle timeout
- **WHEN** `detection.idle_timeout` is not specified
- **THEN** idle timeout defaults to 2 seconds

#### Scenario: Default global patterns
- **WHEN** `detection.prompt_endings` is not specified
- **THEN** prompt endings defaults to empty list

#### Scenario: Default app patterns
- **WHEN** `detection.apps` is not specified
- **THEN** default app patterns are used for claude and aider

