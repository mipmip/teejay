## Why

Pressing Escape quits the app because the bubbles `list` component has default quit keybindings that include `esc`. This conflicts with Escape being used to close popups (edit, delete, browse, configure, quick-answer). Users instinctively press Escape to dismiss a popup but if the keypress reaches the list component, it quits the entire app.

## What Changes

- Disable the default `Quit` keybinding on all `list.New` instances so that `esc` no longer triggers `tea.Quit` via the list component
- Only `q` and `ctrl+c` should exit the application (already handled explicitly at the top of the key handler)

## Capabilities

### New Capabilities

_None — this is a bugfix to existing key handling._

### Modified Capabilities

_None — no spec-level behavior changes, only fixing unintended quit behavior from a library default._

## Impact

- `internal/ui/app.go`: Disable default quit keybindings on `list.New` instances (main list and browser list) so `esc` doesn't propagate as a quit command through the list component
