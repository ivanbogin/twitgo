Feature: Listing tweets

  Scenario: Get JSON list of tweets sorted by date descending
    Given the system knows about the following tweets:
      | body         | datetime      |
      | First tweet  | 1416445412987 |
      | Second tweet | 1416445411987 |
    When the client requests GET /tweets
    Then the response should be JSON:
      """
      [
        {"body": "Second tweet", "datetime": "1416445411987"},
        {"body": "First tweet", "datetime": "1416445412987"}
      ]
      """
