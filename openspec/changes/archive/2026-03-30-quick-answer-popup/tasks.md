## 1. Foundation: Types and Tmux

- [x] 1.1 Create `internal/prompt/prompt.go` — define `PromptType` enum (Permission, Question, Choice, FreeInput, Unknown) and `PromptInfo` struct (type, question text, options, tool name, tool input summary, tool_use ID)
- [x] 1.2 Add `SendKeys(paneID string, keys string) error` to `internal/tmux/` — wraps `tmux send-keys -t <paneID> <keys> Enter`
- [x] 1.3 Add `GetChildPID(shellPID int, command string) (int, error)` to `internal/tmux/` — finds child process by command name via `pgrep -P`

## 2. Claude Session Reader

- [x] 2.1 Create `internal/prompt/claude.go` — implement `ReadClaudeSession(claudePID int) (*SessionInfo, error)` that reads `~/.claude/sessions/<pid>.json` and returns sessionId + cwd
- [x] 2.2 Implement `FindTranscript(sessionID, cwd string) (string, error)` — derives project hash from cwd, locates the `.jsonl` transcript file
- [x] 2.3 Implement `ReadLastAssistant(transcriptPath string) (*TranscriptEntry, error)` — reads from end of file, finds last assistant message, extracts stop_reason and tool_use content
- [x] 2.4 Implement `ParsePrompt(entry *TranscriptEntry) PromptInfo` — converts transcript entry to PromptInfo: maps tool names to Permission/Question/Choice, extracts parameters

## 3. Prompt Recognition

- [x] 3.1 Create `internal/prompt/scrape.go` — implement `ScrapePrompt(capturedContent string) PromptInfo` — basic fallback that extracts last few lines as context, returns FreeInput type
- [x] 3.2 Create `internal/prompt/recognize.go` — implement `Recognize(paneID string, appName string) PromptInfo` — orchestrator that uses Claude reader for "claude" apps, scrape fallback for others

## 4. Status Indicator

- [x] 4.1 Add `promptInfo` field to `paneItem` struct in `internal/ui/app.go`
- [x] 4.2 Add periodic prompt check (~2s tick) for Waiting panes — calls `prompt.Recognize()` and stores result in pane state
- [x] 4.3 Update `browserItemDelegate.Render()` to show `?` (yellow) instead of `●` when paneItem has an actionable prompt (Permission, Question, or Choice)

## 5. Quick-Answer Popup

- [x] 5.1 Add popup state fields to `Model`: `quickAnswering bool`, `quickAnswerPane string`, `quickAnswerPrompt PromptInfo`, `quickAnswerSelected int`, `quickAnswerInput textinput.Model`
- [x] 5.2 Add `space` key handler in `Update()` — checks pane is Waiting, calls `prompt.Recognize()`, opens popup
- [x] 5.3 Implement `renderQuickAnswerPopup()` — adapts display based on prompt type: selectable options for Permission/Choice, text input for Question/FreeInput
- [x] 5.4 Handle popup key events: arrow keys for option selection, text input for free-text, Enter to confirm, Esc to cancel
- [x] 5.5 Implement freshness check on Enter: re-capture pane, re-check Waiting status, (for Claude) verify same tool_use ID. Show "Prompt expired" if stale.
- [x] 5.6 On confirmed send: call `tmux.SendKeys()` with appropriate response, close popup

## 6. Help Text and Polish

- [x] 6.1 Update help footer to include `space: answer` keybinding
- [x] 6.2 Add popup help line: "↑/↓: select • Enter: send • Esc: cancel" (for option mode) or "Enter: send • Esc: cancel" (for text mode)
