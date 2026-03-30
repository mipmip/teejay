## Context

The app uses the `charmbracelet/bubbles` `list` component which ships with default keybindings including `esc` and `q` mapped to `tea.Quit`. The app explicitly handles `q` and `ctrl+c` at the top of its key handler (line 778), but unhandled keys—including `esc`—fall through to `m.list.Update(msg)` (line 959), where the list's default quit binding fires.

Meanwhile, `esc` is used throughout the app to close popups (edit, delete, browse, configure, quick-answer). If a user presses escape when no popup is open (e.g., after a popup was already dismissed), the keypress reaches the list component and quits the app unexpectedly.

## Goals / Non-Goals

**Goals:**
- Prevent `esc` from quitting the app via the list component's default keybindings
- Keep `q` and `ctrl+c` as the only ways to exit

**Non-Goals:**
- Changing popup dismiss behavior (esc should still close popups)
- Adding new keybindings or quit confirmation dialogs

## Decisions

### Decision 1: Disable list component's default Quit keybinding

**Choice:** After each `list.New(...)` call, disable the `Quit` key in the list's `KeyMap` by setting `l.KeyMap.Quit.SetEnabled(false)`.

**Rationale:**
- The app already handles quitting explicitly (`case "q", "ctrl+c": return m, tea.Quit`), so the list's built-in quit binding is redundant
- Disabling is cleaner than intercepting esc before the list update — it removes the behavior at the source
- This is the standard approach recommended by the bubbles library for custom key handling

**Alternative considered:** Intercepting `esc` in the main key handler before it reaches `m.list.Update()`. Rejected because it would require adding explicit handling for every key the list might interpret, which is fragile.

## Risks / Trade-offs

- [Low] If future list features depend on the Quit keybinding being active, they would silently not work → Mitigation: The app already manages quitting explicitly, so this is intentional.
