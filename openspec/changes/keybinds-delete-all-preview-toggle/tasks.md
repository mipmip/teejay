## 1. Delete All

- [x] 1.1 Add `deletingAll bool` field to Model
- [x] 1.2 Add `D` key handler in `Update()` — when watchlist not empty, set `deletingAll = true`
- [x] 1.3 Add `updateDeletingAll` handler before other modal handlers — on `y`: clear all panes, save watchlist, refresh; on `n`/`Esc`: cancel
- [x] 1.4 Add delete-all confirmation rendering in `View()` — show "Delete all N panes? (y/n)" in the footer area

## 2. Preview Toggle

- [x] 2.1 Add `p` key handler in `Update()` — toggle `m.config.Display.ShowPreview`

## 3. Help Footer

- [x] 3.1 Update help footer text to include `p: preview` and `d/D: delete`

## 4. Documentation

- [x] 4.1 Add `D` and `p` to README Keybindings table
- [x] 4.2 Update CHANGELOG.md under [Unreleased]
