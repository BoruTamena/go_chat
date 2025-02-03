Feature: Login 

I want to test system login as credential user

@success 
Scenario: Login successfully

Given user is a registered user

When the user login with their credential

    |email         |password|
    |user@gmail.com|password|

Then the system should authorize the user 