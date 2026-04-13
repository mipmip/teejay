## Why

Users currently need to manually create `~/.config/teejay/config.yaml` by copying or adapting the example file. A `tj init` command provides guided setup, reducing friction for new users and ensuring the config lands in the correct location with valid settings.

## What Changes

- Add `tj init` subcommand that interactively creates a config file at `~/.config/teejay/config.yaml`
- Interactive wizard prompts for key settings: sound alerts, desktop notifications, default layout (single/columns), and sort order (watchlist/activity)
- If a config file already exists, prompt the user to overwrite, backup, or skip
- Writes a well-commented YAML config file with the chosen settings

## Capabilities

### New Capabilities
- `config-init`: Interactive `tj init` command that guides users through initial configuration setup and writes the config file

### Modified Capabilities

(none)

## Impact

- New subcommand `init` added to the CLI parser in `cmd/tj/main.go`
- New command implementation in `internal/cmd/init.go`
- No changes to existing config loading or runtime behavior
- No new dependencies (uses stdlib for user prompts)
