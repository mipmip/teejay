## ADDED Requirements

### Requirement: Test execution via make

The project SHALL provide a `make test` command that runs all Go tests.

#### Scenario: Run all tests
- **WHEN** running `make test`
- **THEN** all `*_test.go` files are executed via `go test ./...`
- **AND** test results are displayed to stdout

### Requirement: Build via make

The project SHALL provide a `make build` command that compiles the binary.

#### Scenario: Build the binary
- **WHEN** running `make build`
- **THEN** the `tmon` binary is created in the project root
- **AND** the build uses `go build ./cmd/tmon`

### Requirement: Lint via make

The project SHALL provide a `make lint` command that runs static analysis.

#### Scenario: Run linting
- **WHEN** running `make lint`
- **THEN** `go vet ./...` is executed
- **AND** any issues are reported to stdout

### Requirement: Test file convention

Test files SHALL be placed alongside their corresponding source files with `_test.go` suffix.

#### Scenario: Test file location
- **WHEN** a source file exists at `internal/foo/bar.go`
- **THEN** its tests are located at `internal/foo/bar_test.go`
