package endpoint

//------------------------------------------------------------
// This file contains the response type that our system
// returns, functions to encode the system response to http response
//-------------------------------------------------------------

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"pilot-management/domain"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

//------------------------------------------------------------
//	Response Type that our system returns.
//-------------------------------------------------------------

type Response struct {
	Data       interface{} `json:"data"`
	Errors     []string    `json:"errors"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Total uint `json:"total,omitempty"`
	Pages uint `json:"pages,omitempty"`
}

//------------------------------------------------------------
//	Functions to encode the success response.
//-------------------------------------------------------------

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

//------------------------------------------------------------
//	Functions to encode the error response.
//-------------------------------------------------------------

func EncodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	e, ok := err.(validator.ValidationErrors)
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Errors: encodeV10Errors(e)})
		return
	}

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	switch {
	case errors.As(err, &syntaxError):
		msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Errors: []string{msg}})
		return

	case errors.Is(err, io.ErrUnexpectedEOF):
		msg := fmt.Sprintf("Request body contains badly-formed JSON")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Errors: []string{msg}})
		return

	case errors.As(err, &unmarshalTypeError):
		msg := fmt.Sprintf("Request body contains an invalid value for the %q field", unmarshalTypeError.Field)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Errors: []string{msg}})
		return

	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Errors: []string{msg}})
		return

	case errors.Is(err, io.EOF):
		msg := "Request body must not be empty"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Errors: []string{msg}})
		return

	case err.Error() == "http: request body too large":
		msg := "Request body must not be larger than 1MB"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Errors: []string{msg}})
		return

	default:
		e := err.Error()
		if checkForUUIDError(e) {
			msg := "Id is expected to be UUID"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Errors: []string{msg}})
			return
		}

		statusCode := codeFrom(err)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(Response{Errors: []string{e}})
		return
	}
}

//------------------------------------------------------------
//	Function to return the pagination data
//-------------------------------------------------------------

func PaginationData(totalEntries uint, totalPages uint) Pagination {
	return Pagination{
		Total: totalEntries,
		Pages: totalPages,
	}
}

//------------------------------------------------------------
//	Internal helper function
//-------------------------------------------------------------

func codeFrom(err error) int {
	switch err {
	case BadRequestError,
		domain.InvalidPilotStatus:
		return http.StatusBadRequest
	case domain.PilotDoesNotExistError:
		return http.StatusNotFound
	default:
		return http.StatusBadRequest
	}
}

func encodeV10Errors(errs validator.ValidationErrors) []string {
	var errorsSlice []string
	for _, err := range errs {
		errorsSlice = append(errorsSlice, toField(err.Field())+":"+err.Tag())
	}
	return errorsSlice
}

func toField(s string) string {
	field := []byte(s)
	field[0] = field[0] | ('a' - 'A')
	return string(field)
}

func checkForUUIDError(err string) bool {
	myRegex, _ := regexp.Compile("invalid UUID length *")
	return myRegex.MatchString(err)
}
