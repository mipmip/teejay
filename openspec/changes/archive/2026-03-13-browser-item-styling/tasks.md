## 1. Styles

- [x] 1.1 Add browserItemStyle and browserItemSelectedStyle lipgloss styles in `internal/ui/app.go`

## 2. Custom Delegate

- [x] 2.1 Create browserItemDelegate struct implementing list.ItemDelegate interface
- [x] 2.2 Implement Height() returning 3 (title + description + margin)
- [x] 2.3 Implement Spacing() returning 0 (we handle spacing in render)
- [x] 2.4 Implement Update() passing through to default behavior
- [x] 2.5 Implement Render() with background colors based on selection state

## 3. Integration

- [x] 3.1 Update loadSessionList() to use browserItemDelegate instead of default delegate
- [x] 3.2 Update loadPaneListForSession() to use browserItemDelegate
- [x] 3.3 Update mouse click detection in updateBrowsing() to use itemHeight=3

## 4. Testing

- [x] 4.1 Verify items display with dark grey background
- [x] 4.2 Verify selected item displays with lighter background
- [x] 4.3 Verify mouse clicks select correct items with new height
