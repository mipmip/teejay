## 1. Implementation

- [x] 1.1 Add helper function to get selected pane's display name from list items
- [x] 1.2 Update preview title rendering in View() to use display name instead of pane ID

## 2. Testing

- [x] 2.1 Add test case for preview title with custom pane name
- [x] 2.2 Add test case for preview title fallback to pane ID
- [x] 2.3 Run existing tests to ensure no regressions (blocked by incomplete auto-dismiss-tmux-error change - build fails due to missing notInTmuxMsg field)
