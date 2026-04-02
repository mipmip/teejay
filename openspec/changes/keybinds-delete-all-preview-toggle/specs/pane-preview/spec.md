## ADDED Requirements

### Requirement: Toggle preview at runtime

The user SHALL be able to toggle the preview panel visibility at runtime with the `p` keybinding.

#### Scenario: Hide preview
- **WHEN** the preview is visible
- **AND** the user presses `p`
- **THEN** the preview panel SHALL be hidden in both default and multi-column layouts

#### Scenario: Show preview
- **WHEN** the preview is hidden
- **AND** the user presses `p`
- **THEN** the preview panel SHALL be shown (when space allows)
