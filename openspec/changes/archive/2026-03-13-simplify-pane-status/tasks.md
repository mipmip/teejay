# Tasks for simplify-pane-status

## 1. Update status.go - Rename states and simplify indicators

### 1.1 Rename status constants
- [x] In `internal/monitor/status.go`, rename `Running` to `Busy`
- [x] Rename `Ready` to `Waiting`
- [x] Remove `Idle` constant entirely
- [x] Update iota ordering (only two states: Busy=0, Waiting=1)

### 1.2 Update String() method
- [x] Change `Running` case to return `"Busy"`
- [x] Change `Ready` case to return `"Waiting"`
- [x] Remove `Idle` case

### 1.3 Update Indicator() method
- [x] Change `Running` case to use spinner (will delegate to animated)
- [x] Change `Ready` case to return green `"●"`
- [x] Remove `Idle` case
- [x] Note: Static Indicator() can just return the first spinner frame for Busy

### 1.4 Update IndicatorAnimated() method
- [x] Change `Running` case to `Busy` (keep spinner logic)
- [x] Change `Ready` case to `Waiting` (keep green `"●"`)
- [x] Remove `Idle` case

## 2. Update monitor.go - Remove idle counter logic

### 2.1 Remove idle tracking
- [x] Remove `idleThreshold` constant
- [x] Remove `idleCounter` field from `paneState` struct

### 2.2 Simplify Update() method
- [x] Remove all `idleCounter` logic
- [x] Change return values from `Running` to `Busy`
- [x] Change return values from `Ready` to `Waiting`
- [x] Simplify logic: if hasPrompt → Waiting, else → Busy

## 3. Update monitor_test.go - Adjust tests for two-state model

### 3.1 Update status name references
- [x] Replace `Running` with `Busy` in all tests
- [x] Replace `Ready` with `Waiting` in all tests
- [x] Remove `Idle` references

### 3.2 Remove idle transition tests
- [x] Remove `TestUpdateIdleTransition` test
- [x] Remove `TestUpdateIdleToRunning` test
- [x] Update `TestLongContent` to expect Busy instead of Idle

### 3.3 Update indicator tests
- [x] Update `TestStatusIndicator` for two states only
- [x] Update `TestStatusIndicatorAnimated` - remove Idle case
- [x] Update `TestStatusString` - remove Idle case

### 3.4 Update other test expectations
- [x] `TestMultiplePanes` - update comment about Idle

## 4. Update app.go - UI references

### 4.1 Update any status references
- [x] Search for any `Running`, `Ready`, `Idle` references in app.go
- [x] Update status display code if needed

## 5. Run tests and verify

### 5.1 Run tests
- [x] Run `go test ./internal/monitor/...`
- [x] Fix any compilation errors
- [x] Fix any failing tests

### 5.2 Manual verification
- [x] Build and run the app
- [x] Verify busy panes show animated spinner
- [x] Verify waiting panes show green indicator
