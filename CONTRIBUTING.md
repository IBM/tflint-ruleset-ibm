# Contributing to the IBM Cloud TFLint ruleset

Thank you for contributing. This document covers how to report issues and send pull requests that can be reviewed quickly.

## Issues

### Bug reports

- [ ] Test against the latest release.
- [ ] Search for a duplicate report (`bug` label).
- [ ] Include a minimal `.tf` file (secrets removed) and the `tflint` output.
- [ ] Include your `.tflint.hcl` and `tflint --version`.

### Feature requests

- [ ] Search for a duplicate request (`enhancement` label).
- [ ] Describe the IBM Cloud resource or argument, the invalid configuration, and why `terraform apply` would fail.

## Pull requests

### Provider-wide practices (mandatory)

- Format with `gofmt`. CI fails if `gofmt -l .` is non-empty.
- Run `go vet ./...` and `go test ./...`.
- Run `go mod tidy` if you change `go.mod`.
- Do not hand-edit generated files (`rules/*_gen.go`, `rules/*_gen_test.go`, `rules/provider.go`, `docs/rules/*.md`). Change `extractor/`, `generator/`, or `scripts/openapi-overlay.py` and regenerate.
- Keep IBM Cloud Terraform provider source disposable and uncommitted. Point `PROVIDER_SRC` at a local clone; do not vendor it.
- Do not add personal GitHub owners, private clone URLs, or credentials.

### Regenerating

```
make extract-all
make overlay
make generate-all
make fmt vet coverage test build
```

Limit a tight loop with `PACKAGES=vpc`.

### New behavior in the runtime

Handwritten code lives in `ibm/`, `rules/check.go`, `extractor/`, `generator/`, `coverage/`, and `project/`. New validation kinds belong in `rules/check.go` with table-driven tests and a generator mapping. Every new check needs a failing fixture so the rule cannot regress silently.

### Documentation

Schema or argument changes must update generated docs in the same PR (`make generate-all`). README, configuration, and contributing changes belong in the same PR as the behavior they describe.

## Code of conduct

This project follows the [Contributor Covenant](CODE_OF_CONDUCT.md).
