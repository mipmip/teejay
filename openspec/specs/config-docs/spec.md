# config-docs Specification

## Purpose
TBD - created by archiving change config-documentation. Update Purpose after archive.
## Requirements
### Requirement: README Configuration section

The README.md SHALL include a Configuration section documenting all config options.

#### Scenario: Configuration section exists
- **WHEN** a user reads README.md
- **THEN** they find a Configuration section with a table of all options
- **AND** each option shows its type, default value, and description

#### Scenario: App-specific patterns documented
- **WHEN** a user reads the Configuration section
- **THEN** they find documentation explaining how app-specific patterns replace global patterns

### Requirement: Example configuration file

The repository SHALL include a `config.example.yaml` file with all configuration options.

#### Scenario: Example file contains all options
- **WHEN** a user views config.example.yaml
- **THEN** they see all available configuration options
- **AND** each option has a comment explaining its purpose

#### Scenario: Example file shows defaults
- **WHEN** a user views config.example.yaml
- **THEN** they see the default values for each option
- **AND** they see examples of app-specific patterns

### Requirement: Example file is copyable

The example configuration file SHALL be immediately usable when copied.

#### Scenario: User copies example file
- **WHEN** user copies config.example.yaml to ~/.config/teejay/config.yaml
- **THEN** the application loads it without errors
- **AND** the configuration works as documented

