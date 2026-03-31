## ADDED Requirements

### Requirement: Global preview toggle

The preview panel SHALL be globally toggleable via config and CLI flag.

#### Scenario: Preview disabled in default layout
- **WHEN** `display.show_preview` is `false`
- **THEN** the default layout SHALL show a full-width pane list without the right-side preview panel

#### Scenario: Preview disabled in multi-column layout
- **WHEN** `display.show_preview` is `false`
- **THEN** the multi-column layout SHALL NOT show the bottom preview panel regardless of available space

#### Scenario: Preview enabled (default)
- **WHEN** `display.show_preview` is `true` (or unspecified)
- **THEN** preview panels SHALL render as before (when space allows)

#### Scenario: CLI flag overrides config
- **WHEN** config has `display.show_preview: true`
- **AND** the user runs `tj --no-preview`
- **THEN** the preview SHALL be hidden

#### Scenario: CLI flag enables preview
- **WHEN** config has `display.show_preview: false`
- **AND** the user runs `tj --preview`
- **THEN** the preview SHALL be shown (when space allows)
