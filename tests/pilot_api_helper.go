package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"pilot-management/endpoint"

	"github.com/DATA-DOG/godog/gherkin"
)

const (
	host = "http://localhost"
	port = ":8002"
)

const (
	statusEndpoint    = "/supply/pilots/status"
	createEndpoint    = "/supply/pilots"
	pilotByIDEndpoint = "/supply/pilots/%s"
)

type pilotAPIHelper struct {
	requestBody   *gherkin.DocString
	requestRecord Request
	response      *http.Response
	decodedBody   endpoint.Response
	decodedPilot  endpoint.PilotView
}

// NewPilotAPIHelper creates and returns a new pilot helper value
func NewPilotAPIHelper() *pilotAPIHelper {
	helper := new(pilotAPIHelper)
	helper.response = new(http.Response)
	helper.requestBody = new(gherkin.DocString)
	return helper
}

func (step *pilotAPIHelper) aPilotIsPresentInTheSystem() error {
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

	return step.sendRequestWithBody("createPilot", &gherkinDocString)
}

func (step *pilotAPIHelper) sendCreatePilotRequest() (*http.Response, error) {
	req := Request{
		Method: http.MethodPost,
		Uri:    fmt.Sprint(host, port, createEndpoint),
		Body:   getGherkinStringAsReader(step.requestBody),
		Header: http.Header{"Content-Type": []string{"application/json"}},
	}
	return req.Send()
}

func (step *pilotAPIHelper) sendUpdatePilotRequest() (*http.Response, error) {
	if err := step.decodeBodyAsPilot(); err != nil {
		return nil, err
	}

	req := Request{
		Method: http.MethodPatch,
		Uri:    fmt.Sprintf("%s%s"+pilotByIDEndpoint, host, port, step.decodedPilot.Id),
		Body:   getGherkinStringAsReader(step.requestBody),
		Header: http.Header{"Content-Type": []string{"application/json"}},
	}
	return req.Send()
}

func (step *pilotAPIHelper) getPilotRequestValidID() error {
	if err := step.decodeBodyAsPilot(); err != nil {
		return err
	}

	resp, err := step.getPilotRequest(step.decodedPilot.Id)
	if err != nil {
		return err
	}
	step.rememberResponse(resp)
	return nil
}
func (step *pilotAPIHelper) getPilotRequestInvalidID(invalidID string) error {

	resp, err := step.getPilotRequest(invalidID)
	if err != nil {
		return err
	}
	step.rememberResponse(resp)
	return nil
}

func (step *pilotAPIHelper) getPilotRequest(ID string) (*http.Response, error) {
	//log.Println(fmt.Sprintf("%s%s"+updateEndpoint, host, port, decodedPilot.Id))
	req := Request{
		Method: http.MethodGet,
		Uri:    fmt.Sprintf("%s%s"+pilotByIDEndpoint, host, port, ID),
		Body:   nil,
		Header: http.Header{"Content-Type": []string{"application/json"}},
	}
	return req.Send()
}

func (step *pilotAPIHelper) sendRequest(requestName string) error {
	resp, err := step.determineRequestFunc(requestName)()
	step.rememberResponse(resp)
	return err
}

func (step *pilotAPIHelper) sendRequestWithBody(requestName string, body *gherkin.DocString) error {
	step.rememberRequestBody(body)
	resp, err := step.determineRequestFunc(requestName)()
	step.rememberResponse(resp)

	return err
}

func (step *pilotAPIHelper) rememberRequestBody(body *gherkin.DocString) {
	step.requestBody = body
}

func (step *pilotAPIHelper) rememberResponse(resp *http.Response) {
	step.response = resp
}

func (step *pilotAPIHelper) determineRequestFunc(requestName string) func() (*http.Response, error) {
	switch requestName {
	case "createPilot":
		return step.sendCreatePilotRequest
	case "updatePilot":
		return step.sendUpdatePilotRequest
	default:
		return func() (*http.Response, error) {
			return nil, errors.New("invalid request name")
		}
	}
}

func (step *pilotAPIHelper) validateStatusCode(code int) error {
	return validateStatusCode(step.response, code)
}

func (step *pilotAPIHelper) validateResponseErrorBody(errorMessage *gherkin.DocString) error {
	if err := step.decodeBodyAsPilot(); err != nil {
		log.Println("Failed to decode body", err)
		return err
	}

	errs := strings.Split(errorMessage.Content, ",")
	log.Printf("step decoded body %+v", step.decodedBody)
	for i, v := range errs {
		if v != step.decodedBody.Errors[i] {
			return fmt.Errorf("the error response doesn't match. want:%s,got:%s", v, step.decodedBody.Errors[i])
		}
	}

	return nil
}

func (step *pilotAPIHelper) validateResponseBody() error {
	if err := step.decodeBodyAsPilot(); err != nil {
		log.Println("Failed to decode body", err)
		return err
	}

	var requestValue endpoint.CreatePilotRequest
	if err := step.loadRequestAsStruct(&requestValue); err != nil {
		log.Println("Failed to decode request", err)
		return err
	}

	if !compareReqWithResponse(requestValue, step.decodedPilot) {
		log.Printf("Request: %+v\n", requestValue)
		log.Printf("Decoded body: %+v\n", step.decodedBody)
		log.Printf("Response: %+v\n", step.decodedPilot)
		return errors.New("request and response are not equal")
	}

	return nil
}

func (step *pilotAPIHelper) decodeBodyAsPilot() error {
	happyBody := new(struct {
		Data       endpoint.PilotView  `json:"data"`
		Errors     []string            `json:"errors"`
		Pagination endpoint.Pagination `json:"pagination"`
	})
	log.Println("Response: ", step.response)
	if err := json.NewDecoder(step.response.Body).Decode(happyBody); err != nil {
		return err
	}
	step.decodedBody = endpoint.Response{
		Data:       happyBody.Data,
		Errors:     happyBody.Errors,
		Pagination: happyBody.Pagination,
	}
	step.decodedPilot = happyBody.Data
	return nil
}

func (step *pilotAPIHelper) loadRequestAsStruct(pilot interface{}) error {
	return json.Unmarshal([]byte(step.requestBody.Content), pilot)
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
	endpoint.StartApp(port)
}

// isServiceHosted checks if the server is already running
// by sending a request to the status endpoint
func isServiceHosted() error {
	var err error
	requestRecord := Request{
		Method: http.MethodGet,
		Uri:    host + port + statusEndpoint,
		Body:   nil,
		Header: nil,
	}

	resp, err := requestRecord.Send()
	if err != nil {
		return err
	}

	return validateStatusCode(resp, 200)
}
