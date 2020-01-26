Feature: Create Pilot
    this feature explains the behavior of creating a Pilot

    @happy
    Scenario: Create pilot request
        Given the service is hosted
        When the user sends a request to "createPilot" with body
        """
        {
        	"userId" : "mraj89",
        	"codeName" : "Mohanraj",
        	"supplierId" : "Sup123",
        	"marketId" : "Mar123",
        	"serviceId" : "Serv123"
        }
        """
        Then the response should be 200
        And the response should have the requested pilot data

    @error
    Scenario Outline: Create pilot request with invalid request
        Given the service is hosted
        When the user sends a request to "createPilot" with body
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
            | 123 | "SAD"| "A" | "W" | "2"   | userId Expected string But Got number |