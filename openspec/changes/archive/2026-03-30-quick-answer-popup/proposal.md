## Why

When monitoring multiple coding agents, users must switch tmux panes to answer questions or approve tool use. This breaks flow — especially for quick interactions like "y/n" permission prompts or selecting from a numbered menu. Users need a way to respond to waiting agents directly from teejay without leaving the monitoring view.

## What Changes

- Add a quick-answer popup (`space` key) that lets users respond to a waiting pane without switching to it
- Add a prompt recognition system that determines *what* a waiting agent is asking for, using Claude Code's session transcript files for structured data (with screen-scraping fallback for other agents)
- Add a `?` status indicator for panes that are waiting with a specific question or choice (vs `●` for idle waiting)
- Send user responses to the target pane via `tmux send-keys`
- Perform a freshness check (re-capture + re-parse) before sending to prevent answering stale prompts

### Scope

- Claude Code only for structured recognition (transcript-based). Other agents get basic free-text input via the popup.
- No cross-pane triage/cycling (future proposal)

### Recognition strategy

For Claude Code panes, use the session transcript (`.jsonl`) for structured prompt data rather than parsing terminal output. The chain: tmux pane → child PID → `~/.claude/sessions/<pid>.json` → session UUID → transcript `.jsonl` → last assistant message. This is version-resistant since it reads a data format, not rendered UI.

For non-Claude agents: fall back to screen scraping the captured pane content (existing `tmux capture-pane` output) to extract basic prompt context.

## Capabilities

### New Capabilities

- `prompt-recognition`: Detects what a waiting pane is asking for — permission prompt, multiple choice, free-text question, or idle. Extracts question text and options.
- `quick-answer-popup`: A popup triggered by `space` that shows the detected prompt and lets users respond via selectable options or free-text input. Sends the response via `tmux send-keys`.
- `claude-session-reader`: Reads Claude Code session transcript files to extract structured prompt state. Maps tmux panes to Claude sessions via the PID chain.

### Modified Capabilities

- `activity-detection`: The status indicator gains a third visual state (`?`) for panes waiting with a specific question/choice, distinct from idle waiting (`●`)

## Impact

- New package: `internal/prompt/` — prompt recognition, Claude session reader, response sending
- `internal/ui/app.go` — new popup state, `space` keybinding, `?` indicator rendering, help text
- `internal/monitor/` — extended PaneStatus or new PromptInfo attached to pane state
- `internal/tmux/` — new `SendKeys()` function
- Reads from `~/.claude/sessions/` and `~/.claude/projects/` (external filesystem, read-only)
