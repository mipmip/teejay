## Why

Issue #29 (bold pane titles) was attempted before but produced a "strange black bar" — a visual rendering bug that wasn't caught because there are no tests for the rendered output of the item delegate. The UI test coverage is only 25%, mostly smoke tests. Without a way to verify what the rendered output looks like, styling changes are trial-and-error. Adding golden-file snapshot tests for the delegate renderer lets the agent (and CI) verify visual output automatically. Relates to issue #29.

## What Changes

- Add golden-file snapshot testing infrastructure for the `browserItemDelegate.Render()` output
- Create golden files that capture the ANSI-styled output of rendered pane items in various states (selected/unselected, busy/waiting, with/without breadcrumb, with/without alert indicators)
- Apply the bold title fix (using `itemTitleStyle`) and verify it doesn't produce visual artifacts via the new snapshot tests
- Add an `-update` flag to regenerate golden files when intentional changes are made

## Capabilities

### New Capabilities
- `delegate-snapshot-tests`: Golden-file snapshot tests for the watchlist item delegate renderer

### Modified Capabilities
- `watchlist-item-delegate`: Pane title text is rendered with bold styling

## Impact

- `internal/ui/app_test.go`: New snapshot test functions for delegate rendering
- `internal/ui/testdata/`: Golden files for expected rendered output
- `internal/ui/app.go`: Apply `itemTitleStyle` to title text in delegate Render()
