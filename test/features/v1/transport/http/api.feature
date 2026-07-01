Feature: Server

  Server allows users to get different status codes.

  Scenario Outline: Set valid status code
    When I request to set the code <code>
    Then I should receive a response with <code> and "<body>"

    Examples:
      | code | body                      |
      | 200  | 200 OK                    |
      | 400  | 400 Bad Request           |
      | 500  | 500 Internal Server Error |
      | 999  | 999                       |

  Scenario: Set valid status code without sleep
    When I request to set the code 200 without sleep
    Then I should receive a response with 200 and "200 OK"

  Scenario Outline: Set valid status code with method
    When I request to set the code 503 with "<method>"
    Then I should receive a response with 503 and "503 Service Unavailable"

    Examples:
      | method |
      | POST   |
      | PUT    |
      | PATCH  |
      | DELETE |

  Scenario: Set valid sleep
    When I request to set the code 200 and sleep "50ms"
    Then I should receive a response with 200 and "200 OK"
    And I should receive the response in at least 50 ms

  Scenario: Set redirect location
    When I request to set the code 302 and location "/redirected"
    Then I should receive a response with 302 and "302 Found"
    And I should receive a location "/redirected"

  Scenario: Set retry after
    When I request to set the code 503 and retry after "500ms"
    Then I should receive a response with 503 and "503 Service Unavailable"
    And I should receive a retry after "1"

  Scenario: Set extra response headers
    When I request to set the code 429 and headers
      | X-Rate-Limit-Remaining | 42  |
      | X-Rate-Limit-Limit     | 100 |
    Then I should receive a response with 429 and "429 Too Many Requests"
    And I should receive a header "X-Rate-Limit-Remaining" "42"
    And I should receive a header "X-Rate-Limit-Limit" "100"

  Scenario: Reject sleep above maximum
    When I request to set the code 200 and sleep "5m1s"
    Then I should receive a bad request response
    And I should receive the response in less than 1000 ms

  Scenario: Reject redirect location for non-redirect status
    When I request to set the code 200 and location "/redirected"
    Then I should receive a bad request response

  Scenario: Reject unsafe redirect location
    When I request to set the code 302 and location "/redirected%0Abad"
    Then I should receive a bad request response

  Scenario: Reject retry after for non-retry status
    When I request to set the code 200 and retry after "2s"
    Then I should receive a bad request response

  Scenario: Reject malformed extra response header
    When I request to set the code 200 and raw header "X-Trace-Id"
    Then I should receive a bad request response

  Scenario Outline: Reject invalid extra response header
    When I request to set the code 200 and header "<header>" "<value>"
    Then I should receive a bad request response

    Examples:
      | header       | value      |
      | Bad Header   | value      |
      | X-Trace-Id   | abc%0Abad  |
      | Content-Type | text/plain |

  Scenario Outline: Set invalid status code
    When I request to set the invalid code "<code>"
    Then I should receive a bad request response

    Examples:
      | code    |
      | 42      |
      | 100     |
      | 199     |
      | 1000    |
      | invalid |

  Scenario Outline: Set invalid sleep
    When I request to set the code "200" and invalid "<sleep>"
    Then I should receive a bad request response

    Examples:
      | sleep   |
      | invalid |

  Scenario Outline: Set invalid retry after
    When I request to set the code 503 and retry after "<retry_after>"
    Then I should receive a bad request response

    Examples:
      | retry_after |
      | invalid     |
      | 0s          |
      | -1s         |
