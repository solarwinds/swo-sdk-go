# swo-sdk-go

Go SDK for SolarWinds Observability. **Speakeasy-generated** from OpenAPI specs in `platform-apis` — do not manually edit generated internals.

## Build & Test

```bash
go test ./...           # run all tests
go build ./...          # verify compilation
go vet ./...            # static analysis
```

Go 1.22+ required.

## Structure

```
swov1/                  # Main SDK package (github.com/solarwinds/swo-sdk-go/swov1)
  *.go                  # Top-level client and resource files
  models/
    components/         # Request/response types
    operations/         # Operation parameter structs
    apierrors/          # Typed error responses
  internal/
    config/             # SDK configuration
    utils/              # Internal helpers
    hooks/              # Lifecycle hooks
  retry/                # Retry/backoff implementation
  types/                # Shared type definitions
  docs/                 # Generated documentation
  tests/                # Test utilities and mocks
.speakeasy/
  workflow.yaml         # Speakeasy generation config — edit to change generation
```

## Key Details

- **Auth:** HTTP Bearer token via `WithSecurity()` or `SWO_API_TOKEN` env var
- **Default server:** `https://api.na-01.cloud.solarwinds.com` (region configurable)
- **Pagination:** Use `.Next()` on paginated responses
- **Retries:** Exponential backoff by default, configurable per-operation
- **Status:** Pre-production beta — breaking changes possible without major version bump

## Making Changes

Do **not** edit generated files in `swov1/` directly — they get overwritten on regeneration.

To add/change API behavior:
1. Update the TypeSpec definition in `platform-apis`
2. Compile to OpenAPI spec
3. Run Speakeasy generation (`replace_env_in_yaml.py` handles env var substitution in specs)

## CI

GitHub Actions + Speakeasy CI for automated SDK regeneration on spec changes.
