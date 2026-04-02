# config-docs Specification

## Purpose
TBD - created by archiving change config-documentation. Update Purpose after archive.
## Requirements
### Requirement: README documents all features

The README SHALL document all CLI flags, keybindings, and configuration options.

#### Scenario: CLI flags documented
- **WHEN** a user reads the README
- **THEN** all CLI flags SHALL be listed with descriptions

#### Scenario: Keybindings documented
- **WHEN** a user reads the README
- **THEN** all TUI keybindings SHALL be listed in a reference table

#### Scenario: Config options documented
- **WHEN** a user reads the README
- **THEN** all config file options SHALL be listed including the display section

#### Scenario: Example config is complete
- **WHEN** a user copies config.example.yaml
- **THEN** it SHALL include all available config sections with comments

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

### Requirement: Documentation sync guardrail

The OpenSpec project context SHALL include a rule reminding contributors to update documentation when adding CLI flags, config options, or keybindings.

#### Scenario: New flag added
- **WHEN** a change adds a new CLI flag
- **THEN** the contributor SHALL be reminded to update README.md, config.example.yaml, and printHelp()

