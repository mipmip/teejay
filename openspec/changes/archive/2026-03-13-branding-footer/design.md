## Context

The app currently has no visual branding. The user wants "Terminal Junkie" displayed in neon-style characters at the bottom right, along with the version number. The version is already available via goreleaser's ldflags in `cmd/tj/main.go`.

## Goals / Non-Goals

**Goals:**
- Display "Terminal Junkie" in a stylized font at bottom-right
- Show version number alongside branding
- Keep footer subtle and non-intrusive
- Pass version from main to UI model

**Non-Goals:**
- Animated neon effects (keep it simple/static)
- Configurable branding position
- Custom fonts or external dependencies

## Decisions

### Decision 1: ASCII art style

**Choice:** Use simple block/neon-style ASCII characters, not full figlet fonts

**Rationale:** Full ASCII art banners are too tall and would take up too much vertical space. A compact stylized text fits better in a footer. Can use lipgloss styling (colors, bold) to achieve neon effect.

**Alternative considered:** Figlet/toilet style large ASCII art - rejected due to space constraints.

### Decision 2: Footer position

**Choice:** Absolute position at bottom-right corner using lipgloss.Place()

**Rationale:** Bottom-right is unobtrusive and conventional for version/branding info. Using lipgloss placement keeps it independent of the main layout flow.

### Decision 3: Neon styling

**Choice:** Use lipgloss color styling with bright cyan/magenta gradient effect

**Rationale:** "Neon" can be achieved through color alone - bright cyan (#00FFFF) or magenta (#FF00FF) on dark background gives neon glow feel without animation.

### Decision 4: Version passing

**Choice:** Add Version field to ui.Model, pass from main.go via NewModel()

**Rationale:** Simple and direct. The version variable already exists in main.go.

## Risks / Trade-offs

**[Trade-off]** Static vs animated neon → Static is simpler and less distracting. Real neon animation would require ticker updates and more complexity.

**[Risk]** Footer may overlap content on very small terminals → Can check terminal size and hide footer if too small.
