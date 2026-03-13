## 1. Model State Changes

- [x] 1.1 Add `sessionItem` type implementing `list.Item` with session name and pane count
- [x] 1.2 Add `browsingSession` bool field to Model (true = session list, false = pane list)
- [x] 1.3 Add `selectedSession` string field to Model to track which session was selected
- [x] 1.4 Add `allBrowserPanes` slice to cache filtered panes when browser opens

## 2. Session List Loading

- [x] 2.1 Create `loadSessionList()` method that groups panes by session and populates browserList with sessionItems
- [x] 2.2 Update `loadBrowserPanes()` to cache panes in allBrowserPanes and call loadSessionList()
- [x] 2.3 Set browsingSession=true and selectedSession="" when loading session list

## 3. Pane List Loading

- [x] 3.1 Create `loadPaneListForSession(sessionName)` method that filters allBrowserPanes and populates browserList with browserItems
- [x] 3.2 Update browserList title to show selected session name
- [x] 3.3 Set browsingSession=false and selectedSession to the chosen session

## 4. Navigation Logic

- [x] 4.1 Update `updateBrowsing()` Enter handler: if browsingSession, call loadPaneListForSession; else add pane to watchlist
- [x] 4.2 Update `updateBrowsing()` Escape handler: if browsingSession, close browser; else call loadSessionList to go back
- [x] 4.3 Keep `q` key as unconditional close at any level

## 5. View Updates

- [x] 5.1 Update `renderBrowserPopup()` to show context-aware footer help text based on browsingSession state
- [x] 5.2 Update empty state message to be context-aware (no sessions vs no panes in session)

## 6. Display Format

- [x] 6.1 Update browserItem Title() to show "window.pane command" format instead of session:window.pane
- [x] 6.2 Update browserItem Description() to show pane ID
