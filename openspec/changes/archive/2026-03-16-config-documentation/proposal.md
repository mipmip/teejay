## Why

The new configuration system (`~/.config/teejay/config.yaml`) lacks documentation. Users need to know what options are available, their defaults, and how to customize them. Additionally, there's no example config file to help users get started.

## What Changes

- Add a Configuration section to README.md with a table documenting all config options
- Create `config.example.yaml` in the repository root with commented examples of all options
- Document the app-specific pattern override behavior

## Capabilities

### New Capabilities

- `config-docs`: Documentation for the configuration file in README.md and an example config file

### Modified Capabilities

None

## Impact

- Modified: `README.md` - add Configuration section with options table
- New file: `config.example.yaml` - example configuration with comments
