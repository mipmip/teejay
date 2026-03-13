## 1. Styles

- [ ] 1.1 Add browserItemStyle and browserItemSelectedStyle lipgloss styles in `internal/ui/app.go`

## 2. Custom Delegate

- [ ] 2.1 Create browserItemDelegate struct implementing list.ItemDelegate interface
- [ ] 2.2 Implement Height() returning 3 (title + description + margin)
- [ ] 2.3 Implement Spacing() returning 0 (we handle spacing in render)
- [ ] 2.4 Implement Update() passing through to default behavior
- [ ] 2.5 Implement Render() with background colors based on selection state

## 3. Integration

- [ ] 3.1 Update loadSessionList() to use browserItemDelegate instead of default delegate
- [ ] 3.2 Update loadPaneListForSession() to use browserItemDelegate
- [ ] 3.3 Update mouse click detection in updateBrowsing() to use itemHeight=3

## 4. Testing

- [ ] 4.1 Verify items display with dark grey background
- [ ] 4.2 Verify selected item displays with lighter background
- [ ] 4.3 Verify mouse clicks select correct items with new height
