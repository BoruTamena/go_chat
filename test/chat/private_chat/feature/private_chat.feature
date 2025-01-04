Feature: Private Chat Feature

  As a user, I want to test the private chat functionality 
  to ensure users can exchange messages successfully.

  Scenario: Text friend successfully
    Given two users are online
      | user_id    |
      | 123User1   |
      | 456User2   |
    When the first user texts a friend with a message
      | friend_id  | message        |
      | 456User2 | Hello, how are you?     |
    Then the friend should receive the same message
     |user_id | message_received |
     |456User2 | Hello, how are you?      |
    And the friend texts the user back with a message
      | friend_id  | response        |
      | 123User1 |I'm good, thanks!    |
    Then the user receive the same message from friend
        |user_id | message_received |
        | 123User1 |I'm good, thanks!    |

    # Examples:
    #   | friend_id | user_id    | message         | message_received | response         |
    #   | 456User2  | 123User1   | Hello, how are you? | Hello, how are you? | I'm good, thanks! |
