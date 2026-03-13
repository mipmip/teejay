## Why

"tmon" was a placeholder name during initial development. We're renaming to "Teejay" (CLI: `tj`) which stands for "Terminal Junky" - a name that captures the addictive nature of vibe coding and how this tool fuels that obsession. The new name better reflects the personality and purpose of the tool.

## What Changes

- Rename module from `tmon` to `tj`
- Rename CLI binary from `tmon` to `tj`
- Update all internal import paths from `tmon/...` to `tj/...`
- Move `cmd/tmon/` directory to `cmd/tj/`
- Update README, docs, and any user-facing references
- Update Makefile/build scripts if present

## Capabilities

### New Capabilities

None - this is a rename/rebrand, not new functionality.

### Modified Capabilities

None - this change affects naming and branding only, not behavioral requirements.

## Impact

- **Module**: `go.mod` module name changes from `tmon` to `tj`
- **Binary**: Compiled binary changes from `tmon` to `tj`
- **Imports**: All Go files importing `tmon/...` packages need path updates
- **Directory**: `cmd/tmon/` moves to `cmd/tj/`
- **User workflows**: Users will invoke `tj` instead of `tmon`
- **Documentation**: Any references to "tmon" need updating to "Teejay" or "tj"
