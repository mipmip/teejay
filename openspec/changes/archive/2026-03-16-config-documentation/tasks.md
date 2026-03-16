## 1. Example Configuration File

- [x] 1.1 Create `config.example.yaml` in repository root
- [x] 1.2 Document `detection.idle_timeout` with comment and default value
- [x] 1.3 Document `detection.prompt_endings` with comment and empty default
- [x] 1.4 Document `detection.waiting_strings` with comment and empty default
- [x] 1.5 Document `detection.apps` section with all default app patterns (claude, aider, codex, opencode)
- [x] 1.6 Add comments explaining app-specific patterns replace globals

## 2. README Documentation

- [x] 2.1 Add Configuration section to README.md after Usage section
- [x] 2.2 Add configuration options table with columns: Option, Type, Default, Description
- [x] 2.3 Document `detection.idle_timeout` in table
- [x] 2.4 Document `detection.prompt_endings` in table
- [x] 2.5 Document `detection.waiting_strings` in table
- [x] 2.6 Document `detection.apps.<name>.prompt_endings` in table
- [x] 2.7 Document `detection.apps.<name>.waiting_strings` in table
- [x] 2.8 Add note explaining app-specific patterns replace global patterns
- [x] 2.9 Add example showing how to customize for a new app

## 3. Verification

- [x] 3.1 Verify config.example.yaml is valid YAML
- [x] 3.2 Verify config.example.yaml loads without errors when copied to config path
