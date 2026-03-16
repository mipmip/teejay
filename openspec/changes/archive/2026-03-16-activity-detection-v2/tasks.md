## 1. Configuration System

- [x] 1.1 Add `gopkg.in/yaml.v3` dependency
- [x] 1.2 Create `internal/config/config.go` with Config struct and detection settings
- [x] 1.3 Implement `Load()` function to read from `~/.config/teejay/config.yaml`
- [x] 1.4 Implement default values (2s idle timeout, empty globals, claude/aider app defaults)
- [x] 1.5 Handle missing file (use defaults) and parse errors (log warning, use defaults)

## 2. Monitor Refactor

- [x] 2.1 Update `paneState` struct to include `lastChangeTime`
- [x] 2.2 Update `Monitor` to accept config (idle timeout, patterns)
- [x] 2.3 Change `Update()` signature to accept app name: `Update(paneID, content, appName string)`
- [x] 2.4 Implement app-specific pattern lookup (replace globals when app config exists)
- [x] 2.5 Implement `hasPromptEnding()` to check configured end characters
- [x] 2.6 Implement `hasWaitingString()` to check configured substrings
- [x] 2.7 Implement idle timeout check (content stable for N seconds → Waiting)
- [x] 2.8 Remove hardcoded prompt patterns from `hasPrompt()`

## 3. Integration

- [x] 3.1 Load config in `internal/ui/app.go` New() function
- [x] 3.2 Pass config to Monitor constructor
- [x] 3.3 Update `refreshListWithFrame()` to pass app name to `monitor.Update()`
- [x] 3.4 Update preview tick handler to pass app name

## 4. Testing

- [x] 4.1 Add unit tests for config loading (file exists, missing, malformed)
- [x] 4.2 Add unit tests for default values
- [x] 4.3 Add unit tests for app-specific pattern override behavior
- [x] 4.4 Add unit tests for idle timeout detection
- [x] 4.5 Add unit tests for prompt ending detection
- [x] 4.6 Add unit tests for waiting string detection
- [x] 4.7 Manual test: verify Claude Code detection works out of the box
- [x] 4.8 Manual test: verify idle timeout triggers Waiting after 2s
