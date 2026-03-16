## 1. Model State

- [x] 1.1 Add `browserPreviewContent` string field to Model for storing preview content
- [x] 1.2 Add `browserPreviewErr` error field to Model for storing capture errors

## 2. Preview Capture Logic

- [x] 2.1 Create `captureBrowserPreview()` method that captures content for currently selected browser pane
- [x] 2.2 Call `captureBrowserPreview()` in `loadPaneListForSession()` after loading panes (capture first item)
- [x] 2.3 Call `captureBrowserPreview()` in `updateBrowsing()` when up/down navigation changes selection

## 3. Browser Popup Layout

- [x] 3.1 Update `renderBrowserPopup()` to use split layout (list + preview) when viewing panes
- [x] 3.2 Calculate appropriate widths for list and preview panels based on terminal width
- [x] 3.3 Render preview panel with captured content (or error message if capture failed)
- [x] 3.4 Keep single-panel layout when viewing session list (no preview)

## 4. Testing

- [x] 4.1 Verify browser preview fields exist and are initialized correctly
- [x] 4.2 Run existing tests to ensure no regressions
