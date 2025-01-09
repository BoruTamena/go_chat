Feature: Group Chat Feature 

Background: Group Members
Given all the following users are active
|user_id|group_name|
|123user1  | G1|
|456user2  | G1|
|789user3  | G1|
|231user4  | G2|
@success
Scenario: Text Group

When the user from the group send text message
|user_id    |group_name |message|
|123user1   | G1        |hi there|
Then all  members of  group  should receive the same message
 |group_name| message_received |
 |G1        | hi there|  

And the user from different group should not receive the message 


