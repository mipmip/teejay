## ADDED Requirements

### Requirement: Scan at startup

The application SHALL support scanning for agent panes at startup via CLI flag or config.

#### Scenario: Scan flag triggers startup scan
- **WHEN** the user runs `tj --scan`
- **THEN** agent panes SHALL be scanned and added to the watchlist before the first render

#### Scenario: Config triggers startup scan
- **WHEN** config has `display.scan_on_start: true`
- **THEN** agent panes SHALL be scanned at startup

#### Scenario: No scan by default
- **WHEN** no `--scan` flag is provided and config has `display.scan_on_start: false`
- **THEN** no startup scan SHALL occur

#### Scenario: Startup scan uses same logic as runtime scan
- **WHEN** a startup scan runs
- **THEN** it SHALL use the same agent detection and watchlist update logic as pressing `s`
