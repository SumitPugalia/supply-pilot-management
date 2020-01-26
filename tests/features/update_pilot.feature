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
        	"userId" : <uid>,
        	"codeName" : <code>,
        	"supplierId" : <sid>,
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
            | uid | code | sid | mid | serid | errorMessage                                                                                |
            | ""  | ""   | ""  | ""  | ""    | UserId:required,CodeName:required,SupplierId:required,MarketId:required,ServiceId:required  |
            | 123 | "SAD"| "A" | "W" | "2"   | json: cannot unmarshal number into Go struct field CreatePilotRequest.userId of type string |

