package endpoint

//------------------------------------------------------------
// This file contains the response type that our system
// returns, functions to encode the system response to http response
//-------------------------------------------------------------

import (
	"context"
	"encoding/json"
	"net/http"
	"pilot-management/domain"

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
//	Functions to encode the success/error response.
//-------------------------------------------------------------

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if errs, ok := err.(validator.ValidationErrors); ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Errors: encodeV10Errors(errs)})
		return
	}

	if e, k := err.(*json.UnmarshalTypeError); k {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Errors: encodeUnmarshalTypeErrors(e)})
		return
	}

	statusCode := codeFrom(err)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{Errors: []string{err.Error()}})
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
		return http.StatusInternalServerError
	}
}

func encodeV10Errors(errs validator.ValidationErrors) []string {
	var errorsSlice []string
	for _, err := range errs {
		errorsSlice = append(errorsSlice, err.Field()+":"+err.Tag())
	}
	return errorsSlice
}

func encodeUnmarshalTypeErrors(e *json.UnmarshalTypeError) []string {
	msg := e.Field + " Expected " + e.Type.String() + " But Got " + e.Value
	return []string{msg}
}
