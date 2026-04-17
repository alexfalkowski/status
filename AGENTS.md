# AGENTS.md

## Shared skill

Use the shared `coding-standards` skill from `./bin/skills/coding-standards`
for workflow, review, testing, documentation, and PR conventions.

This file is intentionally limited to repo-specific context so we do not
duplicate the shared guidance.

## Repository purpose

`status` is a small Go service used for HTTP status-code and health endpoint
testing.

- `GET /v1/status/{code}` returns the requested status code.
- Health endpoints are exposed via the shared health module:
  `/healthz`, `/livez`, and `/readyz`.
- The binary entrypoint is `status server`.

## Entry points and layout

- `main.go`: CLI bootstrap.
- `internal/cmd/`: server command registration and module wiring.
- `internal/config/`: application config layered on top of
  `github.com/alexfalkowski/go-service/v2/config`.
- `internal/api/v1/transport/http/`: HTTP route registration for
  `/v1/status/{code}`.
- `internal/health/`: health registration and HTTP observers.
- `test/`: Ruby nonnative/cucumber integration and benchmark harness.
- `test/.config/server.yml`: local config used by `dev`, `features`, and
  `benchmarks`.
- `bin/build/make/*.mak`: shared Makefile fragments used by the root
  `Makefile`.

## Preferred commands

From the repo root, prefer the exposed `make` targets:

- `make help`
- `make dep`
- `make build`
- `make build-test`
- `make lint`
- `make specs`
- `make features`
- `make benchmarks`
- `make sec`
- `make coverage`
- `make dev`

Useful direct run while debugging:

```sh
./status server -i file:test/.config/server.yml
```

## Runtime and test notes

- Go version is `1.26.0` in `go.mod`.
- Ruby test harness dependencies live in `test/Gemfile`.
- `test/nonnative.yml` expects the service on `http://localhost:11000`.
- Test and coverage artifacts are written under `test/reports/`.
- The `/v1/status/{code}` handler also accepts `sleep=<duration>` and parses
  it with `time.ParseDuration`.

## CI signal

Use the shared skill for validation strategy. The repo-specific CI source of
truth is `.circleci/config.yml`.

The main `build-service` job runs:

- `make clean`
- `make dep`
- `make lint`
- `make sec`
- `make features`
- `make benchmarks`
- `make analyse`
- `make coverage`

## Gotchas

- The root `Makefile` depends on the `bin` submodule and includes:
  `bin/build/make/help.mak`, `bin/build/make/http.mak`, and
  `bin/build/make/git.mak`.
- Shared git helper targets are exposed here too; some are destructive
  (`make reset`, `make purge`, branch deletion helpers, force-push helpers).
  Do not use them unless the user explicitly asks.
- `make dev` depends on `air`.
- `make sec` may depend on external tools such as `trivy`.
- `make start` and `make stop` rely on a sibling `../docker` repository via
  shared scripts.
