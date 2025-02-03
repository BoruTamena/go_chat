
Feature: I want to test user registeration

@success
Scenario Outline: successfully user registeration

When users request to register themself with a valid data

    |user_name|email|password|
    |<user_name>|<email>|<password>|

Then the system should register user successfully

Examples:
    |user_name|email|password|
    |boru|user@gmail.com|12345we|


