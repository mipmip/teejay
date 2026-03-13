## 1. Module Rename

- [x] 1.1 Update go.mod module name from `tmon` to `tj`
- [x] 1.2 Run `go mod tidy` to validate module

## 2. Directory Rename

- [x] 2.1 Move `cmd/tmon/` directory to `cmd/tj/`

## 3. Import Path Updates

- [x] 3.1 Update imports in `cmd/tj/main.go` from `tmon/...` to `tj/...`
- [x] 3.2 Update imports in `internal/cmd/add.go`
- [x] 3.3 Update imports in `internal/cmd/add_test.go`
- [x] 3.4 Update imports in `internal/ui/app.go`
- [x] 3.5 Update imports in `internal/ui/app_test.go`
- [x] 3.6 Update imports in `internal/monitor/monitor.go`
- [x] 3.7 Update imports in `internal/monitor/monitor_test.go`
- [x] 3.8 Update imports in `internal/monitor/status.go`
- [x] 3.9 Update imports in `internal/alerts/alerts.go`
- [x] 3.10 Update imports in `internal/alerts/alerts_test.go`
- [x] 3.11 Update imports in remaining `internal/` files

## 4. Build Scripts and Config

- [x] 4.1 Update Makefile if present (binary name references)
- [x] 4.2 Update any CI/CD config files
- [x] 4.3 Update flake.nix if it references tmon

## 5. Documentation

- [x] 5.1 Update README.md references from tmon to tj/Teejay
- [x] 5.2 Update any other documentation files

## 6. Verification

- [x] 6.1 Run `go build ./cmd/tj` to verify build
- [x] 6.2 Run `go test ./...` to verify tests pass
- [x] 6.3 Test binary invocation with `./tj`
