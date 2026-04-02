## Context

Teejay loads config from `~/.config/teejay/config.yaml`, falling back to defaults if absent. Users currently discover settings via `config.example.yaml` in the repo or `--help` flags. There's no built-in way to bootstrap a config file, which is the most common first step after installing.

The existing CLI uses a custom flag parser in `cmd/tj/main.go` that dispatches to command functions in `internal/cmd/`. Commands use `bufio.Reader` on `cmd.Stdin` for interactive input (see `add.go`).

## Goals / Non-Goals

**Goals:**
- Provide `tj init` to interactively create a config file at the default location
- Ask about: sound alerts, desktop notifications, default layout, sort order
- Handle existing config gracefully (prompt to overwrite, backup, or skip)
- Write a well-commented YAML file so users can further customize

**Non-Goals:**
- Full config editor / TUI settings screen (future work)
- Configuring detection patterns or app-specific patterns (too advanced for init)
- Supporting `--config` path override for init (always targets default location)

## Decisions

### 1. Plain stdin prompts, not a Bubbletea TUI

The init command runs as a one-shot CLI interaction using `fmt.Printf` + `bufio.Reader`, matching the pattern in `add.go`. A Bubbletea app would be overkill for 4-5 yes/no questions and would add complexity.

**Alternative considered**: Bubbletea form with styled prompts — rejected because init is a one-time operation and the existing command pattern works well.

### 2. Write commented YAML via Go template

Generate the config file from a Go text template embedded in the init command. This keeps the output format close to `config.example.yaml` while allowing dynamic values from user answers. The template includes comments explaining each setting.

**Alternative considered**: Copy `config.example.yaml` and sed-replace values — rejected because it couples to a file that may not be installed with the binary.

### 3. Config path uses same default as config loader

The init command writes to the same `~/.config/teejay/config.yaml` path that `config.DefaultPath()` resolves. This ensures the generated file is picked up automatically.

### 4. Existing file handling: prompt with three choices

When a config file already exists, prompt the user:
- **Overwrite**: Replace the file
- **Backup**: Rename to `config.yaml.bak` then write new
- **Cancel**: Exit without changes

This avoids data loss while keeping the flow simple.

## Risks / Trade-offs

- [Template drift] The embedded YAML template may diverge from `config.example.yaml` over time → Keep the template minimal (only settings asked about) and reference the example file in a comment for advanced options.
- [No --config support] Init always writes to the default path → Acceptable for v1; users who want a custom path can copy the file manually.
