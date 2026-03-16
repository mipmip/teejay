## Context

The activity detection system now supports configurable patterns and idle timeout via `~/.config/teejay/config.yaml`. Users need documentation to understand and customize these options.

## Goals / Non-Goals

**Goals:**
- Document all configuration options in README.md
- Provide a working example config file
- Make it easy for users to customize detection patterns

**Non-Goals:**
- Auto-generating documentation from code (manual for now)
- Config validation tooling

## Decisions

### 1. Documentation Location: README.md

**Decision:** Add a Configuration section to README.md with a table of options.

**Rationale:**
- Single source of truth for users
- Visible on GitHub
- Standard location for project docs

### 2. Example File: config.example.yaml in repo root

**Decision:** Create `config.example.yaml` at repository root with all options commented.

**Rationale:**
- Common convention (like `.env.example`)
- Users can copy and modify
- Shows all available options with explanations

### 3. Table Format for Options

**Decision:** Use a markdown table with columns: Option, Type, Default, Description

**Rationale:**
- Scannable format
- Shows defaults at a glance
- Easy to maintain

## Risks / Trade-offs

**[Documentation drift]** Docs may get out of sync with code
→ Mitigation: Keep example file and README section close to config.go; update when config changes
