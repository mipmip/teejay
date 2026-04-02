## 1. Core Init Command

- [x] 1.1 Create `internal/cmd/init.go` with `InitConfig()` function implementing the interactive wizard (prompts for sound, notifications, layout, sort order)
- [x] 1.2 Add config file conflict handling: detect existing file, prompt to overwrite/backup/cancel
- [x] 1.3 Implement YAML template generation with comments and user-selected values
- [x] 1.4 Create parent directories (`~/.config/teejay/`) if they don't exist

## 2. CLI Integration

- [x] 2.1 Register `init` subcommand in `cmd/tj/main.go` CLI parser alongside `add`, `del`, `scan`
- [x] 2.2 Update `printHelp()` in `main.go` to include `init` command description

## 3. Tests

- [x] 3.1 Add tests in `internal/cmd/init_test.go` covering: wizard defaults, user input, existing file handling, generated YAML validity

## 4. Documentation

- [x] 4.1 Update README.md with `tj init` command documentation
- [x] 4.2 Update CHANGELOG.md under [Unreleased] with the new init command
