@benchmark
Feature: Benchmark HTTP API
  Make sure these endpoints perform at their best.

  Scenario: Set code in a good time frame and memory.
    When I request with HTTP which performs in 15 ms
    And the process 'server' should consume less than '70mb' of memory
