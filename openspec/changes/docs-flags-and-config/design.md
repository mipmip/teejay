## Context

The README was last updated around v0.2.6. Since then: renamed to Terminal Jockey, added multi-column layout, quick-answer popup, recency colors, activity sort, picker mode, CLI flags, and alt screen. The config.example.yaml only covers detection and alerts — missing the display section entirely.

## Goals / Non-Goals

**Goals:**
- Bring all documentation up to date with current features
- Add a guardrail so future changes update docs

**Non-Goals:**
- Rewriting the README from scratch — keep the existing structure, add missing sections
- Adding a docs generation system — manual docs are fine for this project size

## Decisions

### README structure

Keep existing sections, add/update:
- Fix branding: "Terminal Jockey" (was "Terminal Junky")
- Add "CLI Flags" section after Usage
- Add "Keybindings" section with a table of all keys
- Update "Configuration Options" table with display section
- Update Features list

### Guardrail: OpenSpec context rule

Add a rule to the openspec project context that reminds contributors: "When adding CLI flags, config options, or keybindings, update README.md, config.example.yaml, and printHelp()." This gets included in the instructions for every change.

## Risks / Trade-offs

- **[Low] Docs drift again** — Mitigated by the OpenSpec context rule. Not foolproof but provides a reminder.
