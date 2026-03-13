# Tasks for branding-footer

## 1. Pass version to UI

- [x] 1.1 Add `Version` field to `ui.Model` struct
- [x] 1.2 Change `ui.New()` to `ui.New(version string)` accepting version parameter
- [x] 1.3 Update `cmd/tj/main.go` to pass version to `ui.New(version)`

## 2. Create branding styles

- [x] 2.1 Add neon branding style (bright cyan/magenta color)
- [x] 2.2 Add version style (muted, smaller)

## 3. Render footer in View()

- [x] 3.1 Create `renderFooter()` helper method that returns "Terminal Junkie" + version
- [x] 3.2 Check terminal width - only show footer if width >= 80
- [x] 3.3 Position footer at bottom-right using lipgloss.Place()

## 4. Test and verify

- [x] 4.1 Build and run the app
- [x] 4.2 Verify branding displays in bottom-right
- [x] 4.3 Verify version number shows correctly
- [x] 4.4 Verify footer hides on narrow terminal (< 80 cols)
