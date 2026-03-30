## Context

Teejay monitors tmux panes running coding agents. The existing monitor detects Busy/Waiting state via screen scraping (pattern matching on `tmux capture-pane` output). When a pane is Waiting, the user currently must `tmux switch` to interact with it.

Claude Code writes structured session data to `~/.claude/sessions/<pid>.json` (PID→session UUID mapping) and `~/.claude/projects/<project-hash>/<session-id>.jsonl` (conversation transcript). The last assistant message in the transcript contains the exact tool call or question that Claude is presenting to the user.

## Goals / Non-Goals

**Goals:**
- Let users answer agent prompts without leaving the teejay monitoring view
- Use Claude Code's structured transcript data for reliable prompt recognition
- Provide a fallback free-text input for non-Claude agents
- Show a distinct indicator when a pane has a specific question vs just being idle

**Non-Goals:**
- Supporting agents other than Claude Code for structured recognition (future work)
- Cross-pane triage/queue cycling through waiting panes (future proposal)
- Persisting or logging responses sent via quick-answer
- Modifying Claude Code's configuration or hooks

## Decisions

### Package structure: `internal/prompt/`

Create a new package for prompt recognition and response sending. This keeps the concern separate from the monitor (which handles busy/waiting detection) and the UI (which handles display).

Contents:
- `prompt.go` — PromptInfo type and common interface
- `claude.go` — Claude-specific transcript reader and prompt parser
- `scrape.go` — Screen-scraping fallback for non-Claude agents
- `send.go` — `tmux send-keys` wrapper

**Alternative considered**: Extending `internal/monitor/` — rejected because prompt recognition is a distinct concern from activity detection, and the Claude session reading is substantial enough to warrant separation.

### PID chain for Claude session lookup

To map a tmux pane to Claude's session transcript:

1. Get `pane_pid` from tmux (the shell PID — already available)
2. Find child process where `comm == "claude"` via `pgrep -P <pane_pid>`
3. Read `~/.claude/sessions/<claude_pid>.json` → get `sessionId` and `cwd`
4. Derive project hash from `cwd` (Claude uses path with `/` replaced by `-`, prefixed with `-`)
5. Read last entries from `~/.claude/projects/<project-hash>/<sessionId>.jsonl`

This chain is executed only when:
- The pane is detected as Waiting (by existing monitor)
- The foreground command is "claude"
- The user presses `space` (or during the periodic prompt-check tick)

**Alternative considered**: Using Claude's hooks system to push state — rejected because it requires configuring Claude Code, creating a coupling. Reading the transcript is passive and read-only.

### Layered recognition: transcript for type, screen for options

The transcript and screen serve complementary roles:

**Transcript** (`.jsonl`) determines the **prompt type**:

| Last message `stop_reason` | Content | Prompt Type |
|---|---|---|
| `end_turn` | Text only | `FreeInput` — agent finished, waiting at main prompt |
| `tool_use` | `AskUserQuestion` tool | `Question` or `Choice` — depending on whether options are present |
| `tool_use` | Any other tool | `Permission` — extract tool name and input summary |

The transcript also provides the `tool_use` ID for freshness checking.

**Screen scraping** (`tmux capture-pane`) provides the **actual menu options**:

For Permission and Choice prompts, the actual options are scraped from the rendered terminal output using a regex that matches numbered interactive lists (e.g., `❯ 1. Yes`, `  2. No`). This is version-resistant — whatever Claude Code renders, we display. The transcript's hardcoded fallback options are only used if scraping fails.

**Alternative considered**: Extracting options purely from the transcript — rejected because Claude Code's permission UI shows context-specific options (e.g., "Yes, and don't ask again for chmod:*") that aren't in the transcript data. The screen always has the ground truth.

### Freshness check before sending

Before sending a response via `tmux send-keys`:

1. Re-capture pane content via `tmux capture-pane`
2. Re-run monitor status check — confirm still Waiting
3. For Claude: re-read last transcript entry — confirm same tool_use ID
4. If stale (agent moved on): show "Prompt expired" message, close popup

This prevents sending keystrokes to an agent that has already continued or is now busy.

### Popup UX modes

The quick-answer popup adapts based on prompt type:

**Permission mode**: Shows tool name, key parameters, and selectable options (y/n/a). Arrow keys navigate, Enter sends the selected key.

**Choice mode**: Shows question text and numbered/lettered options. Arrow keys navigate, Enter sends the selection key.

**Free-text mode**: Shows question text (if any) and a text input field. Enter sends the typed text followed by a newline.

All modes include Esc to cancel without sending.

### Indicator: `?` for actionable waiting

Extend the pane item display with a third indicator state:
- `⠹` (spinner) — Busy
- `●` (green dot) — Waiting, idle (no specific question)
- `?` (yellow) — Waiting with a question/choice that needs user input

The `?` indicator is set when prompt recognition finds a Permission, Question, or Choice prompt. It stays `●` for FreeInput (idle at main prompt).

The prompt check runs periodically (every ~2s, not every 100ms tick) since reading transcripts is heavier than screen scraping. It only runs for Waiting panes with `claude` as the foreground command.

### Response sending via `tmux send-keys`

Three send modes in `internal/tmux/sendkeys.go`:

- **`SendArrowAndEnter(paneID, downPresses)`** — For interactive list menus (Permission + Choice). Sends N down-arrow presses to navigate to the desired item, then Enter to confirm. Claude Code's permission and AskUserQuestion prompts both render as interactive lists with `❯` cursor navigation.
- **`SendKeys(paneID, text)`** — For free-text input. Sends the text followed by Enter.
- **`SendRawKey(paneID, key)`** — For single raw keypresses without Enter. Available as fallback.

The prompt check runs asynchronously (via `tea.Cmd`) every ~2s to avoid blocking the UI thread. Transcript tail-reading (last 64KB) keeps I/O lightweight.

## Risks / Trade-offs

- **[Medium] Claude transcript format changes** → The `.jsonl` format is a data serialization, not a UI, so it's more stable than screen output. If it does change, only `internal/prompt/claude.go` needs updating. Mitigation: fail gracefully to free-text mode if parsing fails.
- **[Medium] PID chain race condition** → Between detecting Waiting and reading the transcript, Claude might continue. Mitigation: the freshness check catches this before sending.
- **[Low] Session file cleanup** → Claude may clean up old session files. We only read them for active panes, so this is unlikely to be an issue.
- **[Low] Permission prompt safety** → Users might accidentally approve destructive tool use via the popup. Mitigation: the popup clearly shows the tool name and parameters; the user explicitly confirms. This is the same information shown in the terminal.

## Open Questions

- Should the prompt check tick rate be configurable, or is a fixed ~2s interval sufficient?
- Should the `?` indicator show what *kind* of question it is (permission vs choice vs question), or is a single `?` enough?
