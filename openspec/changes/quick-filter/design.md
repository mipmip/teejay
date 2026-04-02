## Context

The `bubbles/list` component has built-in filtering (`SetFilteringEnabled`), but it's currently disabled. The built-in filter uses fuzzy matching on Title/Description, which isn't ideal — we want case-insensitive substring matching across multiple fields (name, session, window, command).

Instead of enabling the built-in filter, we'll implement our own: filter the items before passing them to `list.SetItems()` in `refreshListWithFrame()`. This gives us full control over matching logic and works seamlessly with both layouts.

## Goals / Non-Goals

**Goals:**
- `/` enters filter mode with a text input in the footer
- Real-time filtering as the user types
- Matches against pane name, session name, window name, and foreground command
- Case-insensitive substring matching
- Filter persists after confirming (Enter), cleared by Esc

**Non-Goals:**
- Regex or fuzzy matching (simple substring is sufficient)
- Saving filter across sessions
- Filtering in the browser popup (only the main pane list)

## Decisions

### Filter state on Model

Add three fields:
- `filtering bool` — true when the filter input is active
- `filterQuery string` — the current (confirmed or in-progress) filter
- `filterInput textinput.Model` — the text input for typing the filter

### Filter in refreshListWithFrame

After building the full item list (and optional sort), apply the filter:

```go
if m.filterQuery != "" {
    filtered := make([]list.Item, 0)
    query := strings.ToLower(m.filterQuery)
    for _, item := range items {
        p := item.(paneItem)
        searchText := strings.ToLower(p.Title() + " " + p.session + " " + p.windowName + " " + p.command)
        if strings.Contains(searchText, query) {
            filtered = append(filtered, item)
        }
    }
    items = filtered
}
```

This runs on every refresh (~100ms) which is fine since it's just string matching on a small list.

### Keybinding flow

1. `/` → enter filter mode, focus `filterInput`, show input in footer
2. Typing → update `filterQuery` live, list re-renders with filter applied
3. `Enter` → confirm filter, exit filter mode (filter stays active, shown in footer as reminder)
4. `Esc` → clear filter, exit filter mode, show all panes

### Footer rendering

When filtering: show the text input in the footer
When filter is active (confirmed): show "Filter: <query>" in the status area with a hint "/ to edit • Esc to clear"
