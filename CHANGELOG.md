# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - TBD

### Added

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

### Technical

- Built with Bubble Tea TUI framework
- Nix flake for reproducible builds
- goreleaser for automated releases
- Multi-platform support (Linux/macOS, amd64/arm64)
