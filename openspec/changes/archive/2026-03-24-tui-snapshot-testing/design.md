## Context

The `browserItemDelegate.Render()` writes styled output to an `io.Writer`. This output contains ANSI escape codes for colors, bold, and background. Currently no tests verify this output — the existing tests only check that `Title()`, `Description()`, and `Indicator()` return correct strings, not what the rendered block looks like.

The standard Go approach for visual output testing is golden files: capture the rendered output to a file, and in subsequent runs compare against it. A `-update` test flag regenerates the golden files when intentional changes are made.

## Goals / Non-Goals

**Goals:**
- Test the full Render() output path including ANSI styling
- Catch regressions like the "black bar" bug from the bold title attempt
- Enable confident styling changes (bold titles, new indicators, etc.)
- Make golden file updates simple with `-update` flag

**Non-Goals:**
- Testing the full `View()` method (too complex, depends on terminal size)
- Screenshot-based visual testing (requires a terminal emulator)
- Testing bubbletea event handling or state transitions (covered by existing tests)

## Decisions

### 1. Golden file approach with `bytes.Buffer` capture
Call `delegate.Render()` with a `bytes.Buffer` as the writer, then compare the buffer contents against a golden file in `testdata/`. The Render method already accepts an `io.Writer`, making this straightforward.

**Alternative considered**: Strip ANSI codes and compare plain text — rejected because the whole point is to verify styling (bold, colors, backgrounds).

### 2. Create a mock `list.Model` for Render() calls
The `Render()` method receives a `list.Model` for width and selected index. Create a minimal list with known width and items to get deterministic output.

### 3. Test flag `-update` for golden file regeneration
Use `flag.Bool("update", false, "update golden files")` in the test file. When run with `go test -update`, write actual output to golden files. Standard Go pattern used by many projects.

### 4. Apply bold title fix using `itemTitleStyle.Render(title)`
In the delegate Render(), wrap the title text with `itemTitleStyle.Render()` which applies `Bold(true)`. The previous attempt likely applied bold to the entire line or content block, causing the background to extend — applying it only to the title text within the line should work correctly.

## Risks / Trade-offs

- **ANSI output may differ across lipgloss versions** → Pin lipgloss version in go.mod (already pinned). Golden files need updating on dependency bumps.
- **Golden files are opaque** → The raw ANSI codes are hard to read. Mitigation: name files descriptively and add a comment in each test explaining what the golden file represents.
