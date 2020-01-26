Feature: Get Pilot
    this feature explains the behavior of Reading 

    @happy
    Scenario: Get(Read) Pilot request
        Given the service is hosted
        And a Pilot is present in the system
        When the user sends a GET request with valid pilot id
        Then the response should be 200
        And the response should have the requested pilot data

    @error
    Scenario Outline: Get(Read) Pilot request
        Given the service is hosted
        When the user sends a GET request with invalid pilot id <invalidid>
        Then the response should be 404
        And the response should have the error message
        """
        pilot does not exist
        """

        Examples:
        | invalidid    |
        | "123456"     | 
        | "werwfdfsfd" |
