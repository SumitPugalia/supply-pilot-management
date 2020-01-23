package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/DATA-DOG/godog/gherkin"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/endpoint"
	"gitlab.intelligentb.com/cafu/supply/pilot-management/pilotmanagement"
)

const (
	host = "http://localhost"
	port = ":8002"
)

const (
	statusEndpoint = "/supply/pilots/status"
	createEndpoint = "/supply/pilots"
	updateEndpoint = "/supply/pilots/%s"
)

var (
	requestBody   *gherkin.DocString
	requestRecord Request
	response      *http.Response
	responseBody  string
	decodedBody   endpoint.Response
	decodedPilot  endpoint.PilotView
)

// isServiceHosted checks if the server is already running
// by sending a request to the status endpoint
func isServiceHosted() error {
	var err error
	requestRecord := Request{
		Method: "GET",
		Uri:    host + port + statusEndpoint,
		Body:   nil,
		Header: nil,
	}

	response, err = requestRecord.Send()
	if err != nil {
		return err
	}

	return validateStatusCode(200)
}

func aPilotIsPresentInTheSystem() error {
	randData := fmt.Sprint(rand.Int())
	gherkinDocString := gherkin.DocString{
		Content: fmt.Sprintf(`{
        	"userId" : "%s",
        	"codeName" : "%s",
        	"supplierId" : "%s",
        	"marketId" : "%s",
        	"serviceId" : "%s"
        }`, randData, randData, randData, randData, randData),
	}

	requestBody = &gherkinDocString

	return sendRequestWithBody("createPilot", requestBody)
}

func sendCreatePilotRequest() (*http.Response, error) {
	req := Request{
		Method: http.MethodPost,
		Uri:    fmt.Sprint(host, port, createEndpoint),
		Body:   getGherkinStringAsReader(requestBody),
		Header: http.Header{"Content-Type": []string{"application/json"}},
	}
	return req.Send()
}

func sendUpdatePilotRequest() (*http.Response, error) {
	req := Request{
		Method: http.MethodPut,
		Uri:    fmt.Sprintf("%s%s"+updateEndpoint, host, port, decodedPilot.Id),
		Body:   getGherkinStringAsReader(requestBody),
		Header: http.Header{"Content-Type": []string{"application/json"}},
	}
	return req.Send()
}

func sendGetPilotRequest() (*http.Response, error) {
	if err := decodeBody(); err != nil {
		return nil, err
	}
	log.Println(fmt.Sprintf("%s%s"+updateEndpoint, host, port, decodedPilot.Id))
	req := Request{
		Method: http.MethodGet,
		Uri:    fmt.Sprintf("%s%s"+updateEndpoint, host, port, decodedPilot.Id),
		Body:   nil,
		Header: http.Header{"Content-Type": []string{"application/json"}},
	}
	return req.Send()
}

func sendRequest(requestName string) error {
	var sendFunc func() (*http.Response, error)
	var err error
	switch requestName {
	case "getPilot":
		sendFunc = sendGetPilotRequest
	default:
		return errors.New("invalid endpoint request")
	}
	response, err = sendFunc()
	return err
}

func sendRequestWithBody(requestName string, body *gherkin.DocString) error {
	var sendFunc func() (*http.Response, error)
	var err error
	requestBody = body
	switch requestName {
	case "createPilot":
		sendFunc = sendCreatePilotRequest
	case "updatePilot":
		sendFunc = sendUpdatePilotRequest
	case "getPilot":
		sendFunc = sendGetPilotRequest
	default:
		return errors.New("invalid endpoint request")
	}
	response, err = sendFunc()
	return err
}

func validateResponseErrorBody(errorMessage *gherkin.DocString) error {
	if err := decodeBody(); err != nil {
		log.Println("Failed to decode body", err)
		return err
	}

	errs := strings.Split(errorMessage.Content, ",")

	for i, v := range errs {
		if v != decodedBody.Errors[i] {
			return fmt.Errorf("the error response is not matching the requested body. want:%s,got:%s", v, decodedBody.Errors[i])
		}
	}

	return nil
}

func validateResponseBody() error {
	if err := decodeBody(); err != nil {
		log.Println("Failed to decode body", err)
		return err
	}

	var requestValue endpoint.CreatePilotRequest
	if err := loadRequestAsStruct(&requestValue); err != nil {
		log.Println("Failed to decode request", err)
		return err
	}

	if !compareReqWithResponse(requestValue, decodedPilot) {
		log.Printf("Request: %+v\n", requestValue)
		log.Printf("Decoded body: %+v\n", decodedBody)
		log.Printf("Response: %+v\n", decodedPilot)
		return errors.New("request and response are not equal")
	}

	return nil
}

func decodeBody() error {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	responseBody = string(bodyBytes)
	log.Printf("Raw response body: %s", responseBody)

	if err = json.NewDecoder(strings.NewReader(responseBody)).Decode(&decodedBody); err != nil {
		return err
	}
	return decodeBodyAsPilot()
}

func decodeBodyAsPilot() error {
	var happyBody struct {
		Data   endpoint.PilotView `json:"data"`
		Errors []string           `json:"errors"`
	}
	if err := json.NewDecoder(strings.NewReader(responseBody)).Decode(&happyBody); err != nil {
		return err
	}
	decodedPilot = happyBody.Data
	return nil
}

func loadRequestAsStruct(pilot interface{}) error {
	return json.NewDecoder(getGherkinStringAsReader(requestBody)).Decode(pilot)
}

func compareReqWithResponse(req endpoint.CreatePilotRequest, resp endpoint.PilotView) bool {
	return req.UserId == resp.UserId &&
		req.CodeName == resp.CodeName &&
		req.SupplierId == resp.SupplierId &&
		req.MarketId == resp.MarketId &&
		req.ServiceId == resp.ServiceId &&
		time.Unix(resp.UpdatedAt, 0).After(time.Now().Add(time.Second*-10)) &&
		time.Unix(resp.CreatedAt, 0).After(time.Now().Add(time.Second*-20))
}

func startApp(port string) {
	pilotmanagement.StartApp(port)
}
