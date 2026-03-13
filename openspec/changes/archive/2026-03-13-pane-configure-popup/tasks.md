## 1. Watchlist Data Model

- [x] 1.1 Add SoundOnReady and NotifyOnReady fields to Pane struct in `internal/watchlist/watchlist.go`
- [x] 1.2 Add SetSound and SetNotify methods to Watchlist
- [x] 1.3 Add tests for new watchlist fields and methods

## 2. Alerts Package

- [x] 2.1 Create `internal/alerts/alerts.go` with PlayBell() function
- [x] 2.2 Add SendNotification(title, message string) function using notify-send
- [x] 2.3 Add tests for alerts package

## 3. Configure Popup UI

- [x] 3.1 Add configuring state and configMenuItem to Model in `internal/ui/app.go`
- [x] 3.2 Add 'c' key handler to open configure popup
- [x] 3.3 Implement updateConfiguring() for menu navigation and actions
- [x] 3.4 Add configure popup rendering in View()
- [x] 3.5 Add configEditingName state for name text input in popup

## 4. Alert Triggering

- [x] 4.1 Track previous status per pane to detect transitions
- [x] 4.2 Call alerts on Running -> Ready transition based on pane settings
- [x] 4.3 Update help footer to include 'c' keybinding

## 5. Testing

- [x] 5.1 Add tests for configure popup state transitions
- [x] 5.2 Verify alert triggering only on status transitions
