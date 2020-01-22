package tests

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"
)

var client http.Client

type Request struct {
	Method string
	Uri    string
	Body   io.Reader
	Header http.Header
}

func (r Request) Send() (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.Uri, r.Body)
	if err != nil {
		return nil, err
	}
	addHeader(req, r.Header) //if any

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getGherkinStringAsReader(gs *gherkin.DocString) io.Reader {
	return strings.NewReader(gs.Content)
}

func validateStatusCode(expected int) error {
	if response.StatusCode != expected {
		return fmt.Errorf("status code did not match. want:%d got:%d", expected, response.StatusCode)
	}
	return nil
}

func addHeader(req *http.Request, header http.Header) {
	for k, v := range header {
		for _, val := range v {
			req.Header.Add(k, val)
		}
	}
}
