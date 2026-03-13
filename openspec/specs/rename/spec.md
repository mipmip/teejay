# rename Specification

## Purpose
TBD - created by archiving change rename-to-teejay. Update Purpose after archive.
## Requirements
### Requirement: Application naming
The application SHALL be named "Teejay" with CLI binary name `tj`.

#### Scenario: Binary invocation
- **WHEN** user types `tj` in terminal
- **THEN** the application launches (previously `tmon`)

#### Scenario: Help displays correct name
- **WHEN** user invokes `tj --help` or `tj -h`
- **THEN** output displays "Teejay" as the application name

### Requirement: Module naming
The Go module SHALL be named `tj` to match the binary name.

#### Scenario: Import paths use new module name
- **WHEN** Go files import internal packages
- **THEN** import paths start with `tj/` (e.g., `tj/internal/ui`)

