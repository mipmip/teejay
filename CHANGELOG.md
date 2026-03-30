# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

- fix: escape no longer quits the app (only `q` and `ctrl+c` quit)
- rename to Terminal Jockey
- multi-column layout mode: press `v` to toggle between list+preview and multi-column view
- left/right arrow keys navigate between columns in multi-column mode
- quick-answer popup: press `space` on a waiting pane to respond without switching
- Claude Code transcript-based prompt recognition (structured, version-resistant)
- `?` indicator for panes waiting with an actionable question (yellow) vs idle `●` (green)
- screen-scraping fallback for non-Claude agents
- freshness check before sending responses to prevent answering stale prompts
- multi-column layout: preview panel shown below columns when vertical space allows

## [0.2.8] - 2026-03-24

- title in bold text #27
- auto hide preview when page too narrow #26
- fix incorrect busy symbol for user active pane #24

## [0.2.7] - 2026-03-20

- fix release script when run in jj wc
- auto scan function for coding agents
- symbols to show notitification settings globally and pane specific
- new breadcrum display of location in tmux

## [0.2.6] - 2026-03-17

- release fix

## [0.2.5] - 2026-03-17

- release fix

## [0.2.4] - 2026-03-17

- release fix

## [0.2.3] - 2026-03-17

- fix build runner for MacOS
- feature: notification-pause-on-focus to prevent notification flooding
 
## [0.2.2] - 2026-03-16

- fix: add sessions were not focussed correctly
- fix: allow adding new panes in smaller screens
- feature: alternative config and watchlist cli options
- feature: add --help

## [0.2.1] - 2026-03-16

- fixes to the release script

## [0.2.0] - 2026-03-16

- Add `tj del` command
- add release script
- Configurable activity detection with idle timeout and app-specific patterns
  (Claude Code, Codex, Aider, etc.)
- Configuration file support (`~/.config/teejay/config.yaml`)
- Global default settings for sound and notification alerts in config
- Native audio playback with 5 selectable notification sounds (chime, bell,
  ping, pop, ding)
- Preview panel in pane browser popup for identifying panes
- Auto-dismiss for temporary error messages (3 second timeout)
- Configuration documentation in README.md
- Example configuration file (`config.example.yaml`)
- Preview panel title now shows pane name instead of pane ID
- Sound alerts use native Go audio instead of terminal bell
- goreleaser configuration for automated releases
- GitHub Actions workflow for release automation
- Version embedding via ldflags (`tj --version`)
- RELEASING.md maintainer documentation

## [0.1.0] - TBD

- TUI dashboard for monitoring multiple tmux panes
- Real-time pane content preview with automatic refresh
- Pane status detection (Idle, Running, Ready)
- Animated status indicators for running panes
- CLI command `tj add` to add current pane to watchlist
- Smart pane naming with auto-detection from running command
- Pane browser with session/pane navigation
- Pane configuration (rename, sound alerts, notifications)
- Keyboard and mouse navigation support
- External watchlist file sync (hot reload)
- Desktop notifications when panes become ready
- Sound alerts when panes become ready
- Built with Bubble Tea TUI framework
- Nix flake for reproducible builds
- goreleaser for automated releases
- Multi-platform support (Linux/macOS, amd64/arm64)
