Feature: Group Chat Feature 

Background: Group Members
Given all the following users are active
|user_id|group_name|
|user1  | G1|
|user2  | G1|
|user3  | G1|
|user4  | G2|

@success
Scenario: Text Group

When the user from the group send text message

Then all  members of  group  should receive the same message

And the user from different group should not receive the message 


