## MODIFIED Requirements

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

## ADDED Requirements

### Requirement: Documentation sync guardrail

The OpenSpec project context SHALL include a rule reminding contributors to update documentation when adding CLI flags, config options, or keybindings.

#### Scenario: New flag added
- **WHEN** a change adds a new CLI flag
- **THEN** the contributor SHALL be reminded to update README.md, config.example.yaml, and printHelp()
