## 1. Temporary Message Infrastructure

- [x] 1.1 Add `temporaryMessage` string field to model (replaces `notInTmuxMsg` bool)
- [x] 1.2 Add `dismissTemporaryMsg` struct type for the dismiss timer message
- [x] 1.3 Create `showTemporaryMessage(msg string)` helper that sets message and returns dismiss command
- [x] 1.4 Create `dismissTemporaryCmd()` function that returns tea.Tick with 3-second delay

## 2. Message Display and Dismissal

- [x] 2.1 Update footer rendering to display `temporaryMessage` when set (with error styling)
- [x] 2.2 Handle `dismissTemporaryMsg` in Update to clear `temporaryMessage` state
- [x] 2.3 Update Esc key handler to clear `temporaryMessage` (early dismiss)

## 3. Migration of Not-in-Tmux Error

- [x] 3.1 Replace `notInTmuxMsg = true` with `showTemporaryMessage()` call in Enter key handler
- [x] 3.2 Remove old `notInTmuxMsg` field and related code

## 4. Testing

- [x] 4.1 Update existing tests to use new `temporaryMessage` field
- [x] 4.2 Verify all tests pass with new infrastructure
