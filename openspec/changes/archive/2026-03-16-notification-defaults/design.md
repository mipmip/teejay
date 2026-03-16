## Context

Currently, notification (`notify_on_ready`) and sound (`sound_on_ready`) settings are per-pane only, stored in watchlist.json as booleans defaulting to false. Users must enable these for each pane individually. There's no way to set defaults that apply to all panes.

## Goals / Non-Goals

**Goals:**
- Allow users to configure global default values for sound and notification alerts
- Per-pane settings override global defaults
- New panes automatically inherit global defaults
- Maintain backward compatibility with existing watchlist.json files

**Non-Goals:**
- Changing the alert mechanisms themselves (bell, notify-send)
- Adding new alert types
- Per-app notification defaults (would add complexity)

## Decisions

### 1. Config Structure: Add `alerts` section to config.yaml

**Decision:** Add a new `alerts` section with `sound_on_ready` and `notify_on_ready` booleans.

```yaml
alerts:
  sound_on_ready: false    # default
  notify_on_ready: false   # default
```

**Rationale:**
- Consistent with existing `detection` section pattern
- Clear and simple for users to configure
- Separates concerns from detection settings

### 2. Override Logic: Tri-state for per-pane settings

**Decision:** Change per-pane settings to pointers (`*bool`) to distinguish "not set" from "set to false".

**Rationale:**
- `nil` = use global default
- `*true` = explicitly enabled
- `*false` = explicitly disabled

**Alternative considered:** Separate "use_default" flag - rejected as more complex.

### 3. UI Behavior: Toggle cycles through states

**Decision:** When toggling, cycle: default → enabled → disabled → default

**Rationale:**
- Users can explicitly override to any state
- Can return to "follow global default" without removing the pane

## Risks / Trade-offs

**[JSON compatibility]** Existing watchlist.json files have `bool` not `*bool`
→ Mitigation: JSON unmarshals missing fields as nil pointer, `false` as `*false`. Existing files continue to work - explicit false values are preserved as overrides.

**[UI complexity]** Three states harder to communicate than two
→ Mitigation: Use clear indicators: `[D]` default, `[✓]` enabled, `[✗]` disabled
