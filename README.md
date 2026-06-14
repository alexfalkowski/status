[![CircleCI](https://circleci.com/gh/alexfalkowski/status.svg?style=svg)](https://circleci.com/gh/alexfalkowski/status)
[![codecov](https://codecov.io/gh/alexfalkowski/status/graph/badge.svg?token=G6T3OIWUFK)](https://codecov.io/gh/alexfalkowski/status)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/status)](https://goreportcard.com/report/github.com/alexfalkowski/status)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/status.svg)](https://pkg.go.dev/github.com/alexfalkowski/status)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# 🩺 Status

`status` is a small Go service for testing HTTP status-code handling and health endpoints.
It is intended as a local, controllable alternative to depending on <https://httpstat.us/>
from automated tests.

> [!NOTE]
> This service is built for test fixtures, smoke tests, and client behavior checks. It is not a general-purpose public status-code proxy.

## 🧭 Background

External status-code services are useful, but their availability can make an
otherwise deterministic test suite fail for reasons outside the project under
test. Running this service locally keeps status-code and observability checks
under your control.

## ⚡ Quick Start

After cloning the repository, initialize the shared `bin` submodule and install
dependencies:

```sh
git submodule update --init
make dep
```

The configured `bin` submodule URL uses GitHub SSH. Configure GitHub SSH access
first, or override the submodule URL to HTTPS before initializing it.

Build and run the service with the local test configuration:

```sh
make build
./status server -config file:test/.config/server.yml
```

The local configuration listens on `tcp://:11000`, which binds port `11000` on
all interfaces. For local requests, use:

```sh
curl -i http://localhost:11000/v1/status/200
curl -i "http://localhost:11000/v1/status/503?sleep=50ms"
curl -i http://localhost:11000/status/healthz
```

> [!IMPORTANT]
> The root `Makefile` includes files from the `bin` submodule. Run `git submodule update --init` before using `make` in a fresh checkout.

> [!TIP]
> If you have `air` installed, `make dev` builds and runs the server in watch mode with `test/.config/server.yml`.

## 🖥️ Server

### 🔢 Status Code

Returns the requested HTTP status code.

#### 📥 Request

```http
GET /v1/status/{code}
GET /v1/status/{code}?sleep=50ms
```

> [!NOTE]
> `code` must parse as an integer from `200` through `999`. Values below `200`, values above `999`, and non-numeric values return `400 Bad Request`.

| Parameter | Location | Required | Description |
| --------- | -------- | -------- | ----------- |
| `code` | Path | Yes | Status code to return. Named codes include their standard reason phrase, such as `200 OK`. |
| `sleep` | Query | No | Delay before returning the response. Parsed with Go's [`time.ParseDuration`](https://pkg.go.dev/time#ParseDuration), for example `50ms`, `1s`, or `2m`. Must be less than or equal to the effective `max_sleep` and short enough for the configured HTTP request timeout. |

> [!CAUTION]
> `sleep` intentionally delays the response. Keep durations short in tests so client timeouts and CI jobs do not wait longer than expected. The checked-in local configuration sets `max_sleep` to `2m` and the HTTP timeout to `5s`, so some accepted sleeps can still outlast the transport timeout.

#### 📤 Response

The response status is the requested code. When HTTP permits a response body,
the body is plain text:

```http
200 OK
```

Status codes that do not permit a response body, such as `204 No Content` and
`304 Not Modified`, return no body.

For codes without a standard reason phrase, the body contains the numeric code:

```http
999
```

Invalid status codes, unparsable `sleep` values, and sleeps above the effective
`max_sleep` return `400 Bad Request`. A `sleep` accepted by `max_sleep` can
still exceed the configured HTTP request or client timeout; in that case the
request may time out or close before the requested status response is returned.

The maximum accepted sleep duration defaults to `5m`. Configure a lower maximum with:

```yaml
max_sleep: 2m
```

Omitting `max_sleep` or setting it to `0` uses the `5m` default. Positive
configured values must be less than or equal to `5m`.

## 💓 Health

The shared health module exposes service-prefixed health, liveness, readiness,
and metrics endpoints over HTTP. With the local `status` service name, use:

| Endpoint | Check | Healthy response |
| -------- | ----- | ---------------- |
| `/status/healthz` | Online connectivity | `SERVING` |
| `/status/livez` | No-op liveness check | `SERVING` |
| `/status/readyz` | No-op readiness check | `SERVING` |
| `/status/metrics` | Prometheus metrics | Metrics including `go_info` |

> [!WARNING]
> `/status/healthz` uses the shared online health registration, which checks external internet connectivity by default. In an offline environment, prefer `/status/livez` or `/status/readyz` for local process checks.

Configure health check timing with:

```yaml
health:
  duration: 1s
  timeout: 1s
```

The repository's local configuration is in `test/.config/server.yml`.

## 🚢 Deployment

The service builds as a single binary and can also be built into a Docker image
through the shared make targets. In production-like environments, run it behind
your normal container orchestration and health-check configuration.

## 🛠️ Development

### 🧱 Structure

The project follows the common Go service layout:

| Path | Purpose |
| ---- | ------- |
| `main.go` | CLI bootstrap. |
| `internal/cmd/` | Server command registration and module wiring. |
| `internal/config/` | Application config layered on `github.com/alexfalkowski/go-service/v2/config`. |
| `internal/api/v1/transport/http/` | HTTP route registration for `/v1/status/{code}`. |
| `internal/health/` | Health registration and HTTP observers. |
| `test/` | Ruby nonnative/cucumber integration tests and benchmark harness. |

### 📦 Dependencies

Install these before running the full local workflow:

- Go `1.26.0`, as declared in `go.mod`.
- Ruby and Bundler for the `test/` harness.
- The `bin` submodule, initialized with `git submodule update --init`. The
  configured submodule URL uses GitHub SSH unless you override it locally.

### 🧰 Commands

Prefer the exposed `make` targets from the repository root:

| Command | Purpose |
| ------- | ------- |
| `make help` | Show available commands. |
| `make dep` | Install Go and Ruby test dependencies. |
| `make build` | Build the release binary named `status`. |
| `make build-test` | Build the test binary with feature tags and coverage instrumentation. |
| `make lint` | Lint Go and the Ruby test harness. |
| `make specs` | Run Go tests with race and coverage reporting. |
| `make features` | Build the feature test binary and run cucumber features; Nonnative starts the service on port `11000`. |
| `make benchmarks` | Build the release binary and run cucumber benchmarks; Nonnative starts the service on port `11000`. |
| `make coverage` | Generate HTML and function coverage reports. |
| `make sec` | Run configured security checks. |
| `make dev` | Run the server in watch mode with `air`. |

> [!CAUTION]
> Some shared git helper targets are intentionally destructive, including `make reset`, `make purge`, and branch deletion helpers. Use `make help` and inspect the target before running shared git workflow commands.

Stop any manually started server before running `make features` or
`make benchmarks`; both targets expect port `11000` to be free for the
Nonnative-managed service process.

### ✅ Validation

The main CircleCI `build-service` job runs:

```sh
make clean
make dep
make lint
make sec
make features
make benchmarks
make analyse
make coverage
```

For a local documentation-only change, `make help` is a useful smoke check that
the documented command surface is still available.
