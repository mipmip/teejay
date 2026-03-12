## 1. Add Tick Message Infrastructure

- [x] 1.1 Define `previewTickMsg` struct as a tea.Msg type in app.go
- [x] 1.2 Create `tickCmd()` function that returns a tea.Cmd using tea.Tick with 100ms interval

## 2. Integrate Ticker with Bubbletea Lifecycle

- [x] 2.1 Modify `Init()` to return `tickCmd()` when watchlist is not empty
- [x] 2.2 Add `previewTickMsg` case in `Update()` that calls `captureSelectedPane()` and returns next `tickCmd()`

## 3. Handle Modal State Guards

- [x] 3.1 Skip refresh in tick handler when `m.editing` is true
- [x] 3.2 Skip refresh in tick handler when `m.deleting` is true
- [x] 3.3 Skip refresh in tick handler when `m.empty` is true or `m.selectedPaneID` is empty

## 4. Verify Implementation

- [x] 4.1 Run the app and confirm preview auto-updates when watching an active pane
- [x] 4.2 Verify edit mode blocks preview refresh
- [x] 4.3 Verify delete confirmation mode blocks preview refresh
