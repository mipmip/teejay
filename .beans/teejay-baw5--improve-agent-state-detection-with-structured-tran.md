---
# teejay-baw5
title: Improve agent state detection with structured transcript watchers
status: todo
type: feature
priority: normal
tags:
    - agent-detection
created_at: 2026-11-13T15:15:08Z
updated_at: 2026-11-13T15:15:08Z
---

## Current state

- Agent panes are detected by process name only: the command must be a key in `detection.apps` (`internal/scan/scan.go`).
- Activity state comes from screen content. `internal/monitor/monitor.go` hashes the ANSI-stripped pane content and returns `Waiting` after `idle_timeout` (default 2s) without a change, plus configurable waiting strings, busy strings and prompt endings.
- Prompt classification exists only for Claude (`internal/prompt/claude.go`). It follows the pane PID to `~/.claude/sessions/<pid>.json`, derives the transcript path, reads the last assistant entry, and classifies it as `FreeInput`, `Permission`, `Question` or `Choice`. Menu options are scraped from the rendered pane. Other agents fall back to `ScrapePrompt`, which treats the last five lines as free input context.

## Limitations

- Only two activity states exist. A finished run and a crashed run both look like `Waiting`.
- The idle timeout creates false `Waiting` states when an agent is stuck, or when a spinner or progress line stops repainting.
- Clocks, spinners and token counters cause false `Busy` states.
- Status depends on rendered output, so every TUI change can break detection.
- No thread identity: two agent threads in one pane collapse into one row.
- No structured status for Codex, OpenCode, Pi, Droid or Amp.
- No liveness signal: a row keeps its last state after the agent process exits.

## Proposal

Keep the prompt classification and menu scraping, because that is richer than what opensessions offers. Replace screen-derived activity for supported agents with agent-native state, the way opensessions does.

1. Add a status enum with `running`, `tool-running`, `waiting`, `done`, `error`, `interrupted` and `stale`. Map the existing `Busy` and `Waiting` indicator behavior onto it so the UI keeps working while richer states become available.
2. Add structured adapters:
   - Claude Code: `~/.claude/projects/*/*.jsonl`. Use the last assistant entry: `stop_reason=end_turn` -> done, `tool_use` or thinking -> running, an interrupt marker -> interrupted.
   - Codex: `$CODEX_HOME/sessions/**/*.jsonl`. Use `task_started`/`user_message` -> running, `task_complete`/final answer -> done, `turn_aborted` -> interrupted.
   - OpenCode: `opencode.db` through `sqlite3`. Use the last message: `finish=tool-calls` -> running, `stop` or `unknown` -> done, `error` -> error, an abort error -> interrupted.
   - Pi, Droid and Amp: JSONL or thread JSON adapters with their native event names.
3. Refine with file age:
   - running with a tool call as last entry and no file change for 3s -> waiting;
   - running or waiting with no file change for 15s -> stale;
   - ignore files older than 5 minutes.
4. Keep the pane process scan only for presence and liveness. Never derive status from pane title or screen content for agents with a native adapter.
5. Poll every 2s and apply a snapshot only when status, thread name, last prompt or project directory changed.
6. Match a snapshot to the pane by working directory, and use the existing PID chain for Claude where possible. Store one row per agent thread, not per pane.
7. Mark a row exited when its pane disappears or its process is gone, instead of leaving it waiting.

## Acceptance criteria

- Structured status and liveness exist for Claude Code, Codex, OpenCode, Pi, Droid and Amp.
- Parser fixtures exist per agent, and threshold behavior is tested with a fake clock.
- A finished, crashed or stuck agent produces distinct states instead of a single `Waiting`.
- Claude permission, question and choice prompts still render with scraped options, and existing prompt tests pass.
- The sidebar still shows a sensible indicator when only the process-name scan matches an agent.
- No status is derived from pane title or screen content for an agent with a native adapter.

## References

opensessions implementation, for comparison:

- `packages/runtime-rs/src/agent_watchers.rs`
- `packages/runtime-rs/src/agent_parsers.rs`
- `packages/runtime-rs/src/tracker.rs`
- `apps/server-rs/src/lib.rs`, `run_agent_watcher_loop`
- `packages/runtime-rs/src/tmux_provider.rs`, `agent_from_pane`
