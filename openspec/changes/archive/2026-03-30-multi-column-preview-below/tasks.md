## 1. Height Calculation

- [x] 1.1 Add a `const minPreviewBelowHeight = 8` for the minimum useful preview height
- [x] 1.2 In `renderMultiColumnLayout()`, calculate `gridHeight` (itemsPerCol * 4) and `remainingHeight` (terminal height - gridHeight - footer)

## 2. Below-Preview Rendering

- [x] 2.1 When `remainingHeight >= minPreviewBelowHeight`, render the preview panel below the column grid — full width, using the existing `m.previewContent`/`m.viewport` with `previewPanelStyle` and `previewTitleStyle`
- [x] 2.2 Set `m.viewport.Height` to `remainingHeight - 4` (subtract border + title) before calling `m.viewport.View()` in the below-preview case
- [x] 2.3 Join the column grid and preview panel vertically with `lipgloss.JoinVertical`

## 3. Resize Handling

- [x] 3.1 Ensure `WindowSizeMsg` handler in multi-column mode sets viewport width to full terminal width (for below-preview)
