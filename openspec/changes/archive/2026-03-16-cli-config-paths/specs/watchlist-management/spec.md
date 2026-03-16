## ADDED Requirements

### Requirement: Load watchlist from custom path

The watchlist module SHALL accept an optional custom path for loading, defaulting to `~/.config/teejay/watchlist.json`.

#### Scenario: Load from default path
- **WHEN** Load() is called without a path argument
- **THEN** watchlist is loaded from `~/.config/teejay/watchlist.json`

#### Scenario: Load from custom path
- **WHEN** Load() is called with a custom path argument
- **THEN** watchlist is loaded from the specified path
- **AND** the watchlist remembers the path for subsequent saves

### Requirement: Save watchlist to custom path

The watchlist SHALL save to the same path it was loaded from.

#### Scenario: Save to default path
- **WHEN** watchlist was loaded from default path
- **AND** Save() is called
- **THEN** watchlist is saved to `~/.config/teejay/watchlist.json`

#### Scenario: Save to custom path
- **WHEN** watchlist was loaded from a custom path
- **AND** Save() is called
- **THEN** watchlist is saved to the custom path (not the default)

#### Scenario: Save creates parent directories
- **WHEN** Save() is called
- **AND** the parent directory does not exist
- **THEN** parent directories are created before saving
