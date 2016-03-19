Feature: List tweets

  Scenario: Get JSON list of tweets sorted by date descending
    Given the system knows about the following tweets:
      | body         | created_at                |
      | First tweet  | 2016-03-19T00:15:33+01:00 |
      | Second tweet | 2016-03-20T00:15:33+01:00 |
    When the client requests GET /tweets
    Then the response should be JSON:
      """
      [
        {"body": "Second tweet", "created_at": "2016-03-20T00:15:33+01:00"},
        {"body": "First tweet", "created_at": "2016-03-19T00:15:33+01:00"}
      ]
      """
