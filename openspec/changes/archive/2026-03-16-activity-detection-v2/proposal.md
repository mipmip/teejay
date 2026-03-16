## Why

Activity detection is unreliable (issue #18). The current implementation only checks for shell prompt characters (`$ > # %`) at end of lines, which fails for:
- Claude Code (shows "? for shortcuts" when idle)
- Custom shell prompts (starship, oh-my-zsh themes)
- TUI applications that don't have traditional prompts

Users report panes showing "busy" when they're actually idle and waiting for input.

## What Changes

- Add configurable idle timeout detection: if content unchanged for N seconds, mark as WAITING
- Add configurable pattern matching: prompt endings and waiting strings, with app-specific overrides
- Introduce `~/.config/teejay/config.yaml` for detection settings
- App-specific patterns REPLACE global patterns (not additive)
- Empty global prompt_endings by default (unreliable), rely on idle timeout + app patterns
- Ship with sensible defaults for Claude Code, Codex, Mistral Vide, Aider and OpenCode

## Capabilities

### New Capabilities

- `config-file`: YAML configuration file (`~/.config/teejay/config.yaml`) for application settings with detection configuration

### Modified Capabilities

- `activity-detection`: Replace simple prompt detection with configurable idle timeout + pattern matching system with app-specific overrides

## Impact

- New file: `internal/config/config.go` - config loading and defaults
- New file: `~/.config/teejay/config.yaml` - user configuration (optional)
- Modified: `internal/monitor/monitor.go` - new detection logic with idle tracking
- Modified: `internal/ui/app.go` - pass app name to monitor, load config
- New dependency: `gopkg.in/yaml.v3` for YAML parsing
