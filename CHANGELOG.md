# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.4] - 2026-03-17

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
