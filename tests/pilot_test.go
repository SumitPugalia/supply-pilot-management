package tests

import (
	"log"
	"time"

	"github.com/DATA-DOG/godog"
)

func FeatureContext(s *godog.Suite) {
	s.BeforeSuite(func() {
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		if err := isServiceHosted(); err != nil {
			go startApp(port)
			time.Sleep(time.Second) //waiting for the server to start
		}
	})

	s.Step(`^the service is hosted$`, isServiceHosted)
	s.Step(`^the user sends a request to "([^"]*)" with body$`, sendRequestWithBody)
	s.Step(`^the response should be (\d+)$`, validateStatusCode)
	s.Step(`^the response should have the input data$`, validateResponseBody)
	s.Step(`^the response should have the error message$`, validateResponseErrorBody)
}
