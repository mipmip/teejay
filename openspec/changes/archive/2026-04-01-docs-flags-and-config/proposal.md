## Why

The README, config.example.yaml, and help text are outdated — they don't document the new CLI flags, display config options, keybindings, or the renamed "Terminal Jockey" branding. Future feature additions risk the same drift. We need a documentation update and a guardrail to keep docs in sync.

## What Changes

- Update README.md with:
  - Correct name: "Terminal Jockey"
  - Full CLI flags reference (alerts, display, mode)
  - Keybindings reference (all current keys)
  - Display config options table (recency_color, sort_by_activity, layout_mode, picker_mode)
- Update config.example.yaml with the new `display` section
- Add a documentation rule to the OpenSpec project context so future changes are reminded to update docs when adding flags, config options, or keybindings

## Capabilities

### New Capabilities

_None_

### Modified Capabilities

- `config-docs`: Documentation updated to reflect all current features

## Impact

- `README.md` — major update
- `config.example.yaml` — add display section
- `openspec/openspec.yaml` or project context — add doc-sync rule
