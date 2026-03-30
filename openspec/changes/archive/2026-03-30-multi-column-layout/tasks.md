## 1. Model State

- [x] 1.1 Add `layoutMode` field to `Model` struct with constants `layoutDefault` and `layoutMultiColumn`

## 2. Keybinding

- [x] 2.1 Add `v` key handler in `Update()` to toggle `layoutMode` between default and multi-column
- [x] 2.2 Update help footer text to include `v: view` keybinding

## 3. Multi-Column Rendering

- [x] 3.1 Create `renderMultiColumnLayout()` method that calculates column count from terminal width and min column width (30), distributes items across columns top-to-bottom/left-to-right, renders each item using the delegate pattern, and joins columns horizontally
- [x] 3.2 Wire `renderMultiColumnLayout()` into `View()` — when `layoutMode == layoutMultiColumn`, call it instead of the default list+preview layout

## 4. Column Navigation

- [x] 4.1 In multi-column mode, intercept left/right arrow keys to navigate between columns (same row), translating between flat list index and (column, row) position
- [x] 4.2 Ensure up/down arrow keys navigate within a column, not across the flat list

## 5. Resize Handling

- [x] 5.1 Ensure `WindowSizeMsg` handler updates list dimensions appropriately for both layout modes
