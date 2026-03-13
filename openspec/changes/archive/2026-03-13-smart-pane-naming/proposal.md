## Why

When adding panes to the watchlist, they currently have no name and display as raw pane IDs (e.g., `%5`). This is cryptic and unhelpful when monitoring multiple panes. Users must manually rename each pane after adding it. The system should intelligently guess a meaningful name from available tmux metadata and only prompt the user when the guessed name is too generic.

## What Changes

- Auto-detect a meaningful name when adding a pane using:
  1. The running command in the pane (e.g., `nvim`, `cargo build`)
  2. The tmux window name if distinctive
  3. The tmux session name if distinctive
- Define a list of "generic" names that should trigger a user prompt (e.g., `bash`, `zsh`, `fish`, `tmux`, `claude`, `opencode`, numeric names)
- When the guessed name is generic, prompt the user to provide a name
- Apply to both `tj add` CLI command and in-app pane browser

## Capabilities

### New Capabilities

- `pane-name-guesser`: Logic to extract and evaluate potential pane names from tmux metadata, filtering out generic names

### Modified Capabilities

- `add-pane`: When adding a pane, the system will now auto-assign a guessed name or prompt for one if generic
- `pane-browser`: When selecting a pane from the browser, the same naming logic applies

## Impact

- **internal/tmux**: May need to fetch additional tmux metadata (window name)
- **internal/cmd/add.go**: Add name-guessing logic and optional interactive prompt
- **internal/ui/app.go**: Update browser add flow to use name-guessing logic
- **internal/watchlist**: `Add()` method may accept optional name parameter
