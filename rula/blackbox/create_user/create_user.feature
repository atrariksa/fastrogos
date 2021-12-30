Feature: create user
  In order to be able to login
  I need to create user

  Scenario: Create New User
    When I send create user request with data by map
      |username |email          |role  |password   |
      |abal     |abal@email.com |ADMIN |mysecretpw |
    Then the response must match data by map
      |code     |message        |
      |00000    |User created   |
