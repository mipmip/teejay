## Context

Current activity detection (`internal/monitor/monitor.go`) uses a simple heuristic: check if the last non-empty line ends with `$ > # %`. This fails for modern shells, TUI apps, and AI coding assistants like Claude Code.

We need a more robust system that combines idle detection with configurable pattern matching.

## Goals / Non-Goals

**Goals:**
- Reliable detection for Claude Code, Aider, and standard shells
- User-configurable patterns and timeouts
- Sensible out-of-the-box defaults
- App-specific pattern overrides

**Non-Goals:**
- Perfect detection for all possible applications
- Complex regex patterns (keep it simple: substrings and end-chars)
- Auto-detection of application types

## Decisions

### 1. Configuration Format: YAML

**Decision:** Use YAML for `~/.config/teejay/config.yaml`

**Rationale:**
- Human-readable, easy to edit
- Good support for nested structures (apps → patterns)
- `gopkg.in/yaml.v3` is well-maintained

**Alternatives considered:**
- JSON: Less readable for lists/nested config
- TOML: Less common in Go ecosystem

### 2. Default Patterns: Empty Globals, App-Specific Defaults

**Decision:** Ship with empty global `prompt_endings` and `waiting_strings`. Provide app-specific defaults for known tools.

**Rationale:**
- Shell prompt detection is unreliable across different setups
- Idle timeout catches most cases
- App-specific patterns are more precise

**Default config:**
```yaml
detection:
  idle_timeout: 2s
  prompt_endings: []
  waiting_strings: []
  apps:
    claude:
      waiting_strings:
        - "? for shortcuts"
    aider:
      waiting_strings:
        - "(Y)es/(N)o"
```

### 3. Pattern Override Behavior: Replace, Not Extend

**Decision:** When an app has config, use ONLY its patterns (ignore globals for that app).

**Rationale:**
- Simpler mental model
- Apps like Claude Code have unique patterns that don't mix with shell patterns
- Users can copy global patterns to app config if needed

### 4. Idle Timeout: Per-Pane Timestamp Tracking

**Decision:** Track `last_change_time` per pane. If content hash unchanged and time exceeds `idle_timeout`, mark WAITING.

**Rationale:**
- Simple to implement
- Works for any application
- Complements pattern matching

**State per pane:**
```go
type paneState struct {
    hash           [32]byte
    lastChangeTime time.Time
    status         PaneStatus
}
```

### 5. App Name Source: tmux pane_current_command

**Decision:** Use `#{pane_current_command}` from tmux to identify the app.

**Rationale:**
- Already available (used for display)
- Matches process name (e.g., "claude", "aider", "fish")
- No additional detection needed

## Risks / Trade-offs

**[Idle timeout false positives]** Long-running output could pause briefly, triggering WAITING
→ Mitigation: 2s default is long enough for most streaming output; users can increase

**[App name mismatch]** Process name might not match config key (e.g., "claude-code" vs "claude")
→ Mitigation: Document common names; users can add aliases in config

**[Config file missing]** First-time users won't have config.yaml
→ Mitigation: Use embedded defaults when file doesn't exist

**[YAML parsing errors]** Malformed config could crash app
→ Mitigation: Log warning and fall back to defaults on parse error
