## Context

The app currently displays "Terminal Junkie" as its branding name in the footer area (`internal/ui/app.go:renderBrandingFooter()`). The name is hardcoded as a string literal. The existing spec at `openspec/specs/branding-footer/spec.md` references the old name throughout.

## Goals / Non-Goals

**Goals:**
- Replace all occurrences of "Terminal Junkie" with "Terminal Jockey" in source code and specs

**Non-Goals:**
- Changing the styling, positioning, or any other branding behavior
- Updating archived change documents (they reflect historical state)

## Decisions

**String replacement only** — The name is a single hardcoded string in `renderBrandingFooter()`. A simple find-and-replace in the source file and spec is sufficient. No architectural changes needed.

## Risks / Trade-offs

- **[Low]** Missed occurrences — Mitigated by grep search; only `internal/ui/app.go` and the live spec contain the name in active code/specs.
