Feature: login
  User login

  Scenario: User Login
    When I send login request with data by map
      |username |password   |
      |abal     |mysecretpw |
    Then the response must match data by map
      |code     |message        |
      |00000    |Login success  |
    And the response data should match json:
      """
        {
          "id":1,
          "username":"abal",
          "email":"abal@email.com",
          "role":"ADMIN"
        }
      """
