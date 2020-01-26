Feature: Update Pilot
    this feature explains the behavior of Updating a Pilot details 

    @happy
    Scenario: Update Pilot request
        Given the service is hosted
        And a Pilot is present in the system
        When the user sends a request to "updatePilot" with body
        """
        {
        	"userId" : "updateUser",
        	"codeName" : "updateCode",
        	"supplierId" : "updateSupplier",
        	"marketId" : "updatemarket",
        	"serviceId" : "updateservice"
        }
        """
        Then the response should be 200
        And the response should have the requested pilot data

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
            | "A" | 123 | "2"   | marketId Expected string But Got number |

