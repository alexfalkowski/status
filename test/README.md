# 🧪 Test Harness

The `test/` directory owns the Ruby Cucumber/Nonnative integration and
benchmark harness for the `status` service. Prefer the root `make` targets for
normal local and CI use because they build the correct Go binary before running
the Ruby harness.

## ▶️ Commands

From the repository root:

```sh
make dep
make features
make benchmarks
```

`make dep` installs the Ruby gems under `test/vendor/bundle`. `make features`
builds the feature test binary and runs non-benchmark Cucumber scenarios.
`make benchmarks` builds the release binary and runs scenarios tagged
`@benchmark`.

Direct harness runs are supported only after the matching root build has
created `./status`:

```sh
make build-test
make -C test features

make build
make -C test benchmarks
```

## ⚙️ Runtime Contract

Nonnative starts the service from `test/nonnative.yml` with:

```sh
../status server -config file:.config/server.yml
```

Run direct harness commands through `make -C test` from the repository root, or
change into `test/` and run the same target there. The Nonnative config uses
paths relative to `test/`. The service listens on `http://localhost:11000`, and
port `11000` must be free before the harness starts.

The local server config is `test/.config/server.yml`. It enables Prometheus
metrics, sets `max_sleep: 2m`, and configures the HTTP transport timeout to
`5s`.

## 📈 Benchmarks

The benchmark feature checks one `GET /v1/status/200?sleep=1ms` request against
the current local thresholds:

- response under `15ms`;
- server process memory below `70mb`.

Treat local benchmark failures as a prompt to inspect load, machine resources,
and generated artifacts before assuming a service regression.

## 🧾 Reports

Cucumber and Nonnative write artifacts under `test/reports/`:

- `index.html`: HTML Cucumber report.
- `nonnative.log`: Nonnative process orchestration log.
- `server.log`: service process log.
- `*.xml`: JUnit-style Cucumber reports.
- `*.cov`: Go coverage output from feature builds.

The root `make coverage` target consumes coverage files from this directory and
writes `test/reports/coverage.html`.
