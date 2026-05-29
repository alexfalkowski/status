Feature: Observability

  Observability is a measure of how well internal states of a system can be inferred by knowledge of its external outputs.

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
