## 1. Model State

- [x] 1.1 Add `filtering bool`, `filterQuery string`, and `filterInput textinput.Model` fields to Model
- [x] 1.2 Initialize `filterInput` in `New()` with placeholder "Filter..."

## 2. Keybinding and Handler

- [x] 2.1 Add `/` key handler: set `filtering = true`, focus `filterInput`, populate with current `filterQuery`
- [x] 2.2 Add `updateFiltering` handler (before other modal handlers): forward keys to `filterInput`, update `filterQuery` live on each keystroke
- [x] 2.3 In `updateFiltering`: Enter confirms (exit filter mode, keep query), Esc clears query and exits filter mode

## 3. Filter Logic

- [x] 3.1 In `refreshListWithFrame()`, after building items (and optional sort), filter by `filterQuery` — case-insensitive substring match against name + session + windowName + command

## 4. Footer Rendering

- [x] 4.1 When `filtering`: show "/ " + filterInput.View() in the footer
- [x] 4.2 When not filtering but `filterQuery != ""`: show "Filter: <query>" + hint "/ to edit • Esc to clear"
- [x] 4.3 When not filtering and Esc is pressed with an active filter: clear `filterQuery`

## 5. Documentation

- [x] 5.1 Add `/` and filter behavior to README Keybindings table
- [x] 5.2 Update help footer to include `/: filter`
- [x] 5.3 Update CHANGELOG.md under [Unreleased]
