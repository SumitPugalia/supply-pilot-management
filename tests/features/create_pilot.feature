Feature: Create Pilot
    this feature explains the behavior of creating a Pilot

    @happy
    Scenario: Create pilot request
        Given the service is hosted
        When the user sends a request to "createPilot" with body
        """
        {
        	"userId" : "f8590d99-4e29-46be-a39a-8aa9574bcb2b",
        	"codeName" : "Mohanraj",
        	"supplierId" : "f8590d99-4e29-46be-a39a-8aa9574bcb2b",
        	"marketId" : "f8590d99-4e29-46be-a39a-8aa9574bcb2b",
        	"serviceId" : "f8590d99-4e29-46be-a39a-8aa9574bcb2b"
        }
        """
        Then the response should be 200
        And the response should have the requested pilot data

    Scenario: Create unknown pilot request
        Given the service is hosted
        When the user sends a request to "createPilot" with body
        """
        {
          "unknownField" : "any field",
          "userId" : "f8590d99-4e29-46be-a39a-8aa9574bcb2b",
          "codeName" : "Mohanraj",
          "supplierId" : "f8590d99-4e29-46be-a39a-8aa9574bcb2b",
          "marketId" : "f8590d99-4e29-46be-a39a-8aa9574bcb2b",
          "serviceId" : "f8590d99-4e29-46be-a39a-8aa9574bcb2b"
        }
        """
        Then the response should be 400
        And the response should have the error message 
        """
        Request body contains unknown field "unknownField"
        """

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
            | ""  | ""   | ""  | ""  | ""    | Id is expected to be UUID |
            | 123 | "SAD"| "A" | "W" | "2"   | Id is expected to be UUID |
            | "123" | "SAD"| "A" | "W" | "2"   | Id is expected to be UUID |
            | "f8590d99-4e29-46be-a39a-8aa9574bcb2b" | 123 | "f8590d99-4e29-46be-a39a-8aa9574bcb2b" | "f8590d99-4e29-46be-a39a-8aa9574bcb2b" | "f8590d99-4e29-46be-a39a-8aa9574bcb2b" | Request body contains an invalid value for the "codeName" field |