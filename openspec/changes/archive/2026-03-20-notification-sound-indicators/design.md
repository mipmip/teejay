## Context

Teejay already has a full alert system: global sound/notification defaults in config, per-pane overrides via three-state toggles (default/enabled/disabled), and five sound types. However, there's no visual indication of these settings anywhere in the UI. Users must open the configure popup per pane to see what's active.

The branding footer currently shows "Terminal Junkie" + version in the bottom-right. Pane items show title + breadcrumb description + status indicator.

## Goals / Non-Goals

**Goals:**
- Show global alert config in the branding footer area as a compact "systray"
- Show per-pane override indicators on watchlist items only when they differ from global defaults
- Use modest, colored Unicode symbols that fit the TUI aesthetic

**Non-Goals:**
- Making indicators interactive/clickable (configuration stays in the existing configure popup)
- Adding new alert types or changing alert behavior
- Showing sound type in the indicator (just on/off)

## Decisions

### 1. Symbol choice: Unicode text symbols, not emoji
Use `♪` for sound and `✉` for notifications. These are compact, render reliably in terminals, and can be colored with lipgloss. Emoji rendering varies across terminals and can cause width calculation issues.

**Alternative considered**: Emoji (🔔📢) — rejected due to inconsistent terminal rendering and double-width character issues in lipgloss width calculations.

### 2. Color coding
- Enabled: bright/vivid color (green `#00FF00` for sound, yellow `#FFD700` for notification)
- Disabled: dim gray (`#555555`)

This provides clear at-a-glance status without being visually noisy.

### 3. Global systray placement: after version, before right edge
Render as `Terminal Junkie v0.2.6 ♪ ✉` in the branding footer. The symbols appear after the version, colored according to global config state. This keeps the systray compact and co-located with existing branding.

**Alternative considered**: Separate line above footer — rejected because it would add visual complexity and change layout calculations.

### 4. Per-pane indicators: on the description/breadcrumb line
Append override indicators to the breadcrumb description: `session > window : process  ♪ ✉`. Only shown when the pane has explicit overrides (not when inheriting global defaults). This avoids visual noise for panes using defaults.

**Alternative considered**: On the title line next to the status indicator — rejected because the title line already has the name + animated spinner, adding more would be cramped.

### 5. Determining "has override"
A pane has an override when `pane.SoundOnReady != nil` or `pane.NotifyOnReady != nil`. The indicator shows the effective state (on/off color) but only appears when the pane has an explicit override set. This leverages the existing three-state system in the watchlist.

## Risks / Trade-offs

- **Narrow terminals may truncate indicators** → The systray symbols are only 4 characters wide (`♪ ✉`). On terminals < 80 cols the entire branding footer is already hidden, so no additional risk.
- **Color may not be visible on some terminal themes** → Using bright distinct colors (green, yellow) that work on both dark and light backgrounds. The symbols themselves are still readable even without color.
