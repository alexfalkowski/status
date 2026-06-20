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

- This service intentionally simulates HTTP errors and response delays for tests.
  The `/v1/status/{code}?sleep=<duration>` path may accept sleeps longer than
  the configured HTTP request timeout when the value is within `max_sleep`. Do
  not flag that mismatch as a reliability gap by default; it is expected
  test-fixture behavior unless the task is specifically about changing delay
  semantics or there is concrete evidence of current workflow breakage.
- `/status/healthz` is intentionally wired to the shared health module's online
  registration, which checks external internet connectivity by default. This is
  expected for this service; do not flag the external egress dependency as an
  issue unless the task is specifically about changing health semantics.
- This repository consumes shared Make targets from the `bin/` submodule. If a
  one-command local CI preflight target is needed, it should be added to the
  shared `bin` Make fragments rather than as a service-local target here. Do not
  flag the lack of a root `verify` or `ci-checks` target as a feature gap by
  default.
- The root `Makefile` intentionally depends directly on initialized shared
  `bin` Make fragments. Fresh checkouts must run `git submodule update --init`
  before `make help` or other Make targets can load. Do not flag the absence of a
  service-local fallback `make help` bootstrap path as a project gap by default.
- The configured `bin` submodule URL intentionally uses GitHub SSH. Fresh
  checkouts need GitHub SSH access or a local HTTPS submodule URL override, as
  documented below. Do not flag the checked-in SSH submodule URL as a setup or
  project gap by default unless this repository has explicitly decided to make
  HTTPS the default.
- CircleCI's `version` job runs the external `package` command from the
  `alexfalkowski/release` image. That release tooling owns GoReleaser config
  validation before publishing. Do not flag the absence of a separate
  repository-local GoReleaser config validation job as a project gap by default
  unless there is concrete evidence that the release image no longer validates
  `.goreleaser.yml`, or this repository has explicitly decided to own a
  pre-release GoReleaser check locally.
- Docker images for this service are released and deployed frequently, so a
  separate repository-local container runtime smoke gate before Docker image
  publication is not required by default. Do not flag the absence of a
  pre-publish Docker `docker run`/health endpoint smoke check as a reliability
  or project gap unless there is concrete evidence of current release breakage,
  or this repository has explicitly decided to own that gate locally.
- Docker image validation jobs intentionally run on non-master branches and are
  not required again before the master `version`/`package` release step. Do not
  flag the lack of master-branch `test-docker-*` gating before release writes
  as a project workflow gap by default.
- The Ruby code under `test/` is a local feature-test harness, not production
  service code. Fixed localhost endpoints in `test/lib/**`, `test/nonnative.yml`,
  and related feature helpers are intentional local harness assumptions unless
  there is concrete evidence of current workflow breakage. Do not flag the lack
  of environment-configurable HTTP or observability endpoints as a feature gap
  by default.
- The supported integration and benchmark entrypoints are the root
  `make features` and `make benchmarks` targets, which build the correct binary
  before delegating into `test/`. Direct `make -C test features` and
  `make -C test benchmarks` are not the default workflow. Do not flag the lack
  of a direct `test/` binary preflight as a project gap by default.
- Feature and benchmark Cucumber runs intentionally share the configured HTML
  report path in `test/.config/cucumber.yml`. Treat the JUnit XML reports and
  coverage files as the durable CI artifacts; do not flag the lack of separate
  feature and benchmark HTML report paths as a project workflow gap by default.
- Ruby runtime selection for the `test/` harness is owned by the external
  service CI image and shared Ruby Make wiring, not by production service code.
  Do not flag the absence of a repository-local `.ruby-version`,
  `.tool-versions`, `mise.toml`, or Gemfile `ruby` directive as a project gap by
  default unless there is concrete evidence that CI no longer supplies the
  expected runtime, or this repository has explicitly decided to own Ruby version
  selection locally for the test harness.

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
