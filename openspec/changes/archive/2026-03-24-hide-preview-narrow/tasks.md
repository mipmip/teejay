## 1. Main View() responsive layout

- [x] 1.1 Add `showPreview` boolean based on `listWidth < 25` in View() and conditionally render single-panel or split layout
- [x] 1.2 When preview is hidden, set `listWidth` to `m.width - 4` for full-width sidebar

## 2. Consistent width calculations

- [x] 2.1 Apply the same breakpoint logic to the WindowSizeMsg handler width calculations
- [x] 2.2 Apply the same breakpoint logic to the mouse click list width calculation

## 3. Verification

- [x] 3.1 Run existing tests to verify no regressions
