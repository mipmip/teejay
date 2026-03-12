## Context

tmon currently has no tests. As we add features like the watchlist and add-pane command, we need testing infrastructure to ensure reliability. Go's built-in `testing` package provides everything needed—no external test frameworks required.

## Goals / Non-Goals

**Goals:**
- Establish Go testing conventions for the project
- Create a Makefile for common dev tasks
- Add placeholder tests for existing packages to establish patterns

**Non-Goals:**
- Achieving high code coverage immediately
- Adding integration tests or end-to-end tests
- Setting up CI/CD (separate concern)

## Decisions

### Decision 1: Use Go's built-in testing package

Use the standard `testing` package with table-driven tests where appropriate.

**Rationale**: Go's testing is excellent out of the box. No need for testify, gomega, or other frameworks—they add dependencies without significant benefit for a project this size.

**Alternative considered**: testify for assertions. Rejected because Go's `if got != want` pattern is clear enough, and we avoid an external dependency.

### Decision 2: Makefile for dev workflow

Create a simple Makefile with targets:
- `make test` - run all tests
- `make build` - build the binary
- `make lint` - run go vet (can add golangci-lint later)

**Rationale**: Makefile is universal, works on all Unix systems, and provides discoverability (`make help`).

### Decision 3: Test file placement

Tests go alongside source files: `foo.go` → `foo_test.go` in the same directory.

**Rationale**: This is Go convention. Tests are in the same package, can test unexported functions, and are easy to find.

## Risks / Trade-offs

- [Trade-off] Starting with minimal tests → Acceptable; infrastructure is more important than coverage right now
- [Risk] Makefile doesn't work on Windows → Low priority; can add PowerShell equivalents later if needed
