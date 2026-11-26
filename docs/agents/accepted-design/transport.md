# Accepted design for transport

These entries are part of the [accepted-design index](../accepted-design.md).

## Simulated response delays

Status deliberately simulates HTTP errors and response delays for tests. The
`/v1/status/{code}` handler parses `sleep` with `time.ParseDuration` and rejects
values above the effective `max_sleep`. A delay within that limit can exceed
the configured HTTP request timeout. The local test configuration uses a `2m`
sleep limit and a `5s` HTTP request timeout.

Treat this as expected fixture behavior in reliability reviews. Revisit the
decision when the task explicitly concerns delay semantics or concrete evidence
shows that the behavior breaks a current workflow.

[Status handler](../../../internal/api/v1/transport/http/http.go) ·
[Sleep limit configuration](../../../internal/config/config.go) ·
[Local test configuration](../../../test/.config/server.yaml)
