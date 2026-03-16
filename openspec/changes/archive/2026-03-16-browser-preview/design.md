## Context

The pane browser currently renders as a centered popup with a single list panel. When users select a pane, they can only see the pane ID and running command - not the actual pane content. The main view already has a working preview system using `viewport.Model` and `tmux.CapturePane()`.

## Goals / Non-Goals

**Goals:**
- Add a preview panel to the pane browser showing selected pane content
- Reuse existing preview capture logic (`tmux.CapturePane`)
- Preview updates as user navigates through pane list
- Keep the popup responsive and well-sized

**Non-Goals:**
- Preview during session selection (only when viewing panes)
- Auto-refresh of browser preview (static capture on selection change is sufficient)
- Scrollable preview viewport (simple truncated content is acceptable)

## Decisions

### Decision 1: Split layout for pane selection view

Use Lipgloss `JoinHorizontal` to create a side-by-side layout: pane list on left, preview on right. Only show preview panel when viewing panes (not during session selection).

**Alternatives considered:**
- Stacked layout (list above preview): Takes more vertical space, less natural
- Preview on separate key press: More friction, loses real-time feedback

**Rationale:** Side-by-side mirrors the main view layout; users are already familiar with it.

### Decision 2: Store browser preview content in Model

Add `browserPreviewContent string` and `browserPreviewErr error` fields to Model, similar to main preview. Capture content when browser selection changes.

**Rationale:** Keeps browser preview state separate from main preview state; avoids conflicts when browser closes.

### Decision 3: Fixed popup dimensions

Increase popup width to accommodate both list and preview panels. Use percentage-based sizing relative to terminal dimensions.

**Rationale:** Ensures consistent experience across terminal sizes; preview needs reasonable width to be useful.

### Decision 4: Capture on navigation

Trigger `tmux.CapturePane` when:
- User navigates up/down in pane list
- User enters pane list from session list (capture first item)

**Rationale:** Keeps preview synchronized with selection; no polling needed.

## Risks / Trade-offs

- **[Risk] Popup too wide for small terminals** → Use minimum widths and gracefully degrade (hide preview if terminal too narrow)
- **[Trade-off] Static preview vs live refresh** → Static is simpler and sufficient for selection purposes
- **[Trade-off] Preview panel takes space from list** → Accept slightly smaller list; users can scroll
