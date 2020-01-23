Feature: Get Pilot
    this feature explains the behavior of Reading 

    @happy
    Scenario: Get(Read) Pilot request
        Given the service is hosted
        And a Pilot is present in the system
        When the user sends a request to "getPilot"
        Then the response should be 200
        And the response should have the requested pilot data

    @error
    Scenario Outline: Get(Read) Pilot request
        Given the service is hosted
        When the user sends a request to "getPilot" with <invalidid>
        Then the response should be 400
        And the response should have the error message

        Examples:
        | invalidid |
        | |
        | 123456 |
        | 
