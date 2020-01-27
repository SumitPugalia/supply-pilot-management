Feature: Update Pilot
    this feature explains the behavior of Updating a Pilot details 

    @happy
    Scenario: Update Pilot request
        Given the service is hosted
        And a Pilot is present in the system
        When the user sends a request to "updatePilot" with body
        """
        {
        	"codeName" : "updateCode",
        	"marketId" : "a535dfef-e5c2-4d2e-ac17-041581cd8471",
        	"serviceId" : "a535dfef-e5c2-4d2e-ac17-041581cd8471"
        }
        """
        Then the response should be 200
        And the response should have the requested pilot data

    @error
    Scenario: Update unknown Pilot request
        Given the service is hosted
        And a Pilot is present in the system
        When the user sends a request to "updatePilot" with body
        """
        {
          "unknownField" : "any field",
          "codeName" : "updateCode",
          "marketId" : "a535dfef-e5c2-4d2e-ac17-041581cd8471",
          "serviceId" : "a535dfef-e5c2-4d2e-ac17-041581cd8471"
        }
        """
        Then the response should be 400
        And the response should have the error message 
        """
        json: unknown field "unknownField"
        """

    @error
    Scenario Outline: Update Pilot request
        Given the service is hosted
        And a Pilot is present in the system
        When the user sends a request to "updatePilot" with body
        """
        {
        	"codeName" : <code>,
        	"marketId" : <mid>,
        	"serviceId" : <serid>
        }
        """
        Then the response should be 400
        And the response should have the error message 
        """
        <errorMessage>
        """
        
        Examples:
            | code | mid | serid | errorMessage                                                                                |
            | "A" | 123 | "2"   | bad request |

