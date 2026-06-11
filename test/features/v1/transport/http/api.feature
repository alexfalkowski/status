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

  Scenario: Set valid sleep
    When I request to set the code 200 and sleep "50ms"
    Then I should receive a response with 200 and "200 OK"
    And I should receive the response in at least 50 ms

  Scenario: Reject sleep above maximum
    When I request to set the code 200 and sleep "5m1s"
    Then I should receive a bad request response
    And I should receive the response in less than 1000 ms

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
