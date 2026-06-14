# AGENTS.md

## Shared skills

This repository uses the shared skills from `bin/skills/`. Read
`bin/AGENTS.md` for the canonical shared skill list and use the smallest
matching skill for the task.

## Repository purpose

`status` is a small Go service used for HTTP status-code and health endpoint
testing.

- `GET /v1/status/{code}` returns the requested status code.
- Health endpoints are exposed via the shared health module under the service
  name prefix: `/status/healthz`, `/status/livez`, and `/status/readyz`.
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
./status server -config file:test/.config/server.yml
```

## Runtime and test notes

- Go version is `1.26.0` in `go.mod`.
- Ruby test harness dependencies live in `test/Gemfile`.
- `test/nonnative.yml` expects the service on `http://localhost:11000`.
- `test/.config/server.yml` listens on `tcp://:11000`, which binds port `11000`
  on all interfaces; use `http://localhost:11000` for local client requests.
  Operation endpoints use the `status` service prefix, for example
  `http://localhost:11000/status/healthz`.
- Test and coverage artifacts are written under `test/reports/`.
- The `/v1/status/{code}` handler also accepts `sleep=<duration>`, parses it
  with `time.ParseDuration`, and rejects values above the effective `max_sleep`.
  Longer sleeps can still exceed the configured HTTP request timeout.

## Intentional design choices

- `/status/healthz` is intentionally wired to the shared health module's online
  registration, which checks external internet connectivity by default. This is
  expected for this service; do not flag the external egress dependency as an
  issue unless the task is specifically about changing health semantics.

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
- The configured `bin` submodule URL uses GitHub SSH; fresh checkouts need
  GitHub SSH access or a local HTTPS submodule URL override.
- Shared git helper targets are exposed here too; some are destructive
  (`make reset`, `make purge`, branch deletion helpers, force-push helpers).
  Do not use them unless the user explicitly asks.
- `make dev` depends on `air`.
- `make sec` may depend on external tools such as `trivy`.
- `make start` and `make stop` rely on a sibling `../docker` repository via
  shared scripts.
