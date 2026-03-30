## ADDED Requirements

### Requirement: Preview in horizontal orientation

The preview panel SHALL support rendering below the content area (horizontal orientation) in addition to the existing side-by-side orientation.

#### Scenario: Below-preview uses full width
- **WHEN** the preview is rendered in horizontal orientation (multi-column mode)
- **THEN** the preview panel SHALL use the full terminal width minus borders

#### Scenario: Below-preview uses remaining height
- **WHEN** the preview is rendered in horizontal orientation
- **THEN** the preview panel height SHALL fill the remaining vertical space after the column grid and footer

#### Scenario: Visual consistency with side preview
- **WHEN** the below-preview is rendered
- **THEN** it SHALL use the same border style and title format as the side preview ("Preview: \<pane name\>")
