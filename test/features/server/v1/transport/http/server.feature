Feature: Server

  Server allows users to get different status codes.

  Scenario Outline: Set valid status code
    When I request to set the code <code> with HTTP
    Then I should receive a response with <code> from HTTP

    Examples:
      | code |
      | 200  |
      | 400  |
      | 500  |

  Scenario Outline: Set invalid status code
    When I request to set the invalid code "<code>" with HTTP
    Then I should receive a bad request response from HTTP

    Examples:
      | code    |
      | invalid |

  Scenario Outline: Set invalid sleep
    When I request to set the code "200" and invalid "<sleep>" for HTTP
    Then I should receive a bad request response from HTTP

    Examples:
      | sleep   |
      | invalid |
