Feature: Observability

  Health checks external connectivity, liveness and readiness use local no-op checks,
  and metrics expose Prometheus data.

  Scenario: Health
    When the system requests the "health"
    Then the system should respond with a healthy status

  Scenario: Liveness
    When the system requests the "liveness"
    Then the system should respond with a healthy status

  Scenario: Readiness
    When the system requests the "readiness"
    Then the system should respond with a healthy status

  Scenario: Metrics
    When the system requests the "metrics"
    Then the system should respond with metrics
