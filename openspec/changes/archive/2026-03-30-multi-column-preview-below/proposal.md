## Why

In multi-column layout mode, the pane items often only fill the top portion of the terminal, leaving significant vertical space unused below. The preview panel — which is hidden in this mode — could be shown below the columns when there's enough room, giving users both the overview and the detail without needing to toggle back.

## What Changes

- In multi-column mode, calculate the vertical space remaining below the pane item grid
- When sufficient space exists (a configurable minimum height), render the preview panel for the selected pane below the columns, spanning the full width
- The preview auto-hides when terminal height shrinks below the threshold
- Preview content continues to update for the selected pane as in default mode

## Capabilities

### New Capabilities

_None_

### Modified Capabilities

- `multi-column-layout`: The multi-column view gains an optional preview panel below the item grid when vertical space allows
- `pane-preview`: Preview can now appear in a horizontal (below) orientation in addition to the vertical (right-side) orientation

## Impact

- `internal/ui/app.go` — `renderMultiColumnLayout()` gains height awareness and conditional preview rendering, viewport dimensions adjusted for the below-preview layout
