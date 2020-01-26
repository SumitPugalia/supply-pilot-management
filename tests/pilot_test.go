package tests

import (
	"log"
	"time"

	"github.com/DATA-DOG/godog"
)

var step *pilotAPIHelper

func FeatureContext(s *godog.Suite) {
	step = NewPilotAPIHelper()
	s.BeforeSuite(func() {
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		if err := isServiceHosted(); err != nil {
			go startApp(port)
			time.Sleep(time.Second) //waiting for the server to start
		}
	})

	s.Step(`^the service is hosted$`, isServiceHosted)
	s.Step(`^the user sends a request to "([^"]*)" with body$`, step.sendRequestWithBody)
	s.Step(`^the response should be (\d+)$`, step.validateStatusCode)
	s.Step(`^the response should have the error message$`, step.validateResponseErrorBody)
	s.Step(`^a Pilot is present in the system$`, step.aPilotIsPresentInTheSystem)
	//s.Step(`^the user sends a request to "([^"]*)"$`, step.sendRequest)
	s.Step(`^the response should have the requested pilot data$`, step.validateResponseBody)
	s.Step(`^the user sends a GET request with invalid pilot id "([^"]*)"$`, step.getPilotRequestInvalidID)
	s.Step(`^the user sends a GET request with valid pilot id$`, step.getPilotRequestValidID)
}
