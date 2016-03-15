Feature: List tweets

  Scenario: Get JSON list of tweets sorted by date descending
    Given the system knows about the following tweets:
      | body         | created_at               |
      | First tweet  | 2014-02-10T10:50:42.389Z |
      | Second tweet | 2014-02-10T10:50:57.240Z |
    When the client requests GET /tweets
    Then the response should be JSON:
      """
      [
        {"body": "Second tweet", "created_at": "2014-02-10T10:50:57.240Z"},
        {"body": "First tweet", "created_at": "2014-02-10T10:50:42.389Z"}
      ]
      """
