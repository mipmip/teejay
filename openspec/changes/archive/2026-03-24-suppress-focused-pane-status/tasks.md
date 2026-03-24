## 1. Model fields for focus tracking

- [x] 1.1 Add `lastActivePaneID string` and `paneFocusLostAt map[string]time.Time` fields to the Model struct
- [x] 1.2 Initialize `paneFocusLostAt` map in the New() constructor

## 2. Focus tracking in tick loop

- [x] 2.1 At the start of the tick handler, get `activePaneID` once and compare with `lastActivePaneID` to detect focus changes
- [x] 2.2 When focus changes from pane A to pane B, record `paneFocusLostAt[A] = time.Now()` and update `lastActivePaneID`

## 3. Skip monitoring for paused panes

- [x] 3.1 In the pane status update loop, skip `monitor.Update()` if the pane is the active pane OR if it's within the 2s grace period — but still capture content for preview
- [x] 3.2 Clean up expired grace period entries (> 2s old) from `paneFocusLostAt`

## 4. Verification

- [x] 4.1 Run existing tests to verify no regressions
