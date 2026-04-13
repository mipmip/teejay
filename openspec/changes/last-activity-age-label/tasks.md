## 1. Core Implementation

- [x] 1.1 Add `compactDuration(d time.Duration) string` helper function that formats elapsed time as "3s", "14m", "2h", "1d"
- [x] 1.2 Update `browserItemDelegate.Render()` to show the elapsed-time label on the title row of waiting panes, right-aligned before the status indicator, with dim styling

## 2. Tests

- [x] 2.1 Add unit tests for `compactDuration` covering seconds, minutes, hours, and days boundaries

## 3. Documentation

- [x] 3.1 Update CHANGELOG.md under [Unreleased] with the new activity age label
