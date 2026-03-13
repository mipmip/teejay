## Context

The TUI has multiple contexts with different keyboard shortcuts:
- Main view (pane list)
- Browser popup (session/pane selection)
- Configure popup (pane settings)
- Delete confirmation
- Edit mode (renaming)

## Goals / Non-Goals

**Goals:**
- Document all keyboard shortcuts in README.md
- Organize by context for clarity
- Use standard markdown table format

**Non-Goals:**
- In-app help screen (future enhancement)
- Customizable keybindings

## Decisions

### Decision 1: Section placement

**Choice:** Add "Keyboard Shortcuts" section after "Usage" section.

**Rationale:**
- Logical flow: install → usage → shortcuts
- Users look for shortcuts after learning basic usage

### Decision 2: Format

**Choice:** Use markdown tables grouped by context.

**Rationale:**
- Tables are scannable
- Grouping by context reduces confusion about when shortcuts apply

## Risks / Trade-offs

**[Risk] Docs get outdated** → Keep shortcuts list minimal, link to code for completeness if needed.
