## Why

As tmon grows with new features (watchlist, add command, etc.), we need a reliable way to verify functionality and prevent regressions. Go has excellent built-in testing support, and establishing the testing pattern now ensures all future code is testable.

## What Changes

- Set up Go testing conventions and patterns for the project
- Create initial test files for existing packages
- Add a Makefile for common development tasks (test, build, lint)

## Capabilities

### New Capabilities
- `test-infrastructure`: Go testing setup with conventions for unit tests, test helpers, and development workflow tooling

### Modified Capabilities
<!-- None -->

## Impact

- New `*_test.go` files alongside source files
- New `Makefile` at project root
- No changes to existing functionality—purely additive
