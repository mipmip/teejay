## 1. Makefile

- [x] 1.1 Create `Makefile` at project root
- [x] 1.2 Add `test` target that runs `go test ./...`
- [x] 1.3 Add `build` target that runs `go build -o tmon ./cmd/tmon`
- [x] 1.4 Add `lint` target that runs `go vet ./...`
- [x] 1.5 Add `help` target that lists available commands

## 2. Initial Test Files

- [x] 2.1 Create `internal/ui/app_test.go` with a placeholder test
- [x] 2.2 Ensure tests pass with `make test`

## 3. Verify

- [x] 3.1 Run `make build` and verify binary is created
- [x] 3.2 Run `make lint` and verify no errors
- [x] 3.3 Run `make test` and verify tests pass
