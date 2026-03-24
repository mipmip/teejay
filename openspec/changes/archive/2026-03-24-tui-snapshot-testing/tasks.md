## 1. Snapshot test infrastructure

- [x] 1.1 Add `-update` flag and golden file helper functions (read/write/compare) in `internal/ui/app_test.go`
- [x] 1.2 Create `internal/ui/testdata/` directory for golden files

## 2. Delegate snapshot tests

- [x] 2.1 Add snapshot test for unselected pane item (busy status, breadcrumb, no alerts)
- [x] 2.2 Add snapshot test for selected pane item
- [x] 2.3 Add snapshot test for pane item with Waiting status (green indicator)
- [x] 2.4 Add snapshot test for pane item with alert override indicators

## 3. Bold title fix

- [x] 3.1 Apply `itemTitleStyle.Render()` to the title text in `browserItemDelegate.Render()` (only the title, not the whole line)
- [x] 3.2 Run snapshot tests with `-update` to capture the new golden files with bold titles
- [x] 3.3 Verify golden files contain bold ANSI escape code (`\x1b[1m`) in the title

## 4. Verification

- [x] 4.1 Run full test suite to verify no regressions
