## Context

The main View() currently always renders a 30/70 split layout. The browser popup already has responsive logic (`showPreview := m.width >= 80`) that hides its preview on narrow terminals. The main view needs similar logic.

At 30% width, the sidebar is ~25 chars when the terminal is ~90 cols wide. Below that, item text gets truncated heavily and the layout is unusable.

## Goals / Non-Goals

**Goals:**
- Hide preview panel when sidebar would be < 25 chars wide
- Give full terminal width (minus borders) to the sidebar when preview is hidden
- Keep all width calculation paths consistent (View, WindowSizeMsg, mouse clicks)

**Non-Goals:**
- Changing the 30/70 ratio when both panels are shown
- Adding a user toggle to show/hide preview manually

## Decisions

### 1. Breakpoint: sidebar width < 25 chars
Calculate `listWidth = m.width*30/100 - 2`. If `listWidth < 25`, set `showPreview = false` and give the sidebar `m.width - 4` (full width minus borders). This threshold means the preview disappears at ~90 cols, which is reasonable.

### 2. Reuse pattern from browser popup
The browser popup uses `showPreview` boolean to branch layout rendering. Apply the same pattern to the main View().

### 3. Skip preview content building when hidden
When `showPreview` is false, skip building the preview panel, viewport content, and preview title — just render the list panel at full width.

## Risks / Trade-offs

- **User may expect preview always visible** → The preview reappears as soon as the terminal is wide enough. The transition is seamless.
